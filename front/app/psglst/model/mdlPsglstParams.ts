// Global
export interface MdlPsglstGlobalSrcprm {
  update_psgdtl: string;
  mnthfl_psgdtl: string;
  datefl_psgdtl: string;
  airlfl_psgdtl: string;
  flnbfl_psgdtl: string;
  depart_psgdtl: string;
  routfl_psgdtl: string;
  pnrcde_psgdtl: string;
  tktnfl_psgdtl: string;
  isitfl_psgdtl: string;
  isittx_psgdtl: string;
  isitir_psgdtl: string;
  nclear_psgdtl: string;
  format_psgdtl: string;
  pagenw_psgdtl: number;
  limitp_psgdtl: number;
  pagenw_errlog: number;
  limitp_errlog: number;
  erdvsn_errlog: string;
}
export interface MdlPsglstPsgdtlSrcprm {
  update_psgdtl: string;
  mnthfl_psgdtl: string;
  datefl_psgdtl: string;
  airlfl_psgdtl: string;
  flnbfl_psgdtl: string;
  depart_psgdtl: string;
  routfl_psgdtl: string;
  pnrcde_psgdtl: string;
  tktnfl_psgdtl: string;
  isitfl_psgdtl: string;
  isittx_psgdtl: string;
  isitir_psgdtl: string;
  nclear_psgdtl: string;
  format_psgdtl: string;
  pagenw_psgdtl: number;
  limitp_psgdtl: number;
}
export interface MdlPsglstErrlogSrcprm {
  pagenw_errlog: number;
  limitp_errlog: number;
  erdvsn_errlog: string;
}

// Passangger list detail
export interface MdlPsglstPsgdtlFrntnd {
  mnfest: string;
  slsrpt: string;
  noterr: string;
  source: string;
  tktnfl: string;
  tktnvc: string;
  pnrcde: string;
  pnritl: string;
  curncy: string;
  ntaffl: number;
  ntafvc: number;
  yqtxfl: number;
  yqtxvc: number;
  frrate: number;
  frbcde: string;
  qsrcrw: string;
  qsrcvc: number;
  frcalc: string;
  ndayfl: string;
  datefl: number;
  datevc: number;
  daterv: number;
  mnthfl: number;
  timefl: number;
  timerv: number;
  timeis: number;
  timecr: number;
  airlfl: string;
  airlvc: string;
  airtyp: string;
  flnbfl: string;
  flnbvc: string;
  flgate: string;
  bookdc: number;
  bookdy: number;
  depart: string;
  arrivl: string;
  routfl: string;
  routvc: string;
  routvf: string;
  routac: string;
  routmx: string;
  routfr: string;
  routfx: string;
  routsg: string;
  linenb: number;
  ckinnb: number;
  gender: string;
  typepx: string;
  seatpx: string;
  groupc: string;
  totpax: number;
  segpnr: string;
  segtkt: string;
  psgrid: string;
  tourcd: string;
  staloc: string;
  stanbr: string;
  wrkloc: string;
  hmeloc: string;
  lniata: string;
  emplid: string;
  nmefst: string;
  nmelst: string;
  cpnbfl: number;
  cpnbvc: number;
  clssfl: string;
  clssvc: string;
  statvc: string;
  cbinfl: string;
  cbinvc: string;
  agtdie: string;
  codels: string;
  isitfl: string;
  isittx: string;
  isitir: string;
  isitct: string;
  isittf: string;
  isitnr: string;
  noteup: string;
  updtby: string;
  prmkey: string;

  // Ancillary
  gpcdae: string;
  sbcdae: string;
  descae: string;
  wgbgae: number;
  qtbgae: number;
  routae: string;
  fareae: number;
  currae: string;
  emdnae: string;

  // Bagtag
  nmbrbt: string;
  qntybt: number;
  wghtbt: number;
  paidbt: number;
  fbavbt: number;
  hfbabt: number;
  qtotbt: number;
  wtotbt: number;
  ptotbt: number;
  ftotbt: number;
  excsbt: number;
  typebt: string;
  coment: string;

  // Outbound
  airlob: string;
  flnbob: string;
  clssob: string;
  routob: string;
  dateob: number;
  timeob: number;

  // Inbound
  airlib: string;
  flnbib: string;
  clssib: string;
  dstrib: string;
  dateib: number;
  timeib: number;

  // Ireg
  codeir: string;
  airlir: string;
  flnbir: string;
  dateir: number;

  // Infant
  tktnif: string;
  cpnbif: number;
  dateif: number;
  clssif: string;
  routif: string;
  statif: string;
  paxsif: string;

  // Cancel bagtag
  airlxt: string;
  dstrxt: string;
  nmbrxt: string;
}

// Search param
export interface MdlPsglstPsgdtlSearch {
  update_psgdtl: string;
  mnthfl_psgdtl: string;
  datefl_psgdtl: string;
  airlfl_psgdtl: string;
  flnbfl_psgdtl: string;
  depart_psgdtl: string;
  routfl_psgdtl: string;
  pnrcde_psgdtl: string;
  tktnfl_psgdtl: string;
  isitfl_psgdtl: string;
  isittx_psgdtl: string;
  isitir_psgdtl: string;
  nclear_psgdtl: string;
  format_psgdtl: string;
  pagenw_psgdtl: number;
  limitp_psgdtl: number;
}
export interface MdlPsglstAcpedtDtbase {
  params: string;
  length: number;
}

// Log action
export interface MdlPsglstActlogDtbase {
  timeup: number;
  dateup: number;
  datefl: number;
  statdt: string;
}

// Log error
export interface MdlPsglstErrlogDtbase {
  prmkey: string;
  erstat: string;
  erpart: string;
  ersrce: string;
  erdtil: string;
  erdvsn: string;
  erignr: string;
  dateup: number;
  timeup: number;
  datefl: number;
  airlfl: string;
  depart: string;
  flnbfl: string;
  flstat: string;
  flhour: number;
  routfl: string;
  updtby: string;
  worker: number;
}
