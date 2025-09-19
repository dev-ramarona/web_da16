package fnc_jeddah

import (
	fnc_global "back/global/function"
	mdl_jeddah "back/jeddah/model"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get flight number format map data from database
func FncJeddahFlnbflSycmap(datefl string) (map[string][]mdl_jeddah.MdlJeddahFlnbdbDtbase, *sync.Map) {

	// Inisialisasi variabel
	fnlmap := map[string][]mdl_jeddah.MdlJeddahFlnbdbDtbase{}
	sycmap := &sync.Map{}

	// Select database and collection
	tablex := fnc_global.Client.Database(fnc_global.Dbases).Collection("jeddah_flnbfl")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	datarw, err := tablex.Find(contxt, bson.M{"$or": []bson.M{
		{"isjedh": "Jeddah"}, {"isjedh": ""},
		{"isjedh": bson.M{"$exists": false}},
	}})
	if err != nil {
		panic(err)
	}
	defer datarw.Close(contxt)

	// Append to slice
	for datarw.Next(contxt) {
		var object mdl_jeddah.MdlJeddahFlnbdbDtbase
		datarw.Decode(&object)
		sycmap.Store(object.Prmkey, object)
		if _, ist := fnlmap[object.Airlfl]; !ist {
			fnlmap[object.Airlfl] = []mdl_jeddah.MdlJeddahFlnbdbDtbase{}
		}
		fnlmap[object.Airlfl] = append(fnlmap[object.Airlfl], object)
	}

	// return data
	return fnlmap, sycmap
}

// Get Agent name not complete format slice data from database
func FncJeddahAgtnmeNullnm(c *gin.Context) {

	// Input data post
	var inpDatain mdl_jeddah.MdlJeddahAllpnrInputx
	if err := c.BindJSON(&inpDatain); err != nil {
		panic(err)
	}

	// Pipeline get the data logic match
	var fstPiplne = bson.A{}
	var wg sync.WaitGroup
	var intTotidx = 0
	var slcAgtnme = []mdl_jeddah.MdlJeddahAgtnmeInputx{}

	// Check if data Route all is isset
	if inpDatain.Srtnul_agtnme == "" {
		fstPiplne = append(fstPiplne, bson.D{{Key: "agtidn",
			Value: ""}})
	} else {
		if inpDatain.Airlfl_agtnme != "" {
			fstPiplne = append(fstPiplne, bson.D{{Key: "airlnf",
				Value: inpDatain.Airlfl_agtnme}})
		}
		if inpDatain.Agtnme_agtnme != "" {
			nowagt := inpDatain.Agtnme_agtnme
			fstPiplne = append(fstPiplne, bson.D{{Key: "$or", Value: bson.A{
				bson.D{{Key: "agtnme", Value: bson.D{{Key: "$regex", Value: nowagt}}}},
				bson.D{{Key: "agtdtl", Value: bson.D{{Key: "$regex", Value: nowagt}}}}}}})
		}
	}

	// Final match pipeline
	trdPiplne := bson.D{}
	if len(fstPiplne) != 0 {
		trdPiplne = bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: fstPiplne}}}}
	} else {
		trdPiplne = bson.D{{Key: "$match", Value: bson.D{}}}
	}

	// Select database and collection
	tablex := fnc_global.Client.Database(fnc_global.Dbases).Collection("jeddah_agentx")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get Total Count Data
	wg.Add(1)
	go func() {
		defer wg.Done()
		nowPillne := mongo.Pipeline{
			trdPiplne,
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
				intTotidx = int(count) // Konversi dari int32 ke int
			}
		}
	}()

	// Get route data
	wg.Add(1)
	go func() {
		defer wg.Done()
		piplne := mongo.Pipeline{trdPiplne,
			bson.D{{Key: "$skip", Value: (max(inpDatain.Pagenw_agtnme, 1) - 1) * inpDatain.Limitp_agtnme}},
			bson.D{{Key: "$limit", Value: inpDatain.Limitp_agtnme}}}
		datarw, err := tablex.Aggregate(contxt, piplne)
		if err != nil {
			panic("fail")
		}
		defer datarw.Close(contxt)

		// Append to slice
		for datarw.Next(contxt) {
			var object mdl_jeddah.MdlJeddahAgtnmeDtbase
			if err := datarw.Decode(&object); err == nil {
				slcAgtnme = append(slcAgtnme, mdl_jeddah.MdlJeddahAgtnmeInputx{
					Agtnme: object.Agtnme, Agtdtl: object.Agtdtl, Agtidn: object.Agtidn,
					Agtnew: "", Rtlsrs: object.Rtlsrs, Airlfl: object.Airlfl,
					Prmkey: object.Prmkey})
			}
		}
	}()

	// Send token to frontend
	wg.Wait()
	c.JSON(http.StatusOK, gin.H{"totdta": intTotidx, "arrdta": slcAgtnme})
}

// Get Agent name not complete format slice data from database
func FncJeddahLogactGetall(c *gin.Context) {

	// Select database and collection
	tablex := fnc_global.Client.Database(fnc_global.Dbases).Collection("jeddah_actlog")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	datarw, err := tablex.Find(contxt, bson.M{})
	if err != nil {
		panic("fail")
	}
	defer datarw.Close(contxt)

	// Append to slice
	var slices = []mdl_jeddah.MdlJeddahLogactDtbase{}
	for datarw.Next(contxt) {
		var object mdl_jeddah.MdlJeddahLogactDtbase
		if err := datarw.Decode(&object); err == nil {
			slices = append(slices, object)
		}
	}

	// Send token to frontend
	c.JSON(http.StatusOK, slices)
}

// Get Agent name not complete format slice data from database
func FncJeddahLogactLstdta() string {

	// Select database and collection
	tablex := fnc_global.Client.Database(fnc_global.Dbases).Collection("jeddah_actlog")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get Highest dateup
	var slices mdl_jeddah.MdlJeddahLogactDtbase
	option := options.FindOne().SetSort(bson.D{{Key: "dateup", Value: -1}})
	err := tablex.FindOne(contxt, bson.M{}, option).Decode(&slices)
	if err != nil {
		nowTimenb := time.Now().AddDate(0, 0, -1).Format("060102")
		return nowTimenb
	}

	// Send token to frontend
	prvDateup, _ := strconv.Atoi(time.Now().AddDate(0, 0, -1).Format("060102"))
	if slices.Dateup >= int32(prvDateup) {
		return strconv.Itoa(int(prvDateup))
	}
	return strconv.Itoa(int(slices.Dateup))
}

