package fncJeddah

import (
	fncGlobal "back/global/function"
	mdlJeddah "back/jeddah/model"
	fncSbrapi "back/sbrapi/function"
	mdlSbrapi "back/sbrapi/model"
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
		totWokrer, nowTotdta := 8, int64(0)
		slcFlnbfl, sycFlnbfl, maxTotdta := FncJeddahFlnbflSycmap(nowTimenw[0:6])
		prvDatefl := FncJeddahActlogLstdta()
		slcDrules := FncJeddahDrulesSycmap()
		sycAgtnme := FncJeddahAgtnmeSycmap()
		sycPnrsmr := FncJeddahPnrsmrSycmap(prvDatefl)
		sycPnrlog := FncJeddahPnrlogSycmap(prvDatefl)
		sycPnrdtl := FncJeddahPnrdtlSycmap(prvDatefl)
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
							sycAgtnme, sycPnrlog, sycPnrdtl, sycFlnbfl, idcAgtnme, idcPnrcde,
							idcFlnbfl, sycPnrsmr, nowTimenw, slcDrules, &nowTotdta, &maxTotdta)
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
		intDatenw, _ := strconv.Atoi(nowTimenw[0:6])
		intTimenw, _ := strconv.Atoi(nowTimenw)
		logAction := mdlJeddah.MdlJeddahActlogDtbase{
			Dateup: int32(intDatenw), Timeup: int64(intTimenw),
			Statdt: "Done"}
		rsupdt := fncGlobal.FncGlobalDtbaseBlkwrt([]mongo.WriteModel{
			mongo.NewUpdateOneModel().
				SetFilter(bson.M{"dateup": logAction.Dateup}).
				SetUpdate(bson.M{"$set": logAction}).
				SetUpsert(true)}, "jeddah_actlog")
		if !rsupdt {
			panic("Error Insert/Update to DB")
		}
	}
}

