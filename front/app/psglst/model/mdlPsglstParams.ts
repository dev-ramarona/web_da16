export interface MdlPsglstEdtprmParams {
  params: string;
  length: number;
}

export interface MdlPsglstActlogParams {
  timeup: number;
  dateup: number;
  statdt: string;
}

export interface MdlPsglstErrlogParams {
  prmkey: string;
  status: string;
  errprt: string;
  errtxt: string;
  datefl: number;
  timeup: number;
  airlnf: string;
  depart: string;
  flnumf: string;
  flstat: string;
  flhour: number;
  routef: string;
  dvsion: string;
  worker: number;
  ignore: string;
  updtby: string;
}

export interface MdlPsglstDetailParams {
  prmkey: string;
  mnthfl: string;
  datefl: string;
  datevc: string;
  timefl: string;
  timevc: string;
  airlfl: string;
  airlvc: string;
  flnbfl: string;
  flnbvc: string;
  flhour: number;
  depart: string;
  arrivl: string;
  airtyp: string;
  routfl: string;
  routvc: string;
  routac: string;
  routcm: string;
  lnenbr: string;
  isflwn: string;
  istrst: string;
  gender: string;
  seatpx: string;
  groupx: string;
  pnrcde: string;
  tktnbr: string;
  Tktprv: string;
  status: string;
  fstnme: string;
  lstnme: string;
  cpnbfl: string;
  cpnbvc: string;
  clssfl: string;
  clssvc: string;
  cabnfl: string;
  cabnvc: string;
  frrate: number;
  framntflflwn: number;
  framntflvcrs: number;
  frtxyqflflwn: number;
  frtxyqflvcrs: number;
  frtxyrflflwn: number;
  frtxyrflvcrs: number;
  agtdie: string;
  ireglr: string;
  istedt: string;
  notedt: string;
  nottxt: string;
  enhkey: string;
  updtby: string;
}

export interface MdlPsglstAllprmSrcprm {
  mnthfl_detail: string;
  datefl_detail: string;
  airlfl_detail: string;
  flnbfl_detail: string;
  depart_detail: string;
  routfl_detail: string;
  isflwn_detail: string;
  istrst_detail: string;
  pnrcde_detail: string;
  tktnbr_detail: string;
  istirg_detail: string;
  mnthfl_others: string;
  datefl_others: string;
  airlfl_others: string;
  flnbfl_others: string;
  depart_others: string;
  routfl_others: string;
  isflwn_others: string;
  istrst_others: string;
  pnrcde_others: string;
  tktnbr_others: string;
  istirg_others: string;
}

export interface MdlPsglstDetailSrcprm {
  mnthfl_detail: string;
  datefl_detail: string;
  airlfl_detail: string;
  flnbfl_detail: string;
  depart_detail: string;
  routfl_detail: string;
  isflwn_detail: string;
  istrst_detail: string;
  pnrcde_detail: string;
  tktnbr_detail: string;
  istirg_detail: string;
}
