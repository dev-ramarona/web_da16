package fncJeddah

import (
	fncGlobal "back/global/function"
	mdlJeddah "back/jeddah/model"
	"context"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get Agent name not complete format slice data from database
func FncJeddahAgtnmeNullnm(c *gin.Context) {

	// Input data post
	var inputx mdlJeddah.MdlJeddahAllprmInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var totalx = 0
	var arrobj = []mdlJeddah.MdlJeddahAgtnmeInputx{}
	var sycwgp sync.WaitGroup

	// Check if data Route all is isset
	if inputx.Srtnul_agtnme == "" {
		mtchdt = append(mtchdt, bson.D{{Key: "agtidn",
			Value: ""}})
	} else {
		if inputx.Airlfl_agtnme != "" {
			mtchdt = append(mtchdt, bson.D{{Key: "airlnf",
				Value: inputx.Airlfl_agtnme}})
		}
		if inputx.Agtnme_agtnme != "" {
			nowagt := inputx.Agtnme_agtnme
			mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: bson.A{
				bson.D{{Key: "agtnme", Value: bson.D{{Key: "$regex", Value: nowagt}}}},
				bson.D{{Key: "agtdtl", Value: bson.D{{Key: "$regex", Value: nowagt}}}}}}})
		}
	}

	// Final match pipeline
	var mtchfn bson.D
	if len(mtchdt) != 0 {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: mtchdt}}}}
	} else {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{}}}
	}

	// Select database and collection
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_agentx")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get Total Count Data
	sycwgp.Add(1)
	go func() {
		defer sycwgp.Done()
		pipeln := mongo.Pipeline{
			mtchfn,
			bson.D{{Key: "$count", Value: "totalCount"}},
		}

		// Find user by username in database
		rawdta, err := tablex.Aggregate(contxt, pipeln)
		if err != nil {
			panic(err)
		}
		defer rawdta.Close(contxt)

		// Store to slice from raw bson
		var slcbsn []bson.M
		if err = rawdta.All(contxt, &slcbsn); err != nil {
			panic(err)
		}

		// Mengambil jumlah dokumen dari hasil
		if len(slcbsn) > 0 {
			if countd, ist := slcbsn[0]["totalCount"].(int32); ist {
				totalx = int(countd) // Konversi dari int32 ke int
			}
		}
	}()

	// Get route data
	sycwgp.Add(1)
	go func() {
		defer sycwgp.Done()
		piplne := mongo.Pipeline{mtchfn,
			bson.D{{Key: "$sort", Value: bson.D{{Key: "agtnme", Value: 1}}}},
			bson.D{{Key: "$skip", Value: (max(inputx.Pagenw_agtnme, 1) - 1) * inputx.Limitp_agtnme}},
			bson.D{{Key: "$limit", Value: inputx.Limitp_agtnme}},
		}
		datarw, err := tablex.Aggregate(contxt, piplne)
		if err != nil {
			panic("fail")
		}
		defer datarw.Close(contxt)

		// Append to slice
		for datarw.Next(contxt) {
			var object mdlJeddah.MdlJeddahAgtnmeDtbase
			if err := datarw.Decode(&object); err == nil {
				arrobj = append(arrobj, mdlJeddah.MdlJeddahAgtnmeInputx{
					Agtnme: object.Agtnme, Agtdtl: object.Agtdtl, Agtidn: object.Agtidn,
					Agtnew: "", Rtlsrs: object.Rtlsrs, Airlfl: object.Airlfl,
					Prmkey: object.Prmkey})
			}
		}
	}()

	// Send token to frontend
	sycwgp.Wait()
	c.JSON(http.StatusOK, gin.H{"totdta": totalx, "arrdta": arrobj})
}

// Get Agent name match search params from database
func FncJeddahAgtnmeSearch(c *gin.Context) {

	// Select database and collection
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_agentx")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	var object = mdlJeddah.MdlJeddahAgtnmeDtbase{
		Agtidn: "0X", Agtdtl: "NON JEDDAH", Rtlsrs: "-",
	}
	var newidn = c.Param("newidn")
	if newidn == "" {
		newidn = "XXXXXXXX"
	}
	var newdtl = c.Param("newdtl")
	if newdtl == "" {
		newdtl = "XXXXXXXX"
	}

	// Find user by username in database
	tablex.FindOne(contxt, bson.M{
		"$or": []bson.M{
			{"agtdtl": bson.M{"$regex": newdtl}},
			{"agtidn": bson.M{"$regex": newidn}}},
	}).Decode(&object)

	// Send token to frontend
	c.JSON(http.StatusOK, object)
}