// Get Agent name match search params from database
func FncJeddahAgtnmeAgtsrc(c *gin.Context) {

	// Select database and collection
	tablex := fnc_global.Client.Database(fnc_global.Dbases).Collection("jeddah_agentx")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	var object mdl_jeddah.MdlJeddahAgtnmeDtbase
	var newidn = c.Param("newidn")
	if newidn == "" {
		newidn = "XXXXXXXX"
	}
	var newdtl = c.Param("newdtl")
	if newdtl == "" {
		newdtl = "XXXXXXXX"
	}
	errorx := tablex.FindOne(contxt, bson.M{
		"$or": []bson.M{
			{"agtdtl": bson.M{"$regex": newdtl}},
			{"agtidn": bson.M{"$regex": newidn}}},
	}).Decode(&object)
	if errorx != nil {
		c.JSON(http.StatusOK, errorx)
	}

	// Send token to frontend
	c.JSON(http.StatusOK, object)
}

// Get Response Update database from input
func FncJeddahAgtnmeUpdate(c *gin.Context) {

	// Bind JSON Body input to variable
	var inpDatain mdl_jeddah.MdlJeddahAgtnmeInputx
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
	update := mdl_jeddah.MdlJeddahAgtnmeDtbase{
		Prmkey: inpDatain.Prmkey,
		Airlfl: airlfl,
		Agtnme: agtnme,
		Agtidn: inpDatain.Newidn,
		Agtdtl: inpDatain.Newdtl,
		Rtlsrs: inpDatain.Rtlsrs,
		Updtby: inpDatain.Updtby,
	}

	// Push Summaryt data
	rslupd := fnc_global.FncGlobalDtbaseBlkwrt([]mongo.WriteModel{
		mongo.NewUpdateOneModel().
			SetFilter(bson.M{"prmkey": update.Prmkey}).
			SetUpdate(bson.M{"$set": update}).
			SetUpsert(true)}, "jeddah_agentx")
	if !rslupd {
		c.JSON(http.StatusInternalServerError, "failed")
	}

	// Declar sync wait group
	syncwg := sync.WaitGroup{}
	syncwg.Add(3)

	// Detail PNR
	go func() {

		// Select database and collection
		defer syncwg.Done()
		tablex := fnc_global.Client.Database(fnc_global.Dbases).Collection("jeddah_pnrdtl")
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
			var object mdl_jeddah.MdlJeddahDtlpnrDtbase
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
				rslupd := fnc_global.FncGlobalDtbaseBlkwrt(uplDtlpnr, "jeddah_pnrdtl")
				uplDtlpnr = []mongo.WriteModel{}
				if !rslupd {
					c.JSON(http.StatusInternalServerError, "failed")
				}
			}
		}

		// Push Summaryt data
		if len(uplDtlpnr) > 0 {
			rslupd := fnc_global.FncGlobalDtbaseBlkwrt(uplDtlpnr, "jeddah_pnrdtl")
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
		tablex := fnc_global.Client.Database(fnc_global.Dbases).Collection("jeddah_pnrsmr")
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
			var object mdl_jeddah.MdlJeddahSmrpnrDtbase
			datarw.Decode(&object)
			object.Agtidn = update.Agtidn
			object.Agtdtl = update.Agtdtl
			object.Rtlsrs = update.Rtlsrs
			if strings.Contains(object.Arrcpn, update.Airlfl+"-") {
				uplSmrpnr = append(uplSmrpnr, mongo.NewUpdateOneModel().
					SetFilter(bson.M{"prmkey": object.Prmkey}).
					SetUpdate(bson.M{"$set": object}).
					SetUpsert(true))
			}

			// Push Summaryt data
			if len(uplSmrpnr) >= lmtdta {
				rslupd := fnc_global.FncGlobalDtbaseBlkwrt(uplSmrpnr, "jeddah_pnrsmr")
				uplSmrpnr = []mongo.WriteModel{}
				if !rslupd {
					c.JSON(http.StatusInternalServerError, "failed")
				}
			}
		}

		// Push Summaryt data
		if len(uplSmrpnr) > 0 {
			rslupd := fnc_global.FncGlobalDtbaseBlkwrt(uplSmrpnr, "jeddah_pnrsmr")
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
	tablex := fnc_global.Client.Database(fnc_global.Dbases).Collection("jeddah_agentx")
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
		var object mdl_jeddah.MdlJeddahAgtnmeDtbase
		datarw.Decode(&object)
		sycmap.Store(object.Prmkey, object)
	}

	// return data
	return sycmap
}

// Get Sync map data LC and PUN prev day
func FncJeddahLcnpunSycmap(prvDateup string) *sync.Map {

	// Inisialisasi variabel
	fnldta := &sync.Map{}
	tmpdta := map[string]map[string]mdl_jeddah.MdlJeddahLogpnrDtbase{}

	// Select database and collection
	tablex := fnc_global.Client.Database(fnc_global.Dbases).Collection("jeddah_pnrlog")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	intDateup, _ := strconv.Atoi(prvDateup)
	datarw, err := tablex.Find(contxt, bson.M{"dateup": intDateup, "lcrpun": "lc"})
	if err != nil {
		panic(err)
	}
	defer datarw.Close(contxt)

	// Append to slice
	for datarw.Next(contxt) {
		var object mdl_jeddah.MdlJeddahLogpnrDtbase
		datarw.Decode(&object)
		syckey := object.Airlfl + object.Flnbfl + object.Depart + strconv.Itoa(int(object.Datefl))
		if _, ist := tmpdta[syckey]; !ist {
			if len(tmpdta) < 1 {
				tmpdta = map[string]map[string]mdl_jeddah.MdlJeddahLogpnrDtbase{}
			}
			tmpdta[syckey] = map[string]mdl_jeddah.MdlJeddahLogpnrDtbase{}
		}
		tmpdta[syckey][object.Pnrcde] = object
	}

	// Push to final data
	for key, val := range tmpdta {
		fnldta.Store(key, val)
	}

	// return data
	return fnldta
}

