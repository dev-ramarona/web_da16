package mdlSbrapi

import "encoding/xml"

// Request
type MdlSbrapiFldtilReqenv struct {
	XMLName xml.Name              `xml:"soap-env:Envelope"`
	Xmlns   string                `xml:"xmlns:soap-env,attr"`
	Header  MdlSbrapiFldtilReqhdr `xml:"soap-env:Header"`
	Body    MdlSbrapiFldtilReqbdy `xml:"soap-env:Body"`
}
type MdlSbrapiFldtilReqhdr struct {
	MessageHeader MdlSbrapiMsghdrMainob `xml:"eb:MessageHeader"`
	Security      MdlSbrapiFldtilReqscr `xml:"wsse:Security"`
}
type MdlSbrapiFldtilReqscr struct {
	BinarySecurityToken MdlSbrapiFldtilReqbst `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string                `xml:"xmlns:wsse,attr"`
}
type MdlSbrapiFldtilReqbst struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}
type MdlSbrapiFldtilReqbdy struct {
	ACS_FlightDetailRQ MdlSbrapiFldtilReqafl `xml:"v3:ACS_FlightDetailRQ"`
}
type MdlSbrapiFldtilReqafl struct {
	XMLName    xml.Name              `xml:"v3:ACS_FlightDetailRQ"`
	Xmlns      string                `xml:"xmlns:v3,attr"`
	Version    string                `xml:"Version,attr"`
	FlightInfo MdlSbrapiFldtilReqinf `xml:"FlightInfo"`
}
type MdlSbrapiFldtilReqinf struct {
	Airline       string `xml:"Airline"`
	Flight        string `xml:"Flight"`
	DepartureDate string `xml:"DepartureDate"`
	Origin        string `xml:"Origin"`
}

// Response
type MdlSbrapiFldtilRspenv struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    MdlSbrapiFldtilRspbdy `xml:"Body"`
}
type MdlSbrapiFldtilRspbdy struct {
	ACS_FlightDetailRS MdlSbrapiFldtilRspacs `xml:"http://services.sabre.com/ACS/BSO/flightDetail/v3 ACS_FlightDetailRS"`
}
type MdlSbrapiFldtilRspacs struct {
	ItineraryResponseList MdlSbrapiFldtilRspitn   `xml:"ItineraryResponseList"`
	LegInfoList           MdlSbrapiFldtilRsplfl   `xml:"LegInfoList"`
	PassengerCounts       []MdlSbrapiFldtilRsppgc `xml:"PassengerCounts"`
}
type MdlSbrapiFldtilRspitn struct {
	ItineraryInfoResponse MdlSbrapiFldtilRspitr `xml:"ItineraryInfoResponse"`
}
type MdlSbrapiFldtilRspitr struct {
	Airline                string                `xml:"Airline"`
	Flight                 string                `xml:"Flight"`
	Origin                 string                `xml:"Origin"`
	DepartureDate          string                `xml:"DepartureDate"`
	DepartureTime          string                `xml:"DepartureTime"`
	ScheduledDepartureDate string                `xml:"ScheduledDepartureDate"`
	ScheduledDepartureTime string                `xml:"ScheduledDepartureTime"`
	EstimatedDepartureDate string                `xml:"EstimatedDepartureDate"`
	EstimatedDepartureTime string                `xml:"EstimatedDepartureTime"`
	DepartureGate          string                `xml:"DepartureGate"`
	BoardingTime           string                `xml:"BoardingTime"`
	ArrivalDate            string                `xml:"ArrivalDate"`
	ArrivalTime            string                `xml:"ArrivalTime"`
	ScheduledArrivalTime   string                `xml:"ScheduledArrivalTime"`
	EstimatedArrivalDate   string                `xml:"EstimatedArrivalDate"`
	EstimatedArrivalTime   string                `xml:"EstimatedArrivalTime"`
	AircraftType           string                `xml:"AircraftType"`
	FlightStatus           string                `xml:"FlightStatus"`
	AircraftConfigNumber   string                `xml:"AircraftConfigNumber"`
	CheckinRuleNumber      string                `xml:"CheckinRuleNumber"`
	SeatConfig             string                `xml:"SeatConfig"`
	BagCheckInOption       string                `xml:"BagCheckInOption"`
	AutoOn                 string                `xml:"AutoOn"`
	HeldSeat               string                `xml:"HeldSeat"`
	FreeTextInfoList       MdlSbrapiFldtilRsptxl `xml:"FreeTextInfoList"`
}
type MdlSbrapiFldtilRsptxl struct {
	FreeTextInfo []MdlSbrapiFldtilRsptxi `xml:"FreeTextInfo"`
}
type MdlSbrapiFldtilRsptxi struct {
	TextLine MdlSbrapiFldtilRsptxt `xml:"TextLine"`
}
type MdlSbrapiFldtilRsptxt struct {
	Text string `xml:"Text"`
}
type MdlSbrapiFldtilRsplfl struct {
	LegInfo []MdlSbrapiFldtilRsplif `xml:"LegInfo"`
}
type MdlSbrapiFldtilRsplif struct {
	LegCity          string `xml:"LegCity"`
	LegStatus        string `xml:"LegStatus"`
	LegDate          string `xml:"LegDate"`
	LegDepartureTime string `xml:"LegDepartureTime"`
	LegApp           string `xml:"LegApp"`
}
type MdlSbrapiFldtilRsppgc struct {
	ClassOfService string `xml:"classOfService,attr"`
	Authorized     int    `xml:"Authorized"`
	Booked         int    `xml:"Booked"`
	Available      int    `xml:"Available"`
}
