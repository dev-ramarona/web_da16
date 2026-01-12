package mdlHoldst

type MdlHoldstErrlogDtbase struct {
	Prmkey string  `json:"prmkey" bson:"prmkey,omitempty"`
	Erstat string  `json:"erstat" bson:"erstat,omitempty"`
	Erpart string  `json:"erpart" bson:"erpart,omitempty"`
	Ersrce string  `json:"ersrce" bson:"ersrce,omitempty"`
	Erdtil string  `json:"erdtil" bson:"erdtil,omitempty"`
	Erdvsn string  `json:"erdvsn" bson:"erdvsn,omitempty"`
	Erignr string  `json:"erignr" bson:"erignr,omitempty"`
	Dateup int32   `json:"dateup" bson:"dateup,omitempty"`
	Timeup int64   `json:"timeup" bson:"timeup,omitempty"`
	Datefl int32   `json:"datefl" bson:"datefl,omitempty"`
	Airlfl string  `json:"airlfl" bson:"airlfl,omitempty"`
	Depart string  `json:"depart" bson:"depart,omitempty"`
	Flnbfl string  `json:"flnbfl" bson:"flnbfl,omitempty"`
	Flstat string  `json:"flstat" bson:"flstat,omitempty"`
	Flhour float64 `json:"flhour" bson:"flhour,omitempty"`
	Routfl string  `json:"routfl" bson:"routfl,omitempty"`
	Updtby string  `json:"updtby" bson:"updtby,omitempty"`
	Worker int32   `json:"worker" bson:"worker,omitempty"`
}
