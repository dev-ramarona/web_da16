package mdlSbrapi

import "encoding/xml"

// Request
type MdlSbrapiFrbaseReqenv struct {
	XMLName xml.Name              `xml:"soap-env:Envelope"`
	Xmlns   string                `xml:"xmlns:soap-env,attr"`
	Header  MdlSbrapiFrbaseReqhdr `xml:"soap-env:Header"`
	Body    MdlSbrapiFrbaseReqbdy `xml:"soap-env:Body"`
}
type MdlSbrapiFrbaseReqhdr struct {
	MessageHeader MdlSbrapiMsghdrMainob `xml:"eb:MessageHeader"`
	Security      MdlSbrapiFrbaseReqscr `xml:"wsse:Security"`
}
type MdlSbrapiFrbaseReqscr struct {
	BinarySecurityToken MdlSbrapiFrbaseReqbst `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string                `xml:"xmlns:wsse,attr"`
}
type MdlSbrapiFrbaseReqbst struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}
type MdlSbrapiFrbaseReqbdy struct {
	FareRQ MdlSbrapiFrbaseReqafl `xml:"FareRQ"`
}
type MdlSbrapiFrbaseReqafl struct {
	XMLName                      xml.Name              `xml:"FareRQ"`
	Xmlns                        string                `xml:"xmlns,attr"`
	Version                      string                `xml:"Version,attr"`
	OptionalQualifiers           MdlSbrapiFrbaseReqopq `xml:"OptionalQualifiers"`
	OriginDestinationInformation MdlSbrapiFrbaseReqodi `xml:"OriginDestinationInformation>FlightSegment"`
}
type MdlSbrapiFrbaseReqopq struct {
	Airline           MdlSbrapiFrbaseReqcde `xml:"FlightQualifiers>VendorPrefs>Airline"`
	PricingQualifiers MdlSbrapiFrbaseReqpcq `xml:"PricingQualifiers"`
	TimeQualifiers    MdlSbrapiFrbaseReqtdo `xml:"TimeQualifiers>TravelDateOptions"`
}
type MdlSbrapiFrbaseReqcde struct {
	Code string `xml:"Code,attr"`
}
type MdlSbrapiFrbaseReqpcq struct {
	CurrencyCode  string                  `xml:"CurrencyCode,attr"`
	JourneyType   MdlSbrapiFrbaseReqjtp   `xml:"JourneyType"`
	PassengerType []MdlSbrapiFrbaseReqpst `xml:"PassengerType"`
}
type MdlSbrapiFrbaseReqjtp struct {
	Code string `xml:"Code"`
}
type MdlSbrapiFrbaseReqpst struct {
	Code string `xml:"Code,attr"`
}
type MdlSbrapiFrbaseReqtdo struct {
	Historical MdlSbrapiFrbaseReqhst `xml:"Historical"`
}
type MdlSbrapiFrbaseReqhst struct {
	TicketingDate string `xml:"TicketingDate"`
	TravelDate    string `xml:"TravelDate"`
}
type MdlSbrapiFrbaseReqodi struct {
	DestinationLocation MdlSbrapiFrbaseReqdst `xml:"DestinationLocation"`
	OriginLocation      MdlSbrapiFrbaseReqdst `xml:"OriginLocation"`
}
type MdlSbrapiFrbaseReqdst struct {
	LocationCode string `xml:"LocationCode,attr"`
}

// Response
type MdlSbrapiFrbaseRspenv struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    MdlSbrapiFrbaseRspbdy `xml:"Body"`
}
type MdlSbrapiFrbaseRspbdy struct {
	FareRS MdlSbrapiFrbaseRspfrs `xml:"http://webservices.sabre.com/sabreXML/2011/10 FareRS"`
}
type MdlSbrapiFrbaseRspfrs struct {
	FareBasis []MdlSbrapiFrbaseRspfbs `xml:"FareBasis"`
}
type MdlSbrapiFrbaseRspfbs struct {
	Code                  string                `xml:"Code,attr"`
	CurrencyCode          string                `xml:"CurrencyCode,attr"`
	RPH                   string                `xml:"RPH,attr"`
	AdditionalInformation MdlSbrapiFrbaseRspadi `xml:"AdditionalInformation"`
	BaseFare              MdlSbrapiFrbaseRspamt `xml:"BaseFare"`
}
type MdlSbrapiFrbaseRspadi struct {
	ResBookDesigCode string                  `xml:"ResBookDesigCode,attr"`
	Cabin            string                  `xml:"Cabin"`
	CabinName        string                  `xml:"CabinName"`
	Airline          string                  `xml:"Airline"`
	Fare             []MdlSbrapiFrbaseRspamt `xml:"Fare"`
	OneWayRoundTrip  MdlSbrapiFrbaseRsport   `xml:"OneWayRoundTrip"`
}
type MdlSbrapiFrbaseRspamt struct {
	Amount float64 `xml:"Amount,attr"`
}
type MdlSbrapiFrbaseRsport struct {
	Ind string `xml:"Ind,attr"`
}
