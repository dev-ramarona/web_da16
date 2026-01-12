package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
	"regexp"
	"strconv"
	"strings"
)

// Get data Reservation PNR froms abre
func FncSbrapiMilegeMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	routac string) ([]mdlPsglst.MdlPsglstMilegeDtbase, error) {

	// Breakdownd route actual
	strDepart, prvDepart := "", ""
	slcArrivl := []mdlSbrapi.MdlSbrapiMilegeReqdlc{}
	slcMilege := []mdlPsglst.MdlPsglstMilegeDtbase{}
	slcDepart := []string{}
	for idx, dstrct := range strings.Split(routac, "-") {
		if idx == 0 {
			strDepart = dstrct
		} else {
			slcDepart = append(slcDepart, prvDepart+"-"+dstrct)
			slcArrivl = append(slcArrivl, mdlSbrapi.MdlSbrapiMilegeReqdlc{
				LocationCode: dstrct, RPH: strconv.Itoa(idx)})
		}
		prvDepart = dstrct
	}

	// Isi struktur data
	rspEnvpnr := mdlSbrapi.MdlSbrapiMilegeRspenv{}
	bdyMilege := mdlSbrapi.MdlSbrapiMilegeReqenv{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdlSbrapi.MdlSbrapiMilegeReqhdr{
			MessageHeader: FncSbrapiMsghdrMainob(fncGlobal.Pcckey, "Calculate Air Mileage",
				"MileageLLSRQ", unqhdr),
			Security: mdlSbrapi.MdlSbrapiMilegeReqscr{
				BinarySecurityToken: mdlSbrapi.MdlSbrapiMilegeReqbst{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: unqhdr.Bsttkn,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdlSbrapi.MdlSbrapiMilegeReqbdy{
			MileageRQ: mdlSbrapi.MdlSbrapiMilegeReqmlq{
				Xmlns:   "http://webservices.sabre.com/sabreXML/2011/10",
				Version: "2.0.0",
				OriginDestinationInformation: mdlSbrapi.MdlSbrapiMilegeReqodi{
					DestinationLocation: slcArrivl,
					OriginLocation: mdlSbrapi.MdlSbrapiMilegeReqolc{
						LocationCode: strDepart,
					},
				},
			},
		},
	}

	// Treatment APO Session
	raw, err := FncSbrapiMsghdrXmldta(bdyMilege)
	if err != nil {
		return slcMilege, err
	}

	// Parsing XML ke dalam struktur Go
	err = xml.Unmarshal([]byte(raw), &rspEnvpnr)
	if err != nil {
		return slcMilege, err
	}
	for i, v := range rspEnvpnr.Body.MileageRS.OriginDestinationInformation.DestinationLocation {
		nowRoutfl := slcDepart[i]
		intRegexp := regexp.MustCompile(`\D`)
		intRepstr := intRegexp.ReplaceAllString(v.TicketedPointMileage, "")
		intMilege, _ := strconv.Atoi(intRepstr)
		slcMilege = append(slcMilege, mdlPsglst.MdlPsglstMilegeDtbase{
			Routfl: nowRoutfl, Milege: int64(intMilege)})
	}
	return slcMilege, nil
}
