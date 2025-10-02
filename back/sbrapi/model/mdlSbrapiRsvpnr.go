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
	GetReservationRS MdlSbrapiRsvpnrRspget `xml:"GetReservationRS"`
}
type MdlSbrapiRsvpnrRspget struct {
	Reservation MdlSbrapiRsvpnrRsprsv `xml:"Reservation"`
}
type MdlSbrapiRsvpnrRsprsv struct {
	BookingDetails       MdlSbrapiRsvpnrRspbok `xml:"BookingDetails"`
	PassengerReservation MdlSbrapiRsvpnrRsppsg `xml:"PassengerReservation"`
	POS                  MdlSbrapiRsvpnrRsppos `xml:"POS"`
}

// Booking details (Details)
type MdlSbrapiRsvpnrRspbok struct {
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
	Segments MdlSbrapiRsvpnrRspsgs `xml:"Segments"`
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
