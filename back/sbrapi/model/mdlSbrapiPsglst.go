package mdlSbrapi

import "encoding/xml"

// Request
type MdlSbrapiPsglstReqenv struct {
	XMLName xml.Name              `xml:"soap-env:Envelope"`
	Xmlns   string                `xml:"xmlns:soap-env,attr"`
	Header  MdlSbrapiPsglstReqhdr `xml:"soap-env:Header"`
	Body    MdlSbrapiPsglstReqbdy `xml:"soap-env:Body"`
}
type MdlSbrapiPsglstReqhdr struct {
	MessageHeader MdlSbrapiMsghdrMainob `xml:"eb:MessageHeader"`
	Security      MdlSbrapiPsglstReqscr `xml:"wsse:Security"`
}
type MdlSbrapiPsglstReqscr struct {
	BinarySecurityToken MdlSbrapiPsglstReqbst `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string                `xml:"xmlns:wsse,attr"`
}
type MdlSbrapiPsglstReqbst struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}
type MdlSbrapiPsglstReqbdy struct {
	GetPassengerListRQ MdlSbrapiPsglstReqafl `xml:"GetPassengerListRQ"`
}
type MdlSbrapiPsglstReqafl struct {
	XMLName            xml.Name              `xml:"GetPassengerListRQ"`
	Xmlns              string                `xml:"xmlns,attr"`
	Version            string                `xml:"version,attr"`
	TimeStamp          string                `xml:"timeStamp,attr"`
	MessageID          string                `xml:"messageID,attr"`
	ServiceOption      string                `xml:"serviceOption,attr"`
	Itinerary          MdlSbrapiPsglstReqitn `xml:"Itinerary"`
	DisplayCodeRequest MdlSbrapiPsglstReqdcr `xml:"DisplayCodeRequest>DisplayCodes"`
}
type MdlSbrapiPsglstReqitn struct {
	Airline       string `xml:"Airline"`
	Flight        string `xml:"Flight"`
	DepartureDate string `xml:"DepartureDate"`
	Origin        string `xml:"Origin"`
}
type MdlSbrapiPsglstReqdcr struct {
	Condition    string   `xml:"condition,attr"`
	DisplayCodes []string `xml:"DisplayCode"`
}

// Response
type MdlSbrapiPsglstRspenv struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    MdlSbrapiPsglstRspbdy `xml:"Body"`
}
type MdlSbrapiPsglstRspbdy struct {
	GetPassengerListRS MdlSbrapiPsglstRspgpl `xml:"http://services.sabre.com/checkin/getPassengerList/v4 GetPassengerListRS"`
}
type MdlSbrapiPsglstRspgpl struct {
	ItineraryInfo     MdlSbrapiPsglstRspitf   `xml:"ItineraryInfo"`
	PassengerInfoList []MdlSbrapiPsglstRsppil `xml:"PassengerInfoList>PassengerInfo"`
}
type MdlSbrapiPsglstRspitf struct {
	Itinerary              MdlSbrapiPsglstRspitn   `xml:"Itinerary"`
	DepartureArrival_Dates MdlSbrapiPsglstRspdad   `xml:"DepartureArrival_Dates"`
	DepartureGate          string                  `xml:"DepartureGate"`
	CabinInfoList          []MdlSbrapiPsglstRspcil `xml:"CabinInfoList>CabinInfo"`
}
type MdlSbrapiPsglstRspitn struct {
	Airline       string `xml:"Airline"`
	Flight        string `xml:"Flight"`
	DepartureDate string `xml:"DepartureDate"`
	Origin        string `xml:"Origin"`
	AircraftType  string `xml:"AircraftType"`
}
type MdlSbrapiPsglstRspdad struct {
	ScheduledDepartureDate string `xml:"Scheduled_DepartureDate"`
	EstimatedDepartureDate string `xml:"Estimated_DepartureDate"`
	DepartureTime          string `xml:"DepartureTime"`
	ScheduledArrivalDate   string `xml:"Scheduled_ArrivalDate"`
	EstimatedArrivalDate   string `xml:"Estimated_ArrivalDate"`
	ArrivalTime            string `xml:"ArrivalTime"`
}
type MdlSbrapiPsglstRspcil struct {
	Cabin string `xml:"Cabin"`
	Count int    `xml:"Count"`
}
type MdlSbrapiPsglstRsppil struct {
	NameDetails      MdlSbrapiPsglstRspnme   `xml:"Name_Details"`
	PNRLocator       string                  `xml:"PNRLocator"`
	PassengerID      string                  `xml:"PassengerID"`
	ThruIndicator    string                  `xml:"ThruIndicator"`
	LineNumber       int                     `xml:"LineNumber"`
	GroupCode        string                  `xml:"GroupCode"`
	BookingClass     string                  `xml:"BookingClass"`
	Cabin            string                  `xml:"Cabin"`
	Seat             string                  `xml:"Seat"`
	Destination      string                  `xml:"Destination"`
	BoardingPassFlag string                  `xml:"BoardingPassFlag"`
	PassengerType    string                  `xml:"PassengerType"`
	BagCount         string                  `xml:"BagCount"`
	Indicators       []string                `xml:"Indicators>Indicator"`
	CheckInInfo      MdlSbrapiPsglstRspcki   `xml:"CheckIn_Info"`
	BoardingInfo     MdlSbrapiPsglstRspbdi   `xml:"Boarding_Info"`
	VCRInfo          MdlSbrapiPsglstRspvci   `xml:"VCR_Info"`
	BaggageRouteList []MdlSbrapiPsglstRspbgr `xml:"BaggageRouteList>BaggageRoute"`
	AEDetailsList    MdlSbrapiPsglstRspaed   `xml:"AEDetailsList>AEDetails"`
	EditCodeList     []string                `xml:"EditCodeList>EditCode"`
	FreeTextInfoList []MdlSbrapiPsglstRspfti `xml:"FreeTextInfoList>FreeTextInfo"`
}
type MdlSbrapiPsglstRspnme struct {
	LastName  string `xml:"LastName"`
	FirstName string `xml:"FirstName"`
}
type MdlSbrapiPsglstRspcki struct {
	CheckInNumber int  `xml:"CheckInNumber"`
	CheckInStatus bool `xml:"CheckInStatus"`
}
type MdlSbrapiPsglstRspbdi struct {
	BoardStatus bool `xml:"BoardStatus"`
}
type MdlSbrapiPsglstRspvci struct {
	VCRData MdlSbrapiPsglstRspvcd `xml:"VCR_Data"`
}
type MdlSbrapiPsglstRspvcd struct {
	CouponNumber int    `xml:"CouponNumber"`
	SerialNumber string `xml:"SerialNumber"`
}
type MdlSbrapiPsglstRspbgr struct {
	SegmentID     int    `xml:"SegmentID"`
	Airline       string `xml:"Airline"`
	Flight        string `xml:"Flight"`
	Origin        string `xml:"Origin"`
	Destination   string `xml:"Destination"`
	DepartureDate string `xml:"DepartureDate"`
	DepartureTime string `xml:"DepartureTime"`
	ArrivalDate   string `xml:"ArrivalDate"`
	ArrivalTime   string `xml:"ArrivalTime"`
	SegmentStatus string `xml:"SegmentStatus"`
	PassengerID   string `xml:"PassengerID"`
}
type MdlSbrapiPsglstRspaed struct {
	ItemID           string                `xml:"ItemID"`
	Code             string                `xml:"Code"`
	ATPCOGroupCode   string                `xml:"ATPCOGroupCode"`
	ATPCOSubCode     string                `xml:"ATPCOSubCode"`
	StatusCode       string                `xml:"StatusCode"`
	Description      string                `xml:"Description"`
	QuantityUsed     int32                 `xml:"QuantityUsed"`
	QuantityBought   int32                 `xml:"QuantityBought"`
	PurchaseDate     string                `xml:"PurchaseDate"`
	PurchaseTime     string                `xml:"PurchaseTime"`
	Disassociated    string                `xml:"Disassociated"`
	MarketingAirline string                `xml:"MarketingAirline"`
	MarketingFlight  string                `xml:"MarketingFlight"`
	Origin           string                `xml:"Origin"`
	Destination      string                `xml:"Destination"`
	OperatingAirline string                `xml:"OperatingAirline"`
	OperatingFlight  string                `xml:"OperatingFlight"`
	PriceDetails     MdlSbrapiPsglstRspprd `xml:"PriceDetails"`
}
type MdlSbrapiPsglstRspprd struct {
	TotalPrice      MdlSbrapiPsglstRspprc `xml:"TotalPrice"`
	BasePrice       MdlSbrapiPsglstRspprc `xml:"BasePrice"`
	EquivalentPrice MdlSbrapiPsglstRspprc `xml:"EquivalentPrice"`
}
type MdlSbrapiPsglstRspprc struct {
	CurrencyCode string  `xml:"currencyCode,attr,omitempty"`
	Value        float64 `xml:",chardata"`
}
type MdlSbrapiPsglstRspfti struct {
	EditCode string `xml:"EditCode"`
	TextLine string `xml:"TextLine>Text"`
}
