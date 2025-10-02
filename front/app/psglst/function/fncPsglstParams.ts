import { MdlPsglstAllprmSrcprm } from "../model/mdlPsglstParams";

// Treatment function params
export function FncPsglstDetailParams(params: MdlPsglstAllprmSrcprm) {
  return {
    mnthfl_detail: params.mnthfl_detail || "",
    datefl_detail: params.datefl_detail || "",
    airlfl_detail: params.airlfl_detail || "",
    flnbfl_detail: params.flnbfl_detail || "",
    depart_detail: params.depart_detail || "",
    routfl_detail: params.routfl_detail || "",
    isflwn_detail: params.isflwn_detail || "",
    istrst_detail: params.istrst_detail || "",
    pnrcde_detail: params.pnrcde_detail || "",
    tktnbr_detail: params.tktnbr_detail || "",
    istirg_detail: params.istirg_detail || "",
    mnthfl_others: params.mnthfl_others || "",
    datefl_others: params.datefl_others || "",
    airlfl_others: params.airlfl_others || "",
    flnbfl_others: params.flnbfl_others || "",
    depart_others: params.depart_others || "",
    routfl_others: params.routfl_others || "",
    isflwn_others: params.isflwn_others || "",
    istrst_others: params.istrst_others || "",
    pnrcde_others: params.pnrcde_others || "",
    tktnbr_others: params.tktnbr_others || "",
    istirg_others: params.istirg_others || "",
  } as MdlPsglstAllprmSrcprm;
}