// Get Response Update database from input
func FncJeddahAgtnmeUpdate(c *gin.Context) {

	// Bind JSON Body input to variable
	var inpDatain mdlJeddah.MdlJeddahAgtnmeInputx
	if err := c.BindJSON(&inpDatain); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
	}

	// Prepare data for
	uplSmrpnr := []mongo.WriteModel{}
	uplDtlpnr := []mongo.WriteModel{}
	// uplLogpnr := []mongo.WriteModel{}
	lmtdta := 50
	agtnme := inpDatain.Prmkey[2:len(inpDatain.Prmkey)]
	airlfl := inpDatain.Prmkey[0:2]
	update := mdlJeddah.MdlJeddahAgtnmeDtbase{
		Prmkey: inpDatain.Prmkey,
		Airlfl: airlfl,
		Agtnme: agtnme,
		Agtidn: inpDatain.Newidn,
		Agtdtl: inpDatain.Newdtl,
		Rtlsrs: inpDatain.Rtlsrs,
		Updtby: inpDatain.Updtby,
	}

	// Push Summaryt data
	rslupd := fncGlobal.FncGlobalDtbaseBlkwrt([]mongo.WriteModel{
		mongo.NewUpdateOneModel().
			SetFilter(bson.M{"prmkey": update.Prmkey}).
			SetUpdate(bson.M{"$set": update}).
			SetUpsert(true)}, "jeddah_agentx")
	if !rslupd {
		c.JSON(http.StatusInternalServerError, "failed")
	}

	// Declar sync wait group
	syncwg := sync.WaitGroup{}
	syncwg.Add(2)

	// Detail PNR
	go func() {

		// Select database and collection
		defer syncwg.Done()
		tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_pnrdtl")
		contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Get route data
		datarw, err := tablex.Find(contxt, bson.M{"agtnme": agtnme, "airlfl": airlfl})
		if err != nil {
			panic(err)
		}
		defer datarw.Close(contxt)

		// Append to slice
		for datarw.Next(contxt) {
			var object mdlJeddah.MdlJeddahPnrdtlDtbase
			datarw.Decode(&object)
			object.Agtidn = update.Agtidn
			object.Agtdtl = update.Agtdtl
			object.Rtlsrs = update.Rtlsrs
			uplDtlpnr = append(uplDtlpnr, mongo.NewUpdateOneModel().
				SetFilter(bson.M{"prmkey": object.Prmkey}).
				SetUpdate(bson.M{"$set": object}).
				SetUpsert(true))

			// Push Summaryt data
			if len(uplDtlpnr) >= lmtdta {
				rslupd := fncGlobal.FncGlobalDtbaseBlkwrt(uplDtlpnr, "jeddah_pnrdtl")
				uplDtlpnr = []mongo.WriteModel{}
				if !rslupd {
					c.JSON(http.StatusInternalServerError, "failed")
				}
			}
		}

		// Push Summaryt data
		if len(uplDtlpnr) > 0 {
			rslupd := fncGlobal.FncGlobalDtbaseBlkwrt(uplDtlpnr, "jeddah_pnrdtl")
			uplDtlpnr = []mongo.WriteModel{}
			if !rslupd {
				c.JSON(http.StatusInternalServerError, "failed")
			}
		}
	}()

	// Summary PNR
	go func() {

		// Select database and collection
		defer syncwg.Done()
		tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_pnrsmr")
		contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Get route data
		datarw, err := tablex.Find(contxt, bson.M{"agtnme": agtnme})
		if err != nil {
			panic(err)
		}
		defer datarw.Close(contxt)

		// Append to slice
		for datarw.Next(contxt) {
			var object mdlJeddah.MdlJeddahPnrsmrDtbase
			datarw.Decode(&object)
			object.Agtidn = update.Agtidn
			object.Agtdtl = update.Agtdtl
			object.Rtlsrs = update.Rtlsrs
			if strings.Contains(object.Arrcpn, update.Airlfl+"-") || object.Arrcpn == "" {
				uplSmrpnr = append(uplSmrpnr, mongo.NewUpdateOneModel().
					SetFilter(bson.M{"prmkey": object.Prmkey}).
					SetUpdate(bson.M{"$set": object}).
					SetUpsert(true))
			}

			// Push Summaryt data
			if len(uplSmrpnr) >= lmtdta {
				rslupd := fncGlobal.FncGlobalDtbaseBlkwrt(uplSmrpnr, "jeddah_pnrsmr")
				uplSmrpnr = []mongo.WriteModel{}
				if !rslupd {
					c.JSON(http.StatusInternalServerError, "failed")
				}
			}
		}

		// Push Summaryt data
		if len(uplSmrpnr) > 0 {
			rslupd := fncGlobal.FncGlobalDtbaseBlkwrt(uplSmrpnr, "jeddah_pnrsmr")
			uplSmrpnr = []mongo.WriteModel{}
			if !rslupd {
				c.JSON(http.StatusInternalServerError, "failed")
			}
		}
	}()

	// Send token to frontend
	syncwg.Wait()
	c.JSON(http.StatusOK, "success")
}

// Get Sync map data agent name
func FncJeddahAgtnmeSycmap() *sync.Map {

	// Select database and collection
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_agentx")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	datarw, err := tablex.Find(contxt, bson.M{})
	if err != nil {
		panic(err)
	}
	defer datarw.Close(contxt)

	// Append to slice
	sycmap := &sync.Map{}
	for datarw.Next(contxt) {
		var object mdlJeddah.MdlJeddahAgtnmeDtbase
		datarw.Decode(&object)
		sycmap.Store(object.Prmkey, object)
	}

	// return data
	return sycmap
}
