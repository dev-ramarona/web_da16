package mdl_global

import (
	"encoding/xml"
)

type ScreenReqEnvlpe struct {
	XMLName xml.Name        `xml:"soap-env:Envelope"`
	Xmlns   string          `xml:"xmlns:soap-env,attr"`
	Header  ScreenReqHeader `xml:"soap-env:Header"`
	Body    ScreenReqBodyxx `xml:"soap-env:Body"`
}

type ScreenReqHeader struct {
	MessageHeader Msghdr          `xml:"eb:MessageHeader"`
	Security      ScreenReqSecrty `xml:"wsse:Security"`
}

type ScreenReqSecrty struct {
	BinarySecurityToken ScreenReqBsctkn `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string          `xml:"xmlns:wsse,attr"`
}

type ScreenReqBsctkn struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}

type ScreenReqBodyxx struct {
	SabreCommandLLSRQ ScreenReqSbrcmd `xml:"SabreCommandLLSRQ"`
}

type ScreenReqSbrcmd struct {
	XMLName      xml.Name        `xml:"SabreCommandLLSRQ"`
	Xmlns        string          `xml:"xmlns,attr"`
	Version      string          `xml:"Version,attr"`
	NumResponses string          `xml:"NumResponses,attr"`
	Request      ScreenReqRquest `xml:"Request"`
}

type ScreenReqRquest struct {
	Output      string `xml:"Output,attr"`
	HostCommand string `xml:"HostCommand"`
}

// Response

type ScreenRspEnvlpe struct {
	XMLName xml.Name        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    ScreenRspBodyxx `xml:"Body"`
}

type ScreenRspBodyxx struct {
	SabreCommandLLSRS ScreenRspSbrcmd `xml:"http://webservices.sabre.com/sabreXML/2011/10 SabreCommandLLSRS"`
}

type ScreenRspSbrcmd struct {
	Response string `xml:"Response"`
}