// Running process jeddah
func FncJeddahPrcessWorker(nbr int, tkn *mdlSbrapi.MdlSbrapiMsghdrParams, swg *sync.WaitGroup,
	jobFlbase <-chan mdlJeddah.MdlJeddahFlnbflDtbase,
	sycAgtnme, sycPnrlog, sycPnrdtl, sycFlnbfl, idcAgtnme, idcPnrcde, idcFlnbfl, sycPnrsmr *sync.Map,
	nowTimenw string, slcDrules []mdlJeddah.MdlJeddahRulesjDtbase, nowTotdta *int64, maxTotdta *float64) {
	var mgomdlAgtnme, mgomdlLcnpun, mgomdlSmrfln []mongo.WriteModel
	var mgomdlSmrpnr, mgomdlPnrdtl, mgomdlFlnbfl []mongo.WriteModel
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
		prvPnrlog := map[string]mdlJeddah.MdlJeddahPnrlogDtbase{}
		if tmpPrvdtl, ist := sycPnrlog.Load(dbsAirlfl + dbsFlnbfl + dbsDepart + strDatenw); ist {
			if getPrvdtl, ist := tmpPrvdtl.(map[string]mdlJeddah.MdlJeddahPnrlogDtbase); ist {
				prvPnrlog = getPrvdtl
				sycPnrlog.Delete(dbsAirlfl + dbsFlnbfl + dbsDepart + strDatenw)
			}
		}

		// Get PNR Detail data
		prvPnrdtl := map[string]mdlJeddah.MdlJeddahPnrdtlDtbase{}
		if tmpPrvdtl, ist := sycPnrdtl.Load(dbsAirlfl + dbsFlnbfl + dbsDepart + strDatenw); ist {
			if getPrvdtl, ist := tmpPrvdtl.(map[string]mdlJeddah.MdlJeddahPnrdtlDtbase); ist {
				prvPnrdtl = getPrvdtl
				sycPnrdtl.Delete(dbsAirlfl + dbsFlnbfl + dbsDepart + strDatenw)
			}
		}

		// Check the date is the same or greater than today
		if dbsDatefl < int32(intTimepv) {
			for _, pnrdtl := range prvPnrdtl {

				// Get agent name Lcnpun
				FncJeddahAgtgetParams(&pnrdtl, true, idcAgtnme, &fnlAgtnme)

				// Push to Summary PNR
				nmodelPnrdtl := mongo.NewUpdateOneModel().
					SetFilter(bson.M{"prmkey": pnrdtl.Prmkey}).
					SetUpdate(bson.M{"$set": pnrdtl}).
					SetUpsert(true)
				mgomdlPnrdtl = append(mgomdlPnrdtl, nmodelPnrdtl)
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
			fnlIsjedh, tmpAllpnr := false, map[string]bool{}
			var tmpSmrpnr, tmpPnrdtl, tmpLcnpun, tmpFlnbfl []mongo.WriteModel
			slcOutllc, errOutllc := fncSbrapi.FncSbrapiLcnpunMainob("LC", *tkn, objParams)
			if errOutllc == nil {
				for _, lcnpun := range slcOutllc {

					// Delcare Pnr log data
					tmpAllpnr[lcnpun.Pnrcde] = true
					objParams.Pnrcde = lcnpun.Pnrcde
					fnlPnrdtl := mdlJeddah.MdlJeddahPnrdtlDtbase{
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
						fnlPnrdtl.Totisd = valIssued
						fnlPnrdtl.Totbok = nowBooked
					}
					intTotisd += valIssued

					// Cek from data Summary PNR and get last remove data
					if prv, ist := prvPnrdtl[lcnpun.Pnrcde]; ist {
						delete(prvPnrdtl, lcnpun.Pnrcde)
						fnlPnrdtl.Dateup = prv.Dateup
						fnlPnrdtl.Datefl = prv.Datefl
						fnlPnrdtl.Timecr = prv.Timecr
						fnlPnrdtl.Totspl = prv.Totspl
						fnlPnrdtl.Totchg = prv.Totchg
						fnlPnrdtl.Totcxl = prv.Totcxl
						fnlPnrdtl.Toflnm = prv.Toflnm
						fnlPnrdtl.Flstat = prv.Flstat
						fnlPnrdtl.Notedt = prv.Notedt
					} else {

						// Get remarks history split or cancel
						_, cekIsjedh := FncJeddahItrmrlGetapi("", tkn, &fnlPnrdtl, objParams, sycFlnbfl, idcPnrcde,
							idcFlnbfl, sycAgtnme, &tmpSmrpnr, &tmpFlnbfl)

						// cek is Jeddah
						nowIsjedh := false
						if cekIsjedh {
							fnlIsjedh, nowIsjedh, flbase.Isjedh = true, true, "Jeddah"
						}

						// Get agent name Lcnpun
						FncJeddahAgtgetParams(&fnlPnrdtl, nowIsjedh, idcAgtnme, &fnlAgtnme)
					}

					// Cek from data Summary PNR and get last remove data
					if prv, ist := prvPnrlog[lcnpun.Pnrcde]; ist {
						delete(prvPnrlog, lcnpun.Pnrcde)
						fnlIsjedh = true
						nowRemove := prv.Totpax - lcnpun.Totpax
						if nowRemove > 0 {

							// Get remarks only
							FncJeddahItrmrlGetapi("", tkn, &fnlPnrdtl, objParams, sycFlnbfl, idcPnrcde,
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
							timesc = strconv.Itoa(int(fnlPnrdtl.Dateup))
						}

						// Convert time
						param1, _ := time.Parse("060102", timefs)
						param2, _ := time.Parse("060102", timesc)

						// Start logic different time
						if math.Abs(param1.Sub(param2).Hours()/24) <= float64(drules.Rldays) {
							fnlPnrdtl.Drules = int(drules.Rlrate)
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
					intTotcxl += fnlPnrdtl.Totcxl
					intTotspl += fnlPnrdtl.Totspl
					nmodelPnrdtl := mongo.NewUpdateOneModel().
						SetFilter(bson.M{"prmkey": fnlPnrdtl.Prmkey}).
						SetUpdate(bson.M{"$set": fnlPnrdtl}).
						SetUpsert(true)
					tmpPnrdtl = append(tmpPnrdtl, nmodelPnrdtl)
				}
			}

			switch {

			// If flight still available
			case len(slcOutllc) != 0:

				// // Looping LX data (TEMPORARY! DELETE AFTER RUNNING)
				// slcOutlxx, errOutlxx := fncSbrapi.FncSbrapiLcnpunMainob("LX", *tkn, objParams)
				// if errOutlxx == nil {
				// 	for _, lcnpun := range slcOutlxx {
				// 		if _, ist := tmpAllpnr[lcnpun.Pnrcde]; !ist {
				// 			if !strings.Contains(lcnpun.Agtnme, "/") {
				// 				objParams.Pnrcde = lcnpun.Pnrcde
				// 				nowPrmkey := lcnpun.Airlfl + lcnpun.Flnbfl + lcnpun.Depart +
				// 					strconv.Itoa(int(lcnpun.Datefl)) + lcnpun.Pnrcde

				// 				// Push to mongo Lcnpun
				// 				nmodelLcnpun := mongo.NewUpdateOneModel().
				// 					SetFilter(bson.M{"prmkey": lcnpun.Prmkey}).
				// 					SetUpdate(bson.M{"$set": lcnpun}).
				// 					SetUpsert(true)
				// 				tmpLcnpun = append(tmpLcnpun, nmodelLcnpun)

				// 				// Default summary PNR
				// 				defPnrdtl := mdlJeddah.MdlJeddahPnrdtlDtbase{
				// 					Prmkey: nowPrmkey, Airlfl: lcnpun.Airlfl, Flnbfl: lcnpun.Flnbfl,
				// 					Depart: lcnpun.Depart, Routfl: lcnpun.Routfl, Clssfl: lcnpun.Clssfl,
				// 					Datefl: lcnpun.Datefl, Dateup: lcnpun.Dateup, Timeup: lcnpun.Timeup,
				// 					Agtnme: lcnpun.Agtnme, Agtdtl: "", Agtidn: "",
				// 					Pnrcde: lcnpun.Pnrcde, Rtlsrs: "", Toflnm: "",
				// 					Drules: 0, Totisd: 0, Totbok: 0, Totpax: lcnpun.Totpax,
				// 				}

				// 				// Get remarks history split or cancel
				// 				strToflnm, cekIsjedh :=
				// 					FncJeddahItrmrlGetapi("rmv", tkn, &defPnrdtl, objParams, sycFlnbfl, idcPnrcde,
				// 						idcFlnbfl, sycAgtnme, &tmpSmrpnr, &tmpFlnbfl)

				// 				// cek is Jeddah
				// 				nowIsjedh := false
				// 				if cekIsjedh {
				// 					fnlIsjedh, nowIsjedh, flbase.Isjedh = true, true, "Jeddah"
				// 				}

				// 				// Get agent name Lcnpun
				// 				FncJeddahAgtgetParams(&defPnrdtl, nowIsjedh, idcAgtnme, &fnlAgtnme)

				// 				// Response Itenary
				// 				if strToflnm != "" {
				// 					defPnrdtl.Toflnm = strToflnm
				// 					defPnrdtl.Flstat = "Change"
				// 					defPnrdtl.Totchg = lcnpun.Totpax
				// 					defPnrdtl.Totpax = 0
				// 					intTotchg += lcnpun.Totpax
				// 				} else {
				// 					defPnrdtl.Flstat = "Cancel"
				// 					defPnrdtl.Totcxl = lcnpun.Totpax
				// 					defPnrdtl.Totpax = lcnpun.Totpax
				// 					intTotcxl += lcnpun.Totpax
				// 				}

				// 				// Push to Pnr log data
				// 				nmodelPnrdtl := mongo.NewUpdateOneModel().
				// 					SetFilter(bson.M{"prmkey": nowPrmkey}).
				// 					SetUpdate(bson.M{"$set": defPnrdtl}).
				// 					SetUpsert(true)
				// 				tmpPnrdtl = append(tmpPnrdtl, nmodelPnrdtl)
				// 			}
				// 		}
				// 	}
				// }

				// Push new data flight to database
				flbase.Flstat = "Operate"
				nmodelFlnbfl := mongo.NewUpdateOneModel().
					SetFilter(bson.M{"prmkey": flbase.Prmkey}).
					SetUpdate(bson.M{"$set": flbase}).
					SetUpsert(true)
				mgomdlFlnbfl = append(mgomdlFlnbfl, nmodelFlnbfl)

				// If prev pnr isset but now null
				if len(prvPnrlog) > 0 {
					for _, lcnpun := range prvPnrlog {
						objParams.Pnrcde = lcnpun.Pnrcde

						// Get prev summary pnr array split and cancel
						if nowPnrdtl, ist := prvPnrdtl[lcnpun.Pnrcde]; ist {
							delete(prvPnrdtl, lcnpun.Pnrcde)

							// Get remarks history split or cancel
							strToflnm, _ :=
								FncJeddahItrmrlGetapi("rmv", tkn, &nowPnrdtl, objParams, sycFlnbfl, idcPnrcde,
									idcFlnbfl, sycAgtnme, &mgomdlSmrpnr, &mgomdlFlnbfl)

							// Response Itenary
							if strToflnm != "" {
								nowPnrdtl.Toflnm = strToflnm
								nowPnrdtl.Flstat = "Change"
								nowPnrdtl.Totchg = nowPnrdtl.Totpax
								nowPnrdtl.Totpax = 0
								intTotchg += nowPnrdtl.Totpax
							} else {
								nowPnrdtl.Flstat = "Cancel"
								nowPnrdtl.Totcxl = nowPnrdtl.Totpax
								intTotcxl += nowPnrdtl.Totpax
							}

							// Push to Pnr log data
							nmodelPnrdtl := mongo.NewUpdateOneModel().
								SetFilter(bson.M{"prmkey": nowPnrdtl.Prmkey}).
								SetUpdate(bson.M{"$set": nowPnrdtl}).
								SetUpsert(true)
							mgomdlPnrdtl = append(mgomdlPnrdtl, nmodelPnrdtl)
						}
					}
				}

				// Continue if not exist Jeddah route
				if !fnlIsjedh {
					continue outlop
				} else {
					mgomdlSmrpnr = append(mgomdlSmrpnr, tmpSmrpnr...)
					mgomdlLcnpun = append(mgomdlLcnpun, tmpLcnpun...)
					mgomdlPnrdtl = append(mgomdlPnrdtl, tmpPnrdtl...)
					mgomdlFlnbfl = append(mgomdlFlnbfl, tmpFlnbfl...)
				}

			// If flight cancel/change but never in db
			case len(prvPnrdtl) == 0:

				// Looping LDN data
				fnlIsjedh := false
				flbase.Isjedh = "Non Jeddah"
				var tmpSmrpnr, tmpPnrdtl, tmpLcnpun, tmpFlnbfl []mongo.WriteModel
				slcOutldn, errOutldn := fncSbrapi.FncSbrapiLcnpunMainob("LDN", *tkn, objParams)
				if errOutldn == nil {
					for _, lcnpun := range slcOutldn {
						objParams.Pnrcde = lcnpun.Pnrcde
						nowPrmkey := lcnpun.Airlfl + lcnpun.Flnbfl + lcnpun.Depart +
							strconv.Itoa(int(lcnpun.Datefl)) + lcnpun.Pnrcde

						// Push to mongo Lcnpun
						nmodelLcnpun := mongo.NewUpdateOneModel().
							SetFilter(bson.M{"prmkey": lcnpun.Prmkey}).
							SetUpdate(bson.M{"$set": lcnpun}).
							SetUpsert(true)
						tmpLcnpun = append(tmpLcnpun, nmodelLcnpun)

						// Default summary PNR
						defPnrdtl := mdlJeddah.MdlJeddahPnrdtlDtbase{
							Prmkey: nowPrmkey, Airlfl: lcnpun.Airlfl, Flnbfl: lcnpun.Flnbfl,
							Depart: lcnpun.Depart, Routfl: lcnpun.Routfl, Clssfl: lcnpun.Clssfl,
							Datefl: lcnpun.Datefl, Dateup: lcnpun.Dateup, Timeup: lcnpun.Timeup,
							Agtnme: lcnpun.Agtnme, Agtdtl: "", Agtidn: "",
							Pnrcde: lcnpun.Pnrcde, Rtlsrs: "", Toflnm: "",
							Drules: 0, Totisd: 0, Totbok: 0, Totpax: lcnpun.Totpax,
						}

						// Get remarks history split or cancel
						strToflnm, cekIsjedh :=
							FncJeddahItrmrlGetapi("rmv", tkn, &defPnrdtl, objParams, sycFlnbfl, idcPnrcde,
								idcFlnbfl, sycAgtnme, &tmpSmrpnr, &tmpFlnbfl)

						// cek is Jeddah
						nowIsjedh := false
						if cekIsjedh {
							fnlIsjedh, nowIsjedh, flbase.Isjedh = true, true, "Jeddah"
						}

						// Get agent name Lcnpun
						FncJeddahAgtgetParams(&defPnrdtl, nowIsjedh, idcAgtnme, &fnlAgtnme)

						// Response Itenary
						if strToflnm != "" {
							defPnrdtl.Toflnm = strToflnm
							defPnrdtl.Flstat = "Change"
							defPnrdtl.Totchg = lcnpun.Totpax
							defPnrdtl.Totpax = 0
							intTotchg += lcnpun.Totpax
						} else {
							defPnrdtl.Flstat = "Cancel"
							defPnrdtl.Totcxl = lcnpun.Totpax
							defPnrdtl.Totpax = lcnpun.Totpax
							intTotcxl += lcnpun.Totpax
						}

						// Push to Pnr log data
						nmodelPnrdtl := mongo.NewUpdateOneModel().
							SetFilter(bson.M{"prmkey": nowPrmkey}).
							SetUpdate(bson.M{"$set": defPnrdtl}).
							SetUpsert(true)
						tmpPnrdtl = append(tmpPnrdtl, nmodelPnrdtl)
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
					mgomdlPnrdtl = append(mgomdlPnrdtl, tmpPnrdtl...)
					mgomdlFlnbfl = append(mgomdlFlnbfl, tmpFlnbfl...)
				}

			// If flight cancel/change avail in db
			case len(prvPnrdtl) > 0:

				// Get agent name Lcnpun
				for _, pnrdtl := range prvPnrdtl {
					objParams.Pnrcde = pnrdtl.Pnrcde
					if pnrdtl.Agtidn == "" {
						FncJeddahAgtgetParams(&pnrdtl, true, idcAgtnme, &fnlAgtnme)

						// Push to Pnr log data
						nmodelPnrdtl := mongo.NewUpdateOneModel().
							SetFilter(bson.M{"prmkey": pnrdtl.Prmkey}).
							SetUpdate(bson.M{"$set": pnrdtl}).
							SetUpsert(true)
						mgomdlPnrdtl = append(mgomdlPnrdtl, nmodelPnrdtl)

						// Get remarks history split or cancel
						FncJeddahItrmrlGetapi("rmv", tkn, &pnrdtl, objParams, sycFlnbfl, idcPnrcde,
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
			rspBlkwrt := fncGlobal.FncGlobalDtbaseBlkwrt(mgomdlLcnpun, "jeddah_pnrlog")
			if !rspBlkwrt {
				fmt.Println("ERR LOG HERE, CAN'T INPUT LCNPUN")
			}
			mgomdlLcnpun = []mongo.WriteModel{}
		}

		// Push mongo pnrlog
		if len(mgomdlPnrdtl) > lmtdta && len(mgomdlPnrdtl) != 0 {
			rspBlkwrt := fncGlobal.FncGlobalDtbaseBlkwrt(mgomdlPnrdtl, "jeddah_pnrdtl")
			if !rspBlkwrt {
				fmt.Println("ERR LOG HERE, CAN'T INPUT Pnrdtl")
			}
			mgomdlPnrdtl = []mongo.WriteModel{}
		}

		// Push mongo pnrlog
		if len(mgomdlFlnbfl) > lmtdta && len(mgomdlFlnbfl) != 0 {
			rspBlkwrt := fncGlobal.FncGlobalDtbaseBlkwrt(mgomdlFlnbfl, "jeddah_flnbfl")
			if !rspBlkwrt {
				fmt.Println("ERR LOG HERE, CAN'T INPUT FLNbfl")
			}
			mgomdlFlnbfl = []mongo.WriteModel{}
		}

		// Push mongo pnrlog
		if len(mgomdlSmrpnr) > lmtdta && len(mgomdlSmrpnr) != 0 {
			rspBlkwrt := fncGlobal.FncGlobalDtbaseBlkwrt(mgomdlSmrpnr, "jeddah_pnrsmr")
			if !rspBlkwrt {
				fmt.Println("ERR LOG HERE, CAN'T INPUT SMRPNR")
			}
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
	for key, mgo := range map[string][]mongo.WriteModel{
		"jeddah_pnrlog": mgomdlLcnpun,
		"jeddah_agentx": mgomdlAgtnme,
		"jeddah_pnrsmr": mgomdlSmrpnr,
		"jeddah_pnrdtl": mgomdlPnrdtl,
		"jeddah_flnsmr": mgomdlSmrfln,
		"jeddah_flnbfl": mgomdlFlnbfl,
	} {
		if len(mgo) > 0 {
			rspBlkwrt := fncGlobal.FncGlobalDtbaseBlkwrt(mgo, key)
			if !rspBlkwrt {
				fmt.Println("ERR LOG HERE, CAN'T INPUT " + key + " LAST")
			}
		}
	}

}

// Get agent detail and id name
func FncJeddahAgtgetParams(pnrdtl *mdlJeddah.MdlJeddahPnrdtlDtbase, istjed bool,
	idcAgtnme *sync.Map, fnlAgtnme *map[string]mongo.WriteModel) {

	// Get agent name Lcnpun
	keyAgtnme := pnrdtl.Airlfl + pnrdtl.Agtnme
	nowAgtnme := *fnlAgtnme
	tmpAgtnme := mdlJeddah.MdlJeddahAgtnmeDtbase{
		Prmkey: keyAgtnme, Agtnme: pnrdtl.Agtnme, Airlfl: pnrdtl.Airlfl}
	if pnrdtl.Agtidn == "" {
		if !istjed {
			pnrdtl.Agtdtl, tmpAgtnme.Agtdtl = "NON JEDDAH", "NON JEDDAH"
			pnrdtl.Agtidn, tmpAgtnme.Agtidn = "0X", "0X"
			pnrdtl.Rtlsrs, tmpAgtnme.Rtlsrs = "-", "-"
		}
		if _, ist := idcAgtnme.Load(keyAgtnme); !ist {
			if strings.ReplaceAll(pnrdtl.Agtnme, " ", "") != "" {
				mgoUpdate := mongo.NewUpdateOneModel().
					SetFilter(bson.M{"prmkey": keyAgtnme}).
					SetUpdate(bson.M{"$set": tmpAgtnme}).
					SetUpsert(true)
				idcAgtnme.Store(keyAgtnme, true)
				nowAgtnme[pnrdtl.Pnrcde] = mgoUpdate
			}
		}
	}
}

// Get itenary remark and record locator and fligt number
func FncJeddahItrmrlGetapi(rmv string, tkn *mdlSbrapi.MdlSbrapiMsghdrParams,
	fnlPnrdtl *mdlJeddah.MdlJeddahPnrdtlDtbase, objParams mdlSbrapi.MdlSbrapiMsghdrApndix,
	sycFlnbfl, idcPnrcde, idcFlnbfl, sycAgtnme *sync.Map,
	mgomdlSmrpnr, mgomdlFlnbfl *[]mongo.WriteModel,
) (string, bool) {

	// Get agent name Lcnpun
	if fnlPnrdtl.Agtidn == "" {
		keyAgtnme := fnlPnrdtl.Airlfl + fnlPnrdtl.Agtnme
		if rawAgtnme, ist := sycAgtnme.Load(keyAgtnme); ist {
			if getAgtnme, ist := rawAgtnme.(mdlJeddah.MdlJeddahAgtnmeDtbase); ist {
				fnlPnrdtl.Agtdtl = getAgtnme.Agtdtl
				fnlPnrdtl.Agtidn = getAgtnme.Agtidn
				fnlPnrdtl.Rtlsrs = getAgtnme.Rtlsrs
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
			fnlPnrdtl.Totspl = tmpTotspl
			fnlPnrdtl.Arrspl = strings.Join(tmpArrspl, "|")
		}

		// Get interline PNR
		arlIntrln := arrRmrkit.POS.Source.TTYRecordLocator.CRSCode
		pnrIntrln := arrRmrkit.POS.Source.TTYRecordLocator.RecordLocator
		if pnrIntrln != "" {
			if !strings.Contains(fnlPnrdtl.Intrln, pnrIntrln) {
				if fnlPnrdtl.Intrln == "" {
					fnlPnrdtl.Intrln = arlIntrln + "*" + pnrIntrln
				} else {
					fnlPnrdtl.Intrln += "|" + arlIntrln + "*" + pnrIntrln
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
		fnlPnrdtl.Timecr = int64(intTimerw)

		// Declare first blank object PNR summary
		objSmrpnr := mdlJeddah.MdlJeddahPnrsmrDtbase{
			Routfl: "", Dateup: int32(intDatenw), Timeup: int64(intTimenw),
			Agtnme: fnlPnrdtl.Agtnme, Agtdtl: fnlPnrdtl.Agtdtl, Agtidn: fnlPnrdtl.Agtidn,
			Pnrcde: fnlPnrdtl.Pnrcde, Intrln: fnlPnrdtl.Intrln, Rtlsrs: fnlPnrdtl.Rtlsrs,
			Arrcpn: "", Totisd: fnlPnrdtl.Totisd, Totbok: fnlPnrdtl.Totbok,
			Totpax: fnlPnrdtl.Totpax, Totcxl: fnlPnrdtl.Totcxl, Totspl: fnlPnrdtl.Totspl,
			Arrspl: fnlPnrdtl.Arrspl, Notedt: "", Timedp: int64(intTimedp),
			Timerv: int64(intTimerv), Timecr: int64(intTimerw), Agtdie: varBokdtl.CreationAgentID,
			Prmkey: fnlPnrdtl.Pnrcde + pnrTimecr.Format("0601021504"),
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
					if !strings.Contains(fnlPnrdtl.Intrln, pnrIntrln[3:9]) {
						if fnlPnrdtl.Intrln == "" {
							fnlPnrdtl.Intrln = pnrIntrln
						} else {
							fnlPnrdtl.Intrln += "|" + pnrIntrln
						}
					}
				}

				// Get previous data flight
				if cpnDepart == "JED" {
					cpnFltype, cekDepjed = "Incoming", true
				}

				// Match new flight
				if cpnRoutfl == fnlPnrdtl.Routfl && cpnActncd == "HK" {
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
			objSmrpnr.Intrln = fnlPnrdtl.Intrln

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
