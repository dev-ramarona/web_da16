package fncJeddah

import (
	fncGlobal "back/global/function"
	mdlJeddah "back/jeddah/model"
	"bufio"
	"bytes"
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
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get template retail series blank for update
func FncJeddahRtlsrsTmplte(c *gin.Context) {

	// Set header untuk file CSV
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=Template_Update_Retail_Series.csv")
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")

	// Streaming file CSV ke client
	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()
	writer.Write([]string{"Pnrcde:(DO NOT CHANGE)", "Timecr:(DO NOT CHANGE)",
		"[Retail/Series/-]:(CHANGE ONLY WITH THIS FORMAT)", "Intrln"})

	// Select db and context to do
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_pnrsmr")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find user by username in database
	rawDtaset, err := tablex.Find(contxt, bson.D{
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "notedt", Value: ""}},
			bson.D{{Key: "notedt", Value: bson.D{{Key: "$exists", Value: false}}}},
		}},
	})
	if err != nil {
		panic(err)
	}
	defer rawDtaset.Close(contxt)

	// Store to slice from raw bson
	for rawDtaset.Next(contxt) {
		var slcDtaset mdlJeddah.MdlJeddahPnrsmrDtbase
		rawDtaset.Decode(&slcDtaset)
		writer.Write([]string{
			slcDtaset.Pnrcde,
			strconv.Itoa(int(slcDtaset.Timecr)),
			slcDtaset.Rtlsrs,
			slcDtaset.Intrln,
		})
	}
}

// Get Response Update database from input
func FncJeddahRtlsrsUpdate(c *gin.Context) {

	// Bind JSON Body input to variable
	var inputx mdlJeddah.MdlJeddahPnrsmrDtbase
	if err := c.BindJSON(&inputx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
	}

	// Push Summaryt data
	rsupdt := fncGlobal.FncGlobalDtbaseBlkwrt([]mongo.WriteModel{
		mongo.NewUpdateOneModel().
			SetFilter(bson.M{"prmkey": inputx.Prmkey}).
			SetUpdate(bson.M{"$set": inputx}).
			SetUpsert(true)}, "jeddah_pnrsmr")
	if rsupdt != nil {
		panic("Error Insert/Update to DB:" + rsupdt.Error())
	}

	// Update detail PNR
	func() {

		// Select database and collection
		tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_pnrdtl")
		contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Get route data
		datarw, err := tablex.Find(contxt, bson.M{"$and": []bson.M{
			{"pnrcde": inputx.Pnrcde}, {"timecr": inputx.Timecr}}})
		if err != nil {
			panic(err)
		}
		defer datarw.Close(contxt)

		// Append to slice
		pnrdtlMgomdl := []mongo.WriteModel{}
		for datarw.Next(contxt) {
			var object mdlJeddah.MdlJeddahPnrdtlDtbase
			datarw.Decode(&object)
			object.Rtlsrs = inputx.Rtlsrs
			object.Notedt = inputx.Notedt
			pnrdtlMgomdl = append(pnrdtlMgomdl, mongo.NewUpdateOneModel().
				SetFilter(bson.M{"prmkey": object.Prmkey}).
				SetUpdate(bson.M{"$set": object}).
				SetUpsert(true))
		}

		// Update final data
		rsupdt := fncGlobal.FncGlobalDtbaseBlkwrt(pnrdtlMgomdl, "jeddah_pnrdtl")
		if rsupdt != nil {
			panic("Error Insert/Update to DB:" + rsupdt.Error())
		}
	}()

	// Send token to frontend
	c.JSON(http.StatusOK, "success")
}

