package mdlSbrapi

type MdlSbrapiMsghdrMainob struct {
	From           MdlSbrapiMsghdrPartyx `xml:"eb:From"`
	To             MdlSbrapiMsghdrPartyx `xml:"eb:To"`
	CPAId          string                `xml:"eb:CPAId"`
	ConversationId string                `xml:"eb:ConversationId,omitempty"`
	Service        MdlSbrapiMsghdrSrvice `xml:"eb:Service"`
	Action         string                `xml:"eb:Action"`
	MessageData    MdlSbrapiMsghdrMsgdta `xml:"eb:MessageData"`
	XmlnsEb        string                `xml:"xmlns:eb,attr"`
	MustUnderstand string                `xml:"soap-env:mustUnderstand,attr"`
	Version        string                `xml:"eb:version,attr"`
}
type MdlSbrapiMsghdrPartyx struct {
	PartyId string `xml:"eb:PartyId"`
}
type MdlSbrapiMsghdrSrvice struct {
	Type string `xml:"eb:type,attr"`
	Name string `xml:",chardata"`
}
type MdlSbrapiMsghdrMsgdta struct {
	MessageId string `xml:"eb:MessageId,omitempty"`
	Timestamp string `xml:"eb:Timestamp,omitempty"`
}
type MdlSbrapiMsghdrParams struct {
	Bsttkn string `json:"bsttkn" bson:"bsttkn,omitempty"`
	Convid string `json:"convid" bson:"convid,omitempty"`
	Mssgid string `json:"mssgid" bson:"mssgid,omitempty"`
	Timefm string `json:"timefm" bson:"timefm,omitempty"`
}
type MdlSbrapiMsghdrApndix struct {
	Airlfl string
	Clssfl string
	Datefl int32
	Depart string
	Flnbfl string
	Pnrcde string
	Routfl string
}
