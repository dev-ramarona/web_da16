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

	// protect single run
	if fncGlobal.Status.Sbrapi != 0.0 {
		return
	}
	fncGlobal.Status.Sbrapi = 0.01

	// Input data post
	var inpActlog mdlJeddah.MdlJeddahActlogDtbase
	if err := c.BindJSON(&inpActlog); err != nil {
		panic(err)
	}

	// Initialize time now
	strTimenw := time.Now().AddDate(0, 0, -0).Format("0601021504")
	strDatepv := FncJeddahActlogLstdta()
	intDatenw, _ := strconv.Atoi(strTimenw[0:6])
	intTimenw, _ := strconv.Atoi(strTimenw)

	// Initialize Map and Params
	var totWokrer = 8
	var nowTotdta = int64(0)
	var slcFlnbfl, sycFlnbfl, maxTotdta = FncJeddahFlnbflSycmap()
	var slcDrules = FncJeddahDrulesSlcobj()
	var sycAgtnme = FncJeddahAgtnmeSycmap()
	var sycPnrsmr = FncJeddahPnrsmrSycmap(strDatepv)
	var sycPnrlog = FncJeddahPnrlogSycmap(strDatepv)
	var sycPnrdtl = FncJeddahPnrdtlSycmap(strDatepv)
	var idcAgtnme, idcAgtupd, idcPnrsmr, idcFlnbfl *sync.Map
	var airAction = []string{}
	var logAction = mdlJeddah.MdlJeddahActlogDtbase{
		Dateup: int32(intDatenw),
		Timeup: int64(intTimenw),
		Statdt: "Done"}

	// Looping all flight number per airline
	for airlfl, slice1 := range slcFlnbfl {

		// Handle only airline input
		if inpActlog.Airlfl != "" &&
			!strings.Contains(inpActlog.Airlfl, airlfl) {
			continue
		}

		// Get Multiple API sessions/tokens
		slcRspssn, err := fncSbrapi.FncSbrapiCrtssnMultpl(airlfl, totWokrer)
		lgcRspssn := err != nil || slcRspssn[0].Bsttkn == "" || len(slcRspssn) < 1
		if lgcRspssn {
			airAction = append(airAction, airlfl)
			continue
		}

		// Prepare job queue
		jobFlbase := make(chan mdlJeddah.MdlJeddahFlnbflDtbase, 1000)
		var swg sync.WaitGroup

		// Launch 10 workers using 10 tokens
		for i := 0; i < totWokrer; i++ {
			if len(slcRspssn) >= i+1 {
				if slcRspssn[i].Bsttkn != "Failed" {
					swg.Add(1)
					go FncJeddahPrcessWorker(
						&slcRspssn[i],
						&swg,
						jobFlbase,
						sycAgtnme, sycPnrlog, sycPnrdtl, sycFlnbfl, sycPnrsmr,
						idcAgtnme, idcAgtupd, idcPnrsmr, idcFlnbfl,
						strTimenw, strDatepv,
						slcDrules,
						&nowTotdta,
						&maxTotdta)
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
		fmt.Printf("Done airline:%s time:%s \n", airlfl,
			time.Now().Format("06-Jan-02/15:04"))
		fncSbrapi.FncSbrapiClsssnMultpl(slcRspssn)
	}

	// Modify action if error
	if len(airAction) > 0 {
		logAction.Airlfl = strings.Join(airAction, "|")
		logAction.Statdt = "Pending"
	}

	// Final put log action
	rsupdt := fncGlobal.FncGlobalDtbaseBlkwrt(
		[]mongo.WriteModel{mongo.NewUpdateOneModel().
			SetFilter(bson.M{"dateup": logAction.Dateup}).
			SetUpdate(bson.M{"$set": logAction}).
			SetUpsert(true)}, "jeddah_actlog")
	fncGlobal.Status.Sbrapi = 0
	if rsupdt != nil {
		panic("Error Insert/Update to DB:" + rsupdt.Error())
	}
}

// Running process jeddah
func FncJeddahPrcessWorker(
	tkn *mdlSbrapi.MdlSbrapiMsghdrParams,
	swg *sync.WaitGroup,
	jobFlbase <-chan mdlJeddah.MdlJeddahFlnbflDtbase,
	sycAgtnme, sycPnrlog, sycPnrdtl, sycFlnbfl, sycPnrsmr,
	idcAgtnme, idcAgtupd, idcPnrsmr, idcFlnbfl *sync.Map,
	strTimenw, strDatepv string,
	slcDrules []mdlJeddah.MdlJeddahRulesjDtbase,
	nowTotdta *int64,
	maxTotdta *float64) {

	// Declare global variable
	defer swg.Done()
	const blimit = 50
	var mgoAgtnme, mgoLcnpun, mgoSmrfln []mongo.WriteModel
	var mgoPnrsmr, mgoPnrdtl, mgoFlnbfl []mongo.WriteModel

	// Helper to get agent
	getAgtnme := func(airlfl, agtnme string) mdlJeddah.MdlJeddahAgtnmeDtbase {
		keyAgtnme := airlfl + agtnme
		if rawAgtnme, ist := sycAgtnme.Load(keyAgtnme); ist {
			if getAgtnme, mtc := rawAgtnme.(mdlJeddah.MdlJeddahAgtnmeDtbase); mtc {
				return getAgtnme
			}
		}
		return mdlJeddah.MdlJeddahAgtnmeDtbase{}
	}

	// Helper to update agent (temporary)
	prvAgtnme := func(pnrdtl *mdlJeddah.MdlJeddahPnrdtlDtbase,
		agtnme mdlJeddah.MdlJeddahAgtnmeDtbase) {

		// Update agentname pnr detail
		if pnrdtl.Agtidn == "" {
			pnrdtl.Agtdtl = agtnme.Agtdtl
			pnrdtl.Agtidn = agtnme.Agtidn
			pnrdtl.Rtlsrs = agtnme.Rtlsrs
			if pnrdtl.Isjedh == "" {
				pnrdtl.Isjedh = "Jeddah"
				if pnrdtl.Agtidn == "0X" {
					pnrdtl.Isjedh = "Non Jeddah"
				}
			}
			mgoPnrdtl = append(mgoPnrdtl, mongo.NewUpdateOneModel().
				SetFilter(bson.M{"prmkey": pnrdtl.Prmkey}).
				SetUpdate(bson.M{"$set": pnrdtl}).
				SetUpsert(true))
		}

		// Update agentname pnr summary
		strTimecr := strconv.Itoa(int(pnrdtl.Timecr))
		if _, ist := idcAgtupd.Load(pnrdtl.Pnrcde + strTimecr); !ist {
			if nowPnrsmr, ist := sycPnrsmr.Load(pnrdtl.Pnrcde + strTimecr); ist {
				if pnrsmr, mtc := nowPnrsmr.(mdlJeddah.MdlJeddahPnrsmrDtbase); mtc {
					idcAgtupd.Store(pnrdtl.Pnrcde+strTimecr, true)
					if pnrsmr.Agtidn == "" {
						pnrsmr.Agtdtl = agtnme.Agtdtl
						pnrsmr.Agtidn = agtnme.Agtidn
						pnrsmr.Rtlsrs = agtnme.Rtlsrs
						if pnrsmr.Isjedh == "" {
							pnrsmr.Isjedh = "Jeddah"
							if pnrsmr.Agtidn == "0X" {
								pnrsmr.Isjedh = "Non Jeddah"
							}
						}
						mgoPnrsmr = append(mgoPnrsmr, mongo.NewUpdateOneModel().
							SetFilter(bson.M{"prmkey": pnrsmr.Prmkey}).
							SetUpdate(bson.M{"$set": pnrsmr}).
							SetUpsert(true))
					}
				}
			}
		}
	}

	// Helper to get reservation
	getRsrvtn := func(intTimenw, intDatenw int,
		objParams mdlSbrapi.MdlSbrapiMsghdrApndix,
		fnlPnrdtl *mdlJeddah.MdlJeddahPnrdtlDtbase,
		mgoFlnbfl *[]mongo.WriteModel) (string, string, mdlJeddah.MdlJeddahPnrsmrDtbase) {

		// Declare variable and hit API Sabre
		var strToflnm, nowIsjedh = "", "Non Jeddah"
		var fnlPnrsmr mdlJeddah.MdlJeddahPnrsmrDtbase
		if arrRmrkit, err := fncSbrapi.FncSbrapiRsvpnrMainob(*tkn, objParams.Pnrcde,
			[]string{"REMARKS", "ITINERARY", "RECORD_LOCATOR"}); err == nil {

			// Store to variable
			arrItinry := arrRmrkit.PassengerReservation.Segments.Segment
			arrRemark := arrRmrkit.BookingDetails.DivideSplitDetails.Itemslice
			varBokdtl := arrRmrkit.BookingDetails
			varIntrln := arrRmrkit.POS.Source.TTYRecordLocator

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
			arlIntrln := varIntrln.CRSCode
			pnrIntrln := varIntrln.RecordLocator
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
			intTimest := 1000000000
			varTimest := varBokdtl.FlightsRange.Start
			if pnrTimest, err := time.Parse("2006-01-02T15:04:05", varTimest); err == nil {
				rawTimest, _ := strconv.Atoi(pnrTimest.Format("0601021504"))
				intTimest = rawTimest
			}

			// Date formating PNR book detail Arrival
			intTimend := 1000000000
			varTimend := varBokdtl.FlightsRange.End
			if pnrTimend, err := time.Parse("2006-01-02T15:04:05", varTimend); err == nil {
				rawTimend, _ := strconv.Atoi(pnrTimend.Format("0601021504"))
				intTimend = rawTimend
			}

			// Date formating PNR book PNR Create date
			intTimecr := 1000000000
			varTimecr := varBokdtl.SystemCreationTimestamp
			if pnrTimecr, err := time.Parse("2006-01-02T15:04:05", varTimecr); err == nil {
				rawTimerw, _ := strconv.Atoi(pnrTimecr.Format("0601021504"))
				intTimecr = rawTimerw
			}
			fnlPnrdtl.Timecr = int64(intTimecr)
			strTimecr := strconv.Itoa(int(intTimecr))

			// Declare first blank object PNR summary
			fnlPnrsmr = mdlJeddah.MdlJeddahPnrsmrDtbase{
				Prmkey: fnlPnrdtl.Pnrcde + strTimecr,
				Routfl: "",
				Timest: int64(intTimest),
				Timend: int64(intTimend),
				Dateup: int32(intDatenw),
				Timeup: int64(intTimenw),
				Timecr: int64(intTimecr),
				Agtnme: fnlPnrdtl.Agtnme,
				Agtdtl: fnlPnrdtl.Agtdtl,
				Agtidn: fnlPnrdtl.Agtidn,
				Pnrcde: fnlPnrdtl.Pnrcde,
				Intrln: fnlPnrdtl.Intrln,
				Rtlsrs: fnlPnrdtl.Rtlsrs,
				Arrcpn: "",
				Agtdcr: varBokdtl.CreationAgentID,
				Totisd: fnlPnrdtl.Totisd,
				Totbok: fnlPnrdtl.Totbok,
				Totpax: fnlPnrdtl.Totpax,
				Totcxl: fnlPnrdtl.Totcxl,
				Totspl: fnlPnrdtl.Totspl,
				Arrspl: fnlPnrdtl.Arrspl,
				Notedt: "",
			}

			if len(arrItinry) == 0 {
				nowIsjedh = "No Itinerary"
				return strToflnm, nowIsjedh, fnlPnrsmr
			}

			// Looping intenary
			var lstArrivl string
			var rsvFltype = "Outgoing"
			var tmpFlnbfl = map[string]mdlJeddah.MdlJeddahFlnbflDtbase{}
			var arrRoutfl, arrArrcpn []string
			for _, itinry := range arrItinry {

				// Declare variable from itenary
				rawDepart := itinry.Air.DepartureAirport
				rawArrivl := itinry.Air.ArrivalAirport
				strRoutfl := rawDepart + "-" + rawArrivl
				rawAirlfl := itinry.Air.OperatingAirlineCode
				rawActncd := itinry.Air.ActionCode
				rawClssfl := itinry.Air.OperatingClassOfService

				// Format date from itenary PNR
				rawTimefm, _ := time.Parse("2006-01-02T15:04:05", itinry.Air.DepartureDateTime)
				strTimefl := rawTimefm.Format("0601021504")
				strDatefl := strTimefl[0:6]
				intTimefl, _ := strconv.Atoi(strTimefl)
				intDatefl, _ := strconv.Atoi(strDatefl)

				// Format flight number
				rawFlnbfl := itinry.Air.OperatingFlightNumber
				intFlnbfl, err := strconv.Atoi(rawFlnbfl)
				strFlnbfl := strconv.Itoa(intFlnbfl)
				if err != nil {
					strFlnbfl = rawFlnbfl
				}

				// Push to routfl array
				arrRoutfl = append(arrRoutfl, itinry.Air.DepartureAirport)
				lstArrivl = itinry.Air.ArrivalAirport

				// Default now array coupon string
				slcArrcpn := []string{rawAirlfl, strFlnbfl, strRoutfl, strTimefl}
				strArrcpn := strings.Join(slcArrcpn, "-")
				arrArrcpn = append(arrArrcpn, strArrcpn)

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

				// Get information type and jeddah
				if rawDepart == "JED" {
					rsvFltype, nowIsjedh = "Incoming", "Jeddah"
				}

				// Match new flight from change flight
				if strRoutfl == fnlPnrdtl.Routfl && rawActncd == "HK" {
					slcToflnm := []string{strDatefl, rawAirlfl, strFlnbfl, rawClssfl}
					strToflnm = strings.Join(slcToflnm, ":")
				}

				// Push data base flight number jeddah
				cpnPrmkey := strDatefl + rawAirlfl + strFlnbfl + rawDepart
				tmpFlnbfl[cpnPrmkey] = mdlJeddah.MdlJeddahFlnbflDtbase{
					Prmkey: cpnPrmkey,
					Datefl: int32(intDatefl),
					Timefl: int64(intTimefl),
					Dateup: int32(intDatenw),
					Timeup: int64(intTimenw),
					Airlfl: rawAirlfl,
					Flnbfl: strFlnbfl,
					Depart: rawDepart,
					Routfl: strRoutfl,
					Fltype: rsvFltype,
					Updtby: "System",
				}
			}

			// Push final flight number jeddah
			if len(tmpFlnbfl) > 0 {
				for keyFlnbfl, objFlnbfl := range tmpFlnbfl {
					if nowIsjedh == "Jeddah" {
						objFlnbfl.Isjedh = "Jeddah"
						idcFlnbfl.Store(keyFlnbfl, true)
						mdlFlnbfl := mongo.NewUpdateOneModel().
							SetFilter(bson.M{"prmkey": keyFlnbfl}).
							SetUpdate(bson.M{"$set": objFlnbfl}).
							SetUpsert(true)
						*mgoFlnbfl = append(*mgoFlnbfl, mdlFlnbfl)
					}
				}
			}

			// Add final object PNR summary
			arrRoutfl = append(arrRoutfl, lstArrivl)
			fnlPnrsmr.Routfl = strings.Join(arrRoutfl, "-")
			fnlPnrsmr.Arrcpn = strings.Join(arrArrcpn, "|")
			fnlPnrsmr.Intrln = fnlPnrdtl.Intrln
		}

		// Final return
		return strToflnm, nowIsjedh, fnlPnrsmr
	}

	// Helper to get rules
	getRules := func(fnlPnrdtl *mdlJeddah.MdlJeddahPnrdtlDtbase,
		strDatefl string) {
		for _, drules := range slcDrules {

			// Get parameter date
			datest := strDatefl
			datend := strTimenw[0:6]
			if drules.Rlcolm == "dateup" {
				datest = strTimenw[0:6]
				datend = strconv.Itoa(int(fnlPnrdtl.Dateup))
			}

			// Convert time
			param1, _ := time.Parse("060102", datest)
			param2, _ := time.Parse("060102", datend)

			// Start logic different time
			if math.Abs(param1.Sub(param2).Hours()/24) <= float64(drules.Rldays) {
				fnlPnrdtl.Drules = int(drules.Rlrate)
				break
			}
		}
	}

	// Helper Push summary and detail PNR to temporary
	updPnrall := func(fnlPnrdtl *mdlJeddah.MdlJeddahPnrdtlDtbase,
		fnlPnrsmr *mdlJeddah.MdlJeddahPnrsmrDtbase,
		mgoPnrdtl, mgoPnrsmr *[]mongo.WriteModel) {
		*mgoPnrdtl = append(*mgoPnrdtl, mongo.NewUpdateOneModel().
			SetFilter(bson.M{"prmkey": fnlPnrdtl.Prmkey}).
			SetUpdate(bson.M{"$set": fnlPnrdtl}).
			SetUpsert(true))
		if _, ist := idcPnrsmr.Load(fnlPnrsmr.Prmkey); !ist {
			idcPnrsmr.Store(fnlPnrsmr.Prmkey, true)
			*mgoPnrsmr = append(*mgoPnrsmr, mongo.NewUpdateOneModel().
				SetFilter(bson.M{"prmkey": fnlPnrsmr.Prmkey}).
				SetUpdate(bson.M{"$set": fnlPnrsmr}).
				SetUpsert(true))
		}
	}

	// Helper to push agent
	updAgtnme := func(nowIsjedh string,
		nowAgtnme mdlJeddah.MdlJeddahAgtnmeDtbase,
		mapAgtnme *map[string]mongo.WriteModel,
		pnrdtl *mdlJeddah.MdlJeddahPnrdtlDtbase,
		Pnrsmr *mdlJeddah.MdlJeddahPnrsmrDtbase) {
		pnrdtl.Isjedh = nowIsjedh
		Pnrsmr.Isjedh = nowIsjedh
		if pnrdtl.Agtidn == "" {
			keyAgtnme := pnrdtl.Airlfl + pnrdtl.Agtnme
			tmpAgtnme := mdlJeddah.MdlJeddahAgtnmeDtbase{
				Prmkey: keyAgtnme, Airlfl: pnrdtl.Airlfl, Agtnme: pnrdtl.Agtnme}

			// If non jeddah
			if nowIsjedh == "Non Jeddah" {
				tmpAgtnme.Agtdtl = "NON JEDDAH"
				pnrdtl.Agtdtl, Pnrsmr.Agtdtl = "NON JEDDAH", "NON JEDDAH"
				tmpAgtnme.Agtidn, pnrdtl.Agtidn, Pnrsmr.Agtidn = "0X", "0X", "0X"
				tmpAgtnme.Rtlsrs, pnrdtl.Rtlsrs, Pnrsmr.Rtlsrs = "-", "-", "-"
				tmpAgtnme.Updtby = "System"
			} else if nowAgtnme.Agtidn != "" {
				pnrdtl.Agtdtl, Pnrsmr.Agtdtl = nowAgtnme.Agtdtl, nowAgtnme.Agtdtl
				pnrdtl.Agtidn, Pnrsmr.Agtidn = nowAgtnme.Agtidn, nowAgtnme.Agtidn
				pnrdtl.Rtlsrs, Pnrsmr.Rtlsrs = nowAgtnme.Rtlsrs, nowAgtnme.Rtlsrs
				return
			}

			// Store to map agent if not exist
			if nowIsjedh == "Jeddah" {
				if _, ist := idcAgtnme.Load(keyAgtnme); !ist {
					if strings.ReplaceAll(pnrdtl.Agtnme, " ", "") != "" {
						mgoUpdate := mongo.NewUpdateOneModel().
							SetFilter(bson.M{"prmkey": keyAgtnme}).
							SetUpdate(bson.M{"$set": tmpAgtnme}).
							SetUpsert(true)
						idcAgtnme.Store(keyAgtnme, true)
						(*mapAgtnme)[pnrdtl.Pnrcde] = mgoUpdate
					}
				}
			}
		}
	}

	// iterate jobs
	cntdta := 0
	for flbase := range jobFlbase {
		cntdta++
		atomic.AddInt64(nowTotdta, 1)

		// update indicator progress
		if *maxTotdta > 0 {
			fncGlobal.Status.Sbrapi = float64(*nowTotdta) / *maxTotdta * 100
		}

		// prepare locals
		var mapAgtnme = map[string]mongo.WriteModel{}
		var intTotisd, intTotbok, intTotpax, intTotcxl, intTotchg, intTotspl int
		var intDatefl = flbase.Datefl
		var dbsFlnbfl, dbsDepart = flbase.Flnbfl, flbase.Depart
		var dbsRoutfl, dbsAirlfl = flbase.Routfl, flbase.Airlfl
		var objParams = mdlSbrapi.MdlSbrapiMsghdrApndix{
			Airlfl: dbsAirlfl, Datefl: intDatefl,
			Depart: dbsDepart, Flnbfl: dbsFlnbfl, Routfl: dbsRoutfl}

		// Indicator start process
		fmt.Println("Start", cntdta, "-", dbsAirlfl, dbsFlnbfl, intDatefl,
			dbsRoutfl, time.Now().Format("06-01-02/15:04"))

		// Conver String and int date
		strDatefl := strconv.Itoa(int(intDatefl))
		intDatepv, _ := strconv.Atoi(strDatepv)
		intDatenw, _ := strconv.Atoi(strTimenw[0:6])
		intTimenw, _ := strconv.Atoi(strTimenw)

		// load cached previous data pnrdtl
		prvPnrdtl := map[string]mdlJeddah.MdlJeddahPnrdtlDtbase{}
		if tmpPrvdtl, ist := sycPnrdtl.Load(dbsAirlfl + dbsFlnbfl + dbsDepart + strDatefl); ist {
			if getPrvdtl, mtc := tmpPrvdtl.(map[string]mdlJeddah.MdlJeddahPnrdtlDtbase); mtc {
				prvPnrdtl = getPrvdtl
				sycPnrdtl.Delete(dbsAirlfl + dbsFlnbfl + dbsDepart + strDatefl)
			}
		}

		// load cached previous data pnrlog
		prvPnrlog := map[string]mdlJeddah.MdlJeddahPnrlogDtbase{}
		if tmpPrvdtl, ist := sycPnrlog.Load(dbsAirlfl + dbsFlnbfl + dbsDepart + strDatefl); ist {
			if getPrvdtl, mtc := tmpPrvdtl.(map[string]mdlJeddah.MdlJeddahPnrlogDtbase); mtc {
				prvPnrlog = getPrvdtl
				sycPnrlog.Delete(dbsAirlfl + dbsFlnbfl + dbsDepart + strDatefl)
			}
		}

		// If flight date is older than previous day => handle historical updates
		if intDatefl < int32(intDatepv) {
			for _, pnrdtl := range prvPnrdtl {
				if getAgtnme := getAgtnme(dbsAirlfl, pnrdtl.Agtnme); getAgtnme.Agtidn != "" {
					prvAgtnme(&pnrdtl, getAgtnme)
				}
			}

			// flush if needed
			fncGlobal.FncGlobalDtbaseBtcwrt(map[string]*[]mongo.WriteModel{
				"jeddah_pnrdtl": &mgoPnrdtl,
				"jeddah_pnrsmr": &mgoPnrsmr}, blimit)
			continue
		}

		// Helper to create PNR detail object
		var objPnrdtl = func(lcnpun mdlSbrapi.MdlSbrapiLcnpunDtbase,
			getAgtnme mdlJeddah.MdlJeddahAgtnmeDtbase) mdlJeddah.MdlJeddahPnrdtlDtbase {
			fnlPnrdtl := mdlJeddah.MdlJeddahPnrdtlDtbase{
				Prmkey: dbsAirlfl + dbsFlnbfl,
				Airlfl: dbsAirlfl,
				Flnbfl: dbsFlnbfl,
				Depart: dbsDepart,
				Routfl: dbsRoutfl,
				Clssfl: lcnpun.Clssfl,
				Datefl: intDatefl,
				Dateup: int32(intDatenw),
				Timeup: int64(intTimenw),
				Agtnme: lcnpun.Agtnme,
				Agtdtl: getAgtnme.Agtdtl,
				Agtidn: getAgtnme.Agtidn,
				Pnrcde: lcnpun.Pnrcde,
				Rtlsrs: getAgtnme.Rtlsrs,
				Drules: 100,
				Totisd: lcnpun.Totpax,
				Totpax: lcnpun.Totpax,
			}
			return fnlPnrdtl
		}

		// Looping PUN data
		tmpBooked := map[string]int{}
		slcOutpun, errOutpun := fncSbrapi.FncSbrapiLcnpunMainob("PUN", *tkn, objParams)
		if errOutpun == nil {
			for _, lcnpun := range slcOutpun {
				mgoLcnpun = append(mgoLcnpun, mongo.NewUpdateOneModel().
					SetFilter(bson.M{"prmkey": lcnpun.Prmkey}).
					SetUpdate(bson.M{"$set": lcnpun}).
					SetUpsert(true))
				tmpBooked[lcnpun.Pnrcde] = lcnpun.Totpax
			}
		}

		// declare LC Parameter
		var fnlIsjedh string
		var tmpPnrsmr, tmpPnrdtl, tmpLcnpun, tmpFlnbfl []mongo.WriteModel

		// Looping LC data
		slcOutllc, errOutllc := fncSbrapi.FncSbrapiLcnpunMainob("LC", *tkn, objParams)
		if errOutllc == nil {
			for _, lcnpun := range slcOutllc {

				// Delcare Pnr log data
				objParams.Pnrcde = lcnpun.Pnrcde
				getAgtnme := getAgtnme(dbsAirlfl, lcnpun.Agtnme)
				fnlPnrsmr := mdlJeddah.MdlJeddahPnrsmrDtbase{}
				fnlPnrdtl := objPnrdtl(lcnpun, getAgtnme)

				// Cek booked data from PUN
				if nowBooked, ist := tmpBooked[lcnpun.Pnrcde]; ist {
					valIssued := lcnpun.Totpax - nowBooked
					intTotbok += nowBooked
					fnlPnrdtl.Totisd = valIssued
					fnlPnrdtl.Totbok = nowBooked
				}

				// Cek from PNR detail and get previous data
				pnrdtl, istdtl := prvPnrdtl[lcnpun.Pnrcde]
				if istdtl {
					delete(prvPnrdtl, lcnpun.Pnrcde)
					fnlPnrdtl.Totspl = pnrdtl.Totspl
					fnlPnrdtl.Totchg = pnrdtl.Totchg
					fnlPnrdtl.Totcxl = pnrdtl.Totcxl
					fnlPnrdtl.Toflnm = pnrdtl.Toflnm

					// Cek if jeddah
					if pnrdtl.Isjedh == "Jeddah" {
						fnlIsjedh = "Jeddah"
					}

					// Update previous data for makesure
					if getAgtnme.Agtidn != "" {
						prvAgtnme(&fnlPnrdtl, getAgtnme)
					}

					// Cek from PNR log and get previous data
					nowRemove := pnrdtl.Totpax - lcnpun.Totpax
					if nowRemove != 0 {

						// Get Rservation to get itinerary and other
						_, _, fnlPnrsmr = getRsrvtn(intTimenw,
							intDatenw, objParams, &fnlPnrdtl, &tmpFlnbfl)
					}
				}

				// If PNR not found on previous data
				if !istdtl {

					// Get Rservation to get itinerary and other
					_, nowIsjedh, nowPnrsmr := getRsrvtn(intTimenw,
						intDatenw, objParams, &fnlPnrdtl, &tmpFlnbfl)
					fnlPnrsmr = nowPnrsmr
					updAgtnme(nowIsjedh, getAgtnme, &mapAgtnme, &fnlPnrdtl, &fnlPnrsmr)
					if nowIsjedh == "Jeddah" {
						fnlIsjedh = "Jeddah"
					}
				}

				// Get rules and push to batch
				intTotisd += fnlPnrdtl.Totisd
				intTotpax += fnlPnrdtl.Totpax
				intTotcxl += fnlPnrdtl.Totcxl
				intTotspl += fnlPnrdtl.Totspl
				getRules(&fnlPnrdtl, strDatefl)
				tmpLcnpun = append(tmpLcnpun, mongo.NewUpdateOneModel().
					SetFilter(bson.M{"prmkey": lcnpun.Prmkey}).
					SetUpdate(bson.M{"$set": lcnpun}).
					SetUpsert(true))
				updPnrall(&fnlPnrdtl, &fnlPnrsmr, &tmpPnrdtl, &tmpPnrsmr)
			}
		}

		// Start Switch
		switch {

		// If flight still available
		case len(slcOutllc) != 0:
			flbase.Flstat = "Operate"

			// If prev pnr isset but now null
			if len(prvPnrlog) > 0 {
				for _, lcnpun := range prvPnrlog {
					objParams.Pnrcde = lcnpun.Pnrcde

					// Get prev summary pnr array split and cancel
					if nowPnrdtl, ist := prvPnrdtl[lcnpun.Pnrcde]; ist {
						delete(prvPnrdtl, lcnpun.Pnrcde)

						// Get Rservation to get itinerary and other
						strToflnm, nowIsjedh, nowPnrsmr := getRsrvtn(intTimenw,
							intDatenw, objParams, &nowPnrdtl, &tmpFlnbfl)

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

						// Check agent name
						if nowPnrdtl.Agtidn == "" {
							getAgtnme := getAgtnme(dbsAirlfl, lcnpun.Agtnme)
							updAgtnme(nowIsjedh, getAgtnme, &mapAgtnme, &nowPnrdtl, &nowPnrsmr)
						}

						// Push to batch
						updPnrall(&nowPnrdtl, &nowPnrsmr, &tmpPnrdtl, &tmpPnrsmr)
					}
				}
			}

			// Continue if not exist Jeddah route
			if fnlIsjedh == "Non Jeddah" {
				continue
			} else {
				mgoPnrsmr = append(mgoPnrsmr, tmpPnrsmr...)
				mgoLcnpun = append(mgoLcnpun, tmpLcnpun...)
				mgoPnrdtl = append(mgoPnrdtl, tmpPnrdtl...)
				mgoFlnbfl = append(mgoFlnbfl, tmpFlnbfl...)
			}

		// If flight cancel/change but never in db
		case len(prvPnrdtl) == 0:
			flbase.Flstat = "Cancel"

			// declare LC Parameter
			var fnlIsjedh string
			var tmpPnrsmr, tmpPnrdtl, tmpLcnpun, tmpFlnbfl []mongo.WriteModel

			// Looping LC data
			slcOutllc, errOutllc := fncSbrapi.FncSbrapiLcnpunMainob("LDN", *tkn, objParams)
			if errOutllc == nil {
				for _, lcnpun := range slcOutllc {

					// Delcare Pnr log data
					objParams.Pnrcde = lcnpun.Pnrcde
					getAgtnme := getAgtnme(dbsAirlfl, lcnpun.Agtnme)
					nowPnrsmr := mdlJeddah.MdlJeddahPnrsmrDtbase{}
					nowPnrdtl := objPnrdtl(lcnpun, getAgtnme)

					// Get Rservation to get itinerary and other
					strToflnm, nowIsjedh, nowPnrsmr := getRsrvtn(intTimenw,
						intDatenw, objParams, &nowPnrdtl, &tmpFlnbfl)
					if nowIsjedh == "Jeddah" {
						fnlIsjedh = "Jeddah"
					}

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

					// Check agent name
					if nowPnrdtl.Agtidn == "" {
						updAgtnme(nowIsjedh, getAgtnme, &mapAgtnme, &nowPnrdtl, &nowPnrsmr)
					}

					// Push to batch
					tmpLcnpun = append(tmpLcnpun, mongo.NewUpdateOneModel().
						SetFilter(bson.M{"prmkey": lcnpun.Prmkey}).
						SetUpdate(bson.M{"$set": lcnpun}).
						SetUpsert(true))
					updPnrall(&nowPnrdtl, &nowPnrsmr, &tmpPnrdtl, &tmpPnrsmr)
				}
			}

			// Continue if not exist Jeddah route
			if fnlIsjedh == "Non Jeddah" {
				continue
			} else {
				mgoPnrsmr = append(mgoPnrsmr, tmpPnrsmr...)
				mgoLcnpun = append(mgoLcnpun, tmpLcnpun...)
				mgoPnrdtl = append(mgoPnrdtl, tmpPnrdtl...)
				mgoFlnbfl = append(mgoFlnbfl, tmpFlnbfl...)
			}
		}

		// Push new data flight to database
		mgoFlnbfl = append(mgoFlnbfl, mongo.NewUpdateOneModel().
			SetFilter(bson.M{"prmkey": flbase.Prmkey}).
			SetUpdate(bson.M{"$set": flbase}).
			SetUpsert(true))

		// Push to Pnr log data
		smrFlstat := "Operate"
		if intTotpax == 0 {
			smrFlstat = "Cancel"
		}
		nmodelSmrfln := mongo.NewUpdateOneModel().
			SetFilter(bson.M{"prmkey": flbase.Prmkey}).
			SetUpdate(bson.M{"$set": mdlJeddah.MdlJeddahFlnsmrDtbase{
				Prmkey: flbase.Prmkey, Airlfl: dbsAirlfl, Flnbfl: dbsFlnbfl,
				Depart: dbsDepart, Routfl: dbsRoutfl, Datefl: intDatefl,
				Dateup: int32(intDatenw), Timeup: int64(intTimenw),
				Totisd: intTotisd, Totbok: intTotbok, Totpax: intTotpax,
				Totcxl: intTotcxl, Totchg: intTotchg, Totspl: intTotspl,
				Flstat: smrFlstat, Isjedh: fnlIsjedh}}).
			SetUpsert(true)
		mgoSmrfln = append(mgoSmrfln, nmodelSmrfln)

		// Final flush
		for _, agtnme := range mapAgtnme {
			mgoAgtnme = append(mgoAgtnme, agtnme)
		}
		fncGlobal.FncGlobalDtbaseBtcwrt(map[string]*[]mongo.WriteModel{
			"jeddah_pnrdtl": &mgoPnrdtl,
			"jeddah_pnrsmr": &mgoPnrsmr,
			"jeddah_flnbfl": &mgoFlnbfl,
			"jeddah_flnsmr": &mgoSmrfln,
			"jeddah_agentx": &mgoAgtnme,
			"jeddah_lcnpun": &mgoLcnpun}, blimit)

		// Indicator finish process
		fmt.Println("Done", cntdta, "-", dbsAirlfl, dbsFlnbfl, intDatefl,
			dbsRoutfl, time.Now().Format("06-01-02/15:04"))
	}

	// Final flush
	fncGlobal.FncGlobalDtbaseBtcwrt(map[string]*[]mongo.WriteModel{
		"jeddah_pnrdtl": &mgoPnrdtl,
		"jeddah_pnrsmr": &mgoPnrsmr,
		"jeddah_flnbfl": &mgoFlnbfl,
		"jeddah_flnsmr": &mgoSmrfln,
		"jeddah_agentx": &mgoAgtnme,
		"jeddah_lcnpun": &mgoLcnpun}, 0)

}