// Get Sync map data LC and PUN prev day
func FncJeddahDtlpnrSycmap(prvDateup string) *sync.Map {

	// Inisialisasi variabel
	fnldta := &sync.Map{}
	tmpdta := map[string]map[string]mdl_jeddah.MdlJeddahDtlpnrDtbase{}

	// Select database and collection
	tablex := fnc_global.Client.Database(fnc_global.Dbases).Collection("jeddah_pnrdtl")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	intDateup, _ := strconv.Atoi(prvDateup)
	datarw, err := tablex.Find(contxt, bson.M{"datefl": bson.M{"$gte": intDateup}})
	if err != nil {
		panic(err)
	}
	defer datarw.Close(contxt)

	// Append to slice
	for datarw.Next(contxt) {
		var object mdl_jeddah.MdlJeddahDtlpnrDtbase
		datarw.Decode(&object)
		syckey := object.Airlfl + object.Flnbfl + object.Depart + strconv.Itoa(int(object.Datefl))
		if _, ist := tmpdta[syckey]; !ist {
			if len(tmpdta) < 1 {
				tmpdta = map[string]map[string]mdl_jeddah.MdlJeddahDtlpnrDtbase{}
			}
			tmpdta[syckey] = map[string]mdl_jeddah.MdlJeddahDtlpnrDtbase{}
		}
		tmpdta[syckey][object.Pnrcde] = object
	}

	// Push to final data
	for key, val := range tmpdta {
		fnldta.Store(key, val)
	}

	// return data
	return fnldta
}

// Get Sync map data LC and PUN prev day
func FncJeddahDrulesSycmap() []mdl_jeddah.MdlJeddahDrulesDtbase {

	// Inisialisasi variabel
	fnldta := []mdl_jeddah.MdlJeddahDrulesDtbase{}

	// Select database and collection
	tablex := fnc_global.Client.Database(fnc_global.Dbases).Collection("jeddah_drules")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	datarw, err := tablex.Find(contxt, bson.M{},
		options.Find().SetSort(bson.D{{Key: "rllevl", Value: 1}}))
	if err != nil {
		panic(err)
	}
	defer datarw.Close(contxt)

	// Append to slice
	for datarw.Next(contxt) {
		var nowDrules mdl_jeddah.MdlJeddahDrulesDtbase
		datarw.Decode(&nowDrules)
		fnldta = append(fnldta, nowDrules)
	}

	// return data
	return fnldta
}

// Get Detail PNR from database
func FncJeddahDtlpnrFrntnd(c *gin.Context) {

	// Bind JSON Body input to variable
	istDownld := c.Param("downld")
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inpDatain mdl_jeddah.MdlJeddahAllpnrInputx
	if err := c.BindJSON(&inpDatain); err != nil {
		panic(err)
	}

	// Treatment date number
	intDatefl := 0
	if inpDatain.Datefl_pnrdtl != "" {
		strDatefl, _ := time.Parse("2006-01-02", inpDatain.Datefl_pnrdtl)
		intDatefl, _ = strconv.Atoi(strDatefl.Format("060102"))
	}

	// Select db and context to do
	var totidx = 0
	var slcobj = []mdl_jeddah.MdlJeddahDtlpnrDtbase{}
	tablex := fnc_global.Client.Database(fnc_global.Dbases).Collection("jeddah_pnrdtl")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var fstPiplne = bson.A{}
	var fthPiplne = bson.D{{Key: "$sort", Value: bson.D{{Key: "datefl", Value: 1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	if inpDatain.Datefl_pnrdtl != "" {
		csvFilenm = append(csvFilenm, strconv.Itoa(intDatefl))
		fstPiplne = append(fstPiplne, bson.D{{Key: "datefl",
			Value: intDatefl}})
	}
	if inpDatain.Flnbfl_pnrdtl != "" {
		csvFilenm = append(csvFilenm, inpDatain.Flnbfl_pnrdtl)
		fstPiplne = append(fstPiplne, bson.D{{Key: "flnbfl",
			Value: inpDatain.Flnbfl_pnrdtl}})
	}
	if inpDatain.Routfl_pnrdtl != "" {
		csvFilenm = append(csvFilenm, inpDatain.Routfl_pnrdtl)
		fstPiplne = append(fstPiplne, bson.D{{Key: "routfl",
			Value: inpDatain.Routfl_pnrdtl}})
	}
	if inpDatain.Pnrcde_pnrdtl != "" {
		pnrArrays := strings.Split(inpDatain.Pnrcde_pnrdtl, "|")
		csvFilenm = append(csvFilenm, inpDatain.Pnrcde_pnrdtl)
		fstPiplne = append(fstPiplne, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "pnrcde", Value: bson.D{{Key: "$in", Value: pnrArrays}}}},
			// bson.D{{Key: "intrln", Value: bson.D{{Key: "$regex", Value: inpDatain.Pnrcde_pnrdtl}}}},
			bson.D{{Key: "arrspl", Value: bson.D{{Key: "$regex", Value: inpDatain.Pnrcde_pnrdtl}}}},
		}}})

	}
	if inpDatain.Airlfl_pnrdtl != "" {
		csvFilenm = append(csvFilenm, inpDatain.Airlfl_pnrdtl)
		fstPiplne = append(fstPiplne, bson.D{{Key: "airlfl",
			Value: inpDatain.Airlfl_pnrdtl}})
	}
	if inpDatain.Agtnme_pnrdtl != "" {
		csvFilenm = append(csvFilenm, inpDatain.Agtnme_pnrdtl)
		fstPiplne = append(fstPiplne, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "agtnme", Value: bson.D{{Key: "$regex", Value: inpDatain.Agtnme_pnrdtl}}}},
			bson.D{{Key: "agtdtl", Value: bson.D{{Key: "$regex", Value: inpDatain.Agtnme_pnrdtl}}}}}}})
	}
	if inpDatain.Srtspl_pnrdtl != "" {
		csvFilenm = append(csvFilenm, "Split data")
		fstPiplne = append(fstPiplne, bson.D{{Key: "totspl",
			Value: bson.D{{Key: "$gt", Value: 0}}}})
		hghlow := 1
		if inpDatain.Srtspl_pnrdtl == "Highest" {
			hghlow = -1
		}
		fthPiplne = bson.D{{Key: "$sort", Value: bson.D{{Key: "totspl", Value: hghlow}}}}
	}
	if inpDatain.Srtcxl_pnrdtl != "" {
		csvFilenm = append(csvFilenm, "Cancel data")
		fstPiplne = append(fstPiplne, bson.D{{Key: "totcxl",
			Value: bson.D{{Key: "$gt", Value: 0}}}})
		hghlow := 1
		if inpDatain.Srtcxl_pnrdtl == "Highest" {
			hghlow = -1
		}
		fthPiplne = bson.D{{Key: "$sort", Value: bson.D{{Key: "totcxl", Value: hghlow}}}}
	}

	// Final match pipeline
	trdPiplne := bson.D{}
	if len(fstPiplne) != 0 {
		trdPiplne = bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: fstPiplne}}}}
	} else {
		trdPiplne = bson.D{{Key: "$match", Value: bson.D{}}}
	}

	// Logic download data
	if istDownld == "downld" {

		// Set header untuk file CSV
		fnlFilenm := strings.Join(csvFilenm, "_")
		c.Header("Content-Type", "text/csv")
		c.Header("Content-Disposition", "attachment; filename=Jeddah_Detail_PNR_"+fnlFilenm+".csv")
		c.Header("Access-Control-Expose-Headers", "Content-Disposition")

		// Streaming file CSV ke client
		writer := csv.NewWriter(c.Writer)
		defer writer.Flush()
		writer.Write([]string{"Prmkey", "Airlfl", "Flnbfl", "Depart", "Routfl", "Clssfl",
			"Datefl", "Dateup", "Timeup", "Timecr", "Agtnme", "Agtdtl", "Agtidn", "Pnrcde",
			"Intrln", "Rtlsrs", "Toflnm", "Drules", "Totisd", "Totbok", "Totpax", "Totcxl",
			"Totchg", "Totspl", "Arrspl", "Notedt", "Flstat",
		})

		// Get All Match Data
		nowPiplne := mongo.Pipeline{
			trdPiplne,
			fthPiplne,
		}

		// Find user by username in database
		rawDtaset, err := tablex.Aggregate(contxt, nowPiplne)
		if err != nil {
			panic(err)
		}
		defer rawDtaset.Close(contxt)

		// Store to slice from raw bson
		for rawDtaset.Next(contxt) {
			var slcDtaset mdl_jeddah.MdlJeddahDtlpnrDtbase
			rawDtaset.Decode(&slcDtaset)
			writer.Write([]string{
				slcDtaset.Prmkey,
				slcDtaset.Airlfl,
				slcDtaset.Flnbfl,
				slcDtaset.Depart,
				slcDtaset.Routfl,
				slcDtaset.Clssfl,
				strconv.Itoa(int(slcDtaset.Datefl)),
				strconv.Itoa(int(slcDtaset.Dateup)),
				strconv.Itoa(int(slcDtaset.Timeup)),
				strconv.Itoa(int(slcDtaset.Timecr)),
				slcDtaset.Agtnme,
				slcDtaset.Agtdtl,
				slcDtaset.Agtidn,
				slcDtaset.Pnrcde,
				slcDtaset.Intrln,
				slcDtaset.Rtlsrs,
				slcDtaset.Toflnm,
				strconv.Itoa(slcDtaset.Drules),
				strconv.Itoa(slcDtaset.Totisd),
				strconv.Itoa(slcDtaset.Totbok),
				strconv.Itoa(slcDtaset.Totpax),
				strconv.Itoa(slcDtaset.Totcxl),
				strconv.Itoa(slcDtaset.Totchg),
				strconv.Itoa(slcDtaset.Totspl),
				slcDtaset.Arrspl,
				slcDtaset.Notedt,
				slcDtaset.Flstat,
			})
		}
	} else {

		// Get Total Count Data
		wg.Add(1)
		go func() {
			defer wg.Done()
			nowPillne := mongo.Pipeline{
				trdPiplne,
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
					totidx = int(count) // Konversi dari int32 ke int
				}
			}
		}()

		// Get All Match Data
		wg.Add(1)
		go func() {
			defer wg.Done()
			nowPiplne := mongo.Pipeline{
				trdPiplne,
				fthPiplne,
				bson.D{{Key: "$skip", Value: (max(inpDatain.Pagenw_pnrdtl, 1) - 1) * inpDatain.Limitp_pnrdtl}},
				bson.D{{Key: "$limit", Value: inpDatain.Limitp_pnrdtl}},
			}

			// Find user by username in database
			rawDtaset, err := tablex.Aggregate(contxt, nowPiplne)
			if err != nil {
				panic(err)
			}
			defer rawDtaset.Close(contxt)

			// Store to slice from raw bson
			for rawDtaset.Next(contxt) {
				var slcDtaset mdl_jeddah.MdlJeddahDtlpnrDtbase
				rawDtaset.Decode(&slcDtaset)
				slcobj = append(slcobj, slcDtaset)
			}
		}()

		// Waiting until all go done
		wg.Wait()

		// Return final output
		c.JSON(200, gin.H{"totdta": totidx, "arrdta": slcobj})
	}
}

