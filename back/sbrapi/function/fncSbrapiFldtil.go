package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

// Comamand macro Sabre API Sreen
func FncSbrapiFldtilMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	apndix mdlSbrapi.MdlSbrapiMsghdrApndix,
	fllist *mdlPsglst.MdlPsglstFllistDtbase) error {

	// Isi struktur data
	strDatefl := strconv.Itoa(int(apndix.Datefl))
	rawDatefl, _ := time.Parse("060102", strDatefl)
	dmyDatefl := rawDatefl.Format("2006-01-02")
	bdyFldtil := mdlSbrapi.MdlSbrapiFldtilReqenv{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdlSbrapi.MdlSbrapiFldtilReqhdr{
			MessageHeader: FncSbrapiMsghdrMainob(fncGlobal.Pcckey,
				"Get Flight Information", "ACS_FlightDetailRQ", unqhdr),
			Security: mdlSbrapi.MdlSbrapiFldtilReqscr{
				BinarySecurityToken: mdlSbrapi.MdlSbrapiFldtilReqbst{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: unqhdr.Bsttkn,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdlSbrapi.MdlSbrapiFldtilReqbdy{
			ACS_FlightDetailRQ: mdlSbrapi.MdlSbrapiFldtilReqafl{
				Xmlns:   "http://services.sabre.com/ACS/BSO/flightDetail/v3",
				Version: "3.0.0",
				FlightInfo: mdlSbrapi.MdlSbrapiFldtilReqinf{
					Airline:       apndix.Airlfl,
					Flight:        apndix.Flnbfl,
					DepartureDate: dmyDatefl,
					Origin:        apndix.Depart,
				},
			},
		},
	}

	// Treatment APO Session
	raw, err := FncSbrapiMsghdrXmldta(bdyFldtil)
	if err != nil {
		return err
	}

	// Parsing XML ke dalam struktur Go
	rspFldtil := mdlSbrapi.MdlSbrapiFldtilRspenv{}
	err = xml.Unmarshal([]byte(raw), &rspFldtil)
	if err != nil {
		return err
	}

	// Return String
	rawFldtil := rspFldtil.Body.ACS_FlightDetailRS
	FncSbrapiFldtilTrtmnt(rawFldtil, apndix, fllist)
	return nil
}

// Treatment data raw flight list
func FncSbrapiFldtilTrtmnt(rawxml mdlSbrapi.MdlSbrapiFldtilRspacs,
	apndix mdlSbrapi.MdlSbrapiMsghdrApndix, fllist *mdlPsglst.MdlPsglstFllistDtbase,
) {

	// Get aircfraft config and secok
	nowFlinfo := rawxml.ItineraryResponseList.ItineraryInfoResponse
	fllist.Aircnf = nowFlinfo.AircraftConfigNumber
	fllist.Seatcn = nowFlinfo.SeatConfig
	for _, val := range nowFlinfo.FreeTextInfoList.FreeTextInfo {
		nowTxtsck := val.TextLine.Text
		if strings.Contains(nowTxtsck, "SECOK") {
			intSecoks, err := strconv.Atoi(nowTxtsck[len(nowTxtsck)-1:])
			if err == nil {
				fllist.Flrpdc = int32(intSecoks)
				break
			}
		}
	}

	// Get data time arrival
	strTimefl := nowFlinfo.ArrivalDate + "/" + nowFlinfo.ArrivalTime
	fmtTimefl, _ := time.Parse("2006-01-02/03:04PM", strTimefl)
	intTimefl, _ := strconv.Atoi(fmtTimefl.Format("0601021504"))
	fllist.Timerv = int64(intTimefl)

	// Handle authrorize dan bboked
	nowAutbkd := rawxml.PassengerCounts
	for _, val := range nowAutbkd {
		if val.ClassOfService == "C" {
			fllist.Autrzc = int32(val.Authorized)
			fllist.Bookdc = int32(val.Booked)
		} else {
			fllist.Autrzy = int32(val.Authorized)
			fllist.Bookdy = int32(val.Booked)
		}
	}

	// Looping segment leg get routex
	slcRoutmx := []string{}
	slcFlstat := []string{}
	for _, leg := range rawxml.LegInfoList.LegInfo {
		slcRoutmx = append(slcRoutmx, leg.LegCity)
		slcFlstat = append(slcFlstat, leg.LegCity+":"+leg.LegStatus)
	}
	fllist.Routmx = strings.Join(slcRoutmx, "-")
	fllist.Flsarr = strings.Join(slcFlstat, "|")

}
