package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
	"errors"
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
					DocumentNumber:    psglst.Tktnfl,
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
	getTktdoc := rspEnvpnr.Body.GetTicketingDocumentRS.CustomDetails.Ticket

	// Check if data isset
	if getTktdoc.FareCalculation == "" {
		return errors.New("TICKET DATA NIL")
	}

	psglst.Frcalc = getTktdoc.FareCalculation
	psglst.Curncy = getTktdoc.Amounts.CurrencyCode
	for _, cpn := range getTktdoc.ServiceCoupon {
		nowDepart, nowArrivl := cpn.StartLocation, cpn.EndLocation
		keyGettkt := nowDepart + nowArrivl + cpn.MarketingFlightNumber
		mapCountr[keyGettkt] += 1
		mapGettkt[keyGettkt] = cpn
		flwDepart, flwArrivl := cpn.FlownCoupon.DepartureCity, cpn.FlownCoupon.ArrivalCity
		if nowArrivl != psglst.Depart && nowDepart != psglst.Arrivl {
			if nowDepart == psglst.Depart || nowArrivl == psglst.Arrivl {
				mapCountr[keyGettkt] += 5
				if nowDepart == psglst.Depart && nowArrivl == psglst.Arrivl {
					mapCountr[keyGettkt] += 15
				}
			}
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
	}

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
		if len(getFlsgmn.StartLocation) == 3 && len(getFlsgmn.EndLocation) == 3 {
			psglst.Routvc = getFlsgmn.StartLocation + "-" + getFlsgmn.EndLocation
		}
		psglst.Cpnbvc = int32(getFlsgmn.Coupon)
		if psglst.Cpnbfl == 0 {
			psglst.Cpnbfl = int32(getFlsgmn.Coupon)
		}
	}
	return nil

}