// Get Summary PNR from database
func FncJeddahSmrpnrFrntnd(c *gin.Context) {

	// Bind JSON Body input to variable
	istDownld := c.Param("downld")
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inpDatain mdl_jeddah.MdlJeddahAllpnrInputx
	if err := c.BindJSON(&inpDatain); err != nil {
		panic(err)
	}

	// Treatment date number
	intDatefl := 0
	if inpDatain.Datefl_pnrsmr != "" {
		strDatefl, _ := time.Parse("2006-01-02", inpDatain.Datefl_pnrsmr)
		intDatefl, _ = strconv.Atoi(strDatefl.Format("060102"))
	}

	// Select db and context to do
	var totidx = 0
	var slcobj = []mdl_jeddah.MdlJeddahSmrpnrDtbase{}
	tablex := fnc_global.Client.Database(fnc_global.Dbases).Collection("jeddah_pnrsmr")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var fstPiplne = bson.A{}
	var fthPiplne = bson.D{{Key: "$sort", Value: bson.D{{Key: "datefl", Value: 1}}}}
	var wg sync.WaitGroup

	// Special in this Summary PNR match
	if inpDatain.Datefl_pnrsmr != "" || inpDatain.Flnbfl_pnrsmr != "" ||
		inpDatain.Routfl_pnrsmr != "" || inpDatain.Airlfl_pnrsmr != "" {
		nowArrcpn := []string{}
		for _, params := range []string{inpDatain.Airlfl_pnrsmr,
			inpDatain.Flnbfl_pnrsmr, inpDatain.Routfl_pnrsmr,
			strconv.Itoa(intDatefl)} {
			if params != "" {
				nowArrcpn = append(nowArrcpn, params)
			} else {
				nowArrcpn = append(nowArrcpn, ".*")
			}
		}
		strArrcpn := strings.Join(nowArrcpn, "-")
		csvFilenm = append(csvFilenm, strconv.Itoa(intDatefl))
		fstPiplne = append(fstPiplne, bson.D{{Key: "arrcpn",
			Value: bson.D{{Key: "$regex", Value: strArrcpn}}}})
	}

	// Check if data Route all is isset
	if inpDatain.Psdate_pnrsmr == "Hide Past Date" {
		nowTimenb := time.Now().Format("0601020000")
		intTimenb, _ := strconv.Atoi(nowTimenb)
		csvFilenm = append(csvFilenm, inpDatain.Psdate_pnrsmr, nowTimenb)
		fstPiplne = append(fstPiplne, bson.D{{Key: "timedp",
			Value: bson.D{{Key: "$gte", Value: intTimenb}}}})
	}
	if inpDatain.Pnrcde_pnrsmr != "" {
		pnrArrays := strings.Split(inpDatain.Pnrcde_pnrsmr, "|")
		csvFilenm = append(csvFilenm, inpDatain.Pnrcde_pnrsmr)
		fstPiplne = append(fstPiplne, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "pnrcde", Value: bson.D{{Key: "$in", Value: pnrArrays}}}},
			// bson.D{{Key: "intrln", Value: bson.D{{Key: "$regex", Value: inpDatain.Pnrcde_pnrsmr}}}},
			bson.D{{Key: "arrspl", Value: bson.D{{Key: "$regex", Value: inpDatain.Pnrcde_pnrsmr}}}},
		}}})

	}
	if inpDatain.Agtnme_pnrsmr != "" {
		csvFilenm = append(csvFilenm, inpDatain.Agtnme_pnrsmr)
		fstPiplne = append(fstPiplne, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "agtnme", Value: bson.D{{Key: "$regex", Value: inpDatain.Agtnme_pnrsmr}}}},
			bson.D{{Key: "agtdtl", Value: bson.D{{Key: "$regex", Value: inpDatain.Agtnme_pnrsmr}}}}}}})
	}
	if inpDatain.Srtspl_pnrsmr != "" {
		csvFilenm = append(csvFilenm, "Split data")
		fstPiplne = append(fstPiplne, bson.D{{Key: "totspl",
			Value: bson.D{{Key: "$gt", Value: 0}}}})
		hghlow := 1
		if inpDatain.Srtspl_pnrsmr == "Highest" {
			hghlow = -1
		}
		fthPiplne = bson.D{{Key: "$sort", Value: bson.D{{Key: "totspl", Value: hghlow}}}}
	}
	if inpDatain.Srtcxl_pnrsmr != "" {
		csvFilenm = append(csvFilenm, "Cancel data")
		fstPiplne = append(fstPiplne, bson.D{{Key: "totcxl",
			Value: bson.D{{Key: "$gt", Value: 0}}}})
		hghlow := 1
		if inpDatain.Srtcxl_pnrsmr == "Highest" {
			hghlow = -1
		}
		fthPiplne = bson.D{{Key: "$sort", Value: bson.D{{Key: "totcxl", Value: hghlow}}}}
	}
	if istDownld == "rtlsrs" {
		fstPiplne = append(fstPiplne, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "notedt", Value: ""}},
			bson.D{{Key: "notedt", Value: bson.D{{Key: "$exists", Value: false}}}}}}})
	}

	// Final match pipeline
	trdPiplne := bson.D{}
	if len(fstPiplne) != 0 {
		trdPiplne = bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: fstPiplne}}}}
	} else {
		trdPiplne = bson.D{{Key: "$match", Value: bson.D{}}}
	}

	// Logic download data
	if istDownld == "downld" || istDownld == "rtlsrs" {

		// Set header untuk file CSV
		fnlFilenm := strings.Join(csvFilenm, "_")
		c.Header("Content-Type", "text/csv")
		c.Header("Content-Disposition", "attachment; filename=Jeddah_Summary_PNR_"+fnlFilenm+".csv")
		c.Header("Access-Control-Expose-Headers", "Content-Disposition")

		// Streaming file CSV ke client
		writer := csv.NewWriter(c.Writer)
		defer writer.Flush()
		if istDownld == "rtlsrs" {
			writer.Write([]string{"Pnrcde:(DO NOT CHANGE)", "Timecr:(DO NOT CHANGE)",
				"[Retail/Series/-]:(CHANGE ONLY WITH THIS FORMAT)", "Intrln"})
		} else {
			writer.Write([]string{"Prmkey", "Routfl", "Timedp", "Timerv", "Dateup", "Timeup", "Timecr",
				"Agtnme", "Agtdtl", "Agtidn", "Pnrcde", "Intrln", "Rtlsrs", "Arrcpn", "Agtdie", "Totisd",
				"Totbok", "Totpax", "Totcxl", "Totspl", "Arrspl", "Notedt"})
		}

		// Get All Match Data
		nowPiplne := mongo.Pipeline{
			trdPiplne,
			fthPiplne,
		}

		// Find user by username in database
		rawDtaset, err := tablex.Aggregate(contxt, nowPiplne)
		if err != nil {
			panic(err)
		}
		defer rawDtaset.Close(contxt)

		// Store to slice from raw bson
		for rawDtaset.Next(contxt) {
			var slcDtaset mdl_jeddah.MdlJeddahSmrpnrDtbase
			rawDtaset.Decode(&slcDtaset)
			if istDownld == "rtlsrs" {
				writer.Write([]string{
					slcDtaset.Pnrcde,
					strconv.Itoa(int(slcDtaset.Timecr)),
					slcDtaset.Rtlsrs,
					slcDtaset.Intrln,
				})
			} else {
				writer.Write([]string{
					slcDtaset.Prmkey,
					slcDtaset.Routfl,
					strconv.Itoa(int(slcDtaset.Timedp)),
					strconv.Itoa(int(slcDtaset.Timerv)),
					strconv.Itoa(int(slcDtaset.Dateup)),
					strconv.Itoa(int(slcDtaset.Timeup)),
					strconv.Itoa(int(slcDtaset.Timecr)),
					slcDtaset.Agtnme,
					slcDtaset.Agtdtl,
					slcDtaset.Agtidn,
					slcDtaset.Pnrcde,
					slcDtaset.Intrln,
					slcDtaset.Rtlsrs,
					slcDtaset.Arrcpn,
					slcDtaset.Agtdie,
					strconv.Itoa(slcDtaset.Totisd),
					strconv.Itoa(slcDtaset.Totbok),
					strconv.Itoa(slcDtaset.Totpax),
					strconv.Itoa(slcDtaset.Totcxl),
					strconv.Itoa(slcDtaset.Totspl),
					slcDtaset.Arrspl,
					slcDtaset.Notedt,
				})
			}
		}
	} else {

		// Get Total Count Data
		wg.Add(1)
		go func() {
			defer wg.Done()
			nowPillne := mongo.Pipeline{
				trdPiplne,
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
					totidx = int(count) // Konversi dari int32 ke int
				}
			}
		}()

		// Get All Match Data
		wg.Add(1)
		go func() {
			defer wg.Done()
			nowPiplne := mongo.Pipeline{
				trdPiplne,
				fthPiplne,
				bson.D{{Key: "$skip", Value: (max(inpDatain.Pagenw_pnrsmr, 1) - 1) * inpDatain.Limitp_pnrsmr}},
				bson.D{{Key: "$limit", Value: inpDatain.Limitp_pnrsmr}},
			}

			// Find user by username in database
			rawDtaset, err := tablex.Aggregate(contxt, nowPiplne)
			if err != nil {
				panic(err)
			}
			defer rawDtaset.Close(contxt)

			// Store to slice from raw bson
			for rawDtaset.Next(contxt) {
				var slcDtaset mdl_jeddah.MdlJeddahSmrpnrDtbase
				rawDtaset.Decode(&slcDtaset)
				slcobj = append(slcobj, slcDtaset)
			}
		}()

		// Waiting until all go done
		wg.Wait()

		// Return final output
		c.JSON(200, gin.H{"totdta": totidx, "arrdta": slcobj})
	}
}

