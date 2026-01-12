package mdlSbrapi

type MdlSbrapiVavhstDtbase struct {
	Prmkey string `json:"prmkey" bson:"prmkey,omitempty"`
	Squenc string `json:"squenc" bson:"squenc,omitempty"`
	Datefl int32  `json:"datefl" bson:"datefl,omitempty"`
	Timeup int64  `json:"timeup" bson:"timeup,omitempty"`
	Clssfl string `json:"clssfl" bson:"clssfl,omitempty"`
	Depart string `json:"depart" bson:"depart,omitempty"`
	Routfl string `json:"routfl" bson:"routfl,omitempty"`
	Flnbfl string `json:"flnbfl" bson:"flnbfl,omitempty"`
	Statfl string `json:"statfl" bson:"statfl,omitempty"`
	Totpax int64  `json:"totpax" bson:"totpax,omitempty"`
	Lniata string `json:"lniata" bson:"lniata,omitempty"`
	Agtdie string `json:"agtdie" bson:"agtdie,omitempty"`
	Agtcty string `json:"agtcty" bson:"agtcty,omitempty"`
	Airlfl string `json:"airlfl" bson:"airlfl,omitempty"`
}
