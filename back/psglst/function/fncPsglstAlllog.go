package fncPsglst

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"
	"context"
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get Sync map data farebase
func FncPsglstErrlogGetall(c *gin.Context) {

	// Bind JSON Body input to variable
	var inputx mdlPsglst.MdlPsglstPsgdtlInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Select db and context to do
	var totidx = 0
	var slcobj = []mdlPsglst.MdlPsglstErrlogDtbase{}
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("psglst_errlog")
	contxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "datefl", Value: 1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	mtchdt = append(mtchdt, bson.D{{Key: "erstat",
		Value: "Pending"}})
	if inputx.Erdvsn_errlog != "" {
		mtchdt = append(mtchdt, bson.D{{Key: "erdvsn",
			Value: inputx.Erdvsn_errlog}})
	}

	// Final match pipeline
	var mtchfn bson.D
	if len(mtchdt) != 0 {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: mtchdt}}}}
	} else {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{}}}
	}

	// Get Total Count Data
	wg.Add(1)
	go func() {
		defer wg.Done()
		nowPillne := mongo.Pipeline{
			mtchfn,
			bson.D{{Key: "$count", Value: "totalCount"}},
		}

		// Find user by username in database
		rawDtaset, err := tablex.Aggregate(contxt, nowPillne)
		if err != nil {
			panic(err)
		}
		defer rawDtaset.Close(contxt)

		// Store to slice from raw bson
		var slcDtaset []bson.M
		if err = rawDtaset.All(contxt, &slcDtaset); err != nil {
			panic(err)
		}

		// Mengambil jumlah dokumen dari hasil
		if len(slcDtaset) > 0 {
			if count, ok := slcDtaset[0]["totalCount"].(int32); ok {
				totidx = int(count)
			}
		}
	}()

	// Get All Match Data
	wg.Add(1)
	go func() {
		defer wg.Done()
		pipeln := mongo.Pipeline{
			mtchfn,
			sortdt,
			bson.D{{Key: "$skip", Value: (max(inputx.Pagenw_errlog, 1) - 1) * inputx.Limitp_errlog}},
			bson.D{{Key: "$limit", Value: inputx.Limitp_errlog}},
		}

		// Find user by username in database
		rawDtaset, err := tablex.Aggregate(contxt, pipeln)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer rawDtaset.Close(contxt)

		// Store to slice from raw bson
		for rawDtaset.Next(contxt) {
			var slcDtaset mdlPsglst.MdlPsglstErrlogDtbase
			rawDtaset.Decode(&slcDtaset)
			slcobj = append(slcobj, slcDtaset)
		}
	}()

	// Waiting until all go done
	wg.Wait()

	// Return final output
	c.JSON(200, gin.H{"totdta": totidx, "arrdta": slcobj})
}

// Get Sync map data farebase
func FncPsglstErrlogSycmap(intDatefl int32) *sync.Map {

	// Inisialisasi variabel
	fnldta := &sync.Map{}

	// Select database and collection
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("psglst_errlog")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	datarw, err := tablex.Find(contxt, bson.M{"erstat": "Pending", "datefl": intDatefl})
	if err != nil {
		panic(err)
	}
	defer datarw.Close(contxt)

	// Append to slice
	for datarw.Next(contxt) {
		var object mdlPsglst.MdlPsglstErrlogDtbase
		datarw.Decode(&object)
		fnldta.Store(object.Prmkey, object)
	}

	// return data
	return fnldta
}

// Get Sync map data farebase
func FncPsglstErrlogManage(errlog mdlPsglst.MdlPsglstErrlogDtbase,
	isterr bool, sycErrlog *sync.Map) {

	// Set primary key
	var datefl = strconv.Itoa(int(errlog.Datefl))
	var airlfl, flnbfl = errlog.Airlfl, errlog.Flnbfl
	var routfl, erpart = errlog.Routfl, errlog.Erpart
	var depart = ""
	switch erpart {
	case "sssion":
		errlog.Prmkey = erpart + airlfl + datefl
	case "fllstl":
		errlog.Prmkey = erpart + airlfl + depart + datefl + flnbfl
	case "fldtil", "flhour", "psglst", "psgdtl", "fllist":
		errlog.Prmkey = erpart + airlfl + flnbfl + routfl + datefl
		depart = errlog.Routfl[:3]
	case "frbase", "frtaxs":
		errlog.Prmkey = erpart + airlfl + routfl + datefl
		depart = errlog.Routfl[:3]
	}

	// If data not error and complete
	var istclr = false
	errlog.Erstat = "Pending"
	if errsyc, ist := sycErrlog.Load(errlog.Prmkey); ist && !isterr {
		if _, mtc := errsyc.(mdlPsglst.MdlPsglstErrlogDtbase); mtc {
			errlog.Erstat = "Clear"
			istclr = true
			sycErrlog.Delete(errlog.Prmkey)
		}
	}

	// Cek if data ignore error
	if errlog.Erignr == errlog.Prmkey {
		errlog.Erstat = "Ignore"
		istclr = true
		sycErrlog.Delete(errlog.Prmkey)
	}

	// Final put log action
	if isterr || istclr {
		errlog.Depart = depart
		rsupdt := fncGlobal.FncGlobalDtbaseBlkwrt(
			[]mongo.WriteModel{mongo.NewUpdateOneModel().
				SetFilter(bson.M{"prmkey": errlog.Prmkey}).
				SetUpdate(bson.M{"$set": errlog}).
				SetUpsert(true)}, "psglst_errlog")
		if rsupdt != nil {
			panic("Error Insert/Update to DB:" + rsupdt.Error())
		}
	}
}

// Get Agent name not complete format slice data from database
func FncPsglstActlogGetall(c *gin.Context) {

	// Select database and collection
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("psglst_actlog")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	datarw, err := tablex.Find(contxt, bson.M{}, options.Find().
		SetSort(bson.D{{Key: "datefl", Value: -1}}))
	if err != nil {
		panic("fail")
	}
	defer datarw.Close(contxt)

	// Append to slice
	var slices = []mdlPsglst.MdlPsglstActlogDtbase{}
	var intDatefl = []int{}
	var slcDatefl = []string{}
	for datarw.Next(contxt) {
		var object mdlPsglst.MdlPsglstActlogDtbase
		if err := datarw.Decode(&object); err == nil {
			slices = append(slices, object)
			intDatefl = append(intDatefl, int(object.Datefl))
		}
	}
	sort.Ints(intDatefl)
	for _, datefl := range intDatefl {
		fmtDatefl, _ := time.Parse("060102", strconv.Itoa(int(datefl)))
		strDatefl := fmtDatefl.Format("2006-01-02")
		slcDatefl = append(slcDatefl, strDatefl)
	}

	// Send token to frontend
	c.JSON(200, gin.H{"actlog": slices, "datefl": slcDatefl})
}
