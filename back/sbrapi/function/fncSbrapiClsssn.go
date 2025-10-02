package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlSbrapi "back/sbrapi/model"
	"fmt"
	"strings"
)

// Login Sabre and create Session API
func FncSbrapiClsssnMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams) error {

	// Isi struktur data
	bdyClsssn := mdlSbrapi.MdlSbrapiClsssnReqenv{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdlSbrapi.MdlSbrapiClsssnReqhdr{
			MessageHeader: FncSbrapiMsghdrMainob(fncGlobal.Pcckey, "Close Session", "SessionCloseRQ", unqhdr),
			Security: mdlSbrapi.MdlSbrapiClsssnReqscr{
				BinarySecurityToken: mdlSbrapi.MdlSbrapiClsssnReqbst{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: unqhdr.Bsttkn,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdlSbrapi.MdlSbrapiClsssnReqbdy{
			SessionCloseRQ: mdlSbrapi.MdlSbrapiClsssnReqssc{
				POS: mdlSbrapi.MdlSbrapiClsssnReqpos{
					Source: mdlSbrapi.MdlSbrapiClsssnReqsrc{
						PseudoCityCode: fncGlobal.Pcckey,
					},
				},
			},
		},
	}

	// Treatment APO Session
	fnl, err := FncSbrapiMsghdrXmldta(bdyClsssn)
	if !strings.Contains(string(fnl), `status="Approved"`) {
		if err == nil {
			err = fmt.Errorf("tidak Approve")
		}
	}
	return err
}

// Get multiple session
func FncSbrapiClsssnMultpl(arrUnqhdr []mdlSbrapi.MdlSbrapiMsghdrParams) {
	for _, unqhdr := range arrUnqhdr {
		FncSbrapiClsssnMainob(unqhdr)
	}
}
