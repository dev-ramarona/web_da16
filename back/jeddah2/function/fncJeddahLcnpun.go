package fnc_jeddah

import (
	fnc_global "back/global/function"
	mdl_global "back/global/model"
	mdl_jeddah "back/jeddah/model"
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

// Get data enhanced Sabre API
func FncGlobalLcnpunApisbr(bstValue, dbsFlnbfl, dbsDepart, dbsRoutfl,
	dbsAirlfl, dbsDatefl, nowTimefl string) (
	[]mdl_jeddah.MdlJeddahLogpnrDtbase, []mdl_jeddah.MdlJeddahLogpnrDtbase) {
	sbrDaterw, _ := time.Parse("060102", dbsDatefl)
	sbrDatefl := strings.ToUpper(sbrDaterw.Format("02Jan"))

	// Isi struktur data
	bdyPun := mdl_global.ScreenReqEnvlpe{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdl_global.ScreenReqHeader{
			MessageHeader: fnc_global.FncGlobalApisbrMsghdr(fnc_global.Pcckey,
				"Send Sabre Command", "SabreCommandLLSRQ"),
			Security: mdl_global.ScreenReqSecrty{
				BinarySecurityToken: mdl_global.ScreenReqBsctkn{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: bstValue,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdl_global.ScreenReqBodyxx{
			SabreCommandLLSRQ: mdl_global.ScreenReqSbrcmd{
				Xmlns:        "http://webservices.sabre.com/sabreXML/2011/10",
				NumResponses: "1",
				Version:      "2.0.0",
				Request: mdl_global.ScreenReqRquest{
					Output:      "SCREEN",
					HostCommand: "PUN" + dbsFlnbfl + "/" + sbrDatefl + dbsDepart,
				},
			},
		},
	}

	// Isi struktur data
	bdyLlc := mdl_global.ScreenReqEnvlpe{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdl_global.ScreenReqHeader{
			MessageHeader: fnc_global.FncGlobalApisbrMsghdr(fnc_global.Pcckey,
				"Send Sabre Command", "SabreCommandLLSRQ"),
			Security: mdl_global.ScreenReqSecrty{
				BinarySecurityToken: mdl_global.ScreenReqBsctkn{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: bstValue,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdl_global.ScreenReqBodyxx{
			SabreCommandLLSRQ: mdl_global.ScreenReqSbrcmd{
				Xmlns:        "http://webservices.sabre.com/sabreXML/2011/10",
				NumResponses: "1",
				Version:      "2.0.0",
				Request: mdl_global.ScreenReqRquest{
					Output:      "SCREEN",
					HostCommand: "LC" + dbsFlnbfl + "/" + sbrDatefl + dbsDepart,
				},
			},
		},
	}

	// Treatment API Passenger List
	outpun := FncGlobalLcnpunTrtmnt(bdyPun, dbsFlnbfl, dbsDepart, dbsRoutfl,
		dbsAirlfl, dbsDatefl, nowTimefl, "pun")
	llcpun := FncGlobalLcnpunTrtmnt(bdyLlc, dbsFlnbfl, dbsDepart, dbsRoutfl,
		dbsAirlfl, dbsDatefl, nowTimefl, "lc")
	return outpun, llcpun
}

// Get data enhanced Sabre API
func FncGlobalLdnxxxApisbr(bstValue, dbsFlnbfl, dbsDepart, dbsRoutfl,
	dbsAirlfl, dbsDatefl, nowTimefl string) []mdl_jeddah.MdlJeddahLogpnrDtbase {
	sbrDaterw, _ := time.Parse("060102", dbsDatefl)
	sbrDatefl := strings.ToUpper(sbrDaterw.Format("02Jan"))

	// Isi struktur data
	bdyPun := mdl_global.ScreenReqEnvlpe{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdl_global.ScreenReqHeader{
			MessageHeader: fnc_global.FncGlobalApisbrMsghdr(fnc_global.Pcckey, "Send Sabre Command", "SabreCommandLLSRQ"),
			Security: mdl_global.ScreenReqSecrty{
				BinarySecurityToken: mdl_global.ScreenReqBsctkn{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: bstValue,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdl_global.ScreenReqBodyxx{
			SabreCommandLLSRQ: mdl_global.ScreenReqSbrcmd{
				Xmlns:        "http://webservices.sabre.com/sabreXML/2011/10",
				NumResponses: "1",
				Version:      "2.0.0",
				Request: mdl_global.ScreenReqRquest{
					Output:      "SCREEN",
					HostCommand: "LDN" + dbsFlnbfl + "/" + sbrDatefl + dbsDepart,
				},
			},
		},
	}

	// Treatment API Passenger List
	outpun := FncGlobalLcnpunTrtmnt(bdyPun, dbsFlnbfl, dbsDepart,
		dbsRoutfl, dbsAirlfl, dbsDatefl, nowTimefl, "ldn")
	return outpun
}

// Function Treatment for API LC AND PUN
func FncGlobalLcnpunTrtmnt(bdyLcnpun mdl_global.ScreenReqEnvlpe,
	dbsFlnbfl, dbsDepart, dbsRoutfl, dbsAirlfl, dbsDatefl, nowTimefl, lcrpun string,
) []mdl_jeddah.MdlJeddahLogpnrDtbase {

	// Declare first output
	var rawLcnpun mdl_global.ScreenRspEnvlpe
	var results []mdl_jeddah.MdlJeddahLogpnrDtbase

	// Baca respons
	rspLcnpun, err := fnc_global.FncGlobalApisbrXmldta(bdyLcnpun)
	if err != nil {
		return results
	}

	// Parsing XML ke dalam struktur Go
	err = xml.Unmarshal([]byte(rspLcnpun), &rawLcnpun)
	if err != nil {
		return results
	}

	outlne := strings.Split(rawLcnpun.Body.SabreCommandLLSRS.Response, "\n")
	for _, outrow := range outlne {
		if len(outrow) <= 6 {
			continue
		}
		slcrow := strings.Split(outrow, ".")
		clnrow := []string{}
		for _, row := range slcrow {
			if strings.TrimSpace(row) != "" {
				clnrow = append(clnrow, row)
			}
		}

		// Conver to number
		intDatefl, _ := strconv.Atoi(dbsDatefl)
		intDateup, _ := strconv.Atoi(nowTimefl[0:6])
		intTimeup, _ := strconv.Atoi(nowTimefl)

		// Push to database LC AND PUN
		if len(clnrow) >= 3 {
			totpax, _ := strconv.Atoi(strings.TrimSpace(clnrow[0][3:6]))
			agtnme := strings.TrimSpace(clnrow[0][6:len(clnrow[0])])
			pnrcde := clnrow[2]
			clssfl := clnrow[1][:1]
			if len(clnrow) == 4 {
				pnrcde = clnrow[3]
				clssfl = clnrow[2]
			}
			results = append(results, mdl_jeddah.MdlJeddahLogpnrDtbase{
				Prmkey: lcrpun + dbsAirlfl + dbsFlnbfl + dbsDepart + dbsDatefl + pnrcde + nowTimefl[0:6],
				Airlfl: dbsAirlfl,
				Lcrpun: lcrpun,
				Totpax: totpax,
				Flnbfl: dbsFlnbfl,
				Depart: dbsDepart,
				Routfl: dbsRoutfl,
				Clssfl: strings.TrimSpace(clssfl),
				Datefl: int32(intDatefl),
				Dateup: int32(intDateup),
				Timeup: int64(intTimeup),
				Agtnme: agtnme,
				// Agtdtl: "",
				// Agtidn: "",
				// Rtlsrs: "",
				Pnrcde: pnrcde,
			})
		}
	}

	// Return final data
	return results
}
