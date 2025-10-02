package mdlSbrapi

import "encoding/xml"

// Request
type MdlSbrapiCrtssnReqenv struct {
	XMLName xml.Name              `xml:"soap-env:Envelope"`
	Xmlns   string                `xml:"xmlns:soap-env,attr"`
	Header  MdlSbrapiCrtssnReqhdr `xml:"soap-env:Header"`
	Body    MdlSbrapiCrtssnReqbdy `xml:"soap-env:Body"`
}
type MdlSbrapiCrtssnReqhdr struct {
	MessageHeader MdlSbrapiMsghdrMainob `xml:"eb:MessageHeader"`
	Security      MdlSbrapiCrtssnReqscr `xml:"wsse:Security"`
}
type MdlSbrapiCrtssnReqscr struct {
	UsernameToken MdlSbrapiCrtssnRequsr `xml:"wsse:UsernameToken"`
	XmlnsWsse     string                `xml:"xmlns:wsse,attr"`
}
type MdlSbrapiCrtssnRequsr struct {
	Username     string `xml:"wsse:Username"`
	Password     string `xml:"wsse:Password"`
	Organization string `xml:"Organization"`
	Domain       string `xml:"Domain"`
}
type MdlSbrapiCrtssnReqbdy struct {
	SessionCreateRQ MdlSbrapiCrtssnReqcrt `xml:"sws:SessionCreateRQ"`
}
type MdlSbrapiCrtssnReqcrt struct {
	Version string                `xml:"Version,attr"`
	POS     MdlSbrapiCrtssnReqpos `xml:"POS"`
	Xmlns   string                `xml:"xmlns:sws,attr"`
}
type MdlSbrapiCrtssnReqpos struct {
	Source MdlSbrapiCrtssnReqsrc `xml:"Source"`
}
type MdlSbrapiCrtssnReqsrc struct {
	PseudoCityCode string `xml:"PseudoCityCode,attr"`
}

// Response
type MdlSbrapiCrtssnRspenv struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  MdlSbrapiCrtssnRsphdr `xml:"Header"`
}
type MdlSbrapiCrtssnRsphdr struct {
	Security MdlSbrapiCrtssnRspscr `xml:"http://schemas.xmlsoap.org/ws/2002/12/secext Security"`
}
type MdlSbrapiCrtssnRspscr struct {
	BinarySecurityToken MdlSbrapiCrtssnRsptkn `xml:"BinarySecurityToken"`
}
type MdlSbrapiCrtssnRsptkn struct {
	ValueType    string `xml:"valueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}
