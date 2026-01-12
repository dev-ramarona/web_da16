package mdlSbrapi

import "encoding/xml"

// Request
type MdlSbrapiDspvcrReqenv struct {
	XMLName xml.Name              `xml:"soap-env:Envelope"`
	Xmlns   string                `xml:"xmlns:soap-env,attr"`
	Header  MdlSbrapiDspvcrReqhdr `xml:"soap-env:Header"`
	Body    MdlSbrapiDspvcrReqbdy `xml:"soap-env:Body"`
}
type MdlSbrapiDspvcrReqhdr struct {
	MessageHeader MdlSbrapiMsghdrMainob `xml:"eb:MessageHeader"`
	Security      MdlSbrapiDspvcrReqscr `xml:"wsse:Security"`
}
type MdlSbrapiDspvcrReqscr struct {
	BinarySecurityToken MdlSbrapiDspvcrReqbst `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string                `xml:"xmlns:wsse,attr"`
}
type MdlSbrapiDspvcrReqbst struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}
type MdlSbrapiDspvcrReqbdy struct {
	VCR_DisplayRQ MdlSbrapiDspvcrReqafl `xml:"VCR_DisplayRQ"`
}
type MdlSbrapiDspvcrReqafl struct {
	XMLName       xml.Name              `xml:"VCR_DisplayRQ"`
	Xmlns         string                `xml:"xmlns,attr"`
	Version       string                `xml:"Version,attr"`
	SearchOptions MdlSbrapiDspvcrReqsrc `xml:"SearchOptions"`
}
type MdlSbrapiDspvcrReqsrc struct {
	Ticketing MdlSbrapiDspvcrReqtkt `xml:"Ticketing"`
}
type MdlSbrapiDspvcrReqtkt struct {
	ETicketNumber    string                `xml:"eTicketNumber,attr"`
	OperatingAirline MdlSbrapiDspvcrReqopa `xml:"OperatingAirline"`
}
type MdlSbrapiDspvcrReqopa struct {
	Code string `xml:"Code,attr"`
}

// Response
type MdlSbrapiDspvcrRspenv struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    MdlSbrapiDspvcrRspbdy `xml:"Body"`
}
type MdlSbrapiDspvcrRspbdy struct {
	VCR_DisplayRS MdlSbrapiDspvcrRspvcr `xml:"VCR_DisplayRS"`
}
type MdlSbrapiDspvcrRspvcr struct {
	TicketingInfo MdlSbrapiDspvcrRsptkt `xml:"TicketingInfos>TicketingInfo"`
}
type MdlSbrapiDspvcrRsptkt struct {
	CouponData MdlSbrapiDspvcrRspcpd `xml:"CouponData"`
	Ticketing  MdlSbrapiDspvcrRsptkx `xml:"Ticketing"`
}
type MdlSbrapiDspvcrRspcpd struct {
	AirItineraryPricingInfo MdlSbrapiDspvcrRspapi   `xml:"AirItineraryPricingInfo"`
	Coupon                  []MdlSbrapiDspvcrRspcpn `xml:"Coupons>Coupon"`
}
type MdlSbrapiDspvcrRspapi struct {
	ItinTotalFare     MdlSbrapiDspvcrRspitf `xml:"ItinTotalFare"`
	PTC_FareBreakdown MdlSbrapiDspvcrRspmsd `xml:"PTC_FareBreakdown"`
}
type MdlSbrapiDspvcrRspitf struct {
	BaseFare  MdlSbrapiDspvcrRspbfr `xml:"BaseFare"`
	Taxes     MdlSbrapiDspvcrRsptxs `xml:"Taxes"`
	TotalFare MdlSbrapiDspvcrRspttf `xml:"TotalFare"`
}
type MdlSbrapiDspvcrRspbfr struct {
	Amount       string `xml:"Amount,attr"`
	CurrencyCode string `xml:"CurrencyCode,attr"`
}
type MdlSbrapiDspvcrRsptxs struct {
	Tax []Tax `xml:"Tax"`
}
type Tax struct {
	Amount       string `xml:"Amount,attr"`
	CurrencyCode string `xml:"CurrencyCode,attr"`
	PaidInd      string `xml:"PaidInd,attr"`
	TaxCode      string `xml:"TaxCode,attr"`
}
type MdlSbrapiDspvcrRspttf struct {
	Amount       string `xml:"Amount,attr"`
	CurrencyCode string `xml:"CurrencyCode,attr"`
}
type MdlSbrapiDspvcrRspmsd struct {
	FareCalculation []MdlSbrapiDspvcrRspfrc `xml:"FareCalculation"`
}
type MdlSbrapiDspvcrRspfrc struct {
	Text string `xml:"Text"`
}
type MdlSbrapiDspvcrRspcpn struct {
	StatusCode         string                `xml:"StatusCode,attr"`
	EntitlementNumber  int                   `xml:"EntitlementNumber,attr"`
	FlightSegment      MdlSbrapiDspvcrRspfsg `xml:"FlightSegment"`
	FlownFlightSegment MdlSbrapiDspvcrRspfls `xml:"FlownFlightSegment"`
}
type MdlSbrapiDspvcrRspfsg struct {
	ResBookDesigCode    string                `xml:"ResBookDesigCode,attr"`
	DepartureDateTime   string                `xml:"DepartureDateTime,attr"`
	DestinationLocation MdlSbrapiDspvcrRsparv `xml:"DestinationLocation"`
	OriginLocation      MdlSbrapiDspvcrRspdep `xml:"OriginLocation"`
	MarketingAirline    MdlSbrapiDspvcrRspair `xml:"MarketingAirline"`
}
type MdlSbrapiDspvcrRspfls struct {
	ResBookDesigCode    string                `xml:"ResBookDesigCode,attr"`
	FlightNumber        string                `xml:"FlightNumber,attr"`
	DepartureDateTime   string                `xml:"DepartureDateTime,attr"`
	DestinationLocation MdlSbrapiDspvcrRsparv `xml:"DestinationLocation"`
	OriginLocation      MdlSbrapiDspvcrRspdep `xml:"OriginLocation"`
}
type MdlSbrapiDspvcrRsparv struct {
	LocationCode string `xml:"LocationCode,attr"`
}
type MdlSbrapiDspvcrRspdep struct {
	LocationCode string `xml:"LocationCode,attr"`
}
type MdlSbrapiDspvcrRspair struct {
	Code         string `xml:"Code,attr"`
	FlightNumber string `xml:"FlightNumber,attr"`
}
type MdlSbrapiDspvcrRsptkx struct {
	AccountingCode string                `xml:"AccountingCode,attr"`
	ETicketNumber  string                `xml:"eTicketNumber,attr"`
	IssueDate      string                `xml:"IssueDate,attr"`
	NumCoupons     string                `xml:"NumCoupons,attr"`
	IssuingAgent   string                `xml:"IssuingAgent,attr"`
	ItineraryRef   MdlSbrapiDspvcrRspitn `xml:"ItineraryRef"`
	TicketData     MdlSbrapiDspvcrRsptkd `xml:"TicketData"`
	CustomerInfo   MdlSbrapiDspvcrRspcsi `xml:"CustomerInfo"`
}
type MdlSbrapiDspvcrRspitn struct {
	ID string `xml:"ID,attr"`
}
type MdlSbrapiDspvcrRsptkd struct {
	ExchangeData MdlSbrapiDspvcrRspexd `xml:"ExchangeData"`
}
type MdlSbrapiDspvcrRspexd struct {
	LocationName string `xml:"LocationName,attr"`
}
type MdlSbrapiDspvcrRspcsi struct {
	PersonName MdlSbrapiDspvcrRspprs `xml:"PersonName"`
}
type MdlSbrapiDspvcrRspprs struct {
	Surname string `xml:"Surname"`
}
