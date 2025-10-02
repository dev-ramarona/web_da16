// Allusr
export interface MdlGlobalAllusrParams {
  stfnme: string;
  usrnme: string;
  stfeml: string;
  psswrd: string;
  access: string[];
  keywrd: string[];
}
export interface mdlGlobalAllusrCookie {
  stfnme: string;
  usrnme: string;
  access: string[];
  keywrd: string[];
}
export interface mdlGlobalAlluserFilter {
  keywrd: string;
  output: string;
}
export interface mdlGlobalAlluserStatus {
  keywrd: string;
  output: string;
}

// Applist
export interface MdlGlobalApplstDtbase {
  prmkey: string;
  detail: string;
}

// Status data
export interface MdlGlobalStatusPrcess {
  sbrapi: number;
  action: number;
}