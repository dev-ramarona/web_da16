package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Comamand macro Sabre API Sreen
func FncSbrapiPsgdtaMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	clslvl map[string]mdlPsglst.MdlPsglstClsslvDtbase,
	psgdta *mdlPsglst.MdlPsglstPsgdtlDtbase,
) error {

	// Isi struktur data
	strDatefl := strconv.Itoa(int(psgdta.Datefl))
	rawDatefl, _ := time.Parse("060102", strDatefl)
	dmyDatefl := rawDatefl.Format("2006-01-02")
	bdyPsgdta := mdlSbrapi.MdlSbrapiPsgdtaReqenv{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdlSbrapi.MdlSbrapiPsgdtaReqhdr{
			MessageHeader: FncSbrapiMsghdrMainob(fncGlobal.Pcckey,
				"Get Passenger Data", "GetPassengerDataRQ", unqhdr),
			Security: mdlSbrapi.MdlSbrapiPsgdtaReqscr{
				BinarySecurityToken: mdlSbrapi.MdlSbrapiPsgdtaReqbst{
					ValueType:    "String",
					EncodingType: "wsse:Base64Binary",
					Token:        unqhdr.Bsttkn,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdlSbrapi.MdlSbrapiPsgdtaReqbdy{
			GetPassengerDataRQ: mdlSbrapi.MdlSbrapiPsgdtaReqafl{
				Xmlns:                       "http://services.sabre.com/checkin/getPassengerData/v4",
				ValidateCheckInRequirements: "true",
				Version:                     "4.0.4",
				ItineraryAndPassengerInfo: mdlSbrapi.MdlSbrapiPsgdtaReqiap{
					Itinerary: mdlSbrapi.MdlSbrapiPsgdtaReqitn{
						Airline:       psgdta.Airlfl,
						Flight:        psgdta.Flnbfl,
						DepartureDate: dmyDatefl,
						Origin:        psgdta.Depart,
					},
					Seat: psgdta.Seatpx,
				},
			},
		},
	}

	// Treatment APO Session
	raw, err := FncSbrapiMsghdrXmldta(bdyPsgdta)
	if err != nil {
		return err
	}

	// Parsing XML ke dalam struktur Go
	rspPsgdta := mdlSbrapi.MdlSbrapiPsgdtaRspenv{}
	err = xml.Unmarshal([]byte(raw), &rspPsgdta)
	if err != nil {
		return err
	}

	// Return String
	rawxml := rspPsgdta.Body.GetPassengerDataRS.PassengerDataResponse
	FncSbrapiPsgdtaTrtmnt(rawxml, clslvl, psgdta)
	return nil
}

// Treatment data raw flight list
func FncSbrapiPsgdtaTrtmnt(rawxml mdlSbrapi.MdlSbrapiPsgdtaRsppdr,
	clslvl map[string]mdlPsglst.MdlPsglstClsslvDtbase,
	psgdta *mdlPsglst.MdlPsglstPsgdtlDtbase) {
	if rawxml.PassengerID == "" {
		return
	}
	psgdta.Linenb = int32(rawxml.LineNumber)
	psgdta.Nmelst = rawxml.LastName
	psgdta.Nmefst = rawxml.FirstName
	psgdta.Psgrid = rawxml.PassengerID
	psgdta.Groupc = rawxml.GroupCode

	// Note edit PNR and Ticket
	if psgdta.Pnrcde == "" && rawxml.PNRLocator != "" {
		psgdta.Pnrcde = rawxml.PNRLocator
		fncGlobal.FncGlobalMainprNoterr(&psgdta.Noteup, "PNR SYSTEM")
	}
	if psgdta.Tktnfl == "" && rawxml.VCRNumber != "" {
		psgdta.Tktnfl = rawxml.VCRNumber
		fncGlobal.FncGlobalMainprNoterr(&psgdta.Noteup, "TKT SYSTEM")
	}

	// Get freetext info
	for _, freetx := range rawxml.FreeTextInfoList {
		switch freetx.EditCode {

		// Electronic ticket
		case "ET":
			partsl := strings.Fields(freetx.TextLine)
			if len(partsl) >= 6 {
				psgdta.Tktnvc = partsl[0]
				rawCpnbvc := strings.ReplaceAll(partsl[1], "C", "")
				intCpnbvc, _ := strconv.Atoi(rawCpnbvc)
				psgdta.Cpnbvc = int32(intCpnbvc)
				strDatevc := fncGlobal.FncGlobalMainprDaymnt(partsl[2])
				intDatevc, _ := strconv.Atoi(strDatevc)
				psgdta.Datevc = int32(intDatevc)
				psgdta.Clssvc = partsl[3]
				psgdta.Cbinvc = clslvl[partsl[3]].Cbinfl
				psgdta.Routvc = partsl[4][:3] + "-" + partsl[4][3:]
				psgdta.Statvc = partsl[5]
			}

			// Bagage
		case "BT":
			partsl := strings.Fields(freetx.TextLine)
			regmkg := regexp.MustCompile(`KG$`)
			if rslmkg := regmkg.FindAllString(partsl[0], -1); len(rslmkg) > 0 {
				istQtycek := true
				psgdta.Typebt = "KG"
				for _, prt := range partsl {
					regmnb := regexp.MustCompile(`\d+`)
					if rslmmb := regmnb.FindAllString(prt, -1); len(rslmmb) > 0 {
						if istQtycek {
							istQtycek = false
							intQntybt, _ := strconv.Atoi(rslmmb[0])
							psgdta.Qntybt = int32(intQntybt)
						} else {
							intWghtbt, _ := strconv.Atoi(rslmmb[0])
							psgdta.Wghtbt = int32(intWghtbt)
						}
					}
				}
			} else if len(partsl) == 3 {
				fncGlobal.FncGlobalMainprNoterr(&psgdta.Nmbrbt, partsl[2])
			}

		// Comment
		case "CM":
			fncGlobal.FncGlobalMainprNoterr(&psgdta.Coment, freetx.TextLine)

		// Outbound
		case "OB":
			partsl, rawtme, strdte := strings.Fields(freetx.TextLine), "", ""
			for _, prdata := range partsl {
				switch {
				case regexp.MustCompile(`^[A-Z]{2}|\*[A-Z]{2}$`).MatchString(prdata):
					psgdta.Airlob = strings.ReplaceAll(prdata, "*", "")
				case regexp.MustCompile(`^\d{1,2}[A-Z]{3}(\d{0,2})?$`).MatchString(prdata):
					if len(prdata) >= 6 {
						fmtObdate, _ := time.Parse("2Jan06", prdata)
						strdte = fmtObdate.Format("060102")
					} else if len(prdata) >= 4 {
						strdte = fncGlobal.FncGlobalMainprDaymnt(prdata)
					}
					intObdate, _ := strconv.Atoi(strdte)
					psgdta.Dateob = int32(intObdate)
				case regexp.MustCompile(`^ETD\d{4}$`).MatchString(prdata):
					rawtme = prdata[3:]
				case regexp.MustCompile(`^\d{1,4}$`).MatchString(prdata):
					psgdta.Flnbob = prdata
				case regexp.MustCompile(`^[A-Z]{3}$`).MatchString(prdata):
					if psgdta.Routob == "" {
						psgdta.Routob = prdata
					} else {
						psgdta.Routob += "-" + prdata
					}
				case regexp.MustCompile(`^[A-Z]$`).MatchString(prdata):
					psgdta.Clssob = prdata
				}
			}
			if rawtme != "" {
				intObtime, _ := strconv.Atoi(strdte + rawtme)
				psgdta.Timeob = int64(intObtime)
			}

		// Inbound
		case "IB":
			partsl := strings.Fields(freetx.TextLine)
			if len(partsl) >= 5 {
				psgdta.Airlib = partsl[0]
				psgdta.Flnbib = partsl[1]
				psgdta.Clssib = partsl[2]
				psgdta.Dstrib = partsl[3]
				strIbdate := fncGlobal.FncGlobalMainprDaymnt(partsl[4])
				intIbdate, _ := strconv.Atoi(strIbdate)
				psgdta.Dateib = int32(intIbdate)
				if len(partsl) == 6 {
					intIbdate, _ := strconv.Atoi(strIbdate + partsl[5][3:])
					psgdta.Timeib = int64(intIbdate)
				}
			}

		// Inbound
		case "IR":
			partsl := strings.Fields(freetx.TextLine)
			if len(partsl) >= 2 {
				psgdta.Codeir = partsl[0]
				partwo := strings.Split(partsl[1], "/")
				psgdta.Airlir = partwo[0][:2]
				psgdta.Flnbir = partwo[0][2:]
				strIrdate := fncGlobal.FncGlobalMainprDaymnt(partwo[1])
				intIrdate, _ := strconv.Atoi(strIrdate)
				psgdta.Dateir = int32(intIrdate)
			}

		// Infant electronic ticket
		case "IFET":
			if strings.Contains(freetx.TextLine, "IFET-") {
				partsl := strings.Fields(freetx.TextLine)
				psgdta.Tktnif = partsl[0][5:]
				rawCpnbif := strings.ReplaceAll(partsl[1], "C", "")
				intCpnbif, _ := strconv.Atoi(rawCpnbif)
				psgdta.Cpnbif = int32(intCpnbif)
				strIfdate := fncGlobal.FncGlobalMainprDaymnt(partsl[2])
				intIfdate, _ := strconv.Atoi(strIfdate)
				psgdta.Dateif = int32(intIfdate)
				psgdta.Clssif = partsl[3]
				psgdta.Routif = partsl[4]
				psgdta.Statif = partsl[5]
			} else {
				psgdta.Paxsif = freetx.TextLine
			}

		// Infant electronic ticket
		case "XT":
			partsl := strings.Fields(freetx.TextLine)
			psgdta.Airlxt = partsl[0]
			psgdta.Dstrxt = partsl[1]
			psgdta.Nmbrxt = partsl[2]
		}
	}
}
