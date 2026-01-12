package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
)

// Get data Reservation PNR froms abre
func FncSbrapiRsvpnrMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	pnrcde string, sbarea []string) (
	mdlSbrapi.MdlSbrapiRsvpnrRsprsv, error) {

	// Isi struktur data
	rspRsvpnr := mdlSbrapi.MdlSbrapiRsvpnrRsprsv{}
	rspEnvpnr := mdlSbrapi.MdlSbrapiRsvpnrRspenv{}
	bdyRsvpnr := mdlSbrapi.MdlSbrapiRsvpnrReqenv{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdlSbrapi.MdlSbrapiRsvpnrReqhdr{
			MessageHeader: FncSbrapiMsghdrMainob(fncGlobal.Pcckey, "Retrieve Itinerary", "GetReservationRQ", unqhdr),
			Security: mdlSbrapi.MdlSbrapiRsvpnrReqscr{
				BinarySecurityToken: mdlSbrapi.MdlSbrapiRsvpnrReqbst{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: unqhdr.Bsttkn,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdlSbrapi.MdlSbrapiRsvpnrReqbdy{
			GetReservationRQ: mdlSbrapi.MdlSbrapiRsvpnrReqrsv{
				Xmlns:       "http://webservices.sabre.com/pnrbuilder/v1_19",
				Version:     "1.19.0",
				Locator:     pnrcde,
				RequestType: "Stateless",
				ReturnOptions: mdlSbrapi.MdlSbrapiRsvpnrReqopt{
					SubjectAreas:   mdlSbrapi.MdlSbrapiRsvpnrReqsbj{SubjectArea: sbarea},
					ViewName:       "Simple",
					ResponseFormat: "STL",
				},
			},
		},
	}

	// Treatment APO Session
	raw, err := FncSbrapiMsghdrXmldta(bdyRsvpnr)
	if err != nil {
		return rspRsvpnr, err
	}

	// Parsing XML ke dalam struktur Go
	err = xml.Unmarshal([]byte(raw), &rspEnvpnr)
	if err != nil {
		return rspRsvpnr, err
	}

	// Final return data
	rspRsvpnr = rspEnvpnr.Body.GetReservationRS.Reservation
	return rspRsvpnr, nil

}
