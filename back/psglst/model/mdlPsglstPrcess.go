package mdlPsglst

type MdlPsglstErrlogDtbase struct {
	Prmkey string  `json:"prmkey,omitempty" bson:"prmkey"`
	Errsts string  `json:"Errsts,omitempty" bson:"Errsts"`
	Errprt string  `json:"errprt,omitempty" bson:"errprt"`
	Errout string  `json:"errout,omitempty" bson:"errout"`
	Errtxt string  `json:"errtxt,omitempty" bson:"errtxt"`
	Errdvs string  `json:"Errdvs,omitempty" bson:"Errdvs"`
	Errign string  `json:"errign,omitempty" bson:"errign"`
	Dateup int32   `json:"dateup,omitempty" bson:"dateup"`
	Timeup int64   `json:"timeup,omitempty" bson:"timeup"`
	Datefl int32   `json:"datefl,omitempty" bson:"datefl"`
	Airlfl string  `json:"airlfl,omitempty" bson:"airlfl"`
	Depart string  `json:"depart,omitempty" bson:"depart"`
	Flnbfl string  `json:"flnbfl,omitempty" bson:"flnbfl"`
	Flstat string  `json:"flstat,omitempty" bson:"flstat"`
	Flhour float64 `json:"flhour,omitempty" bson:"flhour"`
	Routfl string  `json:"routfl,omitempty" bson:"routfl"`
	Updtby string  `json:"updtby,omitempty" bson:"updtby"`
	Worker int32   `json:"worker,omitempty" bson:"worker"`
}

type MdlPsglstClsslvDtbase struct {
	Rbdcls string  `json:"rbdcls,omitempty" bson:"rbdcls,omitempty"`
	Lvlcls int32   `json:"lvlcls,omitempty" bson:"lvlcls,omitempty"`
	Cbncls string  `json:"cbncls,omitempty" bson:"cbncls,omitempty"`
	Discnt float64 `json:"discnt,omitempty" bson:"discnt,omitempty"`
}

type MdlPsglstFlhourDtbase struct {
	Prmkey string  `json:"prmkey,omitempty" bson:"prmkey,omitempty"`
	Airlfl string  `json:"airlfl,omitempty" bson:"airlfl,omitempty"`
	Routfl string  `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Flnbfl string  `json:"flnbfl,omitempty" bson:"flnbfl,omitempty"`
	Flhour float64 `json:"flhour,omitempty" bson:"flhour,omitempty"`
	Timefl string  `json:"timefl,omitempty" bson:"timefl,omitempty"`
	Datefl string  `json:"datefl,omitempty" bson:"datefl,omitempty"`
	Dateup int32   `json:"dateup,omitempty" bson:"dateup,omitempty"`
	Timeup int64   `json:"timeup,omitempty" bson:"timeup,omitempty"`
	Airtyp string  `json:"airtyp,omitempty" bson:"airtyp,omitempty"`
	Airmls string  `json:"airmls,omitempty" bson:"airmls,omitempty"`
}
