package fncJeddah

import (
	fncGlobal "back/global/function"
	mdlJeddah "back/jeddah/model"
	fncSbrapi "back/sbrapi/function"
	mdlSbrapi "back/sbrapi/model"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Running main process jeddah
func FncJeddahPrcessMainpg(c *gin.Context) {
	if fncGlobal.Status.Sbrapi == 0.0 {
		fncGlobal.Status.Sbrapi = 0.01

		// Insdicaor Process Start
		nowTimenw := time.Now().AddDate(0, 0, -0).Format("0601021504")
		totWokrer, nowTotdta := 1, int64(0)
		slcFlnbfl, sycFlnbfl, maxTotdta := FncJeddahFlnbflSycmap(nowTimenw[0:6])
		prvDatefl := FncJeddahActlogLstdta()
		slcDrules := FncJeddahDrulesSycmap()
		sycAgtnme := FncJeddahAgtnmeSycmap()
		sycPnrlog := FncJeddahPnrlogSycmap(prvDatefl)
		sycDtlpnr := FncJeddahDtlpnrSycmap(prvDatefl)
		idcAgtnme, idcPnrcde, idcFlnbfl := &sync.Map{}, &sync.Map{}, &sync.Map{}

		// Looping all flight number
		for airlfl, slice1 := range slcFlnbfl {

			// Get 10 API sessions/tokens
			slcRspssn, err := fncSbrapi.FncSbrapiCrtssnMultpl(airlfl, totWokrer)
			lgcRspssn := err != nil || slcRspssn[0].Bsttkn == "" || len(slcRspssn) < 1
			if lgcRspssn {
				fmt.Println(airlfl)
				panic("eror user ga dapet" + err.Error())
			}

			// Prepare job queue
			jobFlbase := make(chan mdlJeddah.MdlJeddahFlnbflDtbase, 1000)
			var swg sync.WaitGroup

			// Launch 10 workers using 10 tokens
			for i := 0; i < totWokrer; i++ {
				if len(slcRspssn) >= i+1 {
					if slcRspssn[i].Bsttkn != "Failed" {
						swg.Add(1)
						go FncJeddahPrcessWorker(i, &slcRspssn[i], &swg, jobFlbase,
							sycAgtnme, sycPnrlog, sycDtlpnr, sycFlnbfl, idcAgtnme,
							idcPnrcde, idcFlnbfl, nowTimenw, slcDrules, &nowTotdta, &maxTotdta)
					}
				}
			}

			// Looping all data
			for _, flbase := range slice1 {
				jobFlbase <- flbase
			}

			// Finish
			close(jobFlbase)
			swg.Wait()
			fmt.Println("Done", airlfl, nowTimenw[0:6])
			fncSbrapi.FncSbrapiClsssnMultpl(slcRspssn)
		}

		// Return to done status
		fncGlobal.Status.Sbrapi = 0

		// Update done again
		// intDatenw, _ := strconv.Atoi(nowTimenw[0:6])
		// intTimenw, _ := strconv.Atoi(nowTimenw)
		// logAction := mdlJeddah.MdlJeddahActlogDtbase{
		// 	Dateup: int32(intDatenw), Timeup: int64(intTimenw),
		// 	Statdt: "Done"}
		// rsupdt := fncGlobal.FncGlobalDtbaseBlkwrt([]mongo.WriteModel{
		// 	mongo.NewUpdateOneModel().
		// 		SetFilter(bson.M{"dateup": logAction.Dateup}).
		// 		SetUpdate(bson.M{"$set": logAction}).
		// 		SetUpsert(true)}, "jeddah_actlog")
		// if !rsupdt {
		// 	panic("Error Insert/Update to DB")
		// }
	}
}

// Running process jeddah
func FncJeddahPrcessWorker(nbr int, tkn *mdlSbrapi.MdlSbrapiMsghdrParams, swg *sync.WaitGroup,
	jobFlbase <-chan mdlJeddah.MdlJeddahFlnbflDtbase,
	sycAgtnme, sycLcnpun, sycDtlpnr, sycFlnbfl, idcAgtnme, idcPnrcde, idcFlnbfl *sync.Map,
	nowTimenw string, slcDrules []mdlJeddah.MdlJeddahRulesjDtbase, nowTotdta *int64, maxTotdta *float64) {
	var mgomdlAgtnme, mgomdlLcnpun, mgomdlSmrfln []mongo.WriteModel
	var mgomdlSmrpnr, mgomdlDtlpnr, mgomdlFlnbfl []mongo.WriteModel
	var lmtdta, cntdta = 50, 0
	defer swg.Done()

	// Looping jobs data
outlop:
	for flbase := range jobFlbase {
		cntdta++
		atomic.AddInt64(nowTotdta, 1)
		fncGlobal.Status.Sbrapi = float64(*nowTotdta) / *maxTotdta * 100
		fnlAgtnme := map[string]mongo.WriteModel{}
		dbsFlnbfl, dbsDatefl, dbsDepart, dbsRoutfl, dbsAirlfl :=
			flbase.Flnbfl, flbase.Datefl, flbase.Depart, flbase.Routfl, flbase.Airlfl
		intTotisd, intTotbok, intTotpax, intTotcxl, intTotchg, intTotspl := 0, 0, 0, 0, 0, 0
		objParams := mdlSbrapi.MdlSbrapiMsghdrApndix{Airlfl: dbsAirlfl, Datefl: dbsDatefl,
			Depart: dbsDepart, Flnbfl: dbsFlnbfl, Routfl: dbsRoutfl}

		// Indicator
		fmt.Println("Start", cntdta, "-", dbsAirlfl, dbsFlnbfl, dbsDatefl,
			dbsRoutfl, time.Now().Format("06-01-02/15:04"))

		// Conver String and int date
		rawTimenw, _ := time.Parse("0601021504", nowTimenw)
		rawTimepv := rawTimenw.AddDate(0, 0, -1)
		strTimepv := rawTimepv.Format("0601021504")
		intTimepv, _ := strconv.Atoi(strTimepv)
		intTimenw, _ := strconv.Atoi(nowTimenw)
		intDatenw, _ := strconv.Atoi(nowTimenw[0:6])
		strDatenw := strconv.Itoa(int(dbsDatefl))

		// Get previous LC and PUN data
		prvLcnpun := map[string]mdlJeddah.MdlJeddahPnrlogDtbase{}
		if tmpPrvdtl, ist := sycLcnpun.Load(dbsAirlfl + dbsFlnbfl + dbsDepart + strDatenw); ist {
			if getPrvdtl, ist := tmpPrvdtl.(map[string]mdlJeddah.MdlJeddahPnrlogDtbase); ist {
				prvLcnpun = getPrvdtl
				sycLcnpun.Delete(dbsAirlfl + dbsFlnbfl + dbsDepart + strDatenw)
			}
		}

		// Get Summary PNR data
		prvDtlpnr := map[string]mdlJeddah.MdlJeddahPnrdtlDtbase{}
		if tmpPrvdtl, ist := sycDtlpnr.Load(dbsAirlfl + dbsFlnbfl + dbsDepart + strDatenw); ist {
			if getPrvdtl, ist := tmpPrvdtl.(map[string]mdlJeddah.MdlJeddahPnrdtlDtbase); ist {
				prvDtlpnr = getPrvdtl
				sycDtlpnr.Delete(dbsAirlfl + dbsFlnbfl + dbsDepart + strDatenw)
			}
		}

		// Check the date is the same or greater than today
		if dbsDatefl < int32(intTimepv) {
			for _, dtlpnr := range prvDtlpnr {

				// Get agent name Lcnpun
				FncJeddahAgtgetParams(&dtlpnr, true, idcAgtnme, &fnlAgtnme)

				// Push to Summary PNR
				nmodelDtlpnr := mongo.NewUpdateOneModel().
					SetFilter(bson.M{"prmkey": dtlpnr.Prmkey}).
					SetUpdate(bson.M{"$set": dtlpnr}).
					SetUpsert(true)
				mgomdlDtlpnr = append(mgomdlDtlpnr, nmodelDtlpnr)
			}
		} else {

			// Looping PUN data
			tmpBooked := map[string]int{}
			slcOutpun, errOutpun := fncSbrapi.FncSbrapiLcnpunMainob("PUN", *tkn, objParams)
			if errOutpun == nil {
				for _, lcnpun := range slcOutpun {

					// Push to mongo Lcnpun
					nmodelLcnpun := mongo.NewUpdateOneModel().
						SetFilter(bson.M{"prmkey": lcnpun.Prmkey}).
						SetUpdate(bson.M{"$set": lcnpun}).
						SetUpsert(true)
					mgomdlLcnpun = append(mgomdlLcnpun, nmodelLcnpun)

					// Push to temp data not issued
					tmpBooked[lcnpun.Pnrcde] = lcnpun.Totpax
				}
			}

			// Looping LC data
			fnlIsjedh := false
			var tmpSmrpnr, tmpDtlpnr, tmpLcnpun, tmpFlnbfl []mongo.WriteModel
			slcOutllc, errOutllc := fncSbrapi.FncSbrapiLcnpunMainob("LC", *tkn, objParams)
			if errOutllc == nil {

				// Looping LX data (TEMPORARY! DELETE AFTER RUNNING)
				slcOutlxx, errOutlxx := fncSbrapi.FncSbrapiLcnpunMainob("LX", *tkn, objParams)
				if errOutlxx == nil {
					for _, lcnpun := range slcOutlxx {
						nowPrmkey := lcnpun.Airlfl + lcnpun.Flnbfl + lcnpun.Depart +
							strconv.Itoa(int(lcnpun.Datefl)) + lcnpun.Pnrcde

						// Push to mongo Lcnpun
						nmodelLcnpun := mongo.NewUpdateOneModel().
							SetFilter(bson.M{"prmkey": lcnpun.Prmkey}).
							SetUpdate(bson.M{"$set": lcnpun}).
							SetUpsert(true)
						tmpLcnpun = append(tmpLcnpun, nmodelLcnpun)

						// Default summary PNR
						defDtlpnr := mdlJeddah.MdlJeddahPnrdtlDtbase{
							Prmkey: nowPrmkey, Airlfl: lcnpun.Airlfl, Flnbfl: lcnpun.Flnbfl,
							Depart: lcnpun.Depart, Routfl: lcnpun.Routfl, Clssfl: lcnpun.Clssfl,
							Datefl: lcnpun.Datefl, Dateup: lcnpun.Dateup, Timeup: lcnpun.Timeup,
							Agtnme: lcnpun.Agtnme, Agtdtl: "", Agtidn: "",
							Pnrcde: lcnpun.Pnrcde, Rtlsrs: "", Toflnm: "",
							Drules: 0, Totisd: 0, Totbok: 0, Totpax: lcnpun.Totpax,
						}

						// Get remarks history split or cancel
						strToflnm, cekIsjedh :=
							FncJeddahItrmrlGetapi("rmv", tkn, &defDtlpnr, objParams, sycFlnbfl, idcPnrcde,
								idcFlnbfl, sycAgtnme, &tmpSmrpnr, &tmpFlnbfl)

						// cek is Jeddah
						nowIsjedh := false
						if cekIsjedh {
							fnlIsjedh, nowIsjedh, flbase.Isjedh = true, true, "Jeddah"
						}

						// Get agent name Lcnpun
						FncJeddahAgtgetParams(&defDtlpnr, nowIsjedh, idcAgtnme, &fnlAgtnme)

						// Response Itenary
						if strToflnm != "" {
							defDtlpnr.Toflnm = strToflnm
							defDtlpnr.Flstat = "Change"
							defDtlpnr.Totchg = lcnpun.Totpax
							defDtlpnr.Totpax = 0
							intTotchg += lcnpun.Totpax
						} else {
							defDtlpnr.Flstat = "Cancel"
							defDtlpnr.Totcxl = lcnpun.Totpax
							defDtlpnr.Totpax = lcnpun.Totpax
							intTotcxl += lcnpun.Totpax
						}

						// Push to Pnr log data
						xxxxxx, _ := json.MarshalIndent(defDtlpnr, "", "  ")
						fmt.Println(string(xxxxxx))
						nmodelDtlpnr := mongo.NewUpdateOneModel().
							SetFilter(bson.M{"prmkey": nowPrmkey}).
							SetUpdate(bson.M{"$set": defDtlpnr}).
							SetUpsert(true)
						tmpDtlpnr = append(tmpDtlpnr, nmodelDtlpnr)
					}
				}

				// Looping LC data
				for _, lcnpun := range slcOutllc {

					// Delcare Pnr log data
					fnlDtlpnr := mdlJeddah.MdlJeddahPnrdtlDtbase{
						Prmkey: dbsAirlfl + dbsFlnbfl + dbsDepart + strDatenw + lcnpun.Pnrcde,
						Airlfl: dbsAirlfl, Flnbfl: dbsFlnbfl, Depart: dbsDepart,
						Clssfl: lcnpun.Clssfl, Datefl: dbsDatefl, Dateup: int32(intDatenw),
						Intrln: "", Timeup: int64(intTimenw), Agtnme: lcnpun.Agtnme,
						Pnrcde: lcnpun.Pnrcde, Drules: 100, Totisd: lcnpun.Totpax,
						Totpax: lcnpun.Totpax, Routfl: dbsRoutfl}

					// Cek booked data from PUN
					intTotpax += lcnpun.Totpax
					valIssued := lcnpun.Totpax
					if nowBooked, ist := tmpBooked[lcnpun.Pnrcde]; ist {
						valIssued = lcnpun.Totpax - nowBooked
						intTotbok += nowBooked
						fnlDtlpnr.Totisd = valIssued
						fnlDtlpnr.Totbok = nowBooked
					}
					intTotisd += valIssued

					// Cek from data Summary PNR and get last remove data
					if prv, ist := prvDtlpnr[lcnpun.Pnrcde]; ist {
						delete(prvDtlpnr, lcnpun.Pnrcde)
						fnlDtlpnr.Dateup = prv.Dateup
						fnlDtlpnr.Datefl = prv.Datefl
						fnlDtlpnr.Timecr = prv.Timecr
						fnlDtlpnr.Totspl = prv.Totspl
						fnlDtlpnr.Totchg = prv.Totchg
						fnlDtlpnr.Totcxl = prv.Totcxl
						fnlDtlpnr.Toflnm = prv.Toflnm
						fnlDtlpnr.Flstat = prv.Flstat
						fnlDtlpnr.Notedt = prv.Notedt
					} else {

						// Get remarks history split or cancel
						_, cekIsjedh := FncJeddahItrmrlGetapi("", tkn, &fnlDtlpnr, objParams, sycFlnbfl, idcPnrcde,
							idcFlnbfl, sycAgtnme, &tmpSmrpnr, &tmpFlnbfl)

						// cek is Jeddah
						nowIsjedh := false
						if cekIsjedh {
							fnlIsjedh, nowIsjedh, flbase.Isjedh = true, true, "Jeddah"
						}

						// Get agent name Lcnpun
						FncJeddahAgtgetParams(&fnlDtlpnr, nowIsjedh, idcAgtnme, &fnlAgtnme)
					}

					// Cek from data Summary PNR and get last remove data
					if prv, ist := prvLcnpun[lcnpun.Pnrcde]; ist {
						delete(prvLcnpun, lcnpun.Pnrcde)
						fnlIsjedh = true
						nowRemove := prv.Totpax - lcnpun.Totpax
						if nowRemove > 0 {

							// Get remarks only
							FncJeddahItrmrlGetapi("", tkn, &fnlDtlpnr, objParams, sycFlnbfl, idcPnrcde,
								idcFlnbfl, sycAgtnme, &tmpSmrpnr, &tmpFlnbfl)

						}
					}

					// Push to Pnr log data
					for _, drules := range slcDrules {

						// Get parameter date
						timefs := strDatenw
						timesc := nowTimenw[0:6]
						if drules.Rlcolm == "dateup" {
							timefs = nowTimenw[0:6]
							timesc = strconv.Itoa(int(fnlDtlpnr.Dateup))
						}

						// Convert time
						param1, _ := time.Parse("060102", timefs)
						param2, _ := time.Parse("060102", timesc)

						// Start logic different time
						if math.Abs(param1.Sub(param2).Hours()/24) <= float64(drules.Rldays) {
							fnlDtlpnr.Drules = int(drules.Rlrate)
							break
						}
					}

					// Push to mongo Lcnpun
					nmodelLcnpun := mongo.NewUpdateOneModel().
						SetFilter(bson.M{"prmkey": lcnpun.Prmkey}).
						SetUpdate(bson.M{"$set": lcnpun}).
						SetUpsert(true)
					tmpLcnpun = append(tmpLcnpun, nmodelLcnpun)

					// Push to Summary PNR
					intTotcxl += fnlDtlpnr.Totcxl
					intTotspl += fnlDtlpnr.Totspl
					nmodelDtlpnr := mongo.NewUpdateOneModel().
						SetFilter(bson.M{"prmkey": fnlDtlpnr.Prmkey}).
						SetUpdate(bson.M{"$set": fnlDtlpnr}).
						SetUpsert(true)
					tmpDtlpnr = append(tmpDtlpnr, nmodelDtlpnr)
				}
			}

			switch {

			// If flight still available
			case len(slcOutllc) != 0:

				// Push new data flight to database
				flbase.Flstat = "Operate"
				nmodelFlnbfl := mongo.NewUpdateOneModel().
					SetFilter(bson.M{"prmkey": flbase.Prmkey}).
					SetUpdate(bson.M{"$set": flbase}).
					SetUpsert(true)
				mgomdlFlnbfl = append(mgomdlFlnbfl, nmodelFlnbfl)

				// If prev pnr isset but now null
				if len(prvLcnpun) > 0 {
					for _, lcnpun := range prvLcnpun {

						// Get prev summary pnr array split and cancel
						if nowDtlpnr, ist := prvDtlpnr[lcnpun.Pnrcde]; ist {
							delete(prvDtlpnr, lcnpun.Pnrcde)

							// Get remarks history split or cancel
							strToflnm, _ :=
								FncJeddahItrmrlGetapi("rmv", tkn, &nowDtlpnr, objParams, sycFlnbfl, idcPnrcde,
									idcFlnbfl, sycAgtnme, &mgomdlSmrpnr, &mgomdlFlnbfl)

							// Response Itenary
							if strToflnm != "" {
								nowDtlpnr.Toflnm = strToflnm
								nowDtlpnr.Flstat = "Change"
								nowDtlpnr.Totchg = nowDtlpnr.Totpax
								nowDtlpnr.Totpax = 0
								intTotchg += nowDtlpnr.Totpax
							} else {
								nowDtlpnr.Flstat = "Cancel"
								nowDtlpnr.Totcxl = nowDtlpnr.Totpax
								intTotcxl += nowDtlpnr.Totpax
							}

							// Push to Pnr log data
							nmodelDtlpnr := mongo.NewUpdateOneModel().
								SetFilter(bson.M{"prmkey": nowDtlpnr.Prmkey}).
								SetUpdate(bson.M{"$set": nowDtlpnr}).
								SetUpsert(true)
							mgomdlDtlpnr = append(mgomdlDtlpnr, nmodelDtlpnr)
						}
					}
				}

				// Continue if not exist Jeddah route
				if !fnlIsjedh {
					continue outlop
				} else {
					mgomdlSmrpnr = append(mgomdlSmrpnr, tmpSmrpnr...)
					mgomdlLcnpun = append(mgomdlLcnpun, tmpLcnpun...)
					mgomdlDtlpnr = append(mgomdlDtlpnr, tmpDtlpnr...)
					mgomdlFlnbfl = append(mgomdlFlnbfl, tmpFlnbfl...)
				}

			// If flight cancel/change but never in db
			case len(prvDtlpnr) == 0:

				// Looping LDN data
				fnlIsjedh := false
				flbase.Isjedh = "Non Jeddah"
				var tmpSmrpnr, tmpDtlpnr, tmpLcnpun, tmpFlnbfl []mongo.WriteModel
				slcOutldn, errOutldn := fncSbrapi.FncSbrapiLcnpunMainob("LDN", *tkn, objParams)
				if errOutldn == nil {
					for _, lcnpun := range slcOutldn {
						nowPrmkey := lcnpun.Airlfl + lcnpun.Flnbfl + lcnpun.Depart +
							strconv.Itoa(int(lcnpun.Datefl)) + lcnpun.Pnrcde

						// Push to mongo Lcnpun
						nmodelLcnpun := mongo.NewUpdateOneModel().
							SetFilter(bson.M{"prmkey": lcnpun.Prmkey}).
							SetUpdate(bson.M{"$set": lcnpun}).
							SetUpsert(true)
						tmpLcnpun = append(tmpLcnpun, nmodelLcnpun)

						// Default summary PNR
						defDtlpnr := mdlJeddah.MdlJeddahPnrdtlDtbase{
							Prmkey: nowPrmkey, Airlfl: lcnpun.Airlfl, Flnbfl: lcnpun.Flnbfl,
							Depart: lcnpun.Depart, Routfl: lcnpun.Routfl, Clssfl: lcnpun.Clssfl,
							Datefl: lcnpun.Datefl, Dateup: lcnpun.Dateup, Timeup: lcnpun.Timeup,
							Agtnme: lcnpun.Agtnme, Agtdtl: "", Agtidn: "",
							Pnrcde: lcnpun.Pnrcde, Rtlsrs: "", Toflnm: "",
							Drules: 0, Totisd: 0, Totbok: 0, Totpax: lcnpun.Totpax,
						}

						// Get remarks history split or cancel
						strToflnm, cekIsjedh :=
							FncJeddahItrmrlGetapi("rmv", tkn, &defDtlpnr, objParams, sycFlnbfl, idcPnrcde,
								idcFlnbfl, sycAgtnme, &tmpSmrpnr, &tmpFlnbfl)

						// cek is Jeddah
						nowIsjedh := false
						if cekIsjedh {
							fnlIsjedh, nowIsjedh, flbase.Isjedh = true, true, "Jeddah"
						}

						// Get agent name Lcnpun
						FncJeddahAgtgetParams(&defDtlpnr, nowIsjedh, idcAgtnme, &fnlAgtnme)

						// Response Itenary
						if strToflnm != "" {
							defDtlpnr.Toflnm = strToflnm
							defDtlpnr.Flstat = "Change"
							defDtlpnr.Totchg = lcnpun.Totpax
							defDtlpnr.Totpax = 0
							intTotchg += lcnpun.Totpax
						} else {
							defDtlpnr.Flstat = "Cancel"
							defDtlpnr.Totcxl = lcnpun.Totpax
							defDtlpnr.Totpax = lcnpun.Totpax
							intTotcxl += lcnpun.Totpax
						}

						// Push to Pnr log data
						nmodelDtlpnr := mongo.NewUpdateOneModel().
							SetFilter(bson.M{"prmkey": nowPrmkey}).
							SetUpdate(bson.M{"$set": defDtlpnr}).
							SetUpsert(true)
						tmpDtlpnr = append(tmpDtlpnr, nmodelDtlpnr)
					}
				}

				// Push new data flight to database and sycmap
				flbase.Flstat = "Cancel"
				nmodelFlnbfl := mongo.NewUpdateOneModel().
					SetFilter(bson.M{"prmkey": flbase.Prmkey}).
					SetUpdate(bson.M{"$set": flbase}).
					SetUpsert(true)
				mgomdlFlnbfl = append(mgomdlFlnbfl, nmodelFlnbfl)

				// Continue if not exist Jeddah route
				if !fnlIsjedh {
					continue outlop
				} else {
					mgomdlSmrpnr = append(mgomdlSmrpnr, tmpSmrpnr...)
					mgomdlLcnpun = append(mgomdlLcnpun, tmpLcnpun...)
					mgomdlDtlpnr = append(mgomdlDtlpnr, tmpDtlpnr...)
					mgomdlFlnbfl = append(mgomdlFlnbfl, tmpFlnbfl...)
				}

			// If flight cancel/change avail in db
			case len(prvDtlpnr) > 0:

				// Get agent name Lcnpun
				for _, dtlpnr := range prvDtlpnr {
					if dtlpnr.Agtidn == "" {
						FncJeddahAgtgetParams(&dtlpnr, true, idcAgtnme, &fnlAgtnme)

						// Push to Pnr log data
						nmodelDtlpnr := mongo.NewUpdateOneModel().
							SetFilter(bson.M{"prmkey": dtlpnr.Prmkey}).
							SetUpdate(bson.M{"$set": dtlpnr}).
							SetUpsert(true)
						mgomdlDtlpnr = append(mgomdlDtlpnr, nmodelDtlpnr)

						// Get remarks history split or cancel
						FncJeddahItrmrlGetapi("rmv", tkn, &dtlpnr, objParams, sycFlnbfl, idcPnrcde,
							idcFlnbfl, sycAgtnme, &mgomdlSmrpnr, &mgomdlFlnbfl)
					}
				}
			}

			// Push to Pnr log data
			flnsmrFlstat := "Operate"
			if intTotpax == 0 {
				flnsmrFlstat = "Cancel"
			}
			nmodelSmrfln := mongo.NewUpdateOneModel().
				SetFilter(bson.M{"prmkey": flbase.Prmkey}).
				SetUpdate(bson.M{"$set": mdlJeddah.MdlJeddahFlnsmrDtbase{
					Prmkey: flbase.Prmkey, Airlfl: dbsAirlfl, Flnbfl: dbsFlnbfl,
					Depart: dbsDepart, Routfl: dbsRoutfl, Datefl: dbsDatefl,
					Dateup: int32(intDatenw), Timeup: int64(intTimenw),
					Totisd: intTotisd, Totbok: intTotbok, Totpax: intTotpax,
					Totcxl: intTotcxl, Totchg: intTotchg, Totspl: intTotspl,
					Flstat: flnsmrFlstat}}).
				SetUpsert(true)
			mgomdlSmrfln = append(mgomdlSmrfln, nmodelSmrfln)
		}

		// Push mongo pnrlog
		if len(mgomdlLcnpun) > lmtdta && len(mgomdlLcnpun) != 0 {
			// rspBlkwrt := fncGlobal.FncGlobalDtbaseBlkwrt(mgomdlLcnpun, "jeddah_pnrlog")
			// if !rspBlkwrt {
			// 	fmt.Println("ERR LOG HERE, CAN'T INPUT LCNPUN")
			// }
			mgomdlLcnpun = []mongo.WriteModel{}
		}

		// Push mongo pnrlog
		if len(mgomdlDtlpnr) > lmtdta && len(mgomdlDtlpnr) != 0 {
			// rspBlkwrt := fncGlobal.FncGlobalDtbaseBlkwrt(mgomdlDtlpnr, "jeddah_pnrdtl")
			// if !rspBlkwrt {
			// 	fmt.Println("ERR LOG HERE, CAN'T INPUT DTLPNR")
			// }
			mgomdlDtlpnr = []mongo.WriteModel{}
		}

		// Push mongo pnrlog
		if len(mgomdlFlnbfl) > lmtdta && len(mgomdlFlnbfl) != 0 {
			// rspBlkwrt := fncGlobal.FncGlobalDtbaseBlkwrt(mgomdlFlnbfl, "jeddah_flnbfl")
			// if !rspBlkwrt {
			// 	fmt.Println("ERR LOG HERE, CAN'T INPUT FLNbfl")
			// }
			mgomdlFlnbfl = []mongo.WriteModel{}
		}

		// Push mongo pnrlog
		if len(mgomdlSmrpnr) > lmtdta && len(mgomdlSmrpnr) != 0 {
			// rspBlkwrt := fncGlobal.FncGlobalDtbaseBlkwrt(mgomdlSmrpnr, "jeddah_pnrsmr")
			// if !rspBlkwrt {
			// 	fmt.Println("ERR LOG HERE, CAN'T INPUT SMRPNR")
			// }
			mgomdlSmrpnr = []mongo.WriteModel{}
		}

		// Indicator and push final agent name
		for _, agtnme := range fnlAgtnme {
			mgomdlAgtnme = append(mgomdlAgtnme, agtnme)
		}
		fmt.Println("Done", cntdta, "-", dbsAirlfl, dbsFlnbfl, dbsDatefl,
			dbsRoutfl, time.Now().Format("06-01-02/15:04"))
	}

	// Push mongo detail
	for _, mgo := range map[string][]mongo.WriteModel{
		"jeddah_pnrlog": mgomdlLcnpun,
		"jeddah_agentx": mgomdlAgtnme,
		"jeddah_pnrsmr": mgomdlSmrpnr,
		"jeddah_pnrdtl": mgomdlDtlpnr,
		"jeddah_flnsmr": mgomdlSmrfln,
		"jeddah_flnbfl": mgomdlFlnbfl,
	} {
		if len(mgo) > 0 {
			// rspBlkwrt := fncGlobal.FncGlobalDtbaseBlkwrt(mgo, key)
			// if !rspBlkwrt {
			// 	fmt.Println("ERR LOG HERE, CAN'T INPUT " + key + " LAST")
			// }
		}
	}

}

// Get agent detail and id name
func FncJeddahAgtgetParams(dtlpnr *mdlJeddah.MdlJeddahPnrdtlDtbase, istjed bool,
	idcAgtnme *sync.Map, fnlAgtnme *map[string]mongo.WriteModel) {

	// Get agent name Lcnpun
	keyAgtnme := dtlpnr.Airlfl + dtlpnr.Agtnme
	nowAgtnme := *fnlAgtnme
	tmpAgtnme := mdlJeddah.MdlJeddahAgtnmeDtbase{
		Prmkey: keyAgtnme, Agtnme: dtlpnr.Agtnme, Airlfl: dtlpnr.Airlfl}
	if dtlpnr.Agtidn == "" {
		if !istjed {
			dtlpnr.Agtdtl, tmpAgtnme.Agtdtl = "NON JEDDAH", "NON JEDDAH"
			dtlpnr.Agtidn, tmpAgtnme.Agtidn = "0X", "0X"
			dtlpnr.Rtlsrs, tmpAgtnme.Rtlsrs = "-", "-"
		}
		if _, ist := idcAgtnme.Load(keyAgtnme); !ist {
			if strings.ReplaceAll(dtlpnr.Agtnme, " ", "") != "" {
				mgoUpdate := mongo.NewUpdateOneModel().
					SetFilter(bson.M{"prmkey": keyAgtnme}).
					SetUpdate(bson.M{"$set": tmpAgtnme}).
					SetUpsert(true)
				idcAgtnme.Store(keyAgtnme, true)
				nowAgtnme[dtlpnr.Pnrcde] = mgoUpdate
			}
		}
	}
}

// Get itenary remark and record locator and fligt number
func FncJeddahItrmrlGetapi(rmv string, tkn *mdlSbrapi.MdlSbrapiMsghdrParams,
	fnlDtlpnr *mdlJeddah.MdlJeddahPnrdtlDtbase, objParams mdlSbrapi.MdlSbrapiMsghdrApndix,
	sycFlnbfl, idcPnrcde, idcFlnbfl, sycAgtnme *sync.Map,
	mgomdlSmrpnr, mgomdlFlnbfl *[]mongo.WriteModel,
) (string, bool) {

	// Get agent name Lcnpun
	if fnlDtlpnr.Agtidn == "" {
		keyAgtnme := fnlDtlpnr.Airlfl + fnlDtlpnr.Agtnme
		if rawAgtnme, ist := sycAgtnme.Load(keyAgtnme); ist {
			if getAgtnme, ist := rawAgtnme.(mdlJeddah.MdlJeddahAgtnmeDtbase); ist {
				fnlDtlpnr.Agtdtl = getAgtnme.Agtdtl
				fnlDtlpnr.Agtidn = getAgtnme.Agtidn
				fnlDtlpnr.Rtlsrs = getAgtnme.Rtlsrs
			}
		}
	}

	// Hit API Sabre
	nowTimenw := time.Now().AddDate(0, 0, -0).Format("0601021504")
	intTimenw, _ := strconv.Atoi(nowTimenw)
	intDatenw, _ := strconv.Atoi(nowTimenw[0:6])
	strToflnm, cekDepjed := "", false
	arrRmrkit, err := fncSbrapi.FncSbrapiRsvpnrMainob(*tkn, objParams,
		[]string{"REMARKS", "ITINERARY", "RECORD_LOCATOR"})
	if err == nil {

		// Store to variable
		arrItinry := arrRmrkit.PassengerReservation.Segments.Segment
		varBokdtl := arrRmrkit.BookingDetails
		arrRemark := arrRmrkit.BookingDetails.DivideSplitDetails.Itemslice
		xxxxxx, _ := json.MarshalIndent(arrRmrkit, "", "  ")
		fmt.Println(fnlDtlpnr.Pnrcde, string(xxxxxx))

		// Get remark split data
		if len(arrRemark) > 0 {
			tmpArrspl, tmpTotspl := []string{}, 0
			for _, remark := range arrRemark {
				if remark.XMLName.Local == "SplitFromRecord" ||
					remark.XMLName.Local == "SplitToRecord" {
					splDaterw := remark.DivideTimestamp
					splPnrcde := remark.RecordLocator
					splPrvpax := remark.OriginalNumberOfPax
					splNewpax := remark.CurrentNumberOfPax
					splDatefm, _ := time.Parse("2006-01-02T15:04:05", splDaterw)
					splTimefl := splDatefm.Format("0601021504")
					splTotspl := splPrvpax - splNewpax
					tmpTotspl += splTotspl
					splString := splPnrcde + ":" + splTimefl + ":" + strconv.Itoa(splTotspl)
					tmpArrspl = append(tmpArrspl, splString)
				}
			}

			// Push to final detail PNR Remark split PNR
			fnlDtlpnr.Totspl = tmpTotspl
			fnlDtlpnr.Arrspl = strings.Join(tmpArrspl, "|")
		}

		// Get interline PNR
		arlIntrln := arrRmrkit.POS.Source.TTYRecordLocator.CRSCode
		pnrIntrln := arrRmrkit.POS.Source.TTYRecordLocator.RecordLocator
		if pnrIntrln != "" {
			if !strings.Contains(fnlDtlpnr.Intrln, pnrIntrln) {
				if fnlDtlpnr.Intrln == "" {
					fnlDtlpnr.Intrln = arlIntrln + "*" + pnrIntrln
				} else {
					fnlDtlpnr.Intrln += "|" + arlIntrln + "*" + pnrIntrln
				}
			}
		}

		// Date formating PNR book detail Departure
		pnrTimedp, err := time.Parse("2006-01-02T15:04:05", varBokdtl.FlightsRange.Start)
		intTimedp := 1000000000
		if err == nil {
			rawTimedp, err := strconv.Atoi(pnrTimedp.Format("0601021504"))
			if err == nil {
				intTimedp = rawTimedp
			}
		}

		// Date formating PNR book detail Arrival
		pnrTimerv, err := time.Parse("2006-01-02T15:04:05", varBokdtl.FlightsRange.End)
		intTimerv := 1000000000
		if err == nil {
			rawTimerv, err := strconv.Atoi(pnrTimerv.Format("0601021504"))
			if err == nil {
				intTimerv = rawTimerv
			}
		}

		// Date formating PNR book PNR Create date
		pnrTimecr, err := time.Parse("2006-01-02T15:04:05", varBokdtl.SystemCreationTimestamp)
		intTimerw := 1000000000
		if err == nil {
			rawTimerw, err := strconv.Atoi(pnrTimecr.Format("0601021504"))
			if err == nil {
				intTimerw = rawTimerw
			}
		}

		// Default data smrpnr
		fnlDtlpnr.Timecr = int64(intTimerw)

		// Declare first blank object PNR summary
		objSmrpnr := mdlJeddah.MdlJeddahPnrsmrDtbase{
			Routfl: "", Dateup: int32(intDatenw), Timeup: int64(intTimenw),
			Agtnme: fnlDtlpnr.Agtnme, Agtdtl: fnlDtlpnr.Agtdtl, Agtidn: fnlDtlpnr.Agtidn,
			Pnrcde: fnlDtlpnr.Pnrcde, Intrln: fnlDtlpnr.Intrln, Rtlsrs: fnlDtlpnr.Rtlsrs,
			Arrcpn: "", Totisd: fnlDtlpnr.Totisd, Totbok: fnlDtlpnr.Totbok,
			Totpax: fnlDtlpnr.Totpax, Totcxl: fnlDtlpnr.Totcxl, Totspl: fnlDtlpnr.Totspl,
			Arrspl: fnlDtlpnr.Arrspl, Notedt: "", Timedp: int64(intTimedp),
			Timerv: int64(intTimerv), Timecr: int64(intTimerw), Agtdie: varBokdtl.CreationAgentID,
			Prmkey: fnlDtlpnr.Pnrcde + pnrTimecr.Format("0601021504"),
		}

		// Looping intenary
		if len(arrItinry) > 0 {

			// Loopin itinerary
			arrRoutfl, arrArrcpn, lstArrivl := []string{}, []string{}, ""
			cpnFltype := "Outgoing"
			tmpFlnbfl := map[string]mdlJeddah.MdlJeddahFlnbflDtbase{}
			for _, itinry := range arrItinry {

				// Declare variable from itenary
				cpnDepart := itinry.Air.DepartureAirport
				cpnArrivl := itinry.Air.ArrivalAirport
				cpnRoutfl := cpnDepart + "-" + cpnArrivl
				cpnAirlfl := itinry.Air.OperatingAirlineCode
				cpnActncd := itinry.Air.ActionCode

				// Format date from itenary PNR
				cpnTimefm, _ := time.Parse("2006-01-02T15:04:05", itinry.Air.DepartureDateTime)
				cpnTimefl := cpnTimefm.Format("0601021504")
				cpnTimint, _ := strconv.Atoi(cpnTimefl)
				cpnDatint, _ := strconv.Atoi(cpnTimefl[0:6])

				// Format flight number
				strFlnbfl := itinry.Air.OperatingFlightNumber
				intFlnbfl, err := strconv.Atoi(strFlnbfl)
				cpnFlnbfl := strconv.Itoa(intFlnbfl)
				if err != nil {
					cpnFlnbfl = strFlnbfl
				}

				// Default now array coupon string
				rawArrcpn := []string{cpnAirlfl, cpnFlnbfl, cpnRoutfl, cpnTimefl}
				strArrcpn := strings.Join(rawArrcpn, "-")
				arrArrcpn = append(arrArrcpn, strArrcpn)

				// Push to routfl array
				arrRoutfl = append(arrRoutfl, itinry.Air.DepartureAirport)
				lstArrivl = itinry.Air.ArrivalAirport

				// Get other interline PNR
				if itinry.Air.AirlineRefId != "" {
					if len(itinry.Air.AirlineRefId) < 7 {
						itinry.Air.AirlineRefId = itinry.Air.AirlineRefId + "PURGED"
					}
					pnrIntrln := itinry.Air.AirlineRefId[2:11]
					if !strings.Contains(fnlDtlpnr.Intrln, pnrIntrln[3:9]) {
						if fnlDtlpnr.Intrln == "" {
							fnlDtlpnr.Intrln = pnrIntrln
						} else {
							fnlDtlpnr.Intrln += "|" + pnrIntrln
						}
					}
				}

				// Get previous data flight
				if cpnDepart == "JED" {
					cpnFltype, cekDepjed = "Incoming", true
				}

				// Match new flight
				if cpnRoutfl == fnlDtlpnr.Routfl && cpnActncd == "HK" {
					strToflnm = cpnTimefl[0:6] + ":" + cpnAirlfl + ":" + cpnFlnbfl
				}

				// Push data base flight number jeddah
				cpnPrmkey := cpnTimefl[0:6] + cpnAirlfl + cpnFlnbfl + cpnDepart
				tmpFlnbfl[cpnPrmkey] = mdlJeddah.MdlJeddahFlnbflDtbase{
					Prmkey: cpnPrmkey,
					Datefl: int32(cpnDatint),
					Timefl: int64(cpnTimint),
					Dateup: int32(intDatenw),
					Timeup: int64(intTimenw),
					Airlfl: cpnAirlfl,
					Flnbfl: cpnFlnbfl,
					Depart: cpnDepart,
					Routfl: cpnRoutfl,
					Fltype: cpnFltype,
					Updtby: "System",
				}
			}

			// Push final flight number jeddah
			if len(tmpFlnbfl) > 0 {
				for keyFlnbfl, objFlnbfl := range tmpFlnbfl {
					if cekDepjed {
						objFlnbfl.Isjedh = "Jeddah"
						idcFlnbfl.Store(keyFlnbfl, true)
						mdlFlnbfl := mongo.NewUpdateOneModel().
							SetFilter(bson.M{"prmkey": keyFlnbfl}).
							SetUpdate(bson.M{"$set": objFlnbfl}).
							SetUpsert(true)
						*mgomdlFlnbfl = append(*mgomdlFlnbfl, mdlFlnbfl)
					}
				}
			}

			// Add final object PNR summary
			arrRoutfl = append(arrRoutfl, lstArrivl)
			objSmrpnr.Routfl = strings.Join(arrRoutfl, "-")
			objSmrpnr.Arrcpn = strings.Join(arrArrcpn, "|")
			objSmrpnr.Intrln = fnlDtlpnr.Intrln

		}

		//Push to summary PNR
		if _, ist := idcPnrcde.Load(objSmrpnr.Prmkey); !ist {
			idcPnrcde.Store(objSmrpnr.Prmkey, true)
			if strToflnm == "" && rmv == "rmv" {
				objSmrpnr.Totcxl = objSmrpnr.Totpax
			} else if rmv == "rmv" {
				idcPnrcde.Delete(objSmrpnr.Prmkey)
			}
			mdlSmrpnr := mongo.NewUpdateOneModel().
				SetFilter(bson.M{"prmkey": objSmrpnr.Prmkey}).
				SetUpdate(bson.M{"$set": objSmrpnr}).
				SetUpsert(true)
			*mgomdlSmrpnr = append(*mgomdlSmrpnr, mdlSmrpnr)
		}
	}

	// Return final data
	return strToflnm, cekDepjed
}
