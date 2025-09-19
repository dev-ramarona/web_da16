package mdl_global

import "encoding/xml"

type SssionReqEnvlpe struct {
	XMLName xml.Name        `xml:"soap-env:Envelope"`
	Xmlns   string          `xml:"xmlns:soap-env,attr"`
	Header  SssionReqHeader `xml:"soap-env:Header"`
	Body    SssionReqBodyxx `xml:"soap-env:Body"`
}

type SssionReqHeader struct {
	MessageHeader Msghdr          `xml:"eb:MessageHeader"`
	Security      SssionReqSecrty `xml:"wsse:Security"`
}

type SssionReqSecrty struct {
	UsernameToken SssionReqUsrtkn `xml:"wsse:UsernameToken"`
	XmlnsWsse     string          `xml:"xmlns:wsse,attr"`
}

type SssionReqUsrtkn struct {
	Username     string `xml:"wsse:Username"`
	Password     string `xml:"wsse:Password"`
	Organization string `xml:"Organization"`
	Domain       string `xml:"Domain"`
}

type SssionReqBodyxx struct {
	SessionCreateRQ SssionReqSsncrt `xml:"sws:SessionCreateRQ"`
}

type SssionReqSsncrt struct {
	Version string          `xml:"Version,attr"`
	POS     SssionReqPosxxx `xml:"POS"`
	Xmlns   string          `xml:"xmlns:sws,attr"`
}

type SssionReqPosxxx struct {
	Source SssionReqSource `xml:"Source"`
}

type SssionReqSource struct {
	PseudoCityCode string `xml:"PseudoCityCode,attr"`
}

// Response

type SssionRspEnvlpe struct {
	XMLName xml.Name        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  SssionRspHeader `xml:"Header"`
}

type SssionRspHeader struct {
	Security SssionRspSecrty `xml:"http://schemas.xmlsoap.org/ws/2002/12/secext Security"`
}

type SssionRspSecrty struct {
	BinarySecurityToken SssionRspScrtkn `xml:"BinarySecurityToken"`
}

type SssionRspScrtkn struct {
	ValueType    string `xml:"valueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}

// Close session

type ClsssnReqEnvlpe struct {
	XMLName xml.Name        `xml:"soap-env:Envelope"`
	Xmlns   string          `xml:"xmlns:soap-env,attr"`
	Header  ClsssnReqHeader `xml:"soap-env:Header"`
	Body    ClsssnReqBodyxx `xml:"soap-env:Body"`
}

type ClsssnReqHeader struct {
	MessageHeader Msghdr          `xml:"eb:MessageHeader"`
	Security      ClsssnReqSecrty `xml:"wsse:Security"`
}

type ClsssnReqSecrty struct {
	BinarySecurityToken ClsssnReqBsctkn `xml:"wsse:BinarySecurityToken"`
	XmlnsWsse           string          `xml:"xmlns:wsse,attr"`
}

type ClsssnReqBsctkn struct {
	ValueType    string `xml:"ValueType,attr"`
	EncodingType string `xml:"EncodingType,attr"`
	Token        string `xml:",chardata"`
}

type ClsssnReqBodyxx struct {
	SessionCloseRQ ClsssnReqFllist `xml:"SessionCloseRQ"`
}

type ClsssnReqFllist struct {
	POS ClsssnReqPosxxx `xml:"POS"`
}

type ClsssnReqPosxxx struct {
	Source ClsssnReqSource `xml:"Source"`
}

type ClsssnReqSource struct {
	PseudoCityCode string `xml:"PseudoCityCode,attr"`
}