// Action upload rtlsrs
func FncJeddahRtlsrsUpload(c *gin.Context) {
	if fncGlobal.Status.Action == 0.0 {
		fncGlobal.Status.Action = 0.1

		// Limit size upload (10 MB)
		const maxupl = 10 << 20
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxupl)

		// Parse multipart form
		mltplr, err := c.Request.MultipartReader()
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		// Declare global variable
		var maxdta, nowdta int64
		var mgodtl, mgosmr []mongo.WriteModel
		errrsp, sycgrp := [][]string{}, sync.WaitGroup{}
		upldby := c.Param("upldby")
		if upldby == "" {
			upldby = "User not found"
		}

		// Looping each part/file
		for {
			nxpart, err := mltplr.NextPart()

			// Check isset next part
			if err == io.EOF {
				break
			}

			// Handle error next part
			if err != nil {
				c.JSON(http.StatusBadRequest, "next part error")
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

			// Count total csv r
			buffer := bytes.Buffer{}
			newlne := io.TeeReader(nxpart, &buffer)

			// Scanner untuk hitung baris
			scannr := bufio.NewScanner(newlne)
			for scannr.Scan() {
				maxdta++
			}
			if err := scannr.Err(); err != nil {
				c.JSON(http.StatusInternalServerError, "Error count row")
				return
			}

			// Siapkan CSV reader
			reader := csv.NewReader(bytes.NewReader(buffer.Bytes()))
			reader.FieldsPerRecord = -1
			reader.LazyQuotes = true
			reader.TrimLeadingSpace = true

			// Read and process CSV file Header
			header, err := reader.Read()
			if err != nil {
				c.JSON(http.StatusInternalServerError, "Error get header")
				continue
			}
			if len(header) > 0 {
				header[0] = strings.TrimPrefix(header[0], "\uFEFF") // hapus BOM
			}

			// Declare scope variable
			sycgrp.Add(1)
			go func() {
				defer sycgrp.Done()
				nowrow, limitz := 1, 100

				// Looping each row
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
					rawPnrcde, rawTimecr, rawRtlsrs := row[0], row[1], row[2]
					tmpNowerr := []string{rawPnrcde, rawTimecr, rawRtlsrs, flname}
					for i := range row {
						row[i] = strings.TrimSpace(row[i])
					}

					// Treatment PNR Code
					if len(rawPnrcde) != 6 {
						tmpNowerr = append(tmpNowerr, fmt.Sprintf("Row %d: Invalid PNR Code '%s'",
							nowrow+1, rawPnrcde))
					}

					// Treatment time create
					intTimecr, err := strconv.Atoi(rawTimecr)
					if err != nil || len(rawTimecr) != 10 {
						tmpNowerr = append(tmpNowerr, fmt.Sprintf("Row %d: Invalid Timecr '%s'",
							nowrow+1, rawTimecr))
					}

					// Treatment Retail/Series
					rilRtlsrs := ""
					mapRtlsrs := map[string]string{"retail": "Retail", "series": "Series", "-": "-"}
					if val, ist := mapRtlsrs[strings.ToLower(strings.TrimSpace(rawRtlsrs))]; ist {
						rilRtlsrs = val
					} else {
						tmpNowerr = append(tmpNowerr, fmt.Sprintf("Row %d: Invalid Retail/Series '%s'",
							nowrow+1, rawRtlsrs))
					}

					// If any error in row, then skip and save to error response
					if len(tmpNowerr) > 4 {
						errrsp = append(errrsp, tmpNowerr)
						continue
					}

					// Update to database PNR Detail
					mgodtl = append(mgodtl, mongo.NewUpdateManyModel().
						SetFilter(bson.M{"pnrcde": rawPnrcde, "timecr": int64(intTimecr)}).
						SetUpdate(bson.M{"$set": bson.M{
							"rtlsrs": rilRtlsrs, "notedt": fmt.Sprintf("Updated %s by %s", rawRtlsrs, upldby),
						}}).
						SetUpsert(true))

					// Push mongo pnrdtl
					if len(mgodtl) > limitz && len(mgodtl) != 0 {
						rsupdt := fncGlobal.FncGlobalDtbaseBlkwrt(mgodtl, "jeddah_pnrdtl")
						if rsupdt != nil {
							panic("Error Insert/Update to DB:" + rsupdt.Error())
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
						rsupdt := fncGlobal.FncGlobalDtbaseBlkwrt(mgosmr, "jeddah_pnrsmr")
						if rsupdt != nil {
							panic("Error Insert/Update to DB:" + rsupdt.Error())
						}
						fmt.Println("Updated batch 100 jeddah_pnrsmr")
						mgosmr = []mongo.WriteModel{}
					}

					// Count row
					nowrow++
					atomic.AddInt64(&nowdta, 1) // aman untuk goroutine
					fncGlobal.Status.Action = float64(nowdta) / float64(maxdta) * 100
				}
			}()
		}

		// Wait until all file clear
		sycgrp.Wait()

		// Push mongo pnrdtl
		if len(mgodtl) > 0 {
			rsupdt := fncGlobal.FncGlobalDtbaseBlkwrt(mgodtl, "jeddah_pnrdtl")
			if rsupdt != nil {
				panic("Error Insert/Update to DB:" + rsupdt.Error())
			}
		}

		// Push mongo pnrsmr
		if len(mgosmr) > 0 {
			rsupdt := fncGlobal.FncGlobalDtbaseBlkwrt(mgosmr, "jeddah_pnrsmr")
			if rsupdt != nil {
				panic("Error Insert/Update to DB:" + rsupdt.Error())
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
			return
		}

		// Respond done
		fncGlobal.Status.Action = 0
		c.JSON(http.StatusOK, fncGlobal.Status)
	}
}
