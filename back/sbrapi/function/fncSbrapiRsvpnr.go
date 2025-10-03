package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
	"fmt"
)

// Get data Reservation PNR froms abre
func FncSbrapiRsvpnrMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	params mdlSbrapi.MdlSbrapiMsghdrApndix, sbarea []string) (
	mdlSbrapi.MdlSbrapiRsvpnrRsprsv, error) {

	// Isi struktur data
	rspRsvpnr := mdlSbrapi.MdlSbrapiRsvpnrRsprsv{}
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
				Locator:     params.Pnrcde,
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
	fmt.Println(string(raw))
	err = xml.Unmarshal([]byte(raw), &rspRsvpnr)
	if err != nil {
		return rspRsvpnr, err
	}

	// Final return data
	return rspRsvpnr, nil

}
