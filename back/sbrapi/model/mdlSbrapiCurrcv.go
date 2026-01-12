package mdlSbrapi

import "encoding/xml"

// Request
type MdlSbrapiCurrcvReqenv struct {
	XMLName xml.Name              `xml:"soap-env:Envelope"`
	Xmlns   string                `xml:"xmlns:soap-env,attr"`
	Header  MdlSbrapiCurrcvReqhdr `xml:"soap-env:Header"`
	Body    MdlSbrapiCurrcvReqbdy `xml:"soap-env:Body"`
}
type MdlSbrapiCurrcvReqhdr struct {
	MessageHeader MdlSbrapiMsghdrMainob `xml:"eb:MessageHeader"`
	Security      MdlSbrapiCurrcvReqscr `xml:"wsse:Security"`
}
type MdlSbrapiCurrcvReqscr struct {
	BinarySecurityToken MdlSbrapiCurrcvReqbst `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string                `xml:"xmlns:wsse,attr"`
}
type MdlSbrapiCurrcvReqbst struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}
type MdlSbrapiCurrcvReqbdy struct {
	DisplayCurrencyRQ MdlSbrapiCurrcvReqafl `xml:"DisplayCurrencyRQ"`
}
type MdlSbrapiCurrcvReqafl struct {
	XMLName      xml.Name `xml:"DisplayCurrencyRQ"`
	Xmlns        string   `xml:"xmlns,attr"`
	Version      string   `xml:"Version,attr"`
	CountryCode  string   `xml:"CountryCode"`
	CurrencyCode string   `xml:"CurrencyCode"`
}

// Response
type MdlSbrapiCurrcvRspenv struct {
	XMLName xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    MdlSbrapiCurrcvRspbdy `xml:"Body"`
}
type MdlSbrapiCurrcvRspbdy struct {
	DisplayCurrencyRS MdlSbrapiCurrcvRspdsp `xml:"http://webservices.sabre.com/sabreXML/2011/10 DisplayCurrencyRS"`
}
type MdlSbrapiCurrcvRspdsp struct {
	Name    string                  `xml:"Name,attr"`
	Country []MdlSbrapiCurrcvRspctr `xml:"Country"`
}
type MdlSbrapiCurrcvRspctr struct {
	CurrencyCode  string  `xml:"CurrencyCode"`
	CurrencyName  string  `xml:"CurrencyName"`
	DecimalPlaces int32   `xml:"DecimalPlaces"`
	Rate          float64 `xml:"Rate"`
}
