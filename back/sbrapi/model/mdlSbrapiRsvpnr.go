package mdlSbrapi

import "encoding/xml"

// Request
type MdlSbrapiRsvpnrReqenv struct {
	XMLName xml.Name              `xml:"soap-env:Envelope"`
	Xmlns   string                `xml:"xmlns:soap-env,attr"`
	Header  MdlSbrapiRsvpnrReqhdr `xml:"soap-env:Header"`
	Body    MdlSbrapiRsvpnrReqbdy `xml:"soap-env:Body"`
}
type MdlSbrapiRsvpnrReqhdr struct {
	MessageHeader MdlSbrapiMsghdrMainob `xml:"eb:MessageHeader"`
	Security      MdlSbrapiRsvpnrReqscr `xml:"wsse:Security"`
}
type MdlSbrapiRsvpnrReqscr struct {
	BinarySecurityToken MdlSbrapiRsvpnrReqbst `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string                `xml:"xmlns:wsse,attr"`
}
type MdlSbrapiRsvpnrReqbst struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}
type MdlSbrapiRsvpnrReqbdy struct {
	GetReservationRQ MdlSbrapiRsvpnrReqrsv `xml:"GetReservationRQ"`
}
type MdlSbrapiRsvpnrReqrsv struct {
	XMLName       xml.Name              `xml:"GetReservationRQ"`
	Xmlns         string                `xml:"xmlns,attr"`
	Version       string                `xml:"Version,attr"`
	Locator       string                `xml:"Locator"`
	RequestType   string                `xml:"RequestType"`
	ReturnOptions MdlSbrapiRsvpnrReqopt `xml:"ReturnOptions"`
}
type MdlSbrapiRsvpnrReqopt struct {
	SubjectAreas   MdlSbrapiRsvpnrReqsbj `xml:"SubjectAreas"`
	ViewName       string                `xml:"ViewName"`
	ResponseFormat string                `xml:"ResponseFormat"`
}
type MdlSbrapiRsvpnrReqsbj struct {
	SubjectArea []string `xml:"SubjectArea"`
}

// Response
type MdlSbrapiRsvpnrRspenv struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    MdlSbrapiRsvpnrRspbdy `xml:"Body"`
}
type MdlSbrapiRsvpnrRspbdy struct {
	GetReservationRS MdlSbrapiRsvpnrRspget `xml:"http://webservices.sabre.com/pnrbuilder/v1_19 GetReservationRS"`
}
type MdlSbrapiRsvpnrRspget struct {
	Reservation MdlSbrapiRsvpnrRsprsv `xml:"Reservation"`
}
type MdlSbrapiRsvpnrRsprsv struct {
	BookingDetails       MdlSbrapiRsvpnrRspbok `xml:"BookingDetails"`
	PassengerReservation MdlSbrapiRsvpnrRsppsg `xml:"PassengerReservation"`
	Remarks              MdlSbrapiRsvpnrRsprmk `xml:"Remarks>Remark"`
	POS                  MdlSbrapiRsvpnrRsppos `xml:"POS"`
}

// Booking details (Details)
type MdlSbrapiRsvpnrRspbok struct {
	RecordLocator           string                `xml:"RecordLocator"`
	SystemCreationTimestamp string                `xml:"SystemCreationTimestamp"`
	CreationAgentID         string                `xml:"CreationAgentID"`
	FlightsRange            MdlSbrapiRsvpnrRspflr `xml:"FlightsRange"`
	DivideSplitDetails      MdlSbrapiRsvpnrRspdvd `xml:"DivideSplitDetails"`
}
type MdlSbrapiRsvpnrRspflr struct {
	Start string `xml:"Start,attr"`
	End   string `xml:"End,attr"`
}
type MdlSbrapiRsvpnrRspdvd struct {
	XMLName   xml.Name                `xml:"DivideSplitDetails"`
	Itemslice []MdlSbrapiRsvpnrRspits `xml:",any"`
}
type MdlSbrapiRsvpnrRspits struct {
	XMLName               xml.Name `xml:""`
	DivideTimestamp       string   `xml:"DivideTimestamp"`
	RecordLocator         string   `xml:"RecordLocator"`
	OriginalNumberOfPax   int      `xml:"OriginalNumberOfPax"`
	CurrentNumberOfPax    int      `xml:"CurrentNumberOfPax"`
	CurrentPassengerNames string   `xml:"CurrentPassengerNames"`
}

// Passenger Reservation (Itenary)
type MdlSbrapiRsvpnrRsppsg struct {
	Segments         MdlSbrapiRsvpnrRspsgs `xml:"Segments"`
	TicketingInfo    MdlSbrapiRsvpnrRsptki `xml:"TicketingInfo"`
	ItineraryPricing MdlSbrapiRsvpnrRspitp `xml:"ItineraryPricing"`
}
type MdlSbrapiRsvpnrRspsgs struct {
	Segment []MdlSbrapiRsvpnrRspseg `xml:"Segment"`
}
type MdlSbrapiRsvpnrRspseg struct {
	Air MdlSbrapiRsvpnrRspair `xml:"Air"`
}
type MdlSbrapiRsvpnrRspair struct {
	DepartureAirport          string `xml:"DepartureAirport"`
	ArrivalAirport            string `xml:"ArrivalAirport"`
	MarketingAirlineCode      string `xml:"MarketingAirlineCode"`
	OperatingAirlineCode      string `xml:"OperatingAirlineCode"`
	OperatingAirlineShortName string `xml:"OperatingAirlineShortName"`
	MarketingClassOfService   string `xml:"MarketingClassOfService"`
	OperatingClassOfService   string `xml:"OperatingClassOfService"`
	MarketingFlightNumber     string `xml:"MarketingFlightNumber"`
	OperatingFlightNumber     string `xml:"OperatingFlightNumber"`
	AirlineRefId              string `xml:"AirlineRefId"`
	DepartureDateTime         string `xml:"DepartureDateTime"`
	ActionCode                string `xml:"ActionCode"`
}

// Remarks (Remarks)
type MdlSbrapiRsvpnrRsprmk struct {
	RemarkLines []string `xml:"RemarkLines>RemarkLine>Text"`
}

// POS (Interline)
type MdlSbrapiRsvpnrRsppos struct {
	Source MdlSbrapiRsvpnrRspsrc `xml:"Source"`
}
type MdlSbrapiRsvpnrRspsrc struct {
	TTYRecordLocator MdlSbrapiRsvpnrRsptty `xml:"TTYRecordLocator"`
}
type MdlSbrapiRsvpnrRsptty struct {
	CRSCode       string `xml:"CRSCode"`
	RecordLocator string `xml:"RecordLocator"`
}

// Ticketing info (Ticketing)
type MdlSbrapiRsvpnrRsptki struct {
	TicketDetails []MdlSbrapiRsvpnrRsptkd `xml:"TicketDetails"`
}
type MdlSbrapiRsvpnrRsptkd struct {
	TransactionIndicator string `xml:"TransactionIndicator"`
	TicketNumber         string `xml:"TicketNumber"`
	PassengerName        string `xml:"PassengerName"`
	AgencyLocation       string `xml:"AgencyLocation"`
	DutyCode             string `xml:"DutyCode"`
	AgentSine            string `xml:"AgentSine"`
	Timestamp            string `xml:"Timestamp"`
}

// TPRICING_INFORMATION
type MdlSbrapiRsvpnrRspitp struct {
	PricedItinerary []MdlSbrapiRsvpnrRspti `xml:"PricedItinerary>AirItineraryPricingInfo"`
}
type MdlSbrapiRsvpnrRspti struct {
	ItinTotalFare     MdlSbrapiRsvpnrRspitf `xml:"ItinTotalFare>Base"`
	PTC_FareBreakdown MdlSbrapiRsvpnrRspptc `xml:"PTC_FareBreakdown"`
}
type MdlSbrapiRsvpnrRspitf struct {
	CurrencyCode string `xml:"currencyCode,attr"`
}
type MdlSbrapiRsvpnrRspptc struct {
	PassengerTypeQuantity MdlSbrapiRsvpnrRspptq   `xml:"PassengerTypeQuantity"`
	FareCalc              string                  `xml:"FareCalc"`
	FareComponent         []MdlSbrapiRsvpnrRspfrc `xml:"FareComponent"`
	FlightSegment         []MdlSbrapiRsvpnrRspfls `xml:"FlightSegment"`
}
type MdlSbrapiRsvpnrRspptq struct {
	Code string `xml:"Code,attr"`
}
type MdlSbrapiRsvpnrRspfrc struct {
	FareBasisCode       string `xml:"FareBasisCode,attr"`
	Amount              string `xml:"Amount,attr"`
	GoverningCarrier    string `xml:"GoverningCarrier,attr"`
	FareComponentNumber int    `xml:"FareComponentNumber,attr"`
}
type MdlSbrapiRsvpnrRspfls struct {
	RPH               int    `xml:"RPH,attr"`
	DepartureDateTime string `xml:"DepartureDateTime"`
	ResBookDesigCode  string `xml:"ResBookDesigCode"`
	FlightNumber      string `xml:"FlightNumber"`
	AirPort           string `xml:"AirPort"`
	OperatingAirline  string `xml:"OperatingAirline"`
	FareBasisCode     string `xml:"FareBasisCode"`
}
