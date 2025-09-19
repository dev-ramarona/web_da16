package fnc_global

import (
	mdl_global "back/global/model"
	"encoding/xml"
	"fmt"
	"strings"
)

// Login Sabre and create Session API
func FncGlobalSssionApisbr(carrierCode string) (string, error) {

	// Isi struktur data
	bdySssion := mdl_global.SssionReqEnvlpe{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdl_global.SssionReqHeader{
			MessageHeader: FncGlobalApisbrMsghdr(Pcckey, "Create Session API", "SessionCreateRQ"),
			Security: mdl_global.SssionReqSecrty{
				UsernameToken: mdl_global.SssionReqUsrtkn{
					Username: Usrnme, Password: Psswrd,
					Organization: carrierCode, Domain: carrierCode,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdl_global.SssionReqBodyxx{
			SessionCreateRQ: mdl_global.SssionReqSsncrt{
				Version: "1.0.0",
				POS:     mdl_global.SssionReqPosxxx{Source: mdl_global.SssionReqSource{PseudoCityCode: Pcckey}},
				Xmlns:   "http://webservices.sabre.com",
			},
		},
	}

	// Declare first output
	var bstSssion string

	// Read response
	rspSession, err := FncGlobalApisbrXmldta(bdySssion)
	if err != nil {
		return bstSssion, err
	}

	// Parsing XML ke dalam struktur Go
	var envlpeSssion mdl_global.SssionRspEnvlpe
	err = xml.Unmarshal([]byte(rspSession), &envlpeSssion)
	if err != nil {
		return bstSssion, err
	}

	// Return non error data
	bstSssion = envlpeSssion.Header.Security.BinarySecurityToken.Token
	return bstSssion, nil

}

// Get multiple session
func FncGlobalSssionMultpl(airline string, count int) ([]string, error) {
	sessions := make([]string, 0, count)
	for i := 0; i < count; i++ {
		token, err := FncGlobalSssionApisbr(airline)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, token)
	}
	return sessions, nil
}

// Login Sabre and create Session API
func FncGlobalClsssnUsrsbr(bstValue string) error {

	// Isi struktur data
	bdyClsssn := mdl_global.ClsssnReqEnvlpe{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdl_global.ClsssnReqHeader{
			MessageHeader: FncGlobalApisbrMsghdr(Pcckey, "Close Session", "SessionCloseRQ"),
			Security: mdl_global.ClsssnReqSecrty{
				BinarySecurityToken: mdl_global.ClsssnReqBsctkn{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: bstValue,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdl_global.ClsssnReqBodyxx{
			SessionCloseRQ: mdl_global.ClsssnReqFllist{
				POS: mdl_global.ClsssnReqPosxxx{
					Source: mdl_global.ClsssnReqSource{
						PseudoCityCode: Pcckey,
					},
				},
			},
		},
	}

	// Treatment APO Session
	fnl, err := FncGlobalApisbrXmldta(bdyClsssn)
	if !strings.Contains(string(fnl), `status="Approved"`) {
		if err == nil {
			err = fmt.Errorf("Tidak Approve")
		}
	}
	return err
}

// Get multiple session
func FncGlobalClsssnMultpl(sssion []string) {
	for _, s := range sssion {
		FncGlobalClsssnUsrsbr(s)
	}
}
