package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
	"fmt"
)

// Comamand macro Sabre API Sreen
func FncSbrapiDspvcrMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	apndix mdlSbrapi.MdlSbrapiMsghdrApndix) (mdlSbrapi.MdlSbrapiDspvcrRsptkt, error) {

	// Isi struktur data
	bdyDspvcr := mdlSbrapi.MdlSbrapiDspvcrReqenv{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdlSbrapi.MdlSbrapiDspvcrReqhdr{
			MessageHeader: FncSbrapiMsghdrMainob(fncGlobal.Pcckey,
				"Get Ticket Info", "VCR_DisplayLLSRQ", unqhdr),
			Security: mdlSbrapi.MdlSbrapiDspvcrReqscr{
				BinarySecurityToken: mdlSbrapi.MdlSbrapiDspvcrReqbst{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: unqhdr.Bsttkn,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdlSbrapi.MdlSbrapiDspvcrReqbdy{
			VCR_DisplayRQ: mdlSbrapi.MdlSbrapiDspvcrReqafl{
				Xmlns:   "http://webservices.sabre.com/sabreXML/2011/10",
				Version: "2.2.2",
				SearchOptions: mdlSbrapi.MdlSbrapiDspvcrReqsrc{
					Ticketing: mdlSbrapi.MdlSbrapiDspvcrReqtkt{
						ETicketNumber: apndix.Tktnfl,
						OperatingAirline: mdlSbrapi.MdlSbrapiDspvcrReqopa{
							Code: apndix.Airlfl,
						},
					},
				},
			},
		},
	}

	// Treatment APO Session
	rawDspvcr := mdlSbrapi.MdlSbrapiDspvcrRsptkt{}
	raw, err := FncSbrapiMsghdrXmldta(bdyDspvcr)
	if err != nil {
		return rawDspvcr, err
	}
	fmt.Println(string(raw), unqhdr)

	// Parsing XML ke dalam struktur Go
	rspDspvcr := mdlSbrapi.MdlSbrapiDspvcrRspenv{}
	err = xml.Unmarshal([]byte(raw), &rspDspvcr)
	if err != nil {
		return rawDspvcr, err
	}

	// Return String
	rawDspvcr = rspDspvcr.Body.VCR_DisplayRS.TicketingInfo
	return rawDspvcr, nil
}
