package mdlSbrapi

import "encoding/xml"

type MdlSbrapiClsssnReqenv struct {
	XMLName xml.Name              `xml:"soap-env:Envelope"`
	Xmlns   string                `xml:"xmlns:soap-env,attr"`
	Header  MdlSbrapiClsssnReqhdr `xml:"soap-env:Header"`
	Body    MdlSbrapiClsssnReqbdy `xml:"soap-env:Body"`
}
type MdlSbrapiClsssnReqhdr struct {
	MessageHeader MdlSbrapiMsghdrMainob `xml:"eb:MessageHeader"`
	Security      MdlSbrapiClsssnReqscr `xml:"wsse:Security"`
}
type MdlSbrapiClsssnReqscr struct {
	BinarySecurityToken MdlSbrapiClsssnReqbst `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string                `xml:"xmlns:wsse,attr"`
}
type MdlSbrapiClsssnReqbst struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}
type MdlSbrapiClsssnReqbdy struct {
	SessionCloseRQ MdlSbrapiClsssnReqssc `xml:"SessionCloseRQ"`
}
type MdlSbrapiClsssnReqssc struct {
	POS MdlSbrapiClsssnReqpos `xml:"POS"`
}
type MdlSbrapiClsssnReqpos struct {
	Source MdlSbrapiClsssnReqsrc `xml:"Source"`
}
type MdlSbrapiClsssnReqsrc struct {
	PseudoCityCode string `xml:"PseudoCityCode,attr"`
}
