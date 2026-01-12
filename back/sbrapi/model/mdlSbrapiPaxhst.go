package mdlSbrapi

type MdlSbrapiPaxhstDtbase struct {
	Prmkey string `json:"prmkey" bson:"prmkey,omitempty"`
	Agtdie string `json:"agtdie" bson:"agtdie,omitempty"`
	Agtcty string `json:"agtcty" bson:"agtcty,omitempty"`
	Datefl int32  `json:"datefl" bson:"datefl,omitempty"`
	Timeup int64  `json:"timeup" bson:"timeup,omitempty"`
	Lniata string `json:"lniata" bson:"lniata,omitempty"`
	Depart string `json:"depart" bson:"depart,omitempty"`
	Flnbfl string `json:"flnbfl" bson:"flnbfl,omitempty"`
	Itemhs string `json:"itemhs" bson:"itemhs,omitempty"`
	Pnrcde string `json:"pnrcde" bson:"pnrcde,omitempty"`
	Nmefst string `json:"nmefst" bson:"nmefst,omitempty"`
	Nmelst string `json:"nmelst" bson:"nmelst,omitempty"`
	Arrivl string `json:"arrivl" bson:"arrivl,omitempty"`
	Seatpx string `json:"seatpx" bson:"seatpx,omitempty"`
	Qntybt string `json:"qntybt" bson:"qntybt,omitempty"`
	Codels string `json:"codels" bson:"codels,omitempty"`
}

type MdlSbrapiPaxhstRawdta struct {
	Prmkey string `json:"prmkey" bson:"prmkey,omitempty"`
	Rawhst string `json:"rawhst" bson:"rawhst,omitempty"`
}
