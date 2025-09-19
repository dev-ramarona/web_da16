package mdl_jeddah

type MdlJeddahSmrpnrDtbase struct {
	Prmkey string `json:"prmkey" bson:"prmkey,omitempty"`
	Routfl string `json:"routfl" bson:"routfl,omitempty"`
	Timedp int64  `json:"timedp" bson:"timedp"`
	Timerv int64  `json:"timerv" bson:"timerv"`
	Dateup int32  `json:"dateup" bson:"dateup"`
	Timeup int64  `json:"timeup" bson:"timeup"`
	Timecr int64  `json:"timecr" bson:"timecr"`
	Agtnme string `json:"agtnme" bson:"agtnme,omitempty"`
	Agtdtl string `json:"agtdtl" bson:"agtdtl,omitempty"`
	Agtidn string `json:"agtidn" bson:"agtidn,omitempty"`
	Pnrcde string `json:"pnrcde" bson:"pnrcde,omitempty"`
	Intrln string `json:"intrln" bson:"intrln,omitempty"`
	Rtlsrs string `json:"rtlsrs" bson:"rtlsrs,omitempty"`
	Arrcpn string `json:"arrcpn" bson:"arrcpn,omitempty"`
	Agtdie string `json:"agtdie" bson:"agtdie,omitempty"`
	Totisd int    `json:"totisd" bson:"totisd"`
	Totbok int    `json:"totbok" bson:"totbok"`
	Totpax int    `json:"totpax" bson:"totpax"`
	Totcxl int    `json:"totcxl" bson:"totcxl"`
	Totspl int    `json:"totspl" bson:"totspl"`
	Arrspl string `json:"arrspl" bson:"arrspl,omitempty"`
	Notedt string `json:"notedt" bson:"notedt,omitempty"`
}

type MdlJeddahDtlpnrDtbase struct {
	Prmkey string `json:"prmkey" bson:"prmkey,omitempty"`
	Airlfl string `json:"airlfl" bson:"airlfl,omitempty"`
	Flnbfl string `json:"flnbfl" bson:"flnbfl,omitempty"`
	Depart string `json:"depart" bson:"depart,omitempty"`
	Routfl string `json:"routfl" bson:"routfl,omitempty"`
	Clssfl string `json:"clssfl" bson:"clssfl,omitempty"`
	Datefl int32  `json:"datefl" bson:"datefl"`
	Dateup int32  `json:"dateup" bson:"dateup"`
	Timeup int64  `json:"timeup" bson:"timeup"`
	Timecr int64  `json:"timecr" bson:"timecr"`
	Agtnme string `json:"agtnme" bson:"agtnme,omitempty"`
	Agtdtl string `json:"agtdtl" bson:"agtdtl,omitempty"`
	Agtidn string `json:"agtidn" bson:"agtidn,omitempty"`
	Pnrcde string `json:"pnrcde" bson:"pnrcde,omitempty"`
	Intrln string `json:"intrln" bson:"intrln,omitempty"`
	Rtlsrs string `json:"rtlsrs" bson:"rtlsrs,omitempty"`
	Toflnm string `json:"toflnm" bson:"toflnm,omitempty"`
	Drules int    `json:"drules" bson:"drules"`
	Totisd int    `json:"totisd" bson:"totisd"`
	Totbok int    `json:"totbok" bson:"totbok"`
	Totpax int    `json:"totpax" bson:"totpax"`
	Totcxl int    `json:"totcxl" bson:"totcxl"`
	Totchg int    `json:"totchg" bson:"totchg"`
	Totspl int    `json:"totspl" bson:"totspl"`
	Arrspl string `json:"arrspl" bson:"arrspl,omitempty"`
	Notedt string `json:"notedt" bson:"notedt,omitempty"`
	Flstat string `json:"flstat" bson:"flstat,omitempty"`
}

