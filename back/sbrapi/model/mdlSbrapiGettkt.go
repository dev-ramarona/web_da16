package mdlSbrapi

import "encoding/xml"

// Request
type MdlSbrapiGettktReqenv struct {
	XMLName xml.Name              `xml:"soap-env:Envelope"`
	Xmlns   string                `xml:"xmlns:soap-env,attr"`
	Header  MdlSbrapiGettktReqhdr `xml:"soap-env:Header"`
	Body    MdlSbrapiGettktReqbdy `xml:"soap-env:Body"`
}
type MdlSbrapiGettktReqhdr struct {
	MessageHeader MdlSbrapiMsghdrMainob `xml:"eb:MessageHeader"`
	Security      MdlSbrapiGettktReqscr `xml:"wsse:Security"`
}
type MdlSbrapiGettktReqscr struct {
	BinarySecurityToken MdlSbrapiGettktReqbst `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string                `xml:"xmlns:wsse,attr"`
}
type MdlSbrapiGettktReqbst struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}
type MdlSbrapiGettktReqbdy struct {
	GetTicketingDocumentRQ MdlSbrapiGettktReqrsv `xml:"ns2:GetTicketingDocumentRQ"`
}
type MdlSbrapiGettktReqrsv struct {
	XMLName          xml.Name              `xml:"ns2:GetTicketingDocumentRQ"`
	XmlnsNs2         string                `xml:"xmlns:ns2,attr"`
	Xmlns            string                `xml:"xmlns,attr"`
	Version          string                `xml:"Version,attr"`
	POS              struct{}              `xml:"POS"`
	SearchParameters MdlSbrapiGettktReqspm `xml:"ns2:SearchParameters"`
}
type MdlSbrapiGettktReqspm struct {
	ResultType            string   `xml:"resultType,attr"`
	TicketingProvider     string   `xml:"ns2:TicketingProvider"`
	DocumentNumber        string   `xml:"ns2:DocumentNumber"`
	CustomResponseDetails []string `xml:"ns2:CustomResponseDetails"`
}

// Response
type MdlSbrapiGettktRspenv struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    MdlSbrapiGettktRspbdy `xml:"Body"`
}
type MdlSbrapiGettktRspbdy struct {
	GetTicketingDocumentRS MdlSbrapiGettktRspget `xml:"GetTicketingDocumentRS"`
}
type MdlSbrapiGettktRspget struct {
	CustomDetails MdlSbrapiGettktRspcsd `xml:"CustomDetails"`
}
type MdlSbrapiGettktRspcsd struct {
	Agent  MdlSbrapiGettktRspagt `xml:"Agent"`
	Ticket MdlSbrapiGettktRsptkt `xml:"Ticket"`
	Number string                `xml:"number,attr"`
}
type MdlSbrapiGettktRspagt struct {
	Duty            string `xml:"duty,attr"`
	Sine            string `xml:"sine,attr"`
	StationLocation string `xml:"StationLocation"`
	StationNumber   string `xml:"StationNumber"`
	WorkLocation    string `xml:"WorkLocation"`
	HomeLocation    string `xml:"HomeLocation"`
	Lniata          string `xml:"Lniata"`
	EmployeeNumber  string `xml:"EmployeeNumber"`
}
type MdlSbrapiGettktRsptkt struct {
	Details         MdlSbrapiGettktRspdtl   `xml:"Details"`
	ServiceCoupon   []MdlSbrapiGettktRspsvc `xml:"ServiceCoupon"`
	Amounts         MdlSbrapiGettktRspbse   `xml:"Amounts>New>Base>Amount"`
	FareCalculation string                  `xml:"FareCalculation>New"`
}
type MdlSbrapiGettktRspdtl struct {
	TourNumber         string                `xml:"TourNumber"`
	Reservation        MdlSbrapiGettktRsprsv `xml:"Reservation>Sabre"`
	LocalIssueDateTime string                `xml:"LocalIssueDateTime"`
}
type MdlSbrapiGettktRsprsv struct {
	CreateDate string `xml:"createDate,attr"`
}
type MdlSbrapiGettktRspsvc struct {
	Coupon                int                   `xml:"coupon,attr"`
	MarketingProvider     string                `xml:"MarketingProvider"`
	MarketingFlightNumber string                `xml:"MarketingFlightNumber"`
	ClassOfService        string                `xml:"ClassOfService"`
	FareBasis             string                `xml:"FareBasis"`
	StartLocation         string                `xml:"StartLocation"`
	StartDateTime         string                `xml:"StartDateTime"`
	EndLocation           string                `xml:"EndLocation"`
	EndDateTime           string                `xml:"EndDateTime"`
	BookingStatus         string                `xml:"BookingStatus"`
	CurrentStatus         string                `xml:"CurrentStatus"`
	PreviousStatus        string                `xml:"PreviousStatus"`
	FlownCoupon           MdlSbrapiGettktRspflc `xml:"FlownCoupon"`
	BagAllowance          string                `xml:"BagAllowance"`
}
type MdlSbrapiGettktRspflc struct {
	MarketingProvider     string `xml:"MarketingProvider"`
	OperatingProvider     string `xml:"OperatingProvider"`
	OperatingFlightNumber string `xml:"OperatingFlightNumber"`
	ClassOfService        string `xml:"ClassOfService"`
	DepartureCity         string `xml:"DepartureCity"`
	DepartureDateTime     string `xml:"DepartureDateTime"`
	ArrivalCity           string `xml:"ArrivalCity"`
	FlightOriginalDate    string `xml:"FlightOriginalDate"`
}
type MdlSbrapiGettktRspbse struct {
	CurrencyCode string `xml:"currencyCode,attr"`
}
