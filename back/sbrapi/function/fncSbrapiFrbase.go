package fncSbrapi

import (
	fncGlobal "back/global/function"
	mdlPsglst "back/psglst/model"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Comamand macro Sabre API Sreen
func FncSbrapiFrbaseMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	apndix mdlSbrapi.MdlSbrapiMsghdrApndix, sycFrbase *sync.Map,
	mapClslvl map[string]mdlPsglst.MdlPsglstClsslvDtbase,
) ([]mongo.WriteModel, error) {

	// Isi struktur data
	strDatefl := strconv.Itoa(int(apndix.Datefl))
	rawDatefl, _ := time.Parse("060102", strDatefl)
	dmyDatefl := rawDatefl.Format("2006-01-02")
	fnlFrbase := []mongo.WriteModel{}
	bdyFrbase := mdlSbrapi.MdlSbrapiFrbaseReqenv{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdlSbrapi.MdlSbrapiFrbaseReqhdr{
			MessageHeader: FncSbrapiMsghdrMainob(fncGlobal.Pcckey,
				"Air Fare By City Pairs", "FareLLSRQ", unqhdr),
			Security: mdlSbrapi.MdlSbrapiFrbaseReqscr{
				BinarySecurityToken: mdlSbrapi.MdlSbrapiFrbaseReqbst{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: unqhdr.Bsttkn,
				}, XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext"}},
		Body: mdlSbrapi.MdlSbrapiFrbaseReqbdy{
			FareRQ: mdlSbrapi.MdlSbrapiFrbaseReqafl{
				Version: "2.9.0",
				Xmlns:   "http://webservices.sabre.com/sabreXML/2011/10",
				OptionalQualifiers: mdlSbrapi.MdlSbrapiFrbaseReqopq{
					Airline: mdlSbrapi.MdlSbrapiFrbaseReqcde{Code: apndix.Airlfl},
					PricingQualifiers: mdlSbrapi.MdlSbrapiFrbaseReqpcq{
						CurrencyCode: "IDR",
						PassengerType: []mdlSbrapi.MdlSbrapiFrbaseReqpst{
							{Code: "ADT"}, {Code: "CHD"}, {Code: "INF"}, {Code: "CNN"}}},
					TimeQualifiers: mdlSbrapi.MdlSbrapiFrbaseReqtdo{
						Historical: mdlSbrapi.MdlSbrapiFrbaseReqhst{
							TicketingDate: dmyDatefl,
							TravelDate:    dmyDatefl}},
				}, OriginDestinationInformation: mdlSbrapi.MdlSbrapiFrbaseReqodi{
					DestinationLocation: mdlSbrapi.MdlSbrapiFrbaseReqdst{
						LocationCode: apndix.Arrivl},
					OriginLocation: mdlSbrapi.MdlSbrapiFrbaseReqdst{
						LocationCode: apndix.Depart},
				},
			},
		},
	}

	// Treatment APO Session
	raw, err := FncSbrapiMsghdrXmldta(bdyFrbase)
	if err != nil {
		return fnlFrbase, err
	}

	// Parsing XML ke dalam struktur Go
	rspFrbase := mdlSbrapi.MdlSbrapiFrbaseRspenv{}
	err = xml.Unmarshal([]byte(raw), &rspFrbase)
	if err != nil {
		return fnlFrbase, err
	}

	// Return String
	rawFrbase := rspFrbase.Body.FareRS.FareBasis
	fnlFrbase = FncSbrapiFrbaseTrtmnt(rawFrbase, apndix, sycFrbase, mapClslvl)
	return fnlFrbase, nil
}

// Treatment data raw flight list
func FncSbrapiFrbaseTrtmnt(rawxml []mdlSbrapi.MdlSbrapiFrbaseRspfbs,
	apndix mdlSbrapi.MdlSbrapiMsghdrApndix, sycFrbase *sync.Map,
	mapClslvl map[string]mdlPsglst.MdlPsglstClsslvDtbase,
) []mongo.WriteModel {

	// Looping xml farebase
	var fnlFrbase = []mongo.WriteModel{}
	for _, frbase := range rawxml {
		nowClssfl := frbase.AdditionalInformation.ResBookDesigCode
		intFrmant := frbase.AdditionalInformation.Fare[0].Amount
		tmpPrmkey := apndix.Airlfl + apndix.Routfl + frbase.Code
		tmpScdkey := apndix.Airlfl + apndix.Routfl + nowClssfl
		if frbase.AdditionalInformation.OneWayRoundTrip.Ind == "R" {
			intFrmant = intFrmant / 2
			tmpScdkey += "RT"
		}

		// Check now than prev frbase
		var intDatenw, _ = strconv.Atoi(time.Now().Format("060102"))
		var nowDatend = int32(intDatenw)
		var nowHstory = string("")
		if val, ist := sycFrbase.Load(tmpPrmkey); ist {
			if get, mtc := val.(mdlPsglst.MdlPsglstFrbaseDtbase); mtc {
				nowDatend, nowHstory = fncGlobal.FncGlobalMainprHstory(get.Frbsbr,
					int32(intFrmant), get.Hstory, get.Datend, int32(intDatenw))
			}
		}

		// Declare fare amount NTA
		getClslvl, ist := mapClslvl[strings.ToUpper(nowClssfl)]
		nowDscont := 0.03
		if ist {
			nowDscont = getClslvl.Clssdc
		}
		intFrbnta := math.Ceil(float64(intFrmant)*(1-nowDscont+0.11)/1.11/1000) * 1000

		// Final result
		nowFrbase := mdlPsglst.MdlPsglstFrbaseDtbase{
			Prmkey: tmpPrmkey,
			Scdkey: tmpScdkey,
			Airlfl: apndix.Airlfl,
			Clssfl: nowClssfl,
			Routfl: apndix.Routfl,
			Frbcde: frbase.Code,
			Frbnta: int32(intFrbnta),
			Frbsbr: int32(intFrmant),
			Datend: nowDatend,
			Hstory: nowHstory,
		}
		if nowFrbase.Frbsbr > 0 && nowFrbase.Prmkey != "" {
			fnlFrbase = append(fnlFrbase, mongo.NewUpdateOneModel().
				SetFilter(bson.M{"prmkey": nowFrbase.Prmkey}).
				SetUpdate(bson.M{"$set": nowFrbase}).
				SetUpsert(true))
			sycFrbase.Store(tmpPrmkey, nowFrbase)
			sycFrbase.Store(tmpScdkey, nowFrbase)
		}
	}

	// Return final data
	return fnlFrbase

}
