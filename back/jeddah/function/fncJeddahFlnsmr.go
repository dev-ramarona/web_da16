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

// Get Summary PNR from database
func FncJeddahFlnsmrGetall(c *gin.Context) {

	// Bind JSON Body input to variable
	istDownld := c.Param("downld")
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inputx mdlJeddah.MdlJeddahAllprmInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Treatment date number
	intDatefl := 0
	if inputx.Datefl_flnsmr != "" {
		strDatefl, _ := time.Parse("2006-01-02", inputx.Datefl_flnsmr)
		intDatefl, _ = strconv.Atoi(strDatefl.Format("060102"))
	}

	// Select db and context to do
	var totidx = 0
	var slcobj = []mdlJeddah.MdlJeddahFlnsmrDtbase{}
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_flnsmr")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "datefl", Value: 1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	if inputx.Datefl_flnsmr != "" {
		csvFilenm = append(csvFilenm, strconv.Itoa(intDatefl))
		mtchdt = append(mtchdt, bson.D{{Key: "datefl",
			Value: intDatefl}})
	}
	if inputx.Airlfl_flnsmr != "" {
		csvFilenm = append(csvFilenm, inputx.Airlfl_flnsmr)
		mtchdt = append(mtchdt, bson.D{{Key: "airlfl",
			Value: inputx.Airlfl_flnsmr}})
	}
	if inputx.Flnbfl_flnsmr != "" {
		csvFilenm = append(csvFilenm, inputx.Flnbfl_flnsmr)
		mtchdt = append(mtchdt, bson.D{{Key: "flnbfl",
			Value: inputx.Flnbfl_flnsmr}})
	}
	if inputx.Routfl_flnsmr != "" {
		csvFilenm = append(csvFilenm, inputx.Routfl_flnsmr)
		mtchdt = append(mtchdt, bson.D{{Key: "routfl",
			Value: inputx.Routfl_flnsmr}})
	}
	if inputx.Srtspl_flnsmr != "" {
		csvFilenm = append(csvFilenm, "Split data")
		mtchdt = append(mtchdt, bson.D{{Key: "totspl",
			Value: bson.D{{Key: "$gt", Value: 0}}}})
		hghlow := 1
		if inputx.Srtspl_flnsmr == "Highest" {
			hghlow = -1
		}
		sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "totspl", Value: hghlow}}}}
	}
	if inputx.Srtcxl_flnsmr != "" {
		csvFilenm = append(csvFilenm, "Cancel data")
		mtchdt = append(mtchdt, bson.D{{Key: "totcxl",
			Value: bson.D{{Key: "$gt", Value: 0}}}})
		hghlow := 1
		if inputx.Srtcxl_flnsmr == "Highest" {
			hghlow = -1
		}
		sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "totcxl", Value: hghlow}}}}
	}
	if inputx.Psdate_flnsmr == "Hide Past Date" {
		nowDatefl := time.Now().Format("060102")
		intDatefl, _ := strconv.Atoi(nowDatefl)
		csvFilenm = append(csvFilenm, inputx.Psdate_flnsmr, nowDatefl)
		mtchdt = append(mtchdt, bson.D{{Key: "datefl",
			Value: bson.D{{Key: "$gte", Value: intDatefl}}}})
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
			var slcDtaset mdlJeddah.MdlJeddahFlnsmrDtbase
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
					totidx = int(count) // Konversi dari int32 ke int
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
				bson.D{{Key: "$skip", Value: (max(inputx.Pagenw_flnsmr, 1) - 1) * inputx.Limitp_flnsmr}},
				bson.D{{Key: "$limit", Value: inputx.Limitp_flnsmr}},
			}

			// Find user by username in database
			rawDtaset, err := tablex.Aggregate(contxt, pipeln)
			if err != nil {
				panic(err)
			}
			defer rawDtaset.Close(contxt)

			// Store to slice from raw bson
			for rawDtaset.Next(contxt) {
				var slcDtaset mdlJeddah.MdlJeddahFlnsmrDtbase
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
