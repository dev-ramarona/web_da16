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
func FncJeddahPnrsmrFrntnd(c *gin.Context) {

	// Bind JSON Body input to variable
	istDownld := c.Param("downld")
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inputx mdlJeddah.MdlJeddahAllprmInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Treatment date number
	intDatefl := 0
	if inputx.Datefl_pnrsmr != "" {
		strDatefl, _ := time.Parse("2006-01-02", inputx.Datefl_pnrsmr)
		intDatefl, _ = strconv.Atoi(strDatefl.Format("060102"))
	}

	// Select db and context to do
	var totidx = 0
	var slcobj = []mdlJeddah.MdlJeddahPnrsmrDtbase{}
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_pnrsmr")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "datefl", Value: 1}}}}
	var sycwgp sync.WaitGroup

	// Special in this Summary PNR match
	if inputx.Datefl_pnrsmr != "" || inputx.Flnbfl_pnrsmr != "" ||
		inputx.Routfl_pnrsmr != "" || inputx.Airlfl_pnrsmr != "" {
		nowArrcpn := []string{}
		for _, params := range []string{inputx.Airlfl_pnrsmr,
			inputx.Flnbfl_pnrsmr, inputx.Routfl_pnrsmr,
			strconv.Itoa(intDatefl)} {
			if params != "" {
				nowArrcpn = append(nowArrcpn, params)
			} else {
				nowArrcpn = append(nowArrcpn, ".*")
			}
		}
		strArrcpn := strings.Join(nowArrcpn, "-")
		csvFilenm = append(csvFilenm, strconv.Itoa(intDatefl))
		mtchdt = append(mtchdt, bson.D{{Key: "arrcpn",
			Value: bson.D{{Key: "$regex", Value: strArrcpn}}}})
	}

	// Check if data Past date is isset
	if inputx.Psdate_pnrsmr == "Hide Past Date" {
		nowTimenb := time.Now().Format("0601020000")
		intTimenb, _ := strconv.Atoi(nowTimenb)
		csvFilenm = append(csvFilenm, inputx.Psdate_pnrsmr, nowTimenb)
		mtchdt = append(mtchdt, bson.D{{Key: "timedp",
			Value: bson.D{{Key: "$gte", Value: intTimenb}}}})
	}
	if inputx.Pnrcde_pnrsmr != "" {
		pnrArrays := strings.Split(inputx.Pnrcde_pnrsmr, "|")
		csvFilenm = append(csvFilenm, inputx.Pnrcde_pnrsmr)
		mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "pnrcde", Value: bson.D{{Key: "$in", Value: pnrArrays}}}},
			bson.D{{Key: "arrspl", Value: bson.D{{Key: "$regex", Value: inputx.Pnrcde_pnrsmr}}}},
		}}})

	}
	if inputx.Agtnme_pnrsmr != "" {
		csvFilenm = append(csvFilenm, inputx.Agtnme_pnrsmr)
		mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "agtnme", Value: bson.D{{Key: "$regex", Value: inputx.Agtnme_pnrsmr}}}},
			bson.D{{Key: "agtdtl", Value: bson.D{{Key: "$regex", Value: inputx.Agtnme_pnrsmr}}}}}}})
	}
	if inputx.Srtspl_pnrsmr != "" {
		csvFilenm = append(csvFilenm, "Split data")
		mtchdt = append(mtchdt, bson.D{{Key: "totspl",
			Value: bson.D{{Key: "$gt", Value: 0}}}})
		hghlow := 1
		if inputx.Srtspl_pnrsmr == "Highest" {
			hghlow = -1
		}
		sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "totspl", Value: hghlow}}}}
	}
	if inputx.Srtcxl_pnrsmr != "" {
		csvFilenm = append(csvFilenm, "Cancel data")
		mtchdt = append(mtchdt, bson.D{{Key: "totcxl",
			Value: bson.D{{Key: "$gt", Value: 0}}}})
		hghlow := 1
		if inputx.Srtcxl_pnrsmr == "Highest" {
			hghlow = -1
		}
		sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "totcxl", Value: hghlow}}}}
	}
	if istDownld == "rtlsrs" {
		mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "notedt", Value: ""}},
			bson.D{{Key: "notedt", Value: bson.D{{Key: "$exists", Value: false}}}}}}})
	}

	// Final match pipeline
	var trdPiplne bson.D
	if len(mtchdt) != 0 {
		trdPiplne = bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: mtchdt}}}}
	} else {
		trdPiplne = bson.D{{Key: "$match", Value: bson.D{}}}
	}

	// Logic download data
	if istDownld == "downld" {

		// Set header untuk file CSV
		fnlFilenm := strings.Join(csvFilenm, "_")
		c.Header("Content-Type", "text/csv")
		c.Header("Content-Disposition", "attachment; filename=Jeddah_Summary_PNR_"+fnlFilenm+".csv")
		c.Header("Access-Control-Expose-Headers", "Content-Disposition")

		// Streaming file CSV ke client
		writer := csv.NewWriter(c.Writer)
		defer writer.Flush()
		writer.Write([]string{"Prmkey", "Routfl", "Timedp", "Timerv", "Dateup", "Timeup", "Timecr",
			"Agtnme", "Agtdtl", "Agtidn", "Pnrcde", "Intrln", "Rtlsrs", "Arrcpn", "Agtdie", "Totisd",
			"Totbok", "Totpax", "Totcxl", "Totspl", "Arrspl", "Notedt"})

		// Get All Match Data
		nowPiplne := mongo.Pipeline{
			trdPiplne,
			sortdt,
		}

		// Find user by username in database
		rawDtaset, err := tablex.Aggregate(contxt, nowPiplne)
		if err != nil {
			panic(err)
		}
		defer rawDtaset.Close(contxt)

		// Store to slice from raw bson
		for rawDtaset.Next(contxt) {
			var slcDtaset mdlJeddah.MdlJeddahPnrsmrDtbase
			rawDtaset.Decode(&slcDtaset)
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
	} else {

		// Get Total Count Data
		sycwgp.Add(1)
		go func() {
			defer sycwgp.Done()
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
		sycwgp.Add(1)
		go func() {
			defer sycwgp.Done()
			nowPiplne := mongo.Pipeline{
				trdPiplne,
				sortdt,
				bson.D{{Key: "$skip", Value: (max(inputx.Pagenw_pnrsmr, 1) - 1) * inputx.Limitp_pnrsmr}},
				bson.D{{Key: "$limit", Value: inputx.Limitp_pnrsmr}},
			}

			// Find user by username in database
			rawDtaset, err := tablex.Aggregate(contxt, nowPiplne)
			if err != nil {
				panic(err)
			}
			defer rawDtaset.Close(contxt)

			// Store to slice from raw bson
			for rawDtaset.Next(contxt) {
				var slcDtaset mdlJeddah.MdlJeddahPnrsmrDtbase
				rawDtaset.Decode(&slcDtaset)
				slcobj = append(slcobj, slcDtaset)
			}
		}()

		// Waiting until all go done
		sycwgp.Wait()

		// Return final output
		c.JSON(200, gin.H{"totdta": totidx, "arrdta": slcobj})
	}
}