// Get Summary PNR from database
func FncJeddahSmrflnFrntnd(c *gin.Context) {

	// Bind JSON Body input to variable
	istDownld := c.Param("downld")
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inpDatain mdl_jeddah.MdlJeddahAllpnrInputx
	if err := c.BindJSON(&inpDatain); err != nil {
		panic(err)
	}

	// Treatment date number
	intDatefl := 0
	if inpDatain.Datefl_flnsmr != "" {
		strDatefl, _ := time.Parse("2006-01-02", inpDatain.Datefl_flnsmr)
		intDatefl, _ = strconv.Atoi(strDatefl.Format("060102"))
	}

	// Select db and context to do
	var totidx = 0
	var slcobj = []mdl_jeddah.MdlJeddahSmrflnDtbase{}
	tablex := fnc_global.Client.Database(fnc_global.Dbases).Collection("jeddah_flnsmr")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var fstPiplne = bson.A{}
	var fthPiplne = bson.D{{Key: "$sort", Value: bson.D{{Key: "datefl", Value: 1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	if inpDatain.Datefl_flnsmr != "" {
		csvFilenm = append(csvFilenm, strconv.Itoa(intDatefl))
		fstPiplne = append(fstPiplne, bson.D{{Key: "datefl",
			Value: intDatefl}})
	}
	if inpDatain.Airlfl_flnsmr != "" {
		csvFilenm = append(csvFilenm, inpDatain.Airlfl_flnsmr)
		fstPiplne = append(fstPiplne, bson.D{{Key: "airlfl",
			Value: inpDatain.Airlfl_flnsmr}})
	}
	if inpDatain.Flnbfl_flnsmr != "" {
		csvFilenm = append(csvFilenm, inpDatain.Flnbfl_flnsmr)
		fstPiplne = append(fstPiplne, bson.D{{Key: "flnbfl",
			Value: inpDatain.Flnbfl_flnsmr}})
	}
	if inpDatain.Routfl_flnsmr != "" {
		csvFilenm = append(csvFilenm, inpDatain.Routfl_flnsmr)
		fstPiplne = append(fstPiplne, bson.D{{Key: "routfl",
			Value: inpDatain.Routfl_flnsmr}})
	}
	if inpDatain.Srtspl_flnsmr != "" {
		csvFilenm = append(csvFilenm, "Split data")
		fstPiplne = append(fstPiplne, bson.D{{Key: "totspl",
			Value: bson.D{{Key: "$gt", Value: 0}}}})
		hghlow := 1
		if inpDatain.Srtspl_flnsmr == "Highest" {
			hghlow = -1
		}
		fthPiplne = bson.D{{Key: "$sort", Value: bson.D{{Key: "totspl", Value: hghlow}}}}
	}
	if inpDatain.Srtcxl_flnsmr != "" {
		csvFilenm = append(csvFilenm, "Cancel data")
		fstPiplne = append(fstPiplne, bson.D{{Key: "totcxl",
			Value: bson.D{{Key: "$gt", Value: 0}}}})
		hghlow := 1
		if inpDatain.Srtcxl_flnsmr == "Highest" {
			hghlow = -1
		}
		fthPiplne = bson.D{{Key: "$sort", Value: bson.D{{Key: "totcxl", Value: hghlow}}}}
	}
	if inpDatain.Psdate_flnsmr == "Hide Past Date" {
		nowDatefl := time.Now().Format("060102")
		intDatefl, _ := strconv.Atoi(nowDatefl)
		csvFilenm = append(csvFilenm, inpDatain.Psdate_flnsmr, nowDatefl)
		fstPiplne = append(fstPiplne, bson.D{{Key: "datefl",
			Value: bson.D{{Key: "$gte", Value: intDatefl}}}})
	}

	// Final match pipeline
	trdPiplne := bson.D{}
	if len(fstPiplne) != 0 {
		trdPiplne = bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: fstPiplne}}}}
	} else {
		trdPiplne = bson.D{{Key: "$match", Value: bson.D{}}}
	}

	// Logic download data
	if istDownld == "downld" {

		// Set header untuk file CSV
		fnlFilenm := strings.Join(csvFilenm, "_")
		c.Header("Content-Type", "text/csv")
		c.Header("Content-Disposition", "attachment; filename=Jeddah_Summary_FlightNumber_"+fnlFilenm+".csv")
		c.Header("Access-Control-Expose-Headers", "Content-Disposition")

		// Streaming file CSV ke client
		writer := csv.NewWriter(c.Writer)
		defer writer.Flush()
		writer.Write([]string{"Prmkey", "Airlfl", "Flnbfl", "Depart", "Routfl",
			"Datefl", "Dateup", "Timeup", "Totisd", "Totbok", "Totpax", "Totcxl",
			"Totchg", "Totspl", "Notedt"})

		// Get All Match Data
		nowPiplne := mongo.Pipeline{
			trdPiplne,
			fthPiplne,
		}

		// Find user by username in database
		rawDtaset, err := tablex.Aggregate(contxt, nowPiplne)
		if err != nil {
			panic(err)
		}
		defer rawDtaset.Close(contxt)

		// Store to slice from raw bson
		for rawDtaset.Next(contxt) {
			var slcDtaset mdl_jeddah.MdlJeddahSmrflnDtbase
			rawDtaset.Decode(&slcDtaset)
			writer.Write([]string{
				slcDtaset.Prmkey,
				slcDtaset.Airlfl,
				slcDtaset.Flnbfl,
				slcDtaset.Depart,
				slcDtaset.Routfl,
				strconv.Itoa(int(slcDtaset.Datefl)),
				strconv.Itoa(int(slcDtaset.Dateup)),
				strconv.Itoa(int(slcDtaset.Timeup)),
				strconv.Itoa(int(slcDtaset.Totisd)),
				strconv.Itoa(int(slcDtaset.Totbok)),
				strconv.Itoa(int(slcDtaset.Totpax)),
				strconv.Itoa(int(slcDtaset.Totcxl)),
				strconv.Itoa(int(slcDtaset.Totchg)),
				strconv.Itoa(int(slcDtaset.Totspl)),
				slcDtaset.Notedt,
			})
		}
	} else {

		// Get Total Count Data
		wg.Add(1)
		go func() {
			defer wg.Done()
			nowPillne := mongo.Pipeline{
				trdPiplne,
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
					totidx = int(count) // Konversi dari int32 ke int
				}
			}
		}()

		// Get All Match Data
		wg.Add(1)
		go func() {
			defer wg.Done()
			nowPiplne := mongo.Pipeline{
				trdPiplne,
				fthPiplne,
				bson.D{{Key: "$skip", Value: (max(inpDatain.Pagenw_flnsmr, 1) - 1) * inpDatain.Limitp_flnsmr}},
				bson.D{{Key: "$limit", Value: inpDatain.Limitp_flnsmr}},
			}

			// Find user by username in database
			rawDtaset, err := tablex.Aggregate(contxt, nowPiplne)
			if err != nil {
				panic(err)
			}
			defer rawDtaset.Close(contxt)

			// Store to slice from raw bson
			for rawDtaset.Next(contxt) {
				var slcDtaset mdl_jeddah.MdlJeddahSmrflnDtbase
				rawDtaset.Decode(&slcDtaset)
				slcobj = append(slcobj, slcDtaset)
			}
		}()

		// Waiting until all go done
		wg.Wait()

		// Return final output
		c.JSON(200, gin.H{"totdta": totidx, "arrdta": slcobj})
	}
}

