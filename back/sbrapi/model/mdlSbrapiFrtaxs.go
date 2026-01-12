package mdlSbrapi

import "encoding/xml"

// Request
type MdlSbrapiFrtaxsReqenv struct {
	XMLName xml.Name              `xml:"soap-env:Envelope"`
	Xmlns   string                `xml:"xmlns:soap-env,attr"`
	Header  MdlSbrapiFrtaxsReqhdr `xml:"soap-env:Header"`
	Body    MdlSbrapiFrtaxsReqbdy `xml:"soap-env:Body"`
}
type MdlSbrapiFrtaxsReqhdr struct {
	MessageHeader MdlSbrapiMsghdrMainob `xml:"eb:MessageHeader"`
	Security      MdlSbrapiFrtaxsReqscr `xml:"wsse:Security"`
}
type MdlSbrapiFrtaxsReqscr struct {
	BinarySecurityToken MdlSbrapiFrtaxsReqbst `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string                `xml:"xmlns:wsse,attr"`
}
type MdlSbrapiFrtaxsReqbst struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}
type MdlSbrapiFrtaxsReqbdy struct {
	AirTaxRQ MdlSbrapiFrtaxsReqatx `xml:"AirTaxRQ"`
}
type MdlSbrapiFrtaxsReqatx struct {
	XMLName        xml.Name              `xml:"AirTaxRQ"`
	Xmlns          string                `xml:"xmlns,attr"`
	Version        string                `xml:"Version,attr"`
	ItineraryInfos MdlSbrapiFrtaxsReqits `xml:"ItineraryInfos>ItineraryInfo>ReservationItems>Item"`
}
type MdlSbrapiFrtaxsReqits struct {
	RPH                string                `xml:"RPH,attr"`
	SalePseudoCityCode string                `xml:"SalePseudoCityCode,attr"`
	TicketingCarrier   string                `xml:"TicketingCarrier,attr"`
	ValidatingCarrier  string                `xml:"ValidatingCarrier,attr"`
	FlightSegment      MdlSbrapiFrtaxsReqfls `xml:"FlightSegment"`
	AirFareInfo        MdlSbrapiFrtaxsReqafi `xml:"AirFareInfo"`
}
type MdlSbrapiFrtaxsReqfls struct {
	ArrivalDateTime   string                `xml:"ArrivalDateTime,attr"`
	DepartureDateTime string                `xml:"DepartureDateTime,attr"`
	FlightNumber      string                `xml:"FlightNumber,attr"`
	ResBookDesigCode  string                `xml:"ResBookDesigCode,attr"`
	DepartureAirport  MdlSbrapiFrtaxsReqdpa `xml:"DepartureAirport"`
	ArrivalAirport    MdlSbrapiFrtaxsReqarv `xml:"ArrivalAirport"`
	Equipment         MdlSbrapiFrtaxsReqeqp `xml:"Equipment"`
	MarketingAirline  MdlSbrapiFrtaxsReqcde `xml:"MarketingAirline"`
	OperatingAirline  MdlSbrapiFrtaxsReqcde `xml:"OperatingAirline"`
}
type MdlSbrapiFrtaxsReqdpa struct {
	CodeContext  string `xml:"CodeContext,attr"`
	LocationCode string `xml:"LocationCode,attr"`
}
type MdlSbrapiFrtaxsReqarv struct {
	CodeContext  string `xml:"CodeContext,attr"`
	LocationCode string `xml:"LocationCode,attr"`
}
type MdlSbrapiFrtaxsReqeqp struct {
	AirEquipType string `xml:"AirEquipType,attr"`
}
type MdlSbrapiFrtaxsReqafi struct {
	PTC_FareBreakdown MdlSbrapiFrtaxsReqfbk `xml:"PTC_FareBreakdown"`
}
type MdlSbrapiFrtaxsReqfbk struct {
	PassengerType MdlSbrapiFrtaxsReqcde `xml:"PassengerType"`
	FareBasisCode string                `xml:"FareBasisCode"`
	PassengerFare MdlSbrapiFrtaxsReqpsf `xml:"PassengerFare"`
}
type MdlSbrapiFrtaxsReqcde struct {
	Code string `xml:"Code,attr"`
}
type MdlSbrapiFrtaxsReqpsf struct {
	BaseFare MdlSbrapiFrtaxsReqbsf `xml:"BaseFare"`
}
type MdlSbrapiFrtaxsReqbsf struct {
	Amount string `xml:"Amount,attr"`
}

// Response
type MdlSbrapiFrtaxsRspenv struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    MdlSbrapiFrtaxsRspbdy `xml:"Body"`
}
type MdlSbrapiFrtaxsRspbdy struct {
	AirTaxRS MdlSbrapiFrtaxsRspatx `xml:"http://webservices.sabre.com/sabreXML/2003/07 AirTaxRS"`
}
type MdlSbrapiFrtaxsRspatx struct {
	ItineraryInfos MdlSbrapiFrtaxsRspiif `xml:"ItineraryInfos"`
}
type MdlSbrapiFrtaxsRspiif struct {
	ItineraryInfo MdlSbrapiFrtaxsRspiti `xml:"ItineraryInfo"`
}
type MdlSbrapiFrtaxsRspiti struct {
	RPH     string                `xml:"RPH,attr"`
	TaxInfo MdlSbrapiFrtaxsRsptxi `xml:"TaxInfo"`
}
type MdlSbrapiFrtaxsRsptxi struct {
	TaxDetails MdlSbrapiFrtaxsRsptxd `xml:"TaxDetails"`
}
type MdlSbrapiFrtaxsRsptxd struct {
	Tax []MdlSbrapiFrtaxsRsptax `xml:"Tax"`
}
type MdlSbrapiFrtaxsRsptax struct {
	TaxCode           string `xml:"TaxCode,attr"`
	Type              string `xml:"Type,attr"`
	Amount            string `xml:"Amount,attr"`
	Currency          string `xml:"Currency,attr"`
	PublishedAmount   string `xml:"PublishedAmount,attr"`
	PublishedCurrency string `xml:"PublishedCurrency,attr"`
	Station           string `xml:"Station,attr"`
	AirlineCode       string `xml:"AirlineCode,attr"`
	Text              string `xml:"Text,attr"`
	SequenceNumber    string `xml:"SequenceNumber,attr"`
}
