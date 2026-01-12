package fncHoldst

import (
	fncGlobal "back/global/function"
	mdlHoldst "back/holdst/model"
	mdlPsglst "back/psglst/model"
	fncSbrapi "back/sbrapi/function"
	mdlSbrapi "back/sbrapi/model"
	"fmt"
	"slices"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Running process hit passanggerlist daily
func FncHoldstPrcessMainpg(c *gin.Context) {

	// protect single run
	if fncGlobal.Status.Sbrapi != 0.0 {
		return
	}
	fncGlobal.Status.Sbrapi = 0.01

	// Bind JSON Body input to variable
	inpErrlog := mdlHoldst.MdlHoldstErrlogDtbase{} //save
	if err := c.BindJSON(&inpErrlog); err != nil {
		panic(err)
	}

	// Declare date format
	strTimenw := time.Now().Format("0601021504")
	// intTimenw, _ := strconv.Atoi(strTimenw)
	// intDatenw, _ := strconv.Atoi(strTimenw[0:6])
	strDatepv := time.Now().AddDate(0, 0, -1)
	if inpErrlog.Datefl != 0 {
		tmpDatefl := strconv.Itoa(int(inpErrlog.Datefl))
		strDatepv, _ = time.Parse("060102", tmpDatefl)
	}
	strDatefl := strDatepv.Format("060102")
	intDatefl, _ := strconv.Atoi(strDatefl)

	// Declare airline
	slcAirlfl := []string{inpErrlog.Airlfl}
	if inpErrlog.Airlfl == "" {
		slcAirlfl = []string{"JT"}
	}

	// Declare Depart
	slcDepart := []string{inpErrlog.Depart}
	if inpErrlog.Depart == "" {
		slcDepart = []string{"CGK"}
	}
	fmt.Println(slcDepart)

	// Declare Flight number
	slcFlnbfl := []string{inpErrlog.Flnbfl}
	if inpErrlog.Flnbfl == "" {
		slcFlnbfl = []string{"All"}
	}

	// Indicator done data
	var nowTotdta = int64(0)
	var maxTotdta = float64(len(slcAirlfl) * len(slcDepart))
	var totWorker = inpErrlog.Worker

	// Looping slice airlines
	for _, airlfl := range slcAirlfl {
		fmt.Println("Processing airline:", airlfl, totWorker)

		// Get Multiple API sessions/tokens
		slcRspssn, err := fncSbrapi.FncSbrapiCrtssnMultpl(airlfl, int(totWorker))
		lgcRspssn := err != nil || slcRspssn[0].Bsttkn == "" || len(slcRspssn) < 1
		if lgcRspssn {
			fmt.Println("Failed")
			continue
		}

		// Prepare job queue
		jobFllist := make(chan mdlPsglst.MdlPsglstFllistDtbase, 1500)
		var swg sync.WaitGroup

		// Launch 10 workers using 10 tokens
		for i := 0; i < int(totWorker); i++ {
			if len(slcRspssn) >= i+1 {
				if slcRspssn[i].Bsttkn != "" {
					swg.Add(1)
					fmt.Println("Success Token-", i)
					go FncHoldstPrcessWorker(slcRspssn[i],
						&swg,
						jobFllist,
						strTimenw)
					continue
				}
				fmt.Println("Failed Token-", i)
			}
		}

		// Looping slice departures
		for _, depart := range slcDepart {

			// update indicator progress
			atomic.AddInt64(&nowTotdta, 1)
			if maxTotdta > 0 {
				fncGlobal.Status.Sbrapi = float64(nowTotdta) / maxTotdta * 100
			}

			// Get API Flight List data
			rawFllist, err := fncSbrapi.FncSbrapiFllistMainob(slcRspssn[0],
				mdlSbrapi.MdlSbrapiMsghdrApndix{Datefl: int32(intDatefl),
					Airlfl: airlfl, Depart: depart})
			if err != nil {
				fmt.Println("faild")
				continue
			}

			// Looping Flight List
			for _, fllist := range rawFllist {

				// Only accept on this route
				if slices.Contains(slcFlnbfl, fllist.Flnbfl) ||
					slices.Contains(slcFlnbfl, "All") {
					jobFllist <- fllist
				}
			}
		}

		// Finish
		close(jobFllist)
		swg.Wait()
		fmt.Printf("Done airline:%s time:%s \n", airlfl,
			time.Now().Format("06-Jan-02/15:04"))
		fncSbrapi.FncSbrapiClsssnMultpl(slcRspssn)
	}

	// Detect error and count it
	fncGlobal.Status.Sbrapi = 0

}

// Running process jeddah
func FncHoldstPrcessWorker(
	tkn mdlSbrapi.MdlSbrapiMsghdrParams,
	swg *sync.WaitGroup,
	jobFllist <-chan mdlPsglst.MdlPsglstFllistDtbase,
	strTimenw string) {

	// Declare global variable
	defer swg.Done()
	var mgoVavhst, mgoPaxhst, mgoRawhst []mongo.WriteModel

	// iterate jobs
	cntdta := 0
	for fllist := range jobFllist {
		cntdta++

		// prepare locals
		var intDatefl = fllist.Datefl
		var dbsFlnbfl, dbsDepart, dbsArrivl = fllist.Flnbfl, fllist.Depart, fllist.Arrivl
		var dbsRoutfl, dbsAirlfl = fllist.Routfl, fllist.Airlfl
		var objParams = mdlSbrapi.MdlSbrapiMsghdrApndix{
			Airlfl: dbsAirlfl, Datefl: intDatefl, Depart: dbsDepart,
			Arrivl: dbsArrivl, Flnbfl: dbsFlnbfl, Routfl: dbsRoutfl}

		// Indicator start process
		fmt.Println("Start", cntdta, "-", dbsAirlfl, dbsFlnbfl, intDatefl,
			dbsRoutfl, time.Now().Format("06-01-02/15:04"))

		// Get data Paxhst
		slcPaxhst, slcRawhst, err := fncSbrapi.FncSbrapiPaxhstMainob(tkn, objParams, 10)
		mapPaxhst := map[string]mdlSbrapi.MdlSbrapiPaxhstDtbase{}
		if err == nil {
			for _, Paxhst := range slcPaxhst {
				mapPaxhst[Paxhst.Prmkey] = Paxhst
				mgoPaxhst = append(mgoPaxhst, mongo.NewUpdateOneModel().
					SetFilter(bson.M{"prmkey": Paxhst.Prmkey}).
					SetUpdate(bson.M{"$set": Paxhst}).
					SetUpsert(true))
			}
			for _, rawhst := range slcRawhst {
				mgoRawhst = append(mgoRawhst, mongo.NewUpdateOneModel().
					SetFilter(bson.M{"prmkey": rawhst.Prmkey}).
					SetUpdate(bson.M{"$set": rawhst}).
					SetUpsert(true))
			}
		}

		// Get data Vavhst
		slcVavhst, err := fncSbrapi.FncSbrapiVavhstMainob(tkn, objParams)
		if err == nil {
			for _, Vavhst := range slcVavhst {
				mgoVavhst = append(mgoVavhst, mongo.NewUpdateOneModel().
					SetFilter(bson.M{"prmkey": Vavhst.Prmkey}).
					SetUpdate(bson.M{"$set": Vavhst}).
					SetUpsert(true))
			}
		}

		// Push to database
		fncGlobal.FncGlobalDtbaseBtcwrt(map[string]*[]mongo.WriteModel{
			"holdst_vavhst": &mgoVavhst,
			"holdst_paxhst": &mgoPaxhst,
			"holdst_rawhst": &mgoRawhst,
		}, 200)
	}

	// Push to database
	fncGlobal.FncGlobalDtbaseBtcwrt(map[string]*[]mongo.WriteModel{
		"holdst_vavhst": &mgoVavhst,
		"holdst_paxhst": &mgoPaxhst,
		"holdst_rawhst": &mgoRawhst,
	}, 0)

}
