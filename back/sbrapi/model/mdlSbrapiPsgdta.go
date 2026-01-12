package mdlSbrapi

import "encoding/xml"

// Request
type MdlSbrapiPsgdtaReqenv struct {
	XMLName xml.Name              `xml:"soap-env:Envelope"`
	Xmlns   string                `xml:"xmlns:soap-env,attr"`
	Header  MdlSbrapiPsgdtaReqhdr `xml:"soap-env:Header"`
	Body    MdlSbrapiPsgdtaReqbdy `xml:"soap-env:Body"`
}
type MdlSbrapiPsgdtaReqhdr struct {
	MessageHeader MdlSbrapiMsghdrMainob `xml:"eb:MessageHeader"`
	Security      MdlSbrapiPsgdtaReqscr `xml:"wsse:Security"`
}
type MdlSbrapiPsgdtaReqscr struct {
	BinarySecurityToken MdlSbrapiPsgdtaReqbst `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string                `xml:"xmlns:wsse,attr"`
}
type MdlSbrapiPsgdtaReqbst struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}
type MdlSbrapiPsgdtaReqbdy struct {
	GetPassengerDataRQ MdlSbrapiPsgdtaReqafl `xml:"v4:GetPassengerDataRQ"`
}
type MdlSbrapiPsgdtaReqafl struct {
	XMLName                     xml.Name              `xml:"v4:GetPassengerDataRQ"`
	Xmlns                       string                `xml:"xmlns:v4,attr"`
	ValidateCheckInRequirements string                `xml:"validateCheckInRequirements,attr"`
	Version                     string                `xml:"version,attr"`
	ItineraryAndPassengerInfo   MdlSbrapiPsgdtaReqiap `xml:"ItineraryAndPassengerInfo"`
}
type MdlSbrapiPsgdtaReqiap struct {
	Itinerary MdlSbrapiPsgdtaReqitn `xml:"Itinerary"`
	Seat      string                `xml:"Seat"`
}
type MdlSbrapiPsgdtaReqitn struct {
	Airline       string `xml:"Airline"`
	Flight        string `xml:"Flight"`
	DepartureDate string `xml:"DepartureDate"`
	Origin        string `xml:"Origin"`
}

// Response
type MdlSbrapiPsgdtaRspenv struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    MdlSbrapiPsgdtaRspbdy `xml:"Body"`
}
type MdlSbrapiPsgdtaRspbdy struct {
	GetPassengerDataRS MdlSbrapiPsgdtaRspgpl `xml:"http://services.sabre.com/checkin/getPassengerData/v4 GetPassengerDataRS"`
}
type MdlSbrapiPsgdtaRspgpl struct {
	PassengerDataResponse MdlSbrapiPsgdtaRsppdr `xml:"PassengerDataResponseList>PassengerDataResponse"`
}
type MdlSbrapiPsgdtaRsppdr struct {
	LineNumber       int                     `xml:"LineNumber"`
	LastName         string                  `xml:"LastName"`
	FirstName        string                  `xml:"FirstName"`
	PassengerID      string                  `xml:"PassengerID"`
	PNRLocator       string                  `xml:"PNRLocator"`
	GroupCode        string                  `xml:"GroupCode"`
	VCRNumber        string                  `xml:"VCRNumber"`
	FreeTextInfoList []MdlSbrapiPsgdtaRspfti `xml:"FreeTextInfoList>FreeTextInfo"`
}
type MdlSbrapiPsgdtaRspfti struct {
	EditCode string `xml:"EditCode"`
	TextLine string `xml:"TextLine>Text"`
}
