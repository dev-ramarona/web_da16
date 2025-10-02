package fncJeddah

import (
	fncGlobal "back/global/function"
	mdlJeddah "back/jeddah/model"
	"context"
	"encoding/csv"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get Detail PNR from database
func FncJeddahPnrdtlGetall(c *gin.Context) {

	// Bind JSON Body input to variable
	istDownld := c.Param("downld")
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inputx mdlJeddah.MdlJeddahAllprmInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Treatment date number
	intDatefl := 0
	if inputx.Datefl_pnrdtl != "" {
		strDatefl, _ := time.Parse("2006-01-02", inputx.Datefl_pnrdtl)
		intDatefl, _ = strconv.Atoi(strDatefl.Format("060102"))
	}

	// Select db and context to do
	var totidx = 0
	var slcobj = []mdlJeddah.MdlJeddahPnrdtlDtbase{}
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_pnrdtl")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "datefl", Value: 1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	if inputx.Datefl_pnrdtl != "" {
		csvFilenm = append(csvFilenm, strconv.Itoa(intDatefl))
		mtchdt = append(mtchdt, bson.D{{Key: "datefl",
			Value: intDatefl}})
	}
	if inputx.Flnbfl_pnrdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Flnbfl_pnrdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "flnbfl",
			Value: inputx.Flnbfl_pnrdtl}})
	}
	if inputx.Routfl_pnrdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Routfl_pnrdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "routfl",
			Value: inputx.Routfl_pnrdtl}})
	}
	if inputx.Pnrcde_pnrdtl != "" {
		pnrArrays := strings.Split(inputx.Pnrcde_pnrdtl, "|")
		csvFilenm = append(csvFilenm, inputx.Pnrcde_pnrdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "pnrcde", Value: bson.D{{Key: "$in", Value: pnrArrays}}}},
			bson.D{{Key: "arrspl", Value: bson.D{{Key: "$regex", Value: inputx.Pnrcde_pnrdtl}}}},
		}}})

	}
	if inputx.Airlfl_pnrdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Airlfl_pnrdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "airlfl",
			Value: inputx.Airlfl_pnrdtl}})
	}
	if inputx.Agtnme_pnrdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Agtnme_pnrdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "agtnme", Value: bson.D{{Key: "$regex", Value: inputx.Agtnme_pnrdtl}}}},
			bson.D{{Key: "agtdtl", Value: bson.D{{Key: "$regex", Value: inputx.Agtnme_pnrdtl}}}}}}})
	}
	if inputx.Srtspl_pnrdtl != "" {
		csvFilenm = append(csvFilenm, "Split data")
		mtchdt = append(mtchdt, bson.D{{Key: "totspl",
			Value: bson.D{{Key: "$gt", Value: 0}}}})
		hghlow := 1
		if inputx.Srtspl_pnrdtl == "Highest" {
			hghlow = -1
		}
		sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "totspl", Value: hghlow}}}}
	}
	if inputx.Srtcxl_pnrdtl != "" {
		csvFilenm = append(csvFilenm, "Cancel data")
		mtchdt = append(mtchdt, bson.D{{Key: "totcxl",
			Value: bson.D{{Key: "$gt", Value: 0}}}})
		hghlow := 1
		if inputx.Srtcxl_pnrdtl == "Highest" {
			hghlow = -1
		}
		sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "totcxl", Value: hghlow}}}}
	}

	// Final match pipeline
	var mtchfn bson.D
	if len(mtchdt) != 0 {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: mtchdt}}}}
	} else {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{}}}
	}

	// Logic download data
	if istDownld == "downld" {
		FncJeddahPnrdtlDownld(c, csvFilenm, inputx, mtchfn, sortdt, tablex, contxt)
	} else {

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
				bson.D{{Key: "$skip", Value: (max(inputx.Pagenw_pnrdtl, 1) - 1) * inputx.Limitp_pnrdtl}},
				bson.D{{Key: "$limit", Value: inputx.Limitp_pnrdtl}},
			}

			// Find user by username in database
			rawDtaset, err := tablex.Aggregate(contxt, pipeln)
			if err != nil {
				panic(err)
			}
			defer rawDtaset.Close(contxt)

			// Store to slice from raw bson
			for rawDtaset.Next(contxt) {
				var slcDtaset mdlJeddah.MdlJeddahPnrdtlDtbase
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

// Download PNR Detail all
func FncJeddahPnrdtlDownld(c *gin.Context, csvFilenm []string, inputx mdlJeddah.MdlJeddahAllprmInputx,
	mtchfn, sortdt bson.D, tablex *mongo.Collection, contxt context.Context) {

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
	pipeln := mongo.Pipeline{
		mtchfn,
		sortdt,
	}

	// Find user by username in database
	rawDtaset, err := tablex.Aggregate(contxt, pipeln)
	if err != nil {
		panic(err)
	}
	defer rawDtaset.Close(contxt)

	// Store to slice from raw bson
	for rawDtaset.Next(contxt) {
		var slcDtaset mdlJeddah.MdlJeddahPnrdtlDtbase
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
}

// Get Sync map data LC and PUN prev day
func FncJeddahDtlpnrSycmap(prvDateup string) *sync.Map {

	// Inisialisasi variabel
	fnldta := &sync.Map{}
	tmpdta := map[string]map[string]mdlJeddah.MdlJeddahPnrdtlDtbase{}

	// Select database and collection
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_pnrdtl")
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
		var object mdlJeddah.MdlJeddahPnrdtlDtbase
		datarw.Decode(&object)
		syckey := object.Airlfl + object.Flnbfl + object.Depart + strconv.Itoa(int(object.Datefl))
		if _, ist := tmpdta[syckey]; !ist {
			if len(tmpdta) < 1 {
				tmpdta = map[string]map[string]mdlJeddah.MdlJeddahPnrdtlDtbase{}
			}
			tmpdta[syckey] = map[string]mdlJeddah.MdlJeddahPnrdtlDtbase{}
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
