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
)

func FncPslgstRsvpnrMainpg(psglst mdlPsglst.MdlPsglstPsgdtlDtbase,
	sycClrpsg, sycNulpsg, sycPnrcde, sycChrter, sycMilege *sync.Map,
	mapCurrcv map[string]mdlPsglst.MdlPsglstCurrcvDtbase,
	sycWgroup *sync.WaitGroup, objtkn mdlSbrapi.MdlSbrapiMsghdrParams,
	airlfl, pnrcde, lstvar string) {
	var cekLstvar, cekChrter, cekIsflwn, cekNonrev, cekTcktng bool
	if sycWgroup != nil {
		defer sycWgroup.Done()
	}

	// DEBUGING
	if lstvar == "last" {
		psglst.Source += "|INTERLINE"
		cekLstvar = true
	}
	if objtkn.Bsttkn == "" {
		psglst.Source += "|TOKEN NIL" + airlfl
	}

	// Check on name XXDHC pax non rev and isit flown
	cekIsflwn = psglst.Isitfl == "F"
	if psglst.Nmefst == "XXDHC" || psglst.Nmelst == "XXDHC" ||
		psglst.Nmefst == "XXSNY" || psglst.Nmelst == "XXSNY" ||
		psglst.Nmefst == "XDHC" || psglst.Nmelst == "XDHC" {
		psglst.Isitnr = "CREW"
		cekNonrev = true
	}

	// Get ticketing from PNR
	slcSbarea := []string{"ITINERARY", "RECORD_LOCATOR", "ANCILLARY", "TICKETING"}

	// Get reservation
	nowRsvpnr := mdlSbrapi.MdlSbrapiRsvpnrRsprsv{}
	if istTcktng, ist := sycPnrcde.Load(pnrcde + airlfl); ist {
		if mtcTcktng, mtc := istTcktng.(mdlSbrapi.MdlSbrapiRsvpnrRsprsv); mtc {
			nowRsvpnr = mtcTcktng
		}
	} else {
		getTcktng, err := fncSbrapi.FncSbrapiRsvpnrMainob(objtkn, pnrcde, slcSbarea)
		if err != nil {
			return
		}
		nowRsvpnr = getTcktng
		sycPnrcde.Store(pnrcde+airlfl, getTcktng)
	}

	// If data not null
	if nowRsvpnr.BookingDetails.RecordLocator != "" {

		// Date formating PNR book PNR Create date
		varTimecr := nowRsvpnr.BookingDetails.SystemCreationTimestamp
		if pnrTimecr, err := time.Parse("2006-01-02T15:04:05", varTimecr); err == nil {
			rawTimerw, _ := strconv.Atoi(pnrTimecr.Format("0601021504"))
			psglst.Timecr = int64(rawTimerw)
		}

		// Get PNR interline
		objPnritl := nowRsvpnr.POS.Source.TTYRecordLocator
		slcPnrtil := strings.Split(psglst.Pnritl, "|")
		if objPnritl.RecordLocator != "" {
			nowPnritl := objPnritl.CRSCode + "*" + objPnritl.RecordLocator
			if !strings.Contains(psglst.Pnritl+psglst.Pnrcde, objPnritl.RecordLocator) {
				slcPnrtil = append(slcPnrtil, nowPnritl)
			}
		}

		// Get PNR interline and itinerary
		slcItinry := nowRsvpnr.PassengerReservation.Segments.Segment
		slcSegpnr := []string{}
		slcRoutsg, lstArrivl := []string{}, ""
		if len(slcItinry) != 0 {
			for idx, itinry := range slcItinry {
				if !slices.Contains([]string{"JT", "ID", "IW", "IU", "OD", "SL"},
					itinry.Air.OperatingAirlineCode) {
					continue
				}

				// PNR Interline
				rawPnritl := itinry.Air.AirlineRefId
				if len(rawPnritl) > 5 {
					if !strings.Contains(psglst.Pnritl+psglst.Pnrcde, rawPnritl[5:]) {
						slcPnrtil = append(slcPnrtil, rawPnritl[2:])
					}
				}

				// Get time flown
				rawTimefl := itinry.Air.DepartureDateTime
				fmtTimefl, _ := time.Parse("2006-01-02T15:04:05", rawTimefl)
				strTimefl := fmtTimefl.Format("0601021504")

				// Itinerary segment
				rawDepart := itinry.Air.DepartureAirport
				rawArrivl := itinry.Air.ArrivalAirport
				rawActncd := itinry.Air.ActionCode
				mktAirlfl := itinry.Air.MarketingAirlineCode
				optAirlfl := itinry.Air.OperatingAirlineCode
				mktFlnbfl := itinry.Air.MarketingFlightNumber
				optFlnbfl := itinry.Air.OperatingFlightNumber
				mktClssfl := itinry.Air.MarketingClassOfService
				optClssfl := itinry.Air.OperatingClassOfService
				fmtSegpnr := fmt.Sprintf("%s-%s-%s-%s-MKT-%s-%s-%s-OPT-%s-%s-%s",
					rawDepart, rawArrivl, rawActncd, strTimefl,
					mktAirlfl, mktFlnbfl, mktClssfl,
					optAirlfl, optFlnbfl, optClssfl)
				lstArrivl = rawArrivl
				if idx <= 1 || len(slcRoutsg) == 0 ||
					slcRoutsg[len(slcRoutsg)-1] != rawDepart {
					slcSegpnr = append(slcSegpnr, fmtSegpnr)
					slcRoutsg = append(slcRoutsg, rawDepart)
				}
			}
			slcRoutsg = append(slcRoutsg, lstArrivl)
			psglst.Routsg = strings.Join(slcRoutsg, "-")
			psglst.Segpnr = strings.Join(slcSegpnr, "|")
		}
		psglst.Pnritl = strings.Join(slcPnrtil, "|")

		// Get ticketing detail for issued date
		var slcTcktng = nowRsvpnr.PassengerReservation.TicketingInfo.TicketDetails
		var mapEmdnae = map[string]bool{}
		if len(slcTcktng) != 0 {
			for _, tcktng := range slcTcktng {

				// Logical gate for ticket number
				if psglst.Tktnfl == "" {
					strFmtnme := (psglst.Nmelst + "     ")[:5]
					strLstnme := (psglst.Nmefst + " ")[:1]
					cncFulln1 := strFmtnme + "/" + strLstnme
					cncFulln2 := psglst.Nmelst + "/" + strLstnme
					if cncFulln1 == tcktng.PassengerName ||
						cncFulln2 == tcktng.PassengerName {

						// Get ticket number blank and emd
						if psglst.Tktnvc == "" && tcktng.TicketNumber[3:4] != "4" {
							psglst.Tktnvc = tcktng.TicketNumber[:13]
						} else if tcktng.TicketNumber[3:4] == "4" {
							mapEmdnae[tcktng.TicketNumber[:13]] = true
						}
					}
				}
			}
		}

		// Get ancillary
		if len(nowRsvpnr.OpenReservationElements) > 0 {
			for _, elm := range nowRsvpnr.OpenReservationElements {
				delete(mapEmdnae, psglst.Emdnae)
				if elm.EMDNumber == "" || elm.ActionCode != "HI" ||
					elm.NameAssociationList.FirstName != psglst.Nmefst ||
					elm.NameAssociationList.LastName != psglst.Nmelst {
					continue
				}

				// Looping segment assoc
				nowRoutae := ""
				for _, cpn := range elm.SegmentAssociationList {
					nowDepart, nowArrivl := cpn.BoardPoint, cpn.OffPoint

					// Compare to flown data
					if nowDepart == psglst.Depart || nowArrivl == psglst.Arrivl {
						nowRoutae = cpn.BoardPoint + "-" + cpn.OffPoint
					}

					// Compare to route vcr
					if len(psglst.Routvc) >= 7 {
						if nowDepart == psglst.Routvc[:3] || nowArrivl == psglst.Routvc[4:] {
							nowRoutae = cpn.BoardPoint + "-" + cpn.OffPoint
						}
					}
				}

				// Push final data if assoc
				nowPaidbt := 1
				regDescae := regexp.MustCompile(`\d+K|\d+ K`)
				rslDescae := regDescae.FindAllString(elm.CommercialName, -1)
				if len(rslDescae) > 0 {
					rawPaidbt := rslDescae[0][:len(rslDescae[0])-1]
					intPaidbt, _ := strconv.Atoi(rawPaidbt)
					nowPaidbt = intPaidbt
				}

				// If get route assoc
				if nowRoutae != "" {
					fncGlobal.FncGlobalMainprNoterr(&psglst.Gpcdae, elm.GroupCode)
					fncGlobal.FncGlobalMainprNoterr(&psglst.Sbcdae, elm.RficSubcode)
					fncGlobal.FncGlobalMainprNoterr(&psglst.Descae, elm.CommercialName)
					psglst.Wgbgae += int32(nowPaidbt)
					psglst.Qtbgae += int32(elm.NumberOfItems)
					psglst.Routae = nowRoutae
					fncGlobal.FncGlobalMainprNoterr(&psglst.Emdnae, elm.EMDNumber)

					// Fare manage
					psglst.Currae = elm.OriginalBasePrice.Currency
					if psglst.Currae != "IDR" {
						if vlx, ist := mapCurrcv[psglst.Currae]; ist {
							cnvFareae := elm.OriginalBasePrice.Price / vlx.Crrate
							psglst.Fareae += cnvFareae
						}
					} else {
						psglst.Fareae += elm.OriginalBasePrice.Price
					}

					// Check group code
					fstGroupc := elm.GroupCode == "BG"
					scdGroupc := elm.GroupCode == "UP"
					trdGroupc := elm.GroupCode == "TS" && psglst.Airlfl != "SL"
					if fstGroupc || scdGroupc || trdGroupc {
						if len(rslDescae) > 0 {
							psglst.Paidbt += int32(nowPaidbt)
						} else {
							psglst.Paidbt += int32(elm.NumberOfItems)
						}
					}

				}
			}

			// Cek emd non exist
			if len(mapEmdnae) > 0 {
				fncGlobal.FncGlobalMainprNoterr(&psglst.Noterr, "EMD DETAIL NIL")
			}
		} else if len(mapEmdnae) > 0 {
			fncGlobal.FncGlobalMainprNoterr(&psglst.Noterr, "EMD DETAIL NIL")
		}
	}

	if psglst.Pnritl == "" {
		cekLstvar = true
	}

	// Get ticketing document
	if psglst.Tktnvc != "" {
		cekTcktng = true
		err := fncSbrapi.FncSbrapiGettktMainob(objtkn, airlfl, &psglst)
		if err != nil {
			psglst.Source += "|" + err.Error()
		} else {
			cekLstvar = true
		}

		// Check non revenue
		tmpNonrev := FncPsglstFrcalcSplitd(&psglst, mapCurrcv, sycMilege, objtkn)
		if tmpNonrev {
			cekNonrev = true
			if psglst.Isitnr == "" {
				psglst.Isitnr = "ZEROFB"
			}
		}
		psglst.Source += "|GETTKT"
	}

	// Check if data clear or not
	if cekChrter || !cekIsflwn || cekNonrev || cekLstvar || cekTcktng {
		istStlerr := true
		mapSuberr := map[string]bool{}
		fncFnlcek := func(params any, noterr, suberr string) {
			valFloatx, mtcFloatx := params.(float64)
			valString, mtcString := params.(string)
			if (mtcFloatx && valFloatx < 1000) || (mtcString && valString == "") {
				istStlerr = false
				mapSuberr[suberr] = true
				fncGlobal.FncGlobalMainprNoterr(&psglst.Noterr, noterr+" NIL")
			}
		}

		// Final check tktnfl
		fncFnlcek(psglst.Tktnvc, "TKTNVC", "MNFEST")
		fncFnlcek(psglst.Pnrcde, "PNRCDE", "MNFEST")
		fncFnlcek(psglst.Timeis, "TIMEIS", "MNFEST")
		fncFnlcek(psglst.Routvc, "ROUTVC", "MNFEST")
		fncFnlcek(psglst.Ntaffl, "NTAFFl", "SLSRPT")
		fncFnlcek(psglst.Ntafvc, "NTAFVC", "SLSRPT")
		fncFnlcek(psglst.Curncy, "CURNCY", "SLSRPT")
		if !istStlerr && cekIsflwn {
			for suberr := range mapSuberr {
				if suberr == "SLSRPT" && !cekNonrev {
					psglst.Slsrpt = "NOT CLEAR"
				}
				if suberr == "MNFEST" {
					psglst.Mnfest = "NOT CLEAR"
				}
			}
		}
		sycClrpsg.Store(psglst.Prmkey, psglst)
	} else {
		sycNulpsg.Store(psglst.Prmkey, psglst)
	}
}