// Push new flight number to database
func FncJeddahAddflnTodtbs(c *gin.Context) {

	// Bind JSON Body input to variable
	var inpDatain mdl_jeddah.MdlJeddahAddflnInpprm
	if err := c.BindJSON(&inpDatain); err != nil {
		panic(err)
	}

	// Treatment date number
	intDatefl, nowDatefl, nowTimenb := 0, 0, 0
	if inpDatain.Datefl_addfln != "" {
		rawDatefl, _ := time.Parse("2006-01-02", inpDatain.Datefl_addfln)
		intDatefl, _ = strconv.Atoi(rawDatefl.Format("060102"))
		intDatefl, _ = strconv.Atoi(rawDatefl.Format("060102"))
		nowDatefl, _ = strconv.Atoi(time.Now().Format("06010"))
		nowTimenb, _ = strconv.Atoi(time.Now().Format("0601021504"))
	}

	// Push to Pnr log data
	rspBlkwrt := fnc_global.FncGlobalDtbaseBlkwrt(
		[]mongo.WriteModel{mongo.NewUpdateOneModel().
			SetFilter(bson.M{"prmkey": ""}).
			SetUpdate(bson.M{"$set": mdl_jeddah.MdlJeddahFlnbdbDtbase{
				Prmkey: "",
				Datefl: int32(intDatefl),
				Dateup: int32(nowDatefl),
				Timeup: int64(nowTimenb),
				Airlfl: inpDatain.Airlfl_addfln,
				Flnbfl: inpDatain.Flnbfl_addfln,
				Depart: inpDatain.Routfl_addfln[0:3],
				Routfl: inpDatain.Routfl_addfln,
				Fltype: "",
				Flstat: "",
			}}).
			SetUpsert(true)}, "jeddah_pnrdtl")
	if !rspBlkwrt {
		fmt.Println("ERR LOG HERE, CAN'T INPUT DTLPNR")
	}
}

