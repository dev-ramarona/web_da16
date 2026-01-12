package fncPsglst

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"
	fncSbrapi "back/sbrapi/function"
	mdlSbrapi "back/sbrapi/model"
	"fmt"
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
	var cekLstvar, cekChrter, cekIsflwn, cekNonrev, cekTcktng, cekFrcalc bool
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

	// Get reservation
	nowRsvpnr := mdlSbrapi.MdlSbrapiRsvpnrRsprsv{}
	if istTcktng, ist := sycPnrcde.Load(pnrcde + airlfl); ist {
		if mtcTcktng, mtc := istTcktng.(mdlSbrapi.MdlSbrapiRsvpnrRsprsv); mtc {
			nowRsvpnr = mtcTcktng
		}
	} else {
		slcSbarea := []string{"TICKETING", "PRICING_INFORMATION", "ITINERARY", "RECORD_LOCATOR"}
		getTcktng, err := fncSbrapi.FncSbrapiRsvpnrMainob(objtkn, pnrcde, slcSbarea)
		if err != nil {
			return
		}
		nowRsvpnr = getTcktng
		sycPnrcde.Store(pnrcde+airlfl, getTcktng)
	}

	// If data null
	if nowRsvpnr.BookingDetails.RecordLocator != "" {

		// Date formating PNR book PNR Create date
		varTimecr := nowRsvpnr.BookingDetails.SystemCreationTimestamp
		if pnrTimecr, err := time.Parse("2006-01-02T15:04:05", varTimecr); err == nil {
			rawTimerw, _ := strconv.Atoi(pnrTimecr.Format("0601021504"))
			psglst.Timecr = int64(rawTimerw)
			psglst.Agtdcr = nowRsvpnr.BookingDetails.CreationAgentID
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
		slcSegmnt := []string{}
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
				fmtSegmnt := fmt.Sprintf("%s-%s-%s-%s-MKT-%s-%s-%s-OPT-%s-%s-%s",
					rawDepart, rawArrivl, rawActncd, strTimefl,
					mktAirlfl, mktFlnbfl, mktClssfl,
					optAirlfl, optFlnbfl, optClssfl)
				lstArrivl = rawArrivl
				if idx <= 1 || len(slcRoutsg) == 0 ||
					slcRoutsg[len(slcRoutsg)-1] != rawDepart {
					slcSegmnt = append(slcSegmnt, fmtSegmnt)
					slcRoutsg = append(slcRoutsg, rawDepart)
				}
			}
			slcRoutsg = append(slcRoutsg, lstArrivl)
			psglst.Routsg = strings.Join(slcRoutsg, "-")
			psglst.Segmnt = strings.Join(slcSegmnt, "|")
		}
		psglst.Pnritl = strings.Join(slcPnrtil, "|")

		// Check last variable
		if !cekLstvar {
			cekLstvar = psglst.Pnritl == ""
			totPnritl := 0
			for _, pnritl := range slcPnrtil {
				if pnritl != "" {
					if slices.Contains([]string{"JT", "ID", "IW", "IU", "OD", "SL"}, pnritl[:2]) {
						totPnritl += 1
					}
				}
			}
			cekLstvar = totPnritl == 0
		}

		// Get ticketing detail for issued date
		var slcTcktng = nowRsvpnr.PassengerReservation.TicketingInfo.TicketDetails
		if len(slcTcktng) != 0 {
			for _, tcktng := range slcTcktng {

				// Logical gate for ticket number
				nowLogicg := tcktng.TicketNumber[:13] == psglst.Tktnfl ||
					tcktng.TicketNumber[:13] == psglst.Tktnvc
				if psglst.Tktnfl == "" && psglst.Tktnvc == "" {
					strFmtnme := (psglst.Nmelst + "     ")[:5]
					strLstnme := (psglst.Nmefst + " ")[:1]
					cncFulln1 := strFmtnme + "/" + strLstnme
					cncFulln2 := psglst.Nmelst + "/" + strLstnme
					if (cncFulln1 == tcktng.PassengerName ||
						cncFulln2 == tcktng.PassengerName) &&
						tcktng.TicketNumber[3:4] != "4" {
						nowLogicg = true
					}
				}

				// Get issued date and agent die
				if nowLogicg {

					// Get agent die
					getAgtlct := tcktng.AgencyLocation
					getAgtdcd := tcktng.DutyCode
					getAgtsne := tcktng.AgentSine
					psglst.Agtdie = getAgtlct + getAgtdcd + getAgtsne

					// Get ticket number blank
					if psglst.Tktnfl == "" && psglst.Tktnvc == "" {
						psglst.Tktnfl = tcktng.TicketNumber[:13]
						psglst.Tktnvc = tcktng.TicketNumber[:13]
					}

					// get time issued
					fmtDateis, _ := time.Parse("2006-01-02T15:04:05", tcktng.Timestamp)
					intDateis, _ := strconv.Atoi(fmtDateis.Add(12 * time.Hour).Format("0601021504"))
					psglst.Timeis = int64(intDateis)
					cekTcktng = true

				}
			}
		}

		// Fare pricing select equal with type pax
		var slcPrcing = nowRsvpnr.PassengerReservation.ItineraryPricing.PricedItinerary
		var nowPrcing = mdlSbrapi.MdlSbrapiRsvpnrRspti{}
		for _, prcing := range slcPrcing {
			if prcing.PTC_FareBreakdown.PassengerTypeQuantity.Code == psglst.Typepx {
				nowPrcing = prcing
				break
			}
		}
		if nowPrcing.PTC_FareBreakdown.FareCalc != "" {

			// Get currency
			if curncy := nowPrcing.ItinTotalFare.CurrencyCode; curncy != "" {
				psglst.Curncy = curncy
			}

			// Looping FlightSegment
			mapFlnbvc := map[string]int{}
			mapFlsgmn := map[string]mdlPsglst.MdlPsglstPsgdtlDtbase{}
			slcFlsgmn := nowPrcing.PTC_FareBreakdown.FlightSegment
			for i := 0; i < len(slcFlsgmn); i++ {
				rawDatefl, _ := time.Parse(slcFlsgmn[i].DepartureDateTime, "2006-01-02T15:04:05")
				datefl := rawDatefl.Format("060102")
				flnbvc := slcFlsgmn[i].FlightNumber
				depart := slcFlsgmn[i].AirPort
				arrivl := ""
				routvc := ""
				if i < len(slcFlsgmn)-1 {
					arrivl = slcFlsgmn[i+1].AirPort
					routvc = depart + "-" + arrivl
				}
				keyseg := datefl + depart + flnbvc
				mapFlnbvc[keyseg] += 1
				mapFlsgmn[keyseg] = mdlPsglst.MdlPsglstPsgdtlDtbase{
					Flnbvc: flnbvc, Depart: depart, Arrivl: arrivl,
					Routvc: routvc, Cpnbvc: int32(slcFlsgmn[i].RPH), Frbcde: slcFlsgmn[i].FareBasisCode}

				// counter value
				if depart != psglst.Arrivl && arrivl != psglst.Depart {
					if depart == psglst.Depart || arrivl == psglst.Arrivl ||
						arrivl == psglst.Depart {
						mapFlnbvc[keyseg] += 5
						if depart == psglst.Depart && arrivl == psglst.Arrivl {
							mapFlnbvc[keyseg] += 15
						}
					}
					if len(psglst.Routvc) >= 7 {
						if depart == psglst.Routvc[:3] ||
							arrivl == psglst.Routvc[4:] || arrivl == psglst.Routvc[:3] {
							mapFlnbvc[keyseg] += 5
							if depart == psglst.Routvc[:3] && arrivl == psglst.Routvc[4:] {
								mapFlnbvc[keyseg] += 15
							}
						}
					}
				}
			}

			// Get data vcr
			hghest := struct {
				key string
				val int
			}{}
			for keyseg, valint := range mapFlnbvc {
				if hghest.val == 0 || hghest.val < valint {
					hghest.key = keyseg
					hghest.val = valint
				}
			}
			if getFlsgmn, ist := mapFlsgmn[hghest.key]; ist {
				psglst.Flnbvc = getFlsgmn.Flnbvc
				psglst.Airlvc = getFlsgmn.Airlvc
				psglst.Frbcde = getFlsgmn.Frbcde
				if psglst.Frbcde == "CHARTER" {
					psglst.Isitct = "CT"
				}

				// Make sure route VCR wrong and get new route VCR
				if len(getFlsgmn.Routvc) == 7 && (psglst.Routvc == "" ||
					psglst.Depart == psglst.Routvc[4:] ||
					psglst.Arrivl == psglst.Routvc[:3]) {
					psglst.Routvc = getFlsgmn.Routvc
				} else if psglst.Routvc == "" {
					cekFrcalc = false
				}

				// Get coupon if empty
				if psglst.Cpnbfl == 0 || psglst.Cpnbvc == 0 {
					psglst.Cpnbvc = getFlsgmn.Cpnbvc
					psglst.Cpnbfl = getFlsgmn.Cpnbvc
				}
			}

			// Get ntafvc
			psglst.Frcalc = nowPrcing.PTC_FareBreakdown.FareCalc
			tmpNonrev := FncPsglstFrcalcSplitd(&psglst, mapCurrcv, sycMilege, objtkn)
			if tmpNonrev {
				cekNonrev, psglst.Isitnr = true, "NONREV"
			}
			cekFrcalc = true
		}
	}

	// Check final route vcr is dif or not
	if len(psglst.Routvc) == 7 {
		if psglst.Depart != psglst.Routvc[:3] &&
			psglst.Arrivl != psglst.Routvc[4:] {
			cekFrcalc = false
		}
	}

	// Check pnr is isset
	if cekIsflwn && psglst.Pnrcde == "" {
		cekLstvar = true
		if psglst.Tktnfl != "" {
			cekIsflwn = true
		}
	}

	// Get ticketing document
	if cekTcktng && cekIsflwn && !cekFrcalc {
		err := fncSbrapi.FncSbrapiGettktMainob(objtkn, airlfl, &psglst)
		if err != nil {
			psglst.Source += "|" + err.Error()
		} else {
			cekLstvar = true
		}

		// Check non revenue
		tmpNonrev := FncPsglstFrcalcSplitd(&psglst, mapCurrcv, sycMilege, objtkn)
		if tmpNonrev {
			cekNonrev, psglst.Isitnr = true, "NONREV"
		}
		psglst.Source += "|GETTKT"
	}

	// Check if data clear or not
	if cekChrter || !cekIsflwn || cekNonrev || cekLstvar || (cekTcktng && cekFrcalc) {
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
		fncFnlcek(psglst.Tktnfl, "TKTNFL", "MNFEST")
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
