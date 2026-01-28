package mdlPsglst

type MdlPsglstPsgdtlAcpedt struct {
	Params string `json:"params,omitempty" bson:"params,omitempty"`
	Length int    `json:"length,omitempty" bson:"length,omitempty"`
}

type MdlPsglstPsgdtlInputx struct {
	Mnthfl_psgdtl string `json:"mnthfl_psgdtl,omitempty" bson:"mnthfl_psgdtl,omitempty"`
	Datefl_psgdtl string `json:"datefl_psgdtl,omitempty" bson:"datefl_psgdtl,omitempty"`
	Airlfl_psgdtl string `json:"airlfl_psgdtl,omitempty" bson:"airlfl_psgdtl,omitempty"`
	Flnbfl_psgdtl string `json:"flnbfl_psgdtl,omitempty" bson:"flnbfl_psgdtl,omitempty"`
	Depart_psgdtl string `json:"depart_psgdtl,omitempty" bson:"depart_psgdtl,omitempty"`
	Routfl_psgdtl string `json:"routfl_psgdtl,omitempty" bson:"routfl_psgdtl,omitempty"`
	Pnrcde_psgdtl string `json:"pnrcde_psgdtl,omitempty" bson:"pnrcde_psgdtl,omitempty"`
	Tktnfl_psgdtl string `json:"tktnfl_psgdtl,omitempty" bson:"tktnfl_psgdtl,omitempty"`
	Isitfl_psgdtl string `json:"isitfl_psgdtl,omitempty" bson:"isitfl_psgdtl,omitempty"`
	Isittx_psgdtl string `json:"isittx_psgdtl,omitempty" bson:"isittx_psgdtl,omitempty"`
	Isitir_psgdtl string `json:"isitir_psgdtl,omitempty" bson:"isitir_psgdtl,omitempty"`
	Nclear_psgdtl string `json:"nclear_psgdtl,omitempty" bson:"nclear_psgdtl,omitempty"`
	Pagenw_psgdtl int    `json:"pagenw_psgdtl,omitempty" bson:"pagenw_psgdtl,omitempty"`
	Limitp_psgdtl int    `json:"limitp_psgdtl,omitempty" bson:"limitp_psgdtl,omitempty"`
	Pagenw_errlog int    `json:"pagenw_errlog,omitempty" bson:"pagenw_errlog,omitempty"`
	Limitp_errlog int    `json:"limitp_errlog,omitempty" bson:"limitp_errlog,omitempty"`
	Erdvsn_errlog string `json:"erdvsn_errlog,omitempty" bson:"erdvsn_errlog,omitempty"`
}

