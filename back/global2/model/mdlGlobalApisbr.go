package mdl_global

type Msghdr struct {
	From           Partyx  `xml:"eb:From"`
	To             Partyx  `xml:"eb:To"`
	CPAId          string  `xml:"eb:CPAId"`
	ConversationId string  `xml:"eb:ConversationId,omitempty"`
	Service        Service `xml:"eb:Service"`
	Action         string  `xml:"eb:Action"`
	MessageData    Msgdta  `xml:"eb:MessageData"`
	XmlnsEb        string  `xml:"xmlns:eb,attr"`
	MustUnderstand string  `xml:"soap-env:mustUnderstand,attr"`
	Version        string  `xml:"eb:version,attr"`
}

type Partyx struct {
	PartyId string `xml:"eb:PartyId"`
}

type Service struct {
	Type string `xml:"eb:type,attr"`
	Name string `xml:",chardata"`
}

type Msgdta struct {
	MessageId string `xml:"eb:MessageId,omitempty"`
	Timestamp string `xml:"eb:Timestamp,omitempty"`
}
