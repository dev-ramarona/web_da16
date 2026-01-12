package mdlSbrapi

import "encoding/xml"

// Request
type MdlSbrapiMilegeReqenv struct {
	XMLName xml.Name              `xml:"soap-env:Envelope"`
	Xmlns   string                `xml:"xmlns:soap-env,attr"`
	Header  MdlSbrapiMilegeReqhdr `xml:"soap-env:Header"`
	Body    MdlSbrapiMilegeReqbdy `xml:"soap-env:Body"`
}
type MdlSbrapiMilegeReqhdr struct {
	MessageHeader MdlSbrapiMsghdrMainob `xml:"eb:MessageHeader"`
	Security      MdlSbrapiMilegeReqscr `xml:"wsse:Security"`
}
type MdlSbrapiMilegeReqscr struct {
	BinarySecurityToken MdlSbrapiMilegeReqbst `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string                `xml:"xmlns:wsse,attr"`
}
type MdlSbrapiMilegeReqbst struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}
type MdlSbrapiMilegeReqbdy struct {
	MileageRQ MdlSbrapiMilegeReqmlq `xml:"MileageRQ"`
}
type MdlSbrapiMilegeReqmlq struct {
	XMLName                      xml.Name              `xml:"MileageRQ"`
	Xmlns                        string                `xml:"xmlns,attr"`
	Version                      string                `xml:"Version,attr"`
	OriginDestinationInformation MdlSbrapiMilegeReqodi `xml:"OriginDestinationInformation"`
}
type MdlSbrapiMilegeReqodi struct {
	DestinationLocation []MdlSbrapiMilegeReqdlc `xml:"DestinationLocation"`
	OriginLocation      MdlSbrapiMilegeReqolc   `xml:"OriginLocation"`
}
type MdlSbrapiMilegeReqdlc struct {
	LocationCode string `xml:"LocationCode,attr"`
	RPH          string `xml:"RPH,attr"`
}
type MdlSbrapiMilegeReqolc struct {
	LocationCode string `xml:"LocationCode,attr"`
}

// Response
type MdlSbrapiMilegeRspenv struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    MdlSbrapiMilegeRspbdy `xml:"Body"`
}
type MdlSbrapiMilegeRspbdy struct {
	MileageRS MdlSbrapiMilegeRspmlg `xml:"MileageRS"`
}
type MdlSbrapiMilegeRspmlg struct {
	OriginDestinationInformation MdlSbrapiMilegeRspodi `xml:"OriginDestinationInformation"`
}
type MdlSbrapiMilegeRspodi struct {
	DestinationLocation []MdlSbrapiMilegeRspdsl `xml:"DestinationLocation"`
}
type MdlSbrapiMilegeRspdsl struct {
	LocationCode         string `xml:"LocationCode,attr"`
	TicketedPointMileage string `xml:"TicketedPointMileage"`
}