type MdlPsglstPsgdtlDtbase struct {
	Mnfest string  `json:"mnfest" bson:"mnfest,omitempty"`
	Slsrpt string  `json:"slsrpt" bson:"slsrpt,omitempty"`
	Noterr string  `json:"noterr" bson:"noterr,omitempty"`
	Source string  `json:"source" bson:"source,omitempty"`
	Tktnfl string  `json:"tktnfl" bson:"tktnfl,omitempty"`
	Tktnvc string  `json:"tktnvc" bson:"tktnvc,omitempty"`
	Pnrcde string  `json:"pnrcde" bson:"pnrcde,omitempty"`
	Pnritl string  `json:"pnritl" bson:"pnritl,omitempty"`
	Curncy string  `json:"curncy" bson:"curncy,omitempty"`
	Ntaffl int32   `json:"ntaffl" bson:"ntaffl"`
	Ntafvc float64 `json:"ntafvc" bson:"ntafvc"`
	Yqtxfl int32   `json:"yqtxfl" bson:"yqtxfl"`
	Yqtxvc float64 `json:"yqtxvc" bson:"yqtxvc"`
	Frrate float64 `json:"frrate" bson:"frrate,omitempty"`
	Frbcde string  `json:"frbcde" bson:"frbcde,omitempty"`
	Qsrcrw string  `json:"qsrcrw" bson:"qsrcrw,omitempty"`
	Qsrcvc float64 `json:"qsrcvc" bson:"qsrcvc,omitempty"`
	Frcalc string  `json:"frcalc" bson:"frcalc,omitempty"`
	Ndayfl string  `json:"ndayfl" bson:"ndayfl,omitempty"`
	Datefl int32   `json:"datefl" bson:"datefl,omitempty"`
	Datevc int32   `json:"datevc" bson:"datevc,omitempty"`
	Daterv int32   `json:"daterv" bson:"daterv,omitempty"`
	Mnthfl int32   `json:"mnthfl" bson:"mnthfl,omitempty"`
	Timefl int64   `json:"timefl" bson:"timefl,omitempty"`
	Timerv int64   `json:"timerv" bson:"timerv,omitempty"`
	Timeis int64   `json:"timeis" bson:"timeis"`
	Timecr int64   `json:"timecr" bson:"timecr"`
	Airlfl string  `json:"airlfl" bson:"airlfl,omitempty"`
	Airlvc string  `json:"airlvc" bson:"airlvc,omitempty"`
	Airtyp string  `json:"airtyp" bson:"airtyp,omitempty"`
	Flnbfl string  `json:"flnbfl" bson:"flnbfl,omitempty"`
	Flnbvc string  `json:"flnbvc" bson:"flnbvc,omitempty"`
	Flgate string  `json:"flgate" bson:"flgate,omitempty"`
	Bookdc int32   `json:"bookdc" bson:"bookdc,omitempty"`
	Bookdy int32   `json:"bookdy" bson:"bookdy,omitempty"`
	Depart string  `json:"depart" bson:"depart,omitempty"`
	Arrivl string  `json:"arrivl" bson:"arrivl,omitempty"`
	Routfl string  `json:"routfl" bson:"routfl,omitempty"`
	Routvc string  `json:"routvc" bson:"routvc,omitempty"`
	Routvf string  `json:"routvf" bson:"routvf,omitempty"`
	Routac string  `json:"routac" bson:"routac,omitempty"`
	Routmx string  `json:"routmx" bson:"routmx,omitempty"`
	Routfr string  `json:"routfr" bson:"routfr,omitempty"`
	Routfx string  `json:"routfx" bson:"routfx,omitempty"`
	Routsg string  `json:"routsg" bson:"routsg,omitempty"`
	Linenb int32   `json:"linenb" bson:"linenb,omitempty"`
	Ckinnb int32   `json:"ckinnb" bson:"ckinnb,omitempty"`
	Gender string  `json:"gender" bson:"gender,omitempty"`
	Typepx string  `json:"typepx" bson:"typepx,omitempty"`
	Seatpx string  `json:"seatpx" bson:"seatpx,omitempty"`
	Groupc string  `json:"groupc" bson:"groupc,omitempty"`
	Totpax int32   `json:"totpax" bson:"totpax,omitempty"`
	Segpnr string  `json:"segpnr" bson:"segpnr,omitempty"`
	Segtkt string  `json:"segtkt" bson:"segtkt,omitempty"`
	Psgrid string  `json:"psgrid" bson:"psgrid,omitempty"`
	Tourcd string  `json:"tourcd" bson:"tourcd,omitempty"`
	Staloc string  `json:"staloc" bson:"staloc,omitempty"`
	Stanbr string  `json:"stanbr" bson:"stanbr,omitempty"`
	Wrkloc string  `json:"wrkloc" bson:"wrkloc,omitempty"`
	Hmeloc string  `json:"hmeloc" bson:"hmeloc,omitempty"`
	Lniata string  `json:"lniata" bson:"lniata,omitempty"`
	Emplid string  `json:"emplid" bson:"emplid,omitempty"`
	Nmefst string  `json:"nmefst" bson:"nmefst,omitempty"`
	Nmelst string  `json:"nmelst" bson:"nmelst,omitempty"`
	Cpnbfl int32   `json:"cpnbfl" bson:"cpnbfl,omitempty"`
	Cpnbvc int32   `json:"cpnbvc" bson:"cpnbvc,omitempty"`
	Clssfl string  `json:"clssfl" bson:"clssfl,omitempty"`
	Clssvc string  `json:"clssvc" bson:"clssvc,omitempty"`
	Statvc string  `json:"statvc" bson:"statvc,omitempty"`
	Cbinfl string  `json:"cbinfl" bson:"cbinfl,omitempty"`
	Cbinvc string  `json:"cbinvc" bson:"cbinvc,omitempty"`
	Agtdie string  `json:"agtdie" bson:"agtdie,omitempty"`
	Codels string  `json:"codels" bson:"codels,omitempty"`
	Isitfl string  `json:"isitfl" bson:"isitfl,omitempty"`
	Isittx string  `json:"isittx" bson:"isittx,omitempty"`
	Isitir string  `json:"isitir" bson:"isitir,omitempty"`
	Isitct string  `json:"isitct" bson:"isitct,omitempty"`
	Isittf string  `json:"isittf" bson:"isittf,omitempty"`
	Isitnr string  `json:"isitnr" bson:"isitnr,omitempty"`
	Noteup string  `json:"noteup" bson:"noteup,omitempty"`
	Updtby string  `json:"updtby" bson:"updtby,omitempty"`
	Prmkey string  `json:"prmkey" bson:"prmkey,omitempty"`

	// Ancillary
	Gpcdae string  `json:"gpcdae" bson:"gpcdae,omitempty"`
	Sbcdae string  `json:"sbcdae" bson:"sbcdae,omitempty"`
	Descae string  `json:"descae" bson:"descae,omitempty"`
	Wgbgae int32   `json:"wgbgae" bson:"wgbgae,omitempty"`
	Qtbgae int32   `json:"qtbgae" bson:"qtbgae,omitempty"`
	Routae string  `json:"routae" bson:"routae,omitempty"`
	Fareae float64 `json:"fareae" bson:"fareae,omitempty"`
	Currae string  `json:"currae" bson:"currae,omitempty"`
	Emdnae string  `json:"emdnae" bson:"emdnae,omitempty"`

	// Bagtag
	Nmbrbt string `json:"nmbrbt" bson:"nmbrbt,omitempty"`
	Qntybt int32  `json:"qntybt" bson:"qntybt,omitempty"`
	Wghtbt int32  `json:"wghtbt" bson:"wghtbt,omitempty"`
	Paidbt int32  `json:"paidbt" bson:"paidbt,omitempty"`
	Fbavbt int32  `json:"fbavbt" bson:"fbavbt,omitempty"`
	Hfbabt int32  `json:"hfbabt" bson:"hfbabt,omitempty"`
	Qtotbt int32  `json:"qtotbt" bson:"qtotbt,omitempty"`
	Wtotbt int32  `json:"wtotbt" bson:"wtotbt,omitempty"`
	Ptotbt int32  `json:"ptotbt" bson:"ptotbt,omitempty"`
	Ftotbt int32  `json:"ftotbt" bson:"ftotbt,omitempty"`
	Excsbt int32  `json:"excsbt" bson:"excsbt,omitempty"`
	Typebt string `json:"typebt" bson:"typebt,omitempty"`
	Coment string `json:"coment" bson:"coment,omitempty"`

	// Outbound
	Airlob string `json:"airlob" bson:"airlob,omitempty"`
	Flnbob string `json:"flnbob" bson:"flnbob,omitempty"`
	Clssob string `json:"clssob" bson:"clssob,omitempty"`
	Routob string `json:"routob" bson:"routob,omitempty"`
	Dateob int32  `json:"dateob" bson:"dateob,omitempty"`
	Timeob int64  `json:"timeob" bson:"timeob,omitempty"`

	// Inbound
	Airlib string `json:"airlib" bson:"airlib,omitempty"`
	Flnbib string `json:"flnbib" bson:"flnbib,omitempty"`
	Clssib string `json:"clssib" bson:"clssib,omitempty"`
	Dstrib string `json:"dstrib" bson:"dstrib,omitempty"`
	Dateib int32  `json:"dateib" bson:"dateib,omitempty"`
	Timeib int64  `json:"timeib" bson:"timeib,omitempty"`

	// Ireg
	Codeir string `json:"codeir" bson:"codeir,omitempty"`
	Airlir string `json:"airlir" bson:"airlir,omitempty"`
	Flnbir string `json:"flnbir" bson:"flnbir,omitempty"`
	Dateir int32  `json:"dateir" bson:"dateir,omitempty"`

	// Infant
	Tktnif string `json:"tktnif" bson:"tktnif,omitempty"`
	Cpnbif int32  `json:"cpnbif" bson:"cpnbif,omitempty"`
	Dateif int32  `json:"dateif" bson:"dateif,omitempty"`
	Clssif string `json:"clssif" bson:"clssif,omitempty"`
	Routif string `json:"routif" bson:"routif,omitempty"`
	Statif string `json:"statif" bson:"statif,omitempty"`
	Paxsif string `json:"paxsif" bson:"paxsif,omitempty"`

	// Cancel bagtag
	Airlxt string `json:"airlxt" bson:"airlxt,omitempty"`
	Dstrxt string `json:"dstrxt" bson:"dstrxt,omitempty"`
	Nmbrxt string `json:"nmbrxt" bson:"nmbrxt,omitempty"`
}
