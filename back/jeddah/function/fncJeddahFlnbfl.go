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
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get flight number format map data from database
func FncJeddahFlnbflSycmap() (map[string][]mdlJeddah.MdlJeddahFlnbflDtbase, *sync.Map, float64) {

	// Inisialisasi variabel
	fnlmap := map[string][]mdlJeddah.MdlJeddahFlnbflDtbase{}
	sycmap := &sync.Map{}

	// Select database and collection
	tablex := fncGlobal.Client.Database(fncGlobal.Dbases).Collection("jeddah_flnbfl")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	// datarw, err := tablex.Find(contxt, bson.M{"$and": []bson.M{
	// 	{"flnbfl": "323"}, {"datefl": 251012}, {"$or": []bson.M{
	// 		{"isjedh": "Jeddah"}, {"isjedh": ""},
	// 		{"isjedh": bson.M{"$exists": false}},
	// 	}}}})
	datarw, err := tablex.Find(contxt, bson.M{"$or": []bson.M{
		{"isjedh": "Jeddah"}, {"isjedh": ""},
		{"isjedh": bson.M{"$exists": false}},
	}})
	if err != nil {
		panic(err)
	}
	defer datarw.Close(contxt)

	// Append to slice
	totdta := 0.0
	for datarw.Next(contxt) {
		var object mdlJeddah.MdlJeddahFlnbflDtbase
		datarw.Decode(&object)
		sycmap.Store(object.Prmkey, object)

		// Filter registered airlines only
		if !slices.Contains([]string{"JT", "ID", "IW", "IU", "OD", "SL"}, object.Airlfl) {
			continue
		}
		if _, ist := fnlmap[object.Airlfl]; !ist {
			fnlmap[object.Airlfl] = []mdlJeddah.MdlJeddahFlnbflDtbase{}
		}
		totdta++
		fnlmap[object.Airlfl] = append(fnlmap[object.Airlfl], object)
	}

	// return data
	return fnlmap, sycmap, totdta
}

// Get Response Update database from input
func FncJeddahFlnbflUpdate(c *gin.Context) {

	// Bind JSON Body input to variable
	var inputx mdlJeddah.MdlJeddahFlnbflInputx
	if err := c.BindJSON(&inputx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
	}

	// Treatment date number
	intDatefl, strDatefl := 0, ""
	intTimeup, _ := strconv.Atoi(time.Now().Format("0601021504"))
	if inputx.Datefl != "" {
		rawDatefl, _ := time.Parse("2006-01-02", inputx.Datefl)
		strDatefl = rawDatefl.Format("060102")
		intDatefl, _ = strconv.Atoi(strDatefl)
	}

	// Push Summaryt data
	prmkey := strDatefl + inputx.Airlfl + inputx.Flnbfl + inputx.Routfl[0:3]
	rslupd := fncGlobal.FncGlobalDtbaseBlkwrt([]mongo.WriteModel{
		mongo.NewUpdateOneModel().
			SetFilter(bson.M{"prmkey": prmkey}).
			SetUpdate(bson.M{"$set": mdlJeddah.MdlJeddahFlnbflDtbase{
				Prmkey: prmkey,
				Datefl: int32(intDatefl),
				Airlfl: inputx.Airlfl,
				Flnbfl: inputx.Flnbfl,
				Routfl: inputx.Routfl,
				Depart: inputx.Routfl[0:3],
				Fltype: inputx.Fltype,
				Timeup: int64(intTimeup),
				Updtby: inputx.Updtby,
			}}).
			SetUpsert(true)}, "jeddah_flnbfl")
	if rslupd != nil {
		c.JSON(http.StatusInternalServerError, "failed"+rslupd.Error())
	}

	// Send token to frontend
	c.JSON(http.StatusOK, "success")
}

// Get template flight number upload
func FncJeddahFlnbflTmplte(c *gin.Context) {

	// Set header untuk file CSV
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=Template_Upload_Flight_Number.csv")
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")

	// Streaming file CSV ke client
	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()
	writer.Write([]string{
		"Date Flight (DD-MMM-YY)", "Airline (AB)", "Flight Number (1/23/456/7890)",
		"Route Flight (ABC-DEF)", "Flight Type (Outgoing/Incoming)"})
	writer.Write([]string{"31-May-25", "JT", "888", "CGK-KNO", "Outgoing"})
}