type MdlJeddahSmrflnDtbase struct {
	Prmkey string `json:"prmkey" bson:"prmkey,omitempty"`
	Airlfl string `json:"airlfl" bson:"airlfl,omitempty"`
	Flnbfl string `json:"flnbfl" bson:"flnbfl,omitempty"`
	Flstat string `json:"flstat" bson:"flstat,omitempty"`
	Depart string `json:"depart" bson:"depart,omitempty"`
	Routfl string `json:"routfl" bson:"routfl,omitempty"`
	Datefl int32  `json:"datefl" bson:"datefl,omitempty"`
	Dateup int32  `json:"dateup" bson:"dateup,omitempty"`
	Timeup int64  `json:"timeup" bson:"timeup,omitempty"`
	Totisd int    `json:"totisd" bson:"totisd"`
	Totbok int    `json:"totbok" bson:"totbok"`
	Totpax int    `json:"totpax" bson:"totpax"`
	Totcxl int    `json:"totcxl" bson:"totcxl"`
	Totchg int    `json:"totchg" bson:"totchg"`
	Totspl int    `json:"totspl" bson:"totspl"`
	Notedt string `json:"notedt" bson:"notedt,omitempty"`
}

type MdlJeddahAllpnrInputx struct {
	Prmkey_pnrdtl string `json:"prmkey_pnrdtl,omitempty"`
	Clssfl_pnrdtl string `json:"clssfl_pnrdtl,omitempty"`
	Airlfl_pnrdtl string `json:"airlfl_pnrdtl,omitempty"`
	Datefl_pnrdtl string `json:"datefl_pnrdtl,omitempty"`
	Pnrcde_pnrdtl string `json:"pnrcde_pnrdtl,omitempty"`
	Flnbfl_pnrdtl string `json:"flnbfl_pnrdtl,omitempty"`
	Routfl_pnrdtl string `json:"routfl_pnrdtl,omitempty"`
	Agtnme_pnrdtl string `json:"agtnme_pnrdtl,omitempty"`
	Srtspl_pnrdtl string `json:"srtspl_pnrdtl,omitempty"`
	Srtcxl_pnrdtl string `json:"srtcxl_pnrdtl,omitempty"`
	Pagenw_pnrdtl int    `json:"pagenw_pnrdtl,omitempty"`
	Limitp_pnrdtl int    `json:"limitp_pnrdtl,omitempty"`
	Prmkey_pnrsmr string `json:"prmkey_pnrsmr,omitempty"`
	Airlfl_pnrsmr string `json:"airlfl_pnrsmr,omitempty"`
	Flnbfl_pnrsmr string `json:"flnbfl_pnrsmr,omitempty"`
	Routfl_pnrsmr string `json:"routfl_pnrsmr,omitempty"`
	Datefl_pnrsmr string `json:"datefl_pnrsmr,omitempty"`
	Pnrcde_pnrsmr string `json:"pnrcde_pnrsmr,omitempty"`
	Agtnme_pnrsmr string `json:"agtnme_pnrsmr,omitempty"`
	Psdate_pnrsmr string `json:"psdate_pnrsmr,omitempty"`
	Srtspl_pnrsmr string `json:"srtspl_pnrsmr,omitempty"`
	Srtcxl_pnrsmr string `json:"srtcxl_pnrsmr,omitempty"`
	Pagenw_pnrsmr int    `json:"pagenw_pnrsmr,omitempty"`
	Limitp_pnrsmr int    `json:"limitp_pnrsmr,omitempty"`
	Prmkey_flnsmr string `json:"prmkey_flnsmr,omitempty"`
	Airlfl_flnsmr string `json:"airlfl_flnsmr,omitempty"`
	Datefl_flnsmr string `json:"datefl_flnsmr,omitempty"`
	Flnbfl_flnsmr string `json:"flnbfl_flnsmr,omitempty"`
	Routfl_flnsmr string `json:"routfl_flnsmr,omitempty"`
	Psdate_flnsmr string `json:"psdate_flnsmr,omitempty"`
	Srtspl_flnsmr string `json:"srtspl_flnsmr,omitempty"`
	Srtcxl_flnsmr string `json:"srtcxl_flnsmr,omitempty"`
	Pagenw_flnsmr int    `json:"pagenw_flnsmr,omitempty"`
	Limitp_flnsmr int    `json:"limitp_flnsmr,omitempty"`
	Airlfl_agtnme string `json:"airlfl_agtnme,omitempty"`
	Agtnme_agtnme string `json:"agtnme_agtnme,omitempty"`
	Srtnul_agtnme string `json:"srtnul_agtnme,omitempty"`
	Pagenw_agtnme int    `json:"pagenw_agtnme,omitempty"`
	Limitp_agtnme int    `json:"limitp_agtnme,omitempty"`
}

