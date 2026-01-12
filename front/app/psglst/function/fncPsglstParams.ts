import { MdlPsglstSrcprmAllprm } from "../model/mdlPsglstParams";

// Treatment function params
export function FncPsglstDetailParams(params: MdlPsglstSrcprmAllprm) {
  return {
    update_psgdtl: params.update_psgdtl || "",
    mnthfl_psgdtl: params.mnthfl_psgdtl || "",
    datefl_psgdtl: params.datefl_psgdtl || "",
    airlfl_psgdtl: params.airlfl_psgdtl || "",
    flnbfl_psgdtl: params.flnbfl_psgdtl || "",
    depart_psgdtl: params.depart_psgdtl || "",
    routfl_psgdtl: params.routfl_psgdtl || "",
    pnrcde_psgdtl: params.pnrcde_psgdtl || "",
    tktnfl_psgdtl: params.tktnfl_psgdtl || "",
    isitfl_psgdtl: params.isitfl_psgdtl || "",
    isittx_psgdtl: params.isittx_psgdtl || "",
    isitir_psgdtl: params.isitir_psgdtl || "",
    nclear_psgdtl: params.nclear_psgdtl || "",
    pagenw_psgdtl: Number(params.pagenw_psgdtl) || 1,
    limitp_psgdtl: Number(params.limitp_psgdtl) || 15,
    erdvsn_errlog: params.erdvsn_errlog || "",
    pagenw_errlog: Number(params.pagenw_errlog) || 1,
    limitp_errlog: Number(params.limitp_errlog) || 5,
  } as MdlPsglstSrcprmAllprm;
}
