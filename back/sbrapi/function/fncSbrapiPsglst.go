package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

// Comamand macro Sabre API Sreen
func FncSbrapiPsglstMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	apndix mdlSbrapi.MdlSbrapiMsghdrApndix,
	mapcur map[string]mdlPsglst.MdlPsglstCurrcvDtbase,
	fllist mdlPsglst.MdlPsglstFllistDtbase,
	clslvl map[string]mdlPsglst.MdlPsglstClsslvDtbase,
) ([]mdlPsglst.MdlPsglstPsgdtlDtbase, error) {

	// Isi struktur data
	strDatefl := strconv.Itoa(int(apndix.Datefl))
	rawDatefl, _ := time.Parse("060102", strDatefl)
	ymdDatefl := rawDatefl.Format("2006-01-02T00:00:00")
	fnlPsglst := []mdlPsglst.MdlPsglstPsgdtlDtbase{}
	bdyPsglst := mdlSbrapi.MdlSbrapiPsglstReqenv{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdlSbrapi.MdlSbrapiPsglstReqhdr{
			MessageHeader: FncSbrapiMsghdrMainob(fncGlobal.Pcckey,
				"Get Passenger List", "GetPassengerListRQ", unqhdr),
			Security: mdlSbrapi.MdlSbrapiPsglstReqscr{
				BinarySecurityToken: mdlSbrapi.MdlSbrapiPsglstReqbst{
					ValueType:    "String",
					EncodingType: "wsse:Base64Binary",
					Token:        unqhdr.Bsttkn,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdlSbrapi.MdlSbrapiPsglstReqbdy{
			GetPassengerListRQ: mdlSbrapi.MdlSbrapiPsglstReqafl{
				Xmlns:         "http://services.sabre.com/checkin/getPassengerList/v4",
				Version:       "4.0.0",
				MessageID:     unqhdr.Mssgid,
				TimeStamp:     ymdDatefl,
				ServiceOption: "Stateless",
				Itinerary: mdlSbrapi.MdlSbrapiPsglstReqitn{
					Airline:       apndix.Airlfl,
					Flight:        apndix.Flnbfl,
					DepartureDate: ymdDatefl[:10],
					Origin:        apndix.Depart,
				},
				DisplayCodeRequest: mdlSbrapi.MdlSbrapiPsglstReqdcr{
					Condition: "OR",
					DisplayCodes: []string{"CM", "BT", "AE", "ET", "OB", "IB",
						"ON", "SS", "IR", "IFET", "XT", "DOCS"},
				},
			},
		},
	}

	// Treatment APO Session
	raw, err := FncSbrapiMsghdrXmldta(bdyPsglst)
	if err != nil {
		return fnlPsglst, err
	}

	// Parsing XML ke dalam struktur Go
	rspPsglst := mdlSbrapi.MdlSbrapiPsglstRspenv{}
	err = xml.Unmarshal([]byte(raw), &rspPsglst)
	if err != nil {
		return fnlPsglst, err
	}

	// Return String
	rawxml := rspPsglst.Body.GetPassengerListRS
	fnlPsglst = FncSbrapiPsglstTrtmnt(rawxml, apndix,
		mapcur, fllist, clslvl)
	return fnlPsglst, nil
}

// Treatment data raw flight list
func FncSbrapiPsglstTrtmnt(rawxml mdlSbrapi.MdlSbrapiPsglstRspgpl,
	apndix mdlSbrapi.MdlSbrapiMsghdrApndix,
	mapcur map[string]mdlPsglst.MdlPsglstCurrcvDtbase,
	fllist mdlPsglst.MdlPsglstFllistDtbase,
	clslvl map[string]mdlPsglst.MdlPsglstClsslvDtbase,
) []mdlPsglst.MdlPsglstPsgdtlDtbase {
	getItnrry := rawxml.ItineraryInfo.Itinerary
	getDatedp := rawxml.ItineraryInfo.DepartureArrival_Dates

	// Treatment date
	// rawDatefl, _ := time.Parse("2006-01-02", getDatedp.ScheduledDepartureDate)
	// strDatefl := rawDatefl.Format("060102")
	// strNdayfl := rawDatefl.Format("Mon")
	// intDatefl, _ := strconv.Atoi(strDatefl)
	// intMnthfl, _ := strconv.Atoi(strDatefl[:4])
	rawDaterv, _ := time.Parse("2006-01-02", getDatedp.ScheduledDepartureDate)
	strDaterv := rawDaterv.Format("060102")
	intDaterv, _ := strconv.Atoi(strDaterv)

	// Treatment time
	strDatefl := strconv.Itoa(int(apndix.Datefl))
	rawTimefl, _ := time.Parse("3:04PM", getDatedp.DepartureTime)
	strTimefl := rawTimefl.Format("1504")
	intTimefl, _ := strconv.Atoi(strDatefl + strTimefl)
	rawTimerv, _ := time.Parse("3:04PM", getDatedp.ArrivalTime)
	strTimerv := rawTimerv.Format("1504")
	intTimerv, _ := strconv.Atoi(strDaterv + strTimerv)

	// Treatment cabin total
	mapCbinfl := map[string]int{"Y": 0, "C": 0}
	for _, cbn := range rawxml.ItineraryInfo.CabinInfoList {
		mapCbinfl[cbn.Cabin] = cbn.Count
	}

	// Declare global variable
	tmpPsglst := mdlPsglst.MdlPsglstPsgdtlDtbase{
		Datefl: apndix.Datefl,
		Daterv: int32(intDaterv),
		Mnthfl: apndix.Mnthfl,
		Timefl: int64(intTimefl),
		Ndayfl: apndix.Ndayfl,
		Timerv: int64(intTimerv),
		Airlfl: getItnrry.Airline,
		Airtyp: getItnrry.AircraftType,
		Flnbfl: getItnrry.Flight,
		Depart: getItnrry.Origin,
		Flgate: rawxml.ItineraryInfo.DepartureGate,
		Bookdc: int32(mapCbinfl["C"]),
		Bookdy: int32(mapCbinfl["Y"])}

	// Looping all passangger list
	slcPsglst := []mdlPsglst.MdlPsglstPsgdtlDtbase{}
	mapQntybt := map[string]int{}
	mapWghtbt := map[string]int{}
	for _, psglst := range rawxml.PassengerInfoList {
		objPsglst := tmpPsglst

		// Store default params
		objPsglst.Arrivl = psglst.Destination
		objPsglst.Routmx = fllist.Routmx
		objPsglst.Linenb = int32(psglst.LineNumber)
		objPsglst.Seatpx = psglst.Seat
		objPsglst.Pnrcde = psglst.PNRLocator
		objPsglst.Tktnfl = psglst.VCRInfo.VCRData.SerialNumber
		objPsglst.Tktnvc = psglst.VCRInfo.VCRData.SerialNumber
		objPsglst.Psgrid = psglst.PassengerID
		objPsglst.Nmefst = psglst.NameDetails.FirstName
		objPsglst.Nmelst = psglst.NameDetails.LastName
		objPsglst.Cpnbfl = int32(psglst.VCRInfo.VCRData.CouponNumber)
		objPsglst.Cpnbvc = int32(psglst.VCRInfo.VCRData.CouponNumber)
		objPsglst.Clssfl = psglst.BookingClass
		objPsglst.Clssvc = psglst.BookingClass
		objPsglst.Cbinfl = psglst.Cabin
		objPsglst.Cbinvc = psglst.Cabin
		objPsglst.Qntybt = psglst.BagCount

		// Get group code and totpax
		objPsglst.Groupc = psglst.GroupCode
		regTotpax := regexp.MustCompile(`\d+`)
		rslTotpax := regTotpax.FindAllString(psglst.GroupCode, -1)
		objPsglst.Totpax = 1
		if len(rslTotpax) > 0 {
			intTotpax, _ := strconv.Atoi(rslTotpax[0])
			objPsglst.Totpax = int32(intTotpax)
		}

		// Ancillary
		func() {

			// Ancillary get price
			nowPrcdtl := psglst.AEDetailsList.PriceDetails
			nowAetotp := 0
			for _, curncy := range []mdlSbrapi.MdlSbrapiPsglstRspprc{
				nowPrcdtl.BasePrice, nowPrcdtl.EquivalentPrice, nowPrcdtl.TotalPrice} {
				if curncy.CurrencyCode == "IDR" {
					nowAetotp = int(curncy.Value)
					break
				}
			}
			if nowAetotp == 0 {
				if val, ist := mapcur[nowPrcdtl.TotalPrice.CurrencyCode]; ist {
					nowAetotp = int(val.Crrate * nowPrcdtl.TotalPrice.Value)
				}
			}

			// Ancillary store final
			objPsglst.Aeitid = psglst.AEDetailsList.ItemID
			objPsglst.Aegrcd = psglst.AEDetailsList.ATPCOGroupCode
			objPsglst.Aesbcd = psglst.AEDetailsList.ATPCOSubCode
			objPsglst.Aedesc = psglst.AEDetailsList.Description
			objPsglst.Aeqtus = psglst.AEDetailsList.QuantityUsed
			objPsglst.Aeqtbg = psglst.AEDetailsList.QuantityBought
			objPsglst.Aetotp = int32(nowAetotp)
		}()

		// Get all code list
		objPsglst.Codels = strings.Join(psglst.EditCodeList, "|")
		objPsglst.Typepx = "ADT"
		if slices.Contains(psglst.EditCodeList, "INF") {
			objPsglst.Gender = "INF"
		} else if slices.Contains(psglst.EditCodeList, "CHD") {
			objPsglst.Gender = "CHD"
		}
		switch {
		case slices.Contains(psglst.EditCodeList, "F"):
			objPsglst.Gender = "Female"
		case slices.Contains(psglst.EditCodeList, "M"):
			objPsglst.Gender = "Male"
		case slices.Contains(psglst.EditCodeList, "CHD"):
			objPsglst.Gender = "Child"
		default:
			objPsglst.Gender = "Doesn't have gender"
		}

		// Is it data
		objPsglst.Isitfl = "N"
		if psglst.BoardingPassFlag == "true" {
			objPsglst.Isitfl = "F"
		}
		objPsglst.Isittx = "N"
		if psglst.ThruIndicator == "!" {
			objPsglst.Isittx = "TX"
		}
		if slices.Contains(psglst.EditCodeList, "IR") {
			objPsglst.Isitir = "IR"
		}
		if objPsglst.Nmefst == "XXDHC" || objPsglst.Nmelst == "XXDHC" {
			objPsglst.Isitnr = "CREW"
		}

		// Get freetext info
		for _, freetx := range psglst.FreeTextInfoList {
			switch freetx.EditCode {

			// Electronic ticket
			case "ET":
				partsl := strings.Fields(freetx.TextLine)
				if len(partsl) >= 6 {
					objPsglst.Tktnvc = partsl[0]
					rawCpnbvc := strings.ReplaceAll(partsl[1], "C", "")
					intCpnbvc, _ := strconv.Atoi(rawCpnbvc)
					objPsglst.Cpnbvc = int32(intCpnbvc)
					strDatevc := fncGlobal.FncGlobalMainprDaymnt(partsl[2])
					intDatevc, _ := strconv.Atoi(strDatevc)
					objPsglst.Datevc = int32(intDatevc)
					objPsglst.Clssvc = partsl[3]
					objPsglst.Cbinvc = clslvl[partsl[3]].Cbinfl
					objPsglst.Routvc = partsl[4][:3] + "-" + partsl[4][3:]
					objPsglst.Statvc = partsl[5]
				}

			// Bagage
			case "BT":
				partsl := strings.Fields(freetx.TextLine)
				if len(partsl) == 3 {
					objPsglst.Nmbrbt = partsl[2]
				} else if len(partsl) == 4 {

					// Quantity
					if partsl[0] != "" {
						objPsglst.Qntybt = partsl[0]
						regQntybt := regexp.MustCompile(`\d+`)
						rslQntybt := regQntybt.FindAllString(partsl[0], -1)
						if len(rslQntybt) > 0 {
							intQntybt, _ := strconv.Atoi(rslQntybt[0])
							mapQntybt[objPsglst.Groupc] += intQntybt
						}
					}

					// Weight
					intBtwght, _ := strconv.Atoi(partsl[2])
					objPsglst.Wghtbt = int32(intBtwght)
					mapWghtbt[objPsglst.Groupc] += int(intBtwght)
					objPsglst.Typebt = partsl[3]
				}

			// Ancillary
			case "AE":
				emdRegexp := regexp.MustCompile(`EMD-?(\d{13})`)
				mtcRegexp := emdRegexp.FindStringSubmatch(freetx.TextLine)
				if len(mtcRegexp) >= 2 {
					objPsglst.Aemdnb = mtcRegexp[1]
				}

			// Comment
			case "CM":
				fncGlobal.FncGlobalMainprNoterr(&objPsglst.Coment, freetx.TextLine)

			// Outbound
			case "OB":
				partsl, rawtme, strdte := strings.Fields(freetx.TextLine), "", ""
				for _, prdata := range partsl {
					switch {
					case regexp.MustCompile(`^[A-Z]{2}|\*[A-Z]{2}$`).MatchString(prdata):
						objPsglst.Airlob = strings.ReplaceAll(prdata, "*", "")
					case regexp.MustCompile(`^\d{1,2}[A-Z]{3}(\d{0,2})?$`).MatchString(prdata):
						if len(prdata) >= 6 {
							fmtObdate, _ := time.Parse("2Jan06", prdata)
							strdte = fmtObdate.Format("060102")
						} else if len(prdata) >= 4 {
							strdte = fncGlobal.FncGlobalMainprDaymnt(prdata)
						}
						intObdate, _ := strconv.Atoi(strdte)
						objPsglst.Dateob = int32(intObdate)
					case regexp.MustCompile(`^ETD\d{4}$`).MatchString(prdata):
						rawtme = prdata[3:]
					case regexp.MustCompile(`^\d{1,4}$`).MatchString(prdata):
						objPsglst.Flnbob = prdata
					case regexp.MustCompile(`^[A-Z]{3}$`).MatchString(prdata):
						if objPsglst.Routob == "" {
							objPsglst.Routob = prdata
						} else {
							objPsglst.Routob += "-" + prdata
						}
					case regexp.MustCompile(`^[A-Z]$`).MatchString(prdata):
						objPsglst.Clssob = prdata
					}
				}
				if rawtme != "" {
					intObtime, _ := strconv.Atoi(strdte + rawtme)
					objPsglst.Timeob = int64(intObtime)
				}

			// Inbound
			case "IB":
				partsl := strings.Fields(freetx.TextLine)
				if len(partsl) >= 5 {
					objPsglst.Airlib = partsl[0]
					objPsglst.Flnbib = partsl[1]
					objPsglst.Clssib = partsl[2]
					objPsglst.Dstrib = partsl[3]
					strIbdate := fncGlobal.FncGlobalMainprDaymnt(partsl[4])
					intIbdate, _ := strconv.Atoi(strIbdate)
					objPsglst.Dateib = int32(intIbdate)
					if len(partsl) == 6 {
						intIbdate, _ := strconv.Atoi(strIbdate + partsl[5][3:])
						objPsglst.Timeib = int64(intIbdate)
					}
				}

			// Inbound
			case "IR":
				partsl := strings.Fields(freetx.TextLine)
				if len(partsl) >= 2 {
					objPsglst.Codeir = partsl[0]
					partwo := strings.Split(partsl[1], "/")
					objPsglst.Airlir = partwo[0][:2]
					objPsglst.Flnbir = partwo[0][2:]
					strIrdate := fncGlobal.FncGlobalMainprDaymnt(partwo[1])
					intIrdate, _ := strconv.Atoi(strIrdate)
					objPsglst.Dateir = int32(intIrdate)
				}

			// Infant electronic ticket
			case "IFET":
				if strings.Contains(freetx.TextLine, "IFET-") {
					partsl := strings.Fields(freetx.TextLine)
					objPsglst.Tktnif = partsl[0][5:]
					rawCpnbif := strings.ReplaceAll(partsl[1], "C", "")
					intCpnbif, _ := strconv.Atoi(rawCpnbif)
					objPsglst.Cpnbif = int32(intCpnbif)
					strIfdate := fncGlobal.FncGlobalMainprDaymnt(partsl[2])
					intIfdate, _ := strconv.Atoi(strIfdate)
					objPsglst.Dateif = int32(intIfdate)
					objPsglst.Clssif = partsl[3]
					objPsglst.Routif = partsl[4]
					objPsglst.Statif = partsl[5]
				} else {
					objPsglst.Paxsif = freetx.TextLine
				}

			// Infant electronic ticket
			case "XT":
				partsl := strings.Fields(freetx.TextLine)
				objPsglst.Airlxt = partsl[0]
				objPsglst.Dstrxt = partsl[1]
				objPsglst.Nmbrxt = partsl[2]
			}
		}

		// Final push to map
		strDatefl := strconv.Itoa(int(objPsglst.Datefl))
		strSeatpx, strPsgrid := objPsglst.Seatpx, objPsglst.Psgrid
		strFlnbfl, strDepart := objPsglst.Flnbfl, objPsglst.Depart
		objPsglst.Prmkey = strDatefl + strFlnbfl + strDepart + strSeatpx + strPsgrid
		slcPsglst = append(slcPsglst, objPsglst)
	}

	// Looping again to push total ebt
	for idx := range slcPsglst {
		nowGroupc := slcPsglst[idx].Groupc
		if nowGroupc == "" {
			slcPsglst[idx].Qtotbt = 0
			intQntybt, err := strconv.Atoi(slcPsglst[idx].Qntybt)
			if err == nil {
				slcPsglst[idx].Qtotbt = int32(intQntybt)
			}
			slcPsglst[idx].Wtotbt = slcPsglst[idx].Wghtbt
		} else {
			if val, ist := mapQntybt[nowGroupc]; ist {
				slcPsglst[idx].Qtotbt = int32(val)
			}
			if val, ist := mapWghtbt[nowGroupc]; ist {
				slcPsglst[idx].Wtotbt = int32(val)
			}
		}

	}

	// Final return
	return slcPsglst
}
