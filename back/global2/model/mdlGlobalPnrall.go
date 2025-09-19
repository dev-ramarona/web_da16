package mdl_global

import "encoding/xml"

type PnrallReqEnvlpe struct {
	XMLName xml.Name        `xml:"soap-env:Envelope"`
	Xmlns   string          `xml:"xmlns:soap-env,attr"`
	Header  PnrallReqHeader `xml:"soap-env:Header"`
	Body    PnrallReqBodyxx `xml:"soap-env:Body"`
}

type PnrallReqHeader struct {
	MessageHeader Msghdr          `xml:"eb:MessageHeader"`
	Security      PnrallReqSecrty `xml:"wsse:Security"`
}

type PnrallReqSecrty struct {
	BinarySecurityToken PnrallReqBsctkn `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string          `xml:"xmlns:wsse,attr"`
}

type PnrallReqBsctkn struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}

type PnrallReqBodyxx struct {
	GetReservationRQ PnrallReqGetRsvtion `xml:"GetReservationRQ"`
}

type PnrallReqGetRsvtion struct {
	XMLName       xml.Name        `xml:"GetReservationRQ"`
	Xmlns         string          `xml:"xmlns,attr"`
	Version       string          `xml:"Version,attr"`
	Locator       string          `xml:"Locator"`
	RequestType   string          `xml:"RequestType"`
	ReturnOptions PnrallReqRtropt `xml:"ReturnOptions"`
}

type PnrallReqRtropt struct {
	SubjectAreas   PnrallReqSbjare `xml:"SubjectAreas"`
	ViewName       string          `xml:"ViewName"`
	ResponseFormat string          `xml:"ResponseFormat"`
}

type PnrallReqSbjare struct {
	SubjectArea []string `xml:"SubjectArea"`
}

// Response Interline

type IntrlnRspEnvlpe struct {
	XMLName xml.Name        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    IntrlnRspBodyxx `xml:"Body"`
}

type IntrlnRspBodyxx struct {
	GetReservationRS IntrlnRspGetRsvtns `xml:"GetReservationRS"`
}

type IntrlnRspGetRsvtns struct {
	Reservation IntrlnRspRsvtion `xml:"Reservation"`
}

type IntrlnRspRsvtion struct {
	POS IntrlnRspPosxxx `xml:"POS"`
}

type IntrlnRspPosxxx struct {
	Source IntrlnRspSource `xml:"Source"`
}

type IntrlnRspSource struct {
	TTYRecordLocator IntrlnRspTtrlct `xml:"TTYRecordLocator"`
}

type IntrlnRspTtrlct struct {
	CRSCode       string `xml:"CRSCode"`
	RecordLocator string `xml:"RecordLocator"`
}

// Response Itenary

type ItnaryRspEnvlpe struct {
	XMLName xml.Name        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    ItnaryRspBodyxx `xml:"Body"`
}

type ItnaryRspBodyxx struct {
	GetReservationRS ItnaryRspGetRsvtns `xml:"GetReservationRS"`
}

type ItnaryRspGetRsvtns struct {
	Reservation ItnaryRspRsvtion `xml:"Reservation"`
}

type ItnaryRspRsvtion struct {
	BookingDetails       RemarksRspBokdtl `xml:"BookingDetails"`
	PassengerReservation ItnaryRspPsgrsv  `xml:"PassengerReservation"`
}

type ItnaryRspPsgrsv struct {
	Segments ItnaryRspSgmnts `xml:"Segments"`
}

type ItnaryRspSgmnts struct {
	Segment []ItnaryRspSegmnt `xml:"Segment"`
}

type ItnaryRspSegmnt struct {
	Air ItnaryRspAirxxx `xml:"Air"`
}

type ItnaryRspAirxxx struct {
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

// Response Remarks

type RemarksRspEnvlpe struct {
	XMLName xml.Name         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    RemarksRspBodyxx `xml:"Body"`
}

type RemarksRspBodyxx struct {
	GetReservationRS RemarksRspRsvtns `xml:"GetReservationRS"`
}

type RemarksRspRsvtns struct {
	Reservation RemarksRspRsvtion `xml:"Reservation"`
}

type RemarksRspRsvtion struct {
	BookingDetails RemarksRspBokdtl `xml:"BookingDetails"`
}

type RemarksRspBokdtl struct {
	SystemCreationTimestamp string           `xml:"SystemCreationTimestamp"`
	CreationAgentID         string           `xml:"CreationAgentID"`
	FlightsRange            RemarksRspFlrnge `xml:"FlightsRange"`
	DivideSplitDetails      RemarksRspSptdtl `xml:"DivideSplitDetails"`
}

type RemarksRspFlrnge struct {
	Start string `xml:"Start,attr"`
	End   string `xml:"End,attr"`
}

type RemarksRspSptdtl struct {
	XMLName   xml.Name           `xml:"DivideSplitDetails"`
	Itemslice []RemarksRspSptrcd `xml:",any"`
}

type RemarksRspSptrcd struct {
	XMLName               xml.Name `xml:""`
	DivideTimestamp       string   `xml:"DivideTimestamp"`
	RecordLocator         string   `xml:"RecordLocator"`
	OriginalNumberOfPax   int      `xml:"OriginalNumberOfPax"`
	CurrentNumberOfPax    int      `xml:"CurrentNumberOfPax"`
	CurrentPassengerNames string   `xml:"CurrentPassengerNames"`
}

// Response Remarks Itinerary and Interline

type ItrmrlRspEnvlpe struct {
	XMLName xml.Name        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    ItrmrlRspBodyxx `xml:"Body"`
}

type ItrmrlRspBodyxx struct {
	GetReservationRS ItrmrlRspRsvtns `xml:"GetReservationRS"`
}

type ItrmrlRspRsvtns struct {
	Reservation ItrmrlRspRsvtion `xml:"Reservation"`
}

type ItrmrlRspRsvtion struct {
	BookingDetails       RemarksRspBokdtl `xml:"BookingDetails"`
	PassengerReservation ItnaryRspPsgrsv  `xml:"PassengerReservation"`
	POS                  IntrlnRspPosxxx  `xml:"POS"`
}
