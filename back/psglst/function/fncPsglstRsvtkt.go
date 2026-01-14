package fncPsglst

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"
	fncSbrapi "back/sbrapi/function"
	mdlSbrapi "back/sbrapi/model"
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
	if psglst.Tktnvc == "" {

		// Get reservation
		nowRsvpnr := mdlSbrapi.MdlSbrapiRsvpnrRsprsv{}
		if istTcktng, ist := sycPnrcde.Load(pnrcde + airlfl); ist {
			if mtcTcktng, mtc := istTcktng.(mdlSbrapi.MdlSbrapiRsvpnrRsprsv); mtc {
				nowRsvpnr = mtcTcktng
			}
		} else {
			slcSbarea := []string{"TICKETING", "ITINERARY", "RECORD_LOCATOR"}
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
			if len(slcItinry) != 0 {
				for _, itinry := range slcItinry {
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
				}
			}
			psglst.Pnritl = strings.Join(slcPnrtil, "|")

			// Get ticketing detail for issued date
			var slcTcktng = nowRsvpnr.PassengerReservation.TicketingInfo.TicketDetails
			if len(slcTcktng) != 0 {
				for _, tcktng := range slcTcktng {

					// Logical gate for ticket number
					nowLogicg := tcktng.TicketNumber[:13] == psglst.Tktnfl
					if psglst.Tktnfl == "" {
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

					// Get ticket number blank
					if psglst.Tktnvc == "" && nowLogicg {
						psglst.Tktnvc = tcktng.TicketNumber[:13]
					}
				}
			}
		}
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