// Get Response Update database from input
func FncJeddahRtlsrsUpdate(c *gin.Context) {

	// Bind JSON Body input to variable
	var inpDatain mdl_jeddah.MdlJeddahSmrpnrDtbase
	if err := c.BindJSON(&inpDatain); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
	}

	// Push Summaryt data
	rslupd := fnc_global.FncGlobalDtbaseBlkwrt([]mongo.WriteModel{
		mongo.NewUpdateOneModel().
			SetFilter(bson.M{"prmkey": inpDatain.Prmkey}).
			SetUpdate(bson.M{"$set": inpDatain}).
			SetUpsert(true)}, "jeddah_pnrsmr")
	if !rslupd {
		c.JSON(http.StatusInternalServerError, "failed")
	}

	// Update detail PNR
	func() {

		// Select database and collection
		tablex := fnc_global.Client.Database(fnc_global.Dbases).Collection("jeddah_pnrdtl")
		contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Get route data
		datarw, err := tablex.Find(contxt, bson.M{"$and": []bson.M{
			{"pnrcde": inpDatain.Pnrcde}, {"timecr": inpDatain.Timecr}}})
		if err != nil {
			panic(err)
		}
		defer datarw.Close(contxt)

		// Append to slice
		pnrdtlMgomdl := []mongo.WriteModel{}
		for datarw.Next(contxt) {
			var object mdl_jeddah.MdlJeddahDtlpnrDtbase
			datarw.Decode(&object)
			object.Rtlsrs = inpDatain.Rtlsrs
			object.Notedt = inpDatain.Notedt
			pnrdtlMgomdl = append(pnrdtlMgomdl, mongo.NewUpdateOneModel().
				SetFilter(bson.M{"prmkey": object.Prmkey}).
				SetUpdate(bson.M{"$set": object}).
				SetUpsert(true))
		}

		// Update final data
		rslupd := fnc_global.FncGlobalDtbaseBlkwrt(pnrdtlMgomdl, "jeddah_pnrdtl")
		if !rslupd {
			c.JSON(http.StatusInternalServerError, "failed")
		}
	}()

	// Send token to frontend
	c.JSON(http.StatusOK, "success")
}

// Status upload rtlsrs
var rtlsrsStatus = "Done"

// Cek status upload rtlsrs
func FncJeddahRtlsrsStatus(c *gin.Context) {
	c.JSON(http.StatusOK, rtlsrsStatus)
}