type MdlJeddahLogpnrDtbase struct {
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
	// Agtdtl string `json:"agtdtl" bson:"agtdtl,omitempty"`
	// Agtidn string `json:"agtidn" bson:"agtidn,omitempty"`
	// Rtlsrs string `json:"rtlsrs" bson:"rtlsrs,omitempty"`
	Pnrcde string `json:"pnrcde" bson:"pnrcde,omitempty"`
}

type MdlJeddahAgtnmeDtbase struct {
	Prmkey string `json:"prmkey,omitempty" bson:"prmkey"`
	Airlfl string `json:"airlfl,omitempty" bson:"airlfl"`
	Agtnme string `json:"agtnme,omitempty" bson:"agtnme"`
	Agtdtl string `json:"agtdtl,omitempty" bson:"agtdtl"`
	Agtidn string `json:"agtidn,omitempty" bson:"agtidn"`
	Rtlsrs string `json:"rtlsrs,omitempty" bson:"rtlsrs"`
	Updtby string `json:"updtby,omitempty" bson:"updtby"`
}

type MdlJeddahAgtnmeInputx struct {
	Prmkey string `json:"prmkey" bson:"prmkey,omitempty"`
	Airlfl string `json:"airlfl" bson:"airlfl,omitempty"`
	Agtnme string `json:"agtnme" bson:"agtnme,omitempty"`
	Agtdtl string `json:"agtdtl" bson:"agtdtl,omitempty"`
	Newdtl string `json:"newdtl" bson:"newdtl,omitempty"`
	Agtidn string `json:"agtidn" bson:"agtidn,omitempty"`
	Newidn string `json:"newidn" bson:"newidn,omitempty"`
	Rtlsrs string `json:"rtlsrs" bson:"rtlsrs,omitempty"`
	Updtby string `json:"updtby" bson:"updtby,omitempty"`
	Agtnew string `json:"agtnew" bson:"agtnew,omitempty"`
}

type MdlJeddahFlnbdbDtbase struct {
	Prmkey string `json:"prmkey,omitempty" bson:"prmkey,omitempty"`
	Datefl int32  `json:"datefl,omitempty" bson:"datefl,omitempty"`
	Timefl int64  `json:"timefl,omitempty" bson:"timefl,omitempty"`
	Dateup int32  `json:"dateup,omitempty" bson:"dateup,omitempty"`
	Timeup int64  `json:"timeup,omitempty" bson:"timeup,omitempty"`
	Airlfl string `json:"airlfl,omitempty" bson:"airlfl,omitempty"`
	Flnbfl string `json:"flnbfl,omitempty" bson:"flnbfl,omitempty"`
	Depart string `json:"depart,omitempty" bson:"depart,omitempty"`
	Routfl string `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Fltype string `json:"fltype,omitempty" bson:"fltype,omitempty"`
	Flstat string `json:"flstat,omitempty" bson:"flstat,omitempty"`
	Isjedh string `json:"isjedh,omitempty" bson:"isjedh,omitempty"`
}

type MdlJeddahDrulesDtbase struct {
	Rllevl int32  `json:"rllevl,omitempty" bson:"rllevl,omitempty"`
	Rldays int32  `json:"rldays,omitempty" bson:"rldays,omitempty"`
	Rlrate int32  `json:"rlrate,omitempty" bson:"rlrate,omitempty"`
	Rlcolm string `json:"rlcolm,omitempty" bson:"rlcolm,omitempty"`
	Rltype string `json:"rltype,omitempty" bson:"rltype,omitempty"`
	Rlnote string `json:"rlnote,omitempty" bson:"rlnote,omitempty"`
}

type MdlJeddahLogactDtbase struct {
	Dateup int32  `json:"dateup,omitempty" bson:"dateup,omitempty"`
	Timeup int64  `json:"timeup,omitempty" bson:"timeup,omitempty"`
	Statdt string `json:"statdt,omitempty" bson:"statdt,omitempty"`
}

type MdlJeddahAddflnInpprm struct {
	Airlfl_addfln string `json:"airlfl_addfln,omitempty" bson:"airlfl_addfln,omitempty"`
	Flnbfl_addfln string `json:"flnbfl_addfln,omitempty" bson:"flnbfl_addfln,omitempty"`
	Datefl_addfln string `json:"datefl_addfln,omitempty" bson:"datefl_addfln,omitempty"`
	Routfl_addfln string `json:"routfl_addfln,omitempty" bson:"routfl_addfln,omitempty"`
	Fltype_addfln string `json:"fltype_addfln,omitempty" bson:"fltype_addfln,omitempty"`
}
