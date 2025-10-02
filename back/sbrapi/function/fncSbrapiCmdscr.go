package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
)

// Comamand macro Sabre API Sreen
func FncSbrapiCmdscrMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams, cmmand string) (string, error) {

	// Isi struktur data
	bdyCmdscr := mdlSbrapi.MdlSbrapiCmdscrReqenv{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdlSbrapi.MdlSbrapiCmdscrReqhdr{
			MessageHeader: FncSbrapiMsghdrMainob(fncGlobal.Pcckey,
				"Send Sabre Command", "SabreCommandLLSRQ", unqhdr),
			Security: mdlSbrapi.MdlSbrapiCmdscrReqscr{
				BinarySecurityToken: mdlSbrapi.MdlSbrapiCmdscrReqbst{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: unqhdr.Bsttkn,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdlSbrapi.MdlSbrapiCmdscrReqbdy{
			SabreCommandLLSRQ: mdlSbrapi.MdlSbrapiCmdscrReqssc{
				Xmlns:        "http://webservices.sabre.com/sabreXML/2011/10",
				NumResponses: "1",
				Version:      "2.0.0",
				Request: mdlSbrapi.MdlSbrapiCmdscrReqreq{
					Output:      "SCREEN",
					HostCommand: cmmand,
				},
			},
		},
	}

	// Treatment APO Session
	raw, err := FncSbrapiMsghdrXmldta(bdyCmdscr)
	if err != nil {
		return "", err
	}

	// Parsing XML ke dalam struktur Go
	rspCmdscr := mdlSbrapi.MdlSbrapiCmdscrRspenv{}
	err = xml.Unmarshal([]byte(raw), &rspCmdscr)
	if err != nil {
		return "", err
	}

	// Return String
	return rspCmdscr.Body.SabreCommandLLSRS.Response, nil
}
