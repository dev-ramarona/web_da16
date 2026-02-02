package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Get data Reservation PNR froms abre
func FncSbrapiGettktMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	airlfl string, psglst *mdlPsglst.MdlPsglstPsgdtlDtbase) error {

	// Isi struktur data
	rspEnvpnr := mdlSbrapi.MdlSbrapiGettktRspenv{}
	bdyGettkt := mdlSbrapi.MdlSbrapiGettktReqenv{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdlSbrapi.MdlSbrapiGettktReqhdr{
			MessageHeader: FncSbrapiMsghdrMainob(fncGlobal.Pcckey, "Get Ticket Doc Details",
				"TicketingDocumentServicesRQ", unqhdr),
			Security: mdlSbrapi.MdlSbrapiGettktReqscr{
				BinarySecurityToken: mdlSbrapi.MdlSbrapiGettktReqbst{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: unqhdr.Bsttkn,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdlSbrapi.MdlSbrapiGettktReqbdy{
			GetTicketingDocumentRQ: mdlSbrapi.MdlSbrapiGettktReqrsv{
				XmlnsNs2: "http://www.sabre.com/ns/Ticketing/DC",
				Xmlns:    "http://services.sabre.com/STL/v01",
				Version:  "3.28.3",
				POS:      struct{}{},
				SearchParameters: mdlSbrapi.MdlSbrapiGettktReqspm{
					ResultType:        "C",
					TicketingProvider: airlfl,
					DocumentNumber:    psglst.Tktnvc,
					CustomResponseDetails: []string{
						"ServiceCoupon",
						"FareCalculation",
						"Amounts",
					},
				},
			},
		},
	}

	// Treatment APO Session
	raw, err := FncSbrapiMsghdrXmldta(bdyGettkt)
	if err != nil {
		return err
	}

	// Parsing XML ke dalam struktur Go
	err = xml.Unmarshal([]byte(raw), &rspEnvpnr)
	if err != nil {
		return err
	}

	// Final return data
	mapCountr := map[string]int{}
	mapGettkt := map[string]mdlSbrapi.MdlSbrapiGettktRspsvc{}
	getTktdoc := rspEnvpnr.Body.GetTicketingDocumentRS.CustomDetails

	// Check if data isset
	if getTktdoc.Ticket.FareCalculation == "" {
		return errors.New("TICKET DATA NIL")
	}

	// Get time format
	strTimecr := getTktdoc.Ticket.Details.Reservation.CreateDate
	if fmtTimecr, err := time.Parse("2006-01-02T15:04:05", strTimecr); err == nil {
		rawTimerw, _ := strconv.Atoi(fmtTimecr.Format("0601021504"))
		if psglst.Timecr == 0 {
			psglst.Timecr = int64(rawTimerw)
		}
	}
	strTimeis := getTktdoc.Ticket.Details.LocalIssueDateTime
	if fmtTimeis, err := time.Parse("2006-01-02T15:04:05", strTimeis); err == nil {
		rawTimerw, _ := strconv.Atoi(fmtTimeis.Format("0601021504"))
		if psglst.Timeis == 0 {
			psglst.Timeis = int64(rawTimerw)
		}
	}

	// Looping and scoring per coupon
	slcSegtkt, slcRoutvf, lstArrivl := []string{}, []string{}, ""
	for _, cpn := range getTktdoc.Ticket.ServiceCoupon {
		nowDepart, nowArrivl := cpn.StartLocation, cpn.EndLocation
		keyGettkt := nowDepart + nowArrivl + cpn.MarketingFlightNumber
		mapCountr[keyGettkt] += 1
		mapGettkt[keyGettkt] = cpn
		flwDepart, flwArrivl := cpn.FlownCoupon.DepartureCity, cpn.FlownCoupon.ArrivalCity
		if nowArrivl != psglst.Depart && nowDepart != psglst.Arrivl {

			// Compare to flown data
			if nowDepart == psglst.Depart || nowArrivl == psglst.Arrivl {
				mapCountr[keyGettkt] += 5
				if nowDepart == psglst.Depart && nowArrivl == psglst.Arrivl {
					mapCountr[keyGettkt] += 15
				}
			}

			// Compare to route vcr
			if len(psglst.Routvc) >= 7 {
				if nowDepart == psglst.Routvc[:3] || nowArrivl == psglst.Routvc[4:] {
					mapCountr[keyGettkt] += 5
					if nowDepart == psglst.Routvc[:3] && nowArrivl == psglst.Routvc[4:] {
						mapCountr[keyGettkt] += 15
					}
				}
				if flwDepart == psglst.Routvc[:3] || flwArrivl == psglst.Routvc[4:] {
					mapCountr[keyGettkt] += 5
					if flwDepart == psglst.Routvc[:3] && flwArrivl == psglst.Routvc[4:] {
						mapCountr[keyGettkt] += 15
					}
				}
			}
		}

		// Get time flown
		rawTimefl := cpn.StartDateTime
		fmtTimefl, _ := time.Parse("2006-01-02T15:04:05", rawTimefl)
		strTimefl := fmtTimefl.Format("0601021504")
		rawTimend := cpn.EndDateTime
		fmtTimend, _ := time.Parse("2006-01-02T15:04:05", rawTimend)
		strTimend := fmtTimend.Format("0601021504")

		// Coupon segment
		rawDepart := cpn.StartLocation
		rawArrivl := cpn.EndLocation
		rawActncd := cpn.CurrentStatus
		mktAirlfl := cpn.MarketingProvider
		mktFlnbfl := cpn.MarketingFlightNumber
		mktClssfl := cpn.ClassOfService
		fmtSegtkt := fmt.Sprintf("%s:%s-%s-%s-%s-%s-%s-%s",
			strTimefl, strTimend, rawDepart, rawArrivl,
			rawActncd, mktAirlfl, mktFlnbfl, mktClssfl)
		lstArrivl = rawArrivl
		slcSegtkt = append(slcSegtkt, fmtSegtkt)
		slcRoutvf = append(slcRoutvf, rawDepart)
	}

	// Push other data
	slcRoutvf = append(slcRoutvf, lstArrivl)
	psglst.Routvf = strings.Join(slcRoutvf, "-")
	psglst.Segtkt = strings.Join(slcSegtkt, "|")
	psglst.Agtdie = getTktdoc.Agent.Duty + getTktdoc.Agent.Sine
	psglst.Frcalc = getTktdoc.Ticket.FareCalculation
	psglst.Curncy = getTktdoc.Ticket.Amounts.CurrencyCode
	psglst.Tourcd = getTktdoc.Ticket.Details.TourNumber
	psglst.Staloc = getTktdoc.Agent.StationLocation
	psglst.Stanbr = getTktdoc.Agent.StationNumber
	psglst.Wrkloc = getTktdoc.Agent.WorkLocation
	psglst.Hmeloc = getTktdoc.Agent.HomeLocation
	psglst.Lniata = getTktdoc.Agent.Lniata
	psglst.Emplid = getTktdoc.Agent.EmployeeNumber

	// Get data vcr
	hghest := struct {
		key string
		val int
	}{}
	for keyseg, valint := range mapCountr {
		if hghest.val == 0 || hghest.val < valint {
			hghest.key = keyseg
			hghest.val = valint
		}
	}
	if getFlsgmn, ist := mapGettkt[hghest.key]; ist {
		psglst.Flnbvc = getFlsgmn.MarketingFlightNumber
		psglst.Airlvc = getFlsgmn.MarketingProvider
		psglst.Frbcde = getFlsgmn.FareBasis
		psglst.Statvc = getFlsgmn.CurrentStatus
		regmnb := regexp.MustCompile(`\d+`)
		if strfba := getFlsgmn.BagAllowance; strfba != "" {
			if rslmmb := regmnb.FindAllString(strfba, -1); len(rslmmb) > 0 {
				intVfbabt, _ := strconv.Atoi(rslmmb[0])
				psglst.Fbavbt = int32(intVfbabt)
				if strings.Contains(strfba, "PC") {
					psglst.Fbavbt = int32(intVfbabt) * 23
				}
			}
		}
		if psglst.Frbcde == "CHARTER" {
			psglst.Isitct = "CT"
		}
		if psglst.Frbcde == "LOA" || psglst.Frbcde == "FOC" {
			psglst.Isitnr = "NONREV"
		}
		if len(getFlsgmn.StartLocation) == 3 && len(getFlsgmn.EndLocation) == 3 {
			psglst.Routvc = getFlsgmn.StartLocation + "-" + getFlsgmn.EndLocation
		}
		psglst.Cpnbvc = int32(getFlsgmn.Coupon)
	}
	return nil

}
