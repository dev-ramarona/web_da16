package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlSbrapi "back/sbrapi/model"
)

// Comamand macro Sabre API Sreen
func FncSbrapiCmdxmlMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams, cmmand string) ([]byte, error) {

	// Isi struktur data
	bdyCmdxml := mdlSbrapi.MdlSbrapiCmdxmlReqenv{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdlSbrapi.MdlSbrapiCmdxmlReqhdr{
			MessageHeader: FncSbrapiMsghdrMainob(fncGlobal.Pcckey,
				"Send Sabre Command", "SabreCommandLLSRQ", unqhdr),
			Security: mdlSbrapi.MdlSbrapiCmdxmlReqscr{
				BinarySecurityToken: mdlSbrapi.MdlSbrapiCmdxmlReqbst{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: unqhdr.Bsttkn,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdlSbrapi.MdlSbrapiCmdxmlReqbdy{
			SabreCommandLLSRQ: mdlSbrapi.MdlSbrapiCmdxmlReqssc{
				Xmlns:        "http://webservices.sabre.com/sabreXML/2011/10",
				NumResponses: "1",
				Version:      "2.0.0",
				Request: mdlSbrapi.MdlSbrapiCmdxmlReqreq{
					Output:      "SDSXML",
					HostCommand: cmmand,
				},
			},
		},
	}

	// Treatment APO Session
	raw, err := FncSbrapiMsghdrXmldta(bdyCmdxml)
	if err != nil {
		return []byte{}, err
	}

	// Return String
	return raw, nil
}