// Action upload rtlsrs
func FncJeddahFlnbflUpload(c *gin.Context) {
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
		var mgofln []mongo.WriteModel
		errrsp, sycgrp := [][]string{}, sync.WaitGroup{}
		upldby := c.Param("upldby")
		if upldby == "" {
			upldby = "User not found"
		}

		// Declare date now variable
		var strTimenw = time.Now().Format("0601021504")
		var intTimenw, _ = strconv.Atoi(strTimenw)
		var intDatenw, _ = strconv.Atoi(strTimenw[0:6])

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
					for i := range row {
						row[i] = strings.TrimSpace(row[i])
					}
					rawDatefl, rawAirlne, rawFlnbfl, rawRoutfl, rawFltype :=
						row[0], row[1], row[2], row[3], row[4]
					fnlDatefl, fnlAirlne, fnlFlnbfl, fnlRoutfl, fnlFltype := 0, "", "", "", ""
					tmpNowerr := []string{rawDatefl, rawAirlne, rawFlnbfl, rawRoutfl, rawFltype}

					// Treatment Date flight
					prsDatefl, erraaa := time.Parse("2-Jan-06", rawDatefl)
					intDatefl, errbbb := strconv.Atoi(prsDatefl.Format("060102"))
					if erraaa != nil || errbbb != nil {
						tmpNowerr = append(tmpNowerr, fmt.Sprintf("Row %d: Invalid Datefl '%s', (%v|%v)",
							nowrow+1, rawDatefl, erraaa, errbbb))
					}
					fnlDatefl = intDatefl

					// Treatment airline
					if len(rawAirlne) != 2 {
						tmpNowerr = append(tmpNowerr, fmt.Sprintf("Row %d: Invalid Airline '%s'",
							nowrow+1, rawAirlne))
					}
					fnlAirlne = rawAirlne

					// Treatment Flight number
					_, errccc := strconv.Atoi(rawFlnbfl)
					if len(rawFlnbfl) > 4 || errccc != nil {
						tmpNowerr = append(tmpNowerr, fmt.Sprintf("Row %d: Invalid Flight Nunmber '%s', (%v)",
							nowrow+1, rawFlnbfl, errccc))
					}
					fnlFlnbfl = rawFlnbfl

					// Treatment Routefl
					rgxRoutfl := regexp.MustCompile(`^[A-Z]{3}-[A-Z]{3}$`)
					if !rgxRoutfl.MatchString(rawRoutfl) {
						tmpNowerr = append(tmpNowerr, fmt.Sprintf("Row %d: Invalid Routefl '%s'",
							nowrow+1, rawFlnbfl))
					}
					fnlRoutfl = rawRoutfl

					// Treatment Flight Type
					mapFltype := map[string]string{"outgoing": "Outgoing", "incoming": "Incoming", "-": "-"}
					mtcFltype, ist := mapFltype[strings.ToLower(strings.TrimSpace(rawFltype))]
					if !ist {
						tmpNowerr = append(tmpNowerr, fmt.Sprintf("Row %d: Invalid Retail/Series '%s'",
							nowrow+1, rawFltype))
					}
					fnlFltype = mtcFltype

					// If any error in row, then skip and save to error response
					if len(tmpNowerr) > 5 {
						errrsp = append(errrsp, tmpNowerr)
						continue
					}

					// Update to database PNR Detail
					prmkey := strconv.Itoa(fnlDatefl) + fnlAirlne + fnlFlnbfl + fnlRoutfl[0:3]
					mgofln = append(mgofln, mongo.NewUpdateManyModel().
						SetFilter(bson.M{"prmkey": prmkey}).
						SetUpdate(bson.M{"$set": mdlJeddah.MdlJeddahFlnbflDtbase{
							Prmkey: prmkey,
							Datefl: int32(fnlDatefl),
							Dateup: int32(intDatenw),
							Timeup: int64(intTimenw),
							Airlfl: fnlAirlne,
							Flnbfl: fnlFlnbfl,
							Depart: fnlRoutfl[0:3],
							Routfl: fnlRoutfl,
							Fltype: fnlFltype,
							Updtby: upldby,
						}}).
						SetUpsert(true))

					// Push mongo pnrdtl
					if len(mgofln) > limitz && len(mgofln) != 0 {
						rsupdt := fncGlobal.FncGlobalDtbaseBlkwrt(mgofln, "jeddah_flnbfl")
						if rsupdt != nil {
							panic("Error Insert/Update to DB:" + rsupdt.Error())
						}
						fmt.Println("Updated batch 100 jeddah_flnbfl")
						mgofln = []mongo.WriteModel{}
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
		if len(mgofln) > 0 {
			rsupdt := fncGlobal.FncGlobalDtbaseBlkwrt(mgofln, "jeddah_flnbfl")
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
