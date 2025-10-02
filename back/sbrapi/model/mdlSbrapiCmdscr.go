package mdlSbrapi

import "encoding/xml"

// Request
type MdlSbrapiCmdscrReqenv struct {
	XMLName xml.Name              `xml:"soap-env:Envelope"`
	Xmlns   string                `xml:"xmlns:soap-env,attr"`
	Header  MdlSbrapiCmdscrReqhdr `xml:"soap-env:Header"`
	Body    MdlSbrapiCmdscrReqbdy `xml:"soap-env:Body"`
}
type MdlSbrapiCmdscrReqhdr struct {
	MessageHeader MdlSbrapiMsghdrMainob `xml:"eb:MessageHeader"`
	Security      MdlSbrapiCmdscrReqscr `xml:"wsse:Security"`
}
type MdlSbrapiCmdscrReqscr struct {
	BinarySecurityToken MdlSbrapiCmdscrReqbst `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string                `xml:"xmlns:wsse,attr"`
}
type MdlSbrapiCmdscrReqbst struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}
type MdlSbrapiCmdscrReqbdy struct {
	SabreCommandLLSRQ MdlSbrapiCmdscrReqssc `xml:"SabreCommandLLSRQ"`
}
type MdlSbrapiCmdscrReqssc struct {
	XMLName      xml.Name              `xml:"SabreCommandLLSRQ"`
	Xmlns        string                `xml:"xmlns,attr"`
	Version      string                `xml:"Version,attr"`
	NumResponses string                `xml:"NumResponses,attr"`
	Request      MdlSbrapiCmdscrReqreq `xml:"Request"`
}
type MdlSbrapiCmdscrReqreq struct {
	Output      string `xml:"Output,attr"`
	HostCommand string `xml:"HostCommand"`
}

// Response
type MdlSbrapiCmdscrRspenv struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    MdlSbrapiCmdscrRspbdy `xml:"Body"`
}
type MdlSbrapiCmdscrRspbdy struct {
	SabreCommandLLSRS MdlSbrapiCmdscrRsprsp `xml:"http://webservices.sabre.com/sabreXML/2011/10 SabreCommandLLSRS"`
}
type MdlSbrapiCmdscrRsprsp struct {
	Response string `xml:"Response"`
}
