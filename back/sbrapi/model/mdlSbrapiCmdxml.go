package mdlSbrapi

import "encoding/xml"

// Request
type MdlSbrapiCmdxmlReqenv struct {
	XMLName xml.Name              `xml:"soap-env:Envelope"`
	Xmlns   string                `xml:"xmlns:soap-env,attr"`
	Header  MdlSbrapiCmdxmlReqhdr `xml:"soap-env:Header"`
	Body    MdlSbrapiCmdxmlReqbdy `xml:"soap-env:Body"`
}
type MdlSbrapiCmdxmlReqhdr struct {
	MessageHeader MdlSbrapiMsghdrMainob `xml:"eb:MessageHeader"`
	Security      MdlSbrapiCmdxmlReqscr `xml:"wsse:Security"`
}
type MdlSbrapiCmdxmlReqscr struct {
	BinarySecurityToken MdlSbrapiCmdxmlReqbst `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string                `xml:"xmlns:wsse,attr"`
}
type MdlSbrapiCmdxmlReqbst struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}
type MdlSbrapiCmdxmlReqbdy struct {
	SabreCommandLLSRQ MdlSbrapiCmdxmlReqssc `xml:"SabreCommandLLSRQ"`
}
type MdlSbrapiCmdxmlReqssc struct {
	XMLName      xml.Name              `xml:"SabreCommandLLSRQ"`
	Xmlns        string                `xml:"xmlns,attr"`
	Version      string                `xml:"Version,attr"`
	NumResponses string                `xml:"NumResponses,attr"`
	Request      MdlSbrapiCmdxmlReqreq `xml:"Request"`
}
type MdlSbrapiCmdxmlReqreq struct {
	Output      string `xml:"Output,attr"`
	HostCommand string `xml:"HostCommand"`
}

// Response
type MdlSbrapiCmdxmlRspenv struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    MdlSbrapiCmdxmlRspbdy `xml:"Body"`
}
type MdlSbrapiCmdxmlRspbdy struct {
	SabreCommandLLSRS MdlSbrapiCmdxmlRsprsp `xml:"http://webservices.sabre.com/sabreXML/2011/10 SabreCommandLLSRS"`
}
type MdlSbrapiCmdxmlRsprsp struct {
	Response string `xml:"Response"`
}
