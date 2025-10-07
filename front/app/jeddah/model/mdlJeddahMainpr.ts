// Global
export interface MdlJeddahInputxAllprm {
  chssmr: string;
  clssfl_pnrdtl: string; airlfl_pnrdtl: string;
  datefl_pnrdtl: string; pnrcde_pnrdtl: string;
  pnrclk_pnrdtl: string; flnbfl_pnrdtl: string;
  routfl_pnrdtl: string; agtnme_pnrdtl: string;
  srtspl_pnrdtl: string; srtcxl_pnrdtl: string;
  pagenw_pnrdtl: number; limitp_pnrdtl: number;
  airlfl_pnrsmr: string; flnbfl_pnrsmr: string;
  routfl_pnrsmr: string; datefl_pnrsmr: string;
  pnrcde_pnrsmr: string; pnrclk_pnrsmr: string;
  agtnme_pnrsmr: string; psdate_pnrsmr: string;
  srtspl_pnrsmr: string; srtcxl_pnrsmr: string;
  pagenw_pnrsmr: number; limitp_pnrsmr: number;
  airlfl_flnsmr: string; datefl_flnsmr: string;
  flnbfl_flnsmr: string; routfl_flnsmr: string;
  psdate_flnsmr: string; srtspl_flnsmr: string;
  srtcxl_flnsmr: string; pagenw_flnsmr: number;
  limitp_flnsmr: number; airlfl_agtnme: string;
  agtnme_agtnme: string; srtnul_agtnme: string;
  pagenw_agtnme: number; limitp_agtnme: number;
}

// Agent name
export interface MdlJeddahAgtnmeDtbase {
  prmkey: string; airlfl: string; agtnme: string;
  agtdtl: string; newdtl: string; newidn: string;
  agtidn: string; rtlsrs: string; updtby: string;
}
export interface MdlJeddahAgtnmeSearch {
  airlfl_agtnme: string;
  agtnme_agtnme: string;
  srtnul_agtnme: string;
  pagenw_agtnme: number;
  limitp_agtnme: number;
}
export interface MdlJeddahAgtnmeDefapi {
  arrdta: MdlJeddahAgtnmeDtbase[];
  totdta: number;
}


// Log action
export interface MdlJeddahParamsActlog {
  timeup: number;
  dateup: number;
  statdt: string;
}

// Add flight number
export interface MdlJeddahFlnbflDtbase {
  datefl: string; airlfl: string; flnbfl: string;
  routfl: string; fltype: string; updtby: string;
}

// PNR detail
export interface MdlJeddahPnrdtlDtbase {
  prmkey: string; airlfl: string; flnbfl: string;
  depart: string; routfl: string; fldtef: string;
  clssfl: string; datefl: number; dateup: number;
  timeup: number; agtnme: string; agtdtl: string;
  agtidn: string; pnrcde: string; rtlsrs: string;
  drules: number; totisd: number; totbok: number;
  totpax: number; totspl: number; totcxl: number;
  arrspl: string; arrcxl: string; notedt: string;
}
export interface MdlJeddahPnrdtlSearch {
  clssfl_pnrdtl: string; airlfl_pnrdtl: string;
  datefl_pnrdtl: string; pnrcde_pnrdtl: string;
  flnbfl_pnrdtl: string; routfl_pnrdtl: string;
  agtnme_pnrdtl: string; srtspl_pnrdtl: string;
  srtcxl_pnrdtl: string; pagenw_pnrdtl: number;
  limitp_pnrdtl: number;
}
export interface MdlJeddahPnrdtlDefapi {
  arrdta: MdlJeddahPnrdtlDtbase[];
  totdta: number;
}

// PNR Summary
export interface MdlJeddahPnrsmrDtbase {
  prmkey: string; routfl: string; timedp: number;
  timerv: number; dateup: number; timeup: number;
  timecr: number; agtnme: string; agtdtl: string;
  agtidn: string; pnrcde: string; intrln: string;
  rtlsrs: string; arrcpn: string; agtdie: string;
  totisd: number; totbok: number; totpax: number;
  totcxl: number; totspl: number; arrspl: string;
  notedt: string;
}
export interface MdlJeddahPnrsmrSearch {
  airlfl_pnrsmr: string; flnbfl_pnrsmr: string;
  routfl_pnrsmr: string; datefl_pnrsmr: string;
  pnrcde_pnrsmr: string; agtnme_pnrsmr: string;
  psdate_pnrsmr: string; srtspl_pnrsmr: string;
  srtcxl_pnrsmr: string; pagenw_pnrsmr: number;
  limitp_pnrsmr: number;
}
export interface MdlJeddahPnrsmrDefapi {
  arrdta: MdlJeddahPnrsmrDtbase[];
  totdta: number;
}

// Fligth Summary
export interface MdlJeddahFlnsmrDtbase {
  prmkey: string; airlfl: string; flnbfl: string;
  flstat: string; depart: string; routfl: string;
  datefl: number; dateup: number; timeup: number;
  totisd: number; totbok: number; totpax: number;
  totcxl: number; totchg: number; totspl: number;
  notedt: string;
}
export interface MdlJeddahFlnsmrSearch {
  airlfl_flnsmr: string; datefl_flnsmr: string;
  flnbfl_flnsmr: string; routfl_flnsmr: string;
  psdate_flnsmr: string; srtspl_flnsmr: string;
  srtcxl_flnsmr: string; pagenw_flnsmr: number;
  limitp_flnsmr: number;
}
export interface MdlJeddahFlnsmrDefapi {
  arrdta: MdlJeddahFlnsmrDtbase[];
  totdta: number;
}







