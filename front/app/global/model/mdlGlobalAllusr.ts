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
