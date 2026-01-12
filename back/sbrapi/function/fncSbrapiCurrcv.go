package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
)

// Comamand macro Sabre API Sreen
func FncSbrapiCurrcvMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
) (map[string]mdlPsglst.MdlPsglstCurrcvDtbase, error) {

	// Isi struktur data
	fnlCurrcv := map[string]mdlPsglst.MdlPsglstCurrcvDtbase{}
	bdyCurrcv := mdlSbrapi.MdlSbrapiCurrcvReqenv{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdlSbrapi.MdlSbrapiCurrcvReqhdr{
			MessageHeader: FncSbrapiMsghdrMainob(fncGlobal.Pcckey,
				"Get Currency Conversion", "DisplayCurrencyLLSRQ", unqhdr),
			Security: mdlSbrapi.MdlSbrapiCurrcvReqscr{
				BinarySecurityToken: mdlSbrapi.MdlSbrapiCurrcvReqbst{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: unqhdr.Bsttkn,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdlSbrapi.MdlSbrapiCurrcvReqbdy{
			DisplayCurrencyRQ: mdlSbrapi.MdlSbrapiCurrcvReqafl{
				Xmlns:        "http://webservices.sabre.com/sabreXML/2011/10",
				Version:      "2.1.0",
				CountryCode:  "ID",
				CurrencyCode: "CUR",
			},
		},
	}

	// Treatment APO Session
	raw, err := FncSbrapiMsghdrXmldta(bdyCurrcv)
	if err != nil {
		return fnlCurrcv, err
	}

	// Parsing XML ke dalam struktur Go
	rspCurrcv := mdlSbrapi.MdlSbrapiCurrcvRspenv{}
	err = xml.Unmarshal([]byte(raw), &rspCurrcv)
	if err != nil {
		return fnlCurrcv, err
	}

	// Return String
	rawCurrcv := rspCurrcv.Body.DisplayCurrencyRS
	fnlCurrcv = FncSbrapiCurrcvTrtmnt(rawCurrcv)
	return fnlCurrcv, nil
}

// Treatment data raw flight list
func FncSbrapiCurrcvTrtmnt(rawxml mdlSbrapi.MdlSbrapiCurrcvRspdsp,
) map[string]mdlPsglst.MdlPsglstCurrcvDtbase {
	fnlCurrcv := map[string]mdlPsglst.MdlPsglstCurrcvDtbase{}

	// Looping all flight list
	for _, currcv := range rawxml.Country {
		fnlCurrcv[currcv.CurrencyCode] = mdlPsglst.MdlPsglstCurrcvDtbase{
			Crctry: rawxml.Name,
			Crcode: currcv.CurrencyCode,
			Crname: currcv.CurrencyName,
			Crrate: currcv.Rate,
		}
	}

	// Final return
	return fnlCurrcv

}