// Action upload rtlsrs
func FncJeddahRtlsrsUpload(c *gin.Context) {
	rtlsrsStatus = "Wait"

	// Limit size upload (10 MB)
	const maxupl = 10 << 20
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxupl)

	// Parse multipart form
	mltplr, err := c.Request.MultipartReader()
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, fmt.Sprintf("multipart error: %v", err))
		return
	}

	// Looping each part/file
	errrsp := [][]string{}
	for {
		nxpart, err := mltplr.NextPart()
		var upldby = c.Param("upldby")
		if upldby == "" {
			upldby = "User not found"
		}

		// Check isset next part
		if err == io.EOF {
			break
		}

		// Handle error next part
		if err != nil {
			log.Printf("NextPart error: %v", err)
			c.String(http.StatusBadRequest, "next part error")
			return
		}

		// Validasi part/file
		if nxpart == nil || nxpart.Header == nil {
			io.Copy(io.Discard, nxpart)
			continue
		}

		//  Check Content-Disposition for get filename
		contds := nxpart.Header.Get("Content-Disposition")
		if contds == "" {
			io.Copy(io.Discard, nxpart)
			continue
		}

		// Check type standard Mime
		_, params, err := mime.ParseMediaType(contds)
		if err != nil {
			io.Copy(io.Discard, nxpart)
			continue
		}

		// Check and Get filename then clean path
		flname := params["filename"]
		if flname == "" {
			io.Copy(io.Discard, nxpart)
			continue
		}
		flname = filepath.Base(flname)

		// Siapkan CSV reader
		reader := csv.NewReader(nxpart)
		reader.FieldsPerRecord = -1
		reader.LazyQuotes = true
		reader.TrimLeadingSpace = true

		// Read and process CSV file Header
		header, err := reader.Read()
		if err != nil {
			log.Printf("error read header for %s: %v", flname, err)
			continue
		}
		if len(header) > 0 {
			header[0] = strings.TrimPrefix(header[0], "\uFEFF") // hapus BOM
		}

		// Looping each row
		mgodtl, mgosmr := []mongo.WriteModel{}, []mongo.WriteModel{}
		nowrow, limitz := 1, 100
		for {

			// Check isset row
			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("csv read error %s: %v", flname, err)
				break
			}

			// Process row trim all fields
			nowerr := []string{row[0], row[1], row[2], flname}
			for i := range row {
				row[i] = strings.TrimSpace(row[i])
			}

			// Treatment PNR Code
			rawPnrcde := row[0]
			if len(rawPnrcde) != 6 {
				nowerr = append(nowerr, fmt.Sprintf("Row %d: Invalid PNR Code '%s'", nowrow+1, rawPnrcde))
			}

			// Treatment time create
			rawTimecr := row[1]
			intTimecr, err := strconv.Atoi(rawTimecr)
			if err != nil || len(rawTimecr) != 10 {
				nowerr = append(nowerr, fmt.Sprintf("Row %d: Invalid Timecr '%s'", nowrow+1, rawTimecr))
			}

			// Treatment Retail/Series
			rilRtlsrs, rawRtlsrs := "", row[2]
			switch strings.ToLower(strings.TrimSpace(rawRtlsrs)) {
			case "retail":
				rilRtlsrs = "Retail"
			case "series":
				rilRtlsrs = "Series"
			case "-":
				rilRtlsrs = "-"
			default:
				nowerr = append(nowerr, fmt.Sprintf("Row %d: Invalid Retail/Series '%s'", nowrow+1, rawRtlsrs))
			}

			// If any error in row, then skip and save to error response
			if len(nowerr) > 4 {
				errrsp = append(errrsp, nowerr)
				continue
			}

			// Update to database PNR Detail
			mgodtl = append(mgodtl, mongo.NewUpdateManyModel().
				SetFilter(bson.M{"pnrcde": row[0], "timecr": int64(intTimecr)}).
				SetUpdate(bson.M{"$set": bson.M{
					"rtlsrs": rilRtlsrs, "notedt": fmt.Sprintf("Updated %s by %s", rawRtlsrs, upldby),
				}}).
				SetUpsert(true))

			// Push mongo pnrdtl
			if len(mgodtl) > limitz && len(mgodtl) != 0 {
				rspons := fnc_global.FncGlobalDtbaseBlkwrt(mgodtl, "jeddah_pnrdtl")
				if !rspons {
					fmt.Println("ERR LOG HERE, CAN'T INPUT LCNPUN")
				}
				fmt.Println("Updated batch 100 jeddah_pnrdtl")
				mgodtl = []mongo.WriteModel{}
			}

			// Update to database PNR Summary
			mgosmr = append(mgosmr, mongo.NewUpdateOneModel().
				SetFilter(bson.M{"prmkey": rawPnrcde + rawTimecr}).
				SetUpdate(bson.M{"$set": bson.M{
					"rtlsrs": rilRtlsrs, "notedt": fmt.Sprintf("Updated %s by %s", rawRtlsrs, upldby),
				}}).
				SetUpsert(true))

			// Push mongo pnrsmr
			if len(mgosmr) > limitz && len(mgosmr) != 0 {
				rspons := fnc_global.FncGlobalDtbaseBlkwrt(mgosmr, "jeddah_pnrsmr")
				if !rspons {
					fmt.Println("ERR LOG HERE, CAN'T INPUT LCNPUN")
				}
				fmt.Println("Updated batch 100 jeddah_pnrsmr")
				mgosmr = []mongo.WriteModel{}
			}

			fmt.Println(rawPnrcde, "|")
			nowrow++
		}

		// Push mongo pnrdtl
		if len(mgodtl) > 0 {
			rspons := fnc_global.FncGlobalDtbaseBlkwrt(mgodtl, "jeddah_pnrdtl")
			if !rspons {
				fmt.Println("ERR LOG HERE, CAN'T INPUT LCNPUN")
			}
		}

		// Push mongo pnrsmr
		if len(mgosmr) > 0 {
			rspons := fnc_global.FncGlobalDtbaseBlkwrt(mgosmr, "jeddah_pnrsmr")
			if !rspons {
				fmt.Println("ERR LOG HERE, CAN'T INPUT LCNPUN")
			}
		}
	}

	// If any error on response
	if len(errrsp) > 0 {

		// Set header untuk file CSV
		c.Header("Content-Type", "text/csv")
		c.Header("Content-Disposition", "attachment; filename=Jeddah_Summary_PNR_Error_Response.csv")
		c.Header("Access-Control-Expose-Headers", "Content-Disposition")

		// Streaming file CSV ke client
		writer := csv.NewWriter(c.Writer)
		defer writer.Flush()
		writer.Write([]string{"Pnrcde:(DO NOT CHANGE)", "Timecr:(DO NOT CHANGE)",
			"[Retail/Series/-]:(CHANGE ONLY WITH THIS FORMAT)",
			"File name", "Err1", "Err2", "Err3", "Err4", "Err5"})
		for _, val := range errrsp {
			writer.Write(val)
		}
	}

	// Respond done
	c.String(http.StatusOK, "done")
	rtlsrsStatus = "Done"
}
