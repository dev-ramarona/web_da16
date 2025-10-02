package mdlSbrapi

type MdlSbrapiLcnpunDtbase struct {
	Prmkey string `json:"prmkey" bson:"prmkey,omitempty"`
	Airlfl string `json:"airlfl" bson:"airlfl,omitempty"`
	Lcrpun string `json:"lcrpun" bson:"lcrpun,omitempty"`
	Totpax int    `json:"totpax" bson:"totpax,omitempty"`
	Flnbfl string `json:"flnbfl" bson:"flnbfl,omitempty"`
	Depart string `json:"depart" bson:"depart,omitempty"`
	Routfl string `json:"routfl" bson:"routfl,omitempty"`
	Clssfl string `json:"clssfl" bson:"clssfl,omitempty"`
	Datefl int32  `json:"datefl" bson:"datefl,omitempty"`
	Dateup int32  `json:"dateup" bson:"dateup,omitempty"`
	Timeup int64  `json:"timeup" bson:"timeup,omitempty"`
	Agtnme string `json:"agtnme" bson:"agtnme,omitempty"`
	Pnrcde string `json:"pnrcde" bson:"pnrcde,omitempty"`
}
