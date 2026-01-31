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
	sycPnrcde, sycChrter, sycFrbase, sycFrtaxs, sycFlhour, sycMilege,
	idcFrbase, idcFrtaxs, sycErrlog *sync.Map,
	slcHfbalv []mdlPsglst.MdlPsglstHfbalvDtbase,
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

	// Declare empty variable final
	totClrpsg := 0
	mgoFrbase, mgoFrtaxs := []mongo.WriteModel{}, []mongo.WriteModel{}
	mgoFlhour, mgoMilege := []mongo.WriteModel{}, []mongo.WriteModel{}
	mgoPsgdtl, mgoPsgsmr := []mongo.WriteModel{}, []mongo.WriteModel{}
	fnlPsglst := []mdlPsglst.MdlPsglstPsgdtlDtbase{}
	mapPaidbt := map[string]int{}
	mapQntybt := map[string]int{}
	mapWghtbt := map[string]int{}
	mapFbavbt := map[string]int{}
	totSmmary := mdlPsglst.MdlPsglstPsgsmrDtbase{
		Mnthfl: objParams.Mnthfl, Datefl: objParams.Datefl,
		Ndayfl: objParams.Ndayfl, Depart: objParams.Depart,
		Airlfl: objParams.Airlfl, Flnbfl: objParams.Flnbfl,
		Prmkey: objParams.Airlfl + objParams.Flnbfl + objParams.Depart +
			strconv.Itoa(int(objParams.Datefl)),
	}

	// Semi final loop and push to final
	sycClrpsg.Range(func(key, val any) bool {
		if mtcPsglst, mtc := val.(mdlPsglst.MdlPsglstPsgdtlDtbase); mtc {

			// Total group summary bg and ae
			if mtcPsglst.Groupc != "" {
				mapPaidbt[mtcPsglst.Groupc] += int(mtcPsglst.Paidbt)
				mapQntybt[mtcPsglst.Groupc] += int(mtcPsglst.Qntybt)
				mapWghtbt[mtcPsglst.Groupc] += int(mtcPsglst.Wghtbt)
				mapFbavbt[mtcPsglst.Groupc] += int(mtcPsglst.Fbavbt)
			}

			// Get segment now
			if mtcPsglst.Segtkt != "" {
				prvTimefl, prvRoutfl, slcSegtkt, mtcSegtkt := "", "", []string{}, false
				sptSegtkt := strings.Split(mtcPsglst.Segtkt, "|")
				fstDepart := strings.Split(sptSegtkt[0], "-")[1]
				lstArrivl := strings.Split(sptSegtkt[len(sptSegtkt)-1], "-")[2]
				istRoutpp := fstDepart == lstArrivl
				for _, segtkt := range sptSegtkt {
					cpntkt := strings.Split(segtkt, "-")
					timecp := strings.Split(cpntkt[0], ":")
					if intime, _ := strconv.Atoi(timecp[0]); int64(intime) == mtcPsglst.Timefl ||
						(mtcPsglst.Routvc == cpntkt[1]+"-"+cpntkt[2] && mtcPsglst.Flnbfl == cpntkt[5]) {
						mtcSegtkt = true
					}

					// Gate logic
					if prvTimefl == "" {
						slcSegtkt = append(slcSegtkt, segtkt)
					} else {
						fmtprv, _ := time.Parse("0601021504", prvTimefl)
						fmtnow, _ := time.Parse("0601021504", timecp[0])
						fmtdif := fmtprv.Sub(fmtnow)
						if fmtdif.Hours() > 24 || (prvRoutfl == cpntkt[1]+"-"+cpntkt[2] && istRoutpp) {
							if mtcSegtkt {
								continue
							}
							slcSegtkt = []string{}
						} else {
							slcSegtkt = append(slcSegtkt, segtkt)
						}

					}

					// Prev time flight or arrival
					prvRoutfl = cpntkt[2] + "-" + cpntkt[1]
					prvTimefl = timecp[0]
					if timecp[1] != "0101010000" {
						prvTimefl = timecp[1]
					}
				}

				// Get highest fba
				strSegtkt := strings.Join(slcSegtkt, "|")
				slcMaxfba := []int{int(mtcPsglst.Fbavbt)}
				bolStpsrc := true
				for _, nowSegtkt := range slcSegtkt {
					for _, hfbalv := range slcHfbalv {

						// Regex Airline
						nowAirlfl := hfbalv.Airlfl
						regAirlfl := regexp.MustCompile("-(" + nowAirlfl + ")-")
						lgcAirlfl := nowAirlfl == "ALL" || regAirlfl.MatchString(nowSegtkt)

						// Regex class flown
						nowClssfl := hfbalv.Clssfl
						regClssfl := regexp.MustCompile(`-(` + nowClssfl + `)$`)
						lgcClssfl := nowClssfl == "ALL" || regClssfl.MatchString(nowSegtkt)

						// Regex route flown
						fncRoutrg := func(dstrct string) string {
							if dstrct != "ALL" {
								return "(" + dstrct + ")"
							}
							return "[A-Z]{3}"
						}
						slcRoutfl := strings.Split(hfbalv.Routfl, "-")
						strRoutrg := fncRoutrg(slcRoutfl[0]) + ".+" + fncRoutrg(slcRoutfl[1])
						regRoutfl := regexp.MustCompile(strRoutrg)
						lgcRoutfl := regRoutfl.MatchString(strSegtkt)

						// Final result
						if lgcAirlfl && lgcClssfl && lgcRoutfl && bolStpsrc {
							slcMaxfba = append(slcMaxfba, int(hfbalv.Hfbabt))
							if hfbalv.Source == "VCR" {
								bolStpsrc = false
								slcMaxfba = []int{int(mtcPsglst.Fbavbt)}
							}
							continue
						}
					}
				}
				if mtcPsglst.Hfbabt == 0 && len(slcMaxfba) > 0 {
					mtcPsglst.Hfbabt = int32(slices.Max(slcMaxfba))
				}
			}
			fnlPsglst = append(fnlPsglst, mtcPsglst)
		}
		return true
	})

	// Looping again final
	for _, psglst := range fnlPsglst {
		if val, ist := mapPaidbt[psglst.Groupc]; ist {
			psglst.Ptotbt = int32(val)
		} else {
			psglst.Ptotbt = psglst.Paidbt
		}
		if val, ist := mapQntybt[psglst.Groupc]; ist {
			psglst.Qtotbt = int32(val)
		} else {
			psglst.Qtotbt = psglst.Qntybt
		}
		if val, ist := mapWghtbt[psglst.Groupc]; ist {
			psglst.Wtotbt = int32(val)
		} else {
			psglst.Wtotbt = psglst.Wghtbt
		}
		if val, ist := mapFbavbt[psglst.Groupc]; ist {
			psglst.Ftotbt = int32(val)
		} else {
			psglst.Ftotbt = psglst.Fbavbt
		}

		// Manage route
		fnlRoutac := psglst.Routfl
		for _, routll := range []string{psglst.Routvc, psglst.Routfr} {
			if len(routll) >= 7 {
				slcRoutmx := strings.Split(psglst.Routmx, "-")
				slcRoutac := []string{}
				varDeprfl, varDeprvc := psglst.Depart, ""
				varArvlfl, varArvlvc := psglst.Arrivl, ""
				if psglst.Isitfl == "F" && len(routll) >= 7 {
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
				psglst.Routac = strings.Join(slcRoutac, "-")
			}

			// Get route actual
			slcRoutac := strings.Split(fnlRoutac, "-")
			if slcPstion := slices.Index(slcRoutac, psglst.Depart); slcPstion != -1 {
				if slcPstion+1 < len(slcRoutac) {
					psglst.Routfl = strings.Join(slcRoutac[slcPstion:slcPstion+2], "-")
				}
			}
			if psglst.Routfl == "" {
				psglst.Routfl = psglst.Depart + "-" + psglst.Arrivl
			}
		}

		// Get rout from fare calc and segment
		slcRoutfr, cekRoutfr, lstRoutsg := []string{}, false, ""
		if len(psglst.Routfr) >= 7 {
			regRoutfr := regexp.MustCompile(fmt.Sprintf("%s.+%s", psglst.Routfr[:3],
				psglst.Routfr[4:]))
			if res := regRoutfr.MatchString(psglst.Routsg); res {

				// Combine route
				slcRoutmx := strings.Split(psglst.Routmx, "-")
				nowRoutmx := slcRoutmx[0] + "-" + slcRoutmx[len(slcRoutmx)-1]
				segFullrt := psglst.Routsg
				if strings.Contains(psglst.Routsg, nowRoutmx) {
					segFullrt = strings.Replace(psglst.Routsg, nowRoutmx, psglst.Routmx, 1)
				}

				// Looping full rout max
				for _, routsg := range strings.Split(segFullrt, "-") {
					if strings.Contains(psglst.Routfx, lstRoutsg+"-"+routsg) {
						if len(slcRoutfr) > 0 && slcRoutfr[len(slcRoutfr)-1] == lstRoutsg {
							slcRoutfr = slcRoutfr[:len(slcRoutfr)-1]
							continue
						}
					}
					if lstRoutsg+"-"+routsg == psglst.Routfr {
						slcRoutfr = []string{lstRoutsg, routsg}
						break
					}
					if psglst.Routfr[4:] == routsg && len(slcRoutfr) > 0 {
						cekRoutfr = true
						slcRoutfr = append(slcRoutfr, routsg)
						break
					}
					lstRoutsg = routsg
					if psglst.Routfr[:3] == routsg || len(slcRoutfr) > 0 {
						slcRoutfr = append(slcRoutfr, routsg)
					}
				}

				// Last push route actual from route fare
				if cekRoutfr && len(slcRoutfr) >= 2 {
					strRoutfr := strings.Join(slcRoutfr, "-")
					if len(strRoutfr) > len(fnlRoutac) {
						psglst.Routac = strRoutfr
						fnlRoutac = strRoutfr
					}
				}
			}
		}

		// Get flown farebase
		mapRoutfb := map[string]string{"routfl": psglst.Airlfl + psglst.Routfl}
		if psglst.Ntafvc == 0 && (psglst.Isitnr == "" || psglst.Frbcde == "HB") {
			nowRoutvc := psglst.Routvc
			if psglst.Routvc == "" {
				nowRoutvc = psglst.Depart + "-" + psglst.Arrivl
			}
			if slices.Contains([]string{"JT", "ID", "IW", "IU", "OD", "SL"}, psglst.Airlvc) {
				mapRoutfb["routvc"] = psglst.Airlvc + nowRoutvc
			} else {
				mapRoutfb["routvc"] = psglst.Airlfl + nowRoutvc
			}
		}
		for key, val := range mapRoutfb {
			for nky, nvl := range map[string]string{
				"FRBCDE": val + psglst.Frbcde,
				"CLSSFL": val + psglst.Clssfl} {
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
							psglst.Ntaffl = mtcFrbase.Frbnta
							if key == "routfl" {
								psglst.Ntaffl = mtcFrbase.Frbnta
							} else {
								psglst.Ntafvc = float64(mtcFrbase.Frbnta)
								psglst.Isittf = nky
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
						Airlfl: psglst.Airlfl, Depart: val[:3], Arrivl: val[4:], Routfl: val}

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
		mapRouttx := map[string]string{"routfl": psglst.Airlfl + psglst.Routfl + psglst.Cbinfl}
		if psglst.Isitnr != "CREW" {
			mapRouttx["routvc"] = psglst.Airlfl + psglst.Routvc + psglst.Cbinvc
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
						psglst.Yqtxfl = mtcFrtaxs.Ftfuel
					} else {
						psglst.Yqtxvc = float64(mtcFrtaxs.Ftfuel)
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
					objParams := mdlSbrapi.MdlSbrapiMsghdrApndix{Airlfl: psglst.Airlfl,
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
			fnlRoutac = psglst.Depart + "-" + psglst.Arrivl
		}
		slcRoutac, totMilege, nowMilege := strings.Split(fnlRoutac, "-"), float64(0), float64(0)
		for i := 0; i < len(slcRoutac)-1; i++ {

			// Reuse func milege
			fncMilege := func(nowRoutac string) bool {
				istMilege, ist := sycMilege.Load(nowRoutac)
				if mtcMilege, mtc := istMilege.(mdlPsglst.MdlPsglstMilegeDtbase); mtc && ist {
					if slcRoutac[i] == psglst.Depart {
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
		psglst.Frrate = nowMilege / totMilege
		valChrter := float64(1)
		if psglst.Isitct == "CT" {
			valChrter = 0
		}
		psglst.Ntaffl = int32(float64(psglst.Ntaffl) * valChrter)
		psglst.Yqtxfl = int32(float64(psglst.Yqtxfl) * valChrter)
		psglst.Ntafvc = psglst.Ntafvc * psglst.Frrate * valChrter
		psglst.Yqtxvc = psglst.Yqtxvc * psglst.Frrate * valChrter
		psglst.Fareae = psglst.Fareae * psglst.Frrate * valChrter

		// Push summary
		totSmmary.Totnta += psglst.Ntafvc
		totSmmary.Tottyq += psglst.Yqtxvc
		totSmmary.Totpax += 1
		totSmmary.Totfae += psglst.Fareae
		totSmmary.Totqfr += psglst.Qsrcvc

		// Push final to database
		mgoPsgdtl = append(mgoPsgdtl, mongo.NewUpdateOneModel().
			SetFilter(bson.M{"prmkey": psglst.Prmkey}).
			SetUpdate(bson.M{"$set": psglst}).
			SetUpsert(true))
		totClrpsg++
	}

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
