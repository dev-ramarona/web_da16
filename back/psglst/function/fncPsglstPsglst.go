package fncPsglst

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"
	fncSbrapi "back/sbrapi/function"
	mdlSbrapi "back/sbrapi/model"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FncPsglstPsglstPrcess(rspPsglst []mdlPsglst.MdlPsglstPsgdtlDtbase,
	nowObjtkn mdlSbrapi.MdlSbrapiMsghdrParams, objParams mdlSbrapi.MdlSbrapiMsghdrApndix,
	sycPnrcde, sycChrter, sycFrbase, sycFrtaxs, sycFlhour, sycMilege, idcFrbase, idcFrtaxs, sycErrlog *sync.Map,
	mapCurrcv map[string]mdlPsglst.MdlPsglstCurrcvDtbase,
	mapClslvl map[string]mdlPsglst.MdlPsglstClsslvDtbase, nowErignr string) (
	[]mongo.WriteModel, []mongo.WriteModel, []mongo.WriteModel,
	[]mongo.WriteModel, []mongo.WriteModel, []mongo.WriteModel) {
	sycWgroup, sycClrpsg, sycNulpsg := &sync.WaitGroup{}, &sync.Map{}, &sync.Map{}
	totPsgdtl := len(rspPsglst)
	for _, psglst := range rspPsglst {

		// Get null data
		if psglst.Tktnfl == "" || psglst.Pnrcde == "" {
			err := fncSbrapi.FncSbrapiPsgdtaMainob(nowObjtkn, mapClslvl, &psglst)
			if err != nil || psglst.Tktnfl == "" || psglst.Pnrcde == "" {
				fncGlobal.FncGlobalMainprNoterr(&psglst.Noterr, "PSGDATA NIL")
			} else {
				fncGlobal.FncGlobalMainprNoterr(&psglst.Noteup, "PSGDATA GET")
			}
		}

		// Running concurency every psglst
		sycWgroup.Add(1)
		go FncPslgstRsvpnrMainpg(psglst,
			sycClrpsg, sycNulpsg, sycPnrcde, sycChrter, sycMilege,
			mapCurrcv,
			sycWgroup, nowObjtkn,
			psglst.Airlfl, psglst.Pnrcde, "")
	}

	// Looping null and wait all goroutine finish
	fnlPrvpsg := map[string]map[string]mdlPsglst.MdlPsglstPsgdtlDtbase{}
	sycWgroup.Wait()
	sycNulpsg.Range(func(key, val any) bool {
		if mtcPsglst, mtc := val.(mdlPsglst.MdlPsglstPsgdtlDtbase); mtc {
			if _, ist := sycClrpsg.Load(key); !ist {
				slcPntitl := strings.Split(mtcPsglst.Pnritl, "|")
				for _, pntitl := range slcPntitl {
					if pntitl == "" {
						continue
					}
					arrPnritl := strings.Split(pntitl, "*")
					nowAirlfl := arrPnritl[0]
					nowPnrcde := arrPnritl[1]
					if fnlPrvpsg[nowAirlfl] == nil {
						fnlPrvpsg[nowAirlfl] = make(map[string]mdlPsglst.MdlPsglstPsgdtlDtbase)
					}
					if slices.Contains([]string{"JT", "ID", "IW", "IU", "OD", "SL"}, nowAirlfl) {
						fnlPrvpsg[nowAirlfl][nowPnrcde+"|"+mtcPsglst.Prmkey] = mtcPsglst
					}
				}
			}
		}
		return true
	})

	// Looping exist data
	for airlfl, fstPrvpsg := range fnlPrvpsg {
		newObjtkn, er1 := fncSbrapi.FncSbrapiCrtssnMainob(airlfl)
		tmpSycwgp := &sync.WaitGroup{}
		if er1 != nil {
			for i := 0; i < 3; i++ {
				getObjtkn, er3 := fncSbrapi.FncSbrapiCrtssnMainob(airlfl)
				if er3 == nil && getObjtkn.Bsttkn != "" {
					newObjtkn = getObjtkn
					break
				}
				time.Sleep(500 * time.Millisecond)
			}
		}
		for rawkey, scdPrvpsg := range fstPrvpsg {
			slckey := strings.Split(rawkey, "|")
			pnrcde, prmkey := slckey[0], slckey[1]
			if _, ist := sycClrpsg.Load(prmkey); !ist {
				tmpSycwgp.Add(1)
				go FncPslgstRsvpnrMainpg(scdPrvpsg,
					sycClrpsg, sycNulpsg, sycPnrcde, sycChrter, sycMilege,
					mapCurrcv,
					tmpSycwgp, newObjtkn,
					airlfl, pnrcde, "last")
			}
		}
		tmpSycwgp.Wait()
		fncSbrapi.FncSbrapiClsssnMainob(newObjtkn)
	}

	// FInal loop and push to database
	totSmmary := mdlPsglst.MdlPsglstPsgsmrDtbase{
		Mnthfl: objParams.Mnthfl, Datefl: objParams.Datefl,
		Ndayfl: objParams.Ndayfl, Depart: objParams.Depart,
		Airlfl: objParams.Airlfl, Flnbfl: objParams.Flnbfl,
		Prmkey: objParams.Airlfl + objParams.Flnbfl + objParams.Depart +
			strconv.Itoa(int(objParams.Datefl)),
	}
	mgoFrbase, mgoFrtaxs := []mongo.WriteModel{}, []mongo.WriteModel{}
	mgoFlhour, mgoMilege := []mongo.WriteModel{}, []mongo.WriteModel{}
	mgoPsgdtl, mgoPsgsmr := []mongo.WriteModel{}, []mongo.WriteModel{}
	totClrpsg := 0
	sycClrpsg.Range(func(key, val any) bool {
		if mtcPsglst, mtc := val.(mdlPsglst.MdlPsglstPsgdtlDtbase); mtc {

			// Manage route
			fnlRoutac := mtcPsglst.Routfl
			for _, routll := range []string{mtcPsglst.Routvc, mtcPsglst.Routfr} {
				if len(routll) >= 7 {
					slcRoutmx := strings.Split(mtcPsglst.Routmx, "-")
					slcRoutac := []string{}
					varDeprfl, varDeprvc := mtcPsglst.Depart, ""
					varArvlfl, varArvlvc := mtcPsglst.Arrivl, ""
					if mtcPsglst.Isitfl == "F" && len(routll) >= 7 {
						varDeprvc = routll[:3]
						varArvlvc = routll[4:]
					}
					for _, dstrct := range slcRoutmx {
						if varDeprfl == dstrct || varDeprvc == dstrct || len(slcRoutac) != 0 {
							slcRoutac = append(slcRoutac, dstrct)
						}
						if dstrct == varArvlfl || dstrct == varArvlvc {
							break
						}
					}
					if len(fnlRoutac) >= len(strings.Join(slcRoutac, "-")) {
						continue
					}
					fnlRoutac = strings.Join(slcRoutac, "-")
					mtcPsglst.Routac = strings.Join(slcRoutac, "-")
				}

				// Get route actual
				slcRoutac := strings.Split(fnlRoutac, "-")
				if slcPstion := slices.Index(slcRoutac, mtcPsglst.Depart); slcPstion != -1 {
					if slcPstion+1 < len(slcRoutac) {
						mtcPsglst.Routfl = strings.Join(slcRoutac[slcPstion:slcPstion+2], "-")
					}
				}
				if mtcPsglst.Routfl == "" {
					mtcPsglst.Routfl = mtcPsglst.Depart + "-" + mtcPsglst.Arrivl
				}
			}

			// Get rout from fare calc and segment
			slcRoutfr, cekRoutfr, lstRoutsg := []string{}, false, ""
			if len(mtcPsglst.Routfr) >= 7 {
				regRoutfr := regexp.MustCompile(fmt.Sprintf("%s.+%s", mtcPsglst.Routfr[:3],
					mtcPsglst.Routfr[4:]))
				if res := regRoutfr.MatchString(mtcPsglst.Routsg); res {

					// Combine route
					slcRoutmx := strings.Split(mtcPsglst.Routmx, "-")
					nowRoutmx := slcRoutmx[0] + "-" + slcRoutmx[len(slcRoutmx)-1]
					segFullrt := mtcPsglst.Routsg
					if strings.Contains(mtcPsglst.Routsg, nowRoutmx) {
						segFullrt = strings.Replace(mtcPsglst.Routsg, nowRoutmx, mtcPsglst.Routmx, 1)
					}

					// Looping full rout max
					for _, routsg := range strings.Split(segFullrt, "-") {
						if strings.Contains(mtcPsglst.Routfx, lstRoutsg+"-"+routsg) {
							if len(slcRoutfr) > 0 && slcRoutfr[len(slcRoutfr)-1] == lstRoutsg {
								slcRoutfr = slcRoutfr[:len(slcRoutfr)-1]
								continue
							}
						}
						if lstRoutsg+"-"+routsg == mtcPsglst.Routfr {
							slcRoutfr = []string{lstRoutsg, routsg}
							break
						}
						if mtcPsglst.Routfr[4:] == routsg && len(slcRoutfr) > 0 {
							cekRoutfr = true
							slcRoutfr = append(slcRoutfr, routsg)
							break
						}
						lstRoutsg = routsg
						if mtcPsglst.Routfr[:3] == routsg || len(slcRoutfr) > 0 {
							slcRoutfr = append(slcRoutfr, routsg)
						}
					}

					// Last push route actual from route fare
					if cekRoutfr && len(slcRoutfr) >= 2 {
						strRoutfr := strings.Join(slcRoutfr, "-")
						if len(strRoutfr) > len(fnlRoutac) {
							mtcPsglst.Routac = strRoutfr
							fnlRoutac = strRoutfr
						}
					}
				}
			}

			// Get flown farebase
			mapRoutfb := map[string]string{"routfl": mtcPsglst.Airlfl + mtcPsglst.Routfl}
			if mtcPsglst.Ntafvc == 0 && (mtcPsglst.Isitnr == "" || mtcPsglst.Frbcde == "HB") {
				nowRoutvc := mtcPsglst.Routvc
				if mtcPsglst.Routvc == "" {
					nowRoutvc = mtcPsglst.Depart + "-" + mtcPsglst.Arrivl
				}
				if slices.Contains([]string{"JT", "ID", "IW", "IU", "OD", "SL"}, mtcPsglst.Airlvc) {
					mapRoutfb["routvc"] = mtcPsglst.Airlvc + nowRoutvc
				} else {
					mapRoutfb["routvc"] = mtcPsglst.Airlfl + nowRoutvc
				}
			}
			for key, val := range mapRoutfb {
				for nky, nvl := range map[string]string{
					"FRBCDE": val + mtcPsglst.Frbcde,
					"CLSSFL": val + mtcPsglst.Clssfl} {
					if len(val) < 7 || (key == "routfl" && nky == "FRBCDE") {
						continue
					}

					// Reuse func frbase
					fncFrbase := func() bool {
						tmpFrbcde := []string{nvl}
						if nky == "FRBCDE" && strings.Contains(nvl, "RT") && len(nvl) >= 9 {
							tmpRoutnw := nvl[2:9]
							newRoutnw := tmpRoutnw[4:] + "-" + tmpRoutnw[:3]
							tmpFrbcde = append(tmpFrbcde, strings.ReplaceAll(nvl, tmpRoutnw, newRoutnw))
						}
						for _, nowFrbcde := range tmpFrbcde {
							istFrbase, ist := sycFrbase.Load(nowFrbcde)
							if mtcFrbase, mtc := istFrbase.(mdlPsglst.MdlPsglstFrbaseDtbase); mtc && ist {
								mtcPsglst.Ntaffl = mtcFrbase.Frbnta
								if key == "routfl" {
									mtcPsglst.Ntaffl = mtcFrbase.Frbnta
								} else {
									mtcPsglst.Ntafvc = float64(mtcFrbase.Frbnta)
									mtcPsglst.Isittf = nky
								}
								return true
							}
						}
						return false
					}

					// Hit API Sabre if null data
					if rspfnc := fncFrbase(); rspfnc {
						break
					} else {
						objParams := mdlSbrapi.MdlSbrapiMsghdrApndix{
							Airlfl: mtcPsglst.Airlfl, Depart: val[:3], Arrivl: val[4:], Routfl: val}

						// Check indicator before hit API
						if _, ist := idcFrbase.Load(val); !ist {
							nowmgo, err := fncSbrapi.FncSbrapiFrbaseMainob(nowObjtkn, objParams, sycFrbase, mapClslvl)
							if err == nil {
								mgoFrbase = append(mgoFrbase, nowmgo...)
								idcFrbase.Store(val, true)
							}
						}

						// Get flown farebase
						if rspfnc := fncFrbase(); rspfnc {
							break
						}
					}
				}
			}

			// Get all taxes
			mapRouttx := map[string]string{"routfl": mtcPsglst.Airlfl + mtcPsglst.Routfl + mtcPsglst.Cbinfl}
			if mtcPsglst.Isitnr != "CREW" {
				mapRouttx["routvc"] = mtcPsglst.Airlfl + mtcPsglst.Routvc + mtcPsglst.Cbinvc
			}
			for key, val := range mapRouttx {
				if len(val) < 7 {
					continue
				}

				// Reuse func frtaxs
				fncFrtaxs := func() bool {
					istFrtaxs, ist := sycFrtaxs.Load(val)
					if mtcFrtaxs, mtc := istFrtaxs.(mdlPsglst.MdlPsglstFrtaxsDtbase); mtc && ist {
						if key == "routfl" {
							mtcPsglst.Yqtxfl = mtcFrtaxs.Ftfuel
						} else {
							mtcPsglst.Yqtxvc = float64(mtcFrtaxs.Ftfuel)
						}
					}
					return ist
				}

				// Hit API Sabre if null data
				if rspfnc := fncFrtaxs(); !rspfnc && len(val) == 10 {
					slcClscbn := []string{"Y", "C"}
					tmpPrmkey := val[:9]
					for _, clscbn := range slcClscbn {
						nowPrmkey := tmpPrmkey + clscbn
						objParams := mdlSbrapi.MdlSbrapiMsghdrApndix{Airlfl: mtcPsglst.Airlfl,
							Depart: val[:3], Arrivl: val[4:], Routfl: val}

						// Check indicator before hit API
						if _, ist := idcFrtaxs.Load(nowPrmkey); !ist {
							nowmgo, err := fncSbrapi.FncSbrapiFrtaxsMainob(nowObjtkn, objParams, sycFrtaxs, clscbn)
							if err == nil {
								mgoFrtaxs = append(mgoFrtaxs, nowmgo...)
								idcFrtaxs.Store(nowPrmkey, true)
							}
						}
					}

					// Get all farebase
					fncFrtaxs()
				}
			}

			// Get final price
			if fnlRoutac == "" {
				fnlRoutac = mtcPsglst.Depart + "-" + mtcPsglst.Arrivl
			}
			slcRoutac, totMilege, nowMilege := strings.Split(fnlRoutac, "-"), float64(0), float64(0)
			for i := 0; i < len(slcRoutac)-1; i++ {

				// Reuse func milege
				fncMilege := func(nowRoutac string) bool {
					istMilege, ist := sycMilege.Load(nowRoutac)
					if mtcMilege, mtc := istMilege.(mdlPsglst.MdlPsglstMilegeDtbase); mtc && ist {
						if slcRoutac[i] == mtcPsglst.Depart {
							nowMilege = float64(mtcMilege.Milege)
						}
						totMilege += float64(mtcMilege.Milege)
					}
					return ist
				}

				// Route milege hit API Sabre if null data
				nowRoutac := slcRoutac[i] + "-" + slcRoutac[i+1]
				if rspfnc := fncMilege(nowRoutac); !rspfnc {
					rspMilege, err := fncSbrapi.FncSbrapiMilegeMainob(nowObjtkn, fnlRoutac)
					if err == nil {
						for _, milege := range rspMilege {
							if _, ist := sycMilege.Load(milege.Routfl); !ist {
								sycMilege.Store(milege.Routfl, milege)
								mgoMilege = append(mgoMilege, mongo.NewUpdateOneModel().
									SetFilter(bson.M{"routfl": milege.Routfl}).
									SetUpdate(bson.M{"$set": milege}).SetUpsert(true))
							}
						}
						fncMilege(nowRoutac)
					}
				}
			}

			// Final rate and price adjustment
			mtcPsglst.Frrate = nowMilege / totMilege
			valChrter := float64(1)
			if mtcPsglst.Isitct == "CT" {
				valChrter = 0
			}
			mtcPsglst.Ntaffl = int32(float64(mtcPsglst.Ntaffl) * valChrter)
			mtcPsglst.Yqtxfl = int32(float64(mtcPsglst.Yqtxfl) * valChrter)
			mtcPsglst.Ntafvc = float64(mtcPsglst.Ntafvc) * mtcPsglst.Frrate * valChrter
			mtcPsglst.Yqtxvc = float64(mtcPsglst.Yqtxvc) * mtcPsglst.Frrate * valChrter

			// Push summary
			totSmmary.Totnta += mtcPsglst.Ntafvc
			totSmmary.Tottyq += mtcPsglst.Yqtxvc
			totSmmary.Totpax += 1

			// Push final to database
			mgoPsgdtl = append(mgoPsgdtl, mongo.NewUpdateOneModel().
				SetFilter(bson.M{"prmkey": mtcPsglst.Prmkey}).
				SetUpdate(bson.M{"$set": mtcPsglst}).
				SetUpsert(true))
			totClrpsg++
		}
		return true
	})

	// Cek different data
	FncPsglstErrlogManage(mdlPsglst.MdlPsglstErrlogDtbase{
		Erpart: "psgdtl", Ersrce: "sbrapi", Erdvsn: "mnfest",
		Dateup: int32(objParams.Dateup), Timeup: int64(objParams.Timeup),
		Datefl: int32(objParams.Datefl), Airlfl: objParams.Airlfl,
		Flnbfl: objParams.Flnbfl, Routfl: objParams.Routfl, Worker: 1, Erignr: nowErignr,
		Paxdif: fmt.Sprintf("%d/%d", totClrpsg, totPsgdtl),
	}, totClrpsg != totPsgdtl, sycErrlog)

	// Return final data
	mgoPsgsmr = append(mgoPsgsmr, mongo.NewUpdateOneModel().
		SetFilter(bson.M{"prmkey": totSmmary.Prmkey}).
		SetUpdate(bson.M{"$set": totSmmary}).
		SetUpsert(true))
	return mgoPsgdtl, mgoPsgsmr, mgoFrbase, mgoFrtaxs, mgoFlhour, mgoMilege
}
