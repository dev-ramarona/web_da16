package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
	"slices"
	"strconv"
	"strings"
	"time"
)

// Comamand macro Sabre API Sreen
func FncSbrapiFllistMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	apndix mdlSbrapi.MdlSbrapiMsghdrApndix) ([]mdlPsglst.MdlPsglstFllistDtbase, error) {

	// Isi struktur data
	strDatefl := strconv.Itoa(int(apndix.Datefl))
	rawDatefl, _ := time.Parse("060102", strDatefl)
	ymdDatefl := rawDatefl.Format("2006-01-02")
	fnlFllist := []mdlPsglst.MdlPsglstFllistDtbase{}
	bdyFllist := mdlSbrapi.MdlSbrapiFllistReqenv{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdlSbrapi.MdlSbrapiFllistReqhdr{
			MessageHeader: FncSbrapiMsghdrMainob(fncGlobal.Pcckey,
				"Search By Specific Airport", "ACS_AirportFlightListRQ", unqhdr),
			Security: mdlSbrapi.MdlSbrapiFllistReqscr{
				BinarySecurityToken: mdlSbrapi.MdlSbrapiFllistReqbst{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: unqhdr.Bsttkn,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdlSbrapi.MdlSbrapiFllistReqbdy{
			ACS_AirportFlightListRQ: mdlSbrapi.MdlSbrapiFllistReqafl{
				Xmlns: "http://services.sabre.com/ACS/BSO/airportFlightList/v3",
				FlightInfo: mdlSbrapi.MdlSbrapiFllistReqinf{
					Airline:       apndix.Airlfl,
					DepartureDate: ymdDatefl,
					Origin:        apndix.Depart,
					DepartureTimeRange: mdlSbrapi.MdlSbrapiFllistReqdtr{
						StartTime: "0000",
						EndTime:   "2359",
					},
				},
			},
		},
	}

	// Treatment APO Session
	raw, err := FncSbrapiMsghdrXmldta(bdyFllist)
	if err != nil {
		return fnlFllist, err
	}

	// Parsing XML ke dalam struktur Go
	rspFllist := mdlSbrapi.MdlSbrapiFllistRspenv{}
	err = xml.Unmarshal([]byte(raw), &rspFllist)
	if err != nil {
		return fnlFllist, err
	}

	// Return String
	rawFllist := rspFllist.Body.ACS_AirportFlightListRS
	fnlFllist = FncSbrapiFllistTrtmnt(rawFllist, apndix)
	return fnlFllist, nil
}

// Treatment data raw flight list
func FncSbrapiFllistTrtmnt(rawxml mdlSbrapi.MdlSbrapiFllistRspfls,
	apndix mdlSbrapi.MdlSbrapiMsghdrApndix) []mdlPsglst.MdlPsglstFllistDtbase {
	fnlFllist := []mdlPsglst.MdlPsglstFllistDtbase{}
	nowDepart := rawxml.Origin

	// Looping all flight list
	for _, fllist := range rawxml.AirportFlightList.AirportFlight {

		// Treatment date
		nowDatefl := strconv.Itoa(int(apndix.Datefl))
		strTimeup := time.Now().Format("0601021504")
		intTimeup, _ := strconv.Atoi(strTimeup)
		strTimefl := fllist.DepartureDate + "/" + fllist.DepartureTime
		fmtTimefl, _ := time.Parse("2006-01-02/03:04PM", strTimefl)
		intTimefl, _ := strconv.Atoi(fmtTimefl.Format("0601021504"))
		strMnthfl := strconv.Itoa(int(apndix.Datefl))[:4]
		intMnthfl, _ := strconv.Atoi(strMnthfl)

		// Treatment route
		viaArrivl := fllist.Destination
		fnlArrivl := fllist.DestinationFinal
		rawRoutac := []string{nowDepart, viaArrivl, fnlArrivl}
		slcRoutac := slices.DeleteFunc(rawRoutac, func(s string) bool {
			return s == ""
		})
		nowArrivl := slcRoutac[1]
		nowRoutfl := nowDepart + "-" + nowArrivl
		nowRoutac := strings.Join(slcRoutac, "-")

		// Treatment flight number
		strFlnbfl := strings.Trim(fllist.Flight, " ")
		intFlnbfl, err := strconv.Atoi(strFlnbfl)
		if err == nil {
			strFlnbfl = strconv.Itoa(intFlnbfl)
		}

		// Final output
		keyFllist := apndix.Airlfl + strFlnbfl + nowDepart + nowDatefl
		fnlFllist = append(fnlFllist, mdlPsglst.MdlPsglstFllistDtbase{
			Prmkey: keyFllist,
			Airlfl: apndix.Airlfl,
			Flnbfl: strFlnbfl,
			Timeup: int64(intTimeup),
			Timefl: int64(intTimefl),
			Timerv: int64(intTimefl),
			Datefl: apndix.Datefl,
			Mnthfl: int32(intMnthfl),
			Flstat: fllist.Status,
			Routfl: nowRoutfl,
			Routac: nowRoutac,
			Routmx: "fldtil",
			Flsarr: "fldtil",
			Flhour: 1,
			Flrpdc: 0,
			Flgate: fllist.DepartureGate,
			Depart: nowDepart,
			Arrivl: nowArrivl,
			Airtyp: fllist.AircraftType,
			Aircnf: "fldtil",
			Seatcn: "fldtil",
		})
	}

	// Final return
	return fnlFllist

}
