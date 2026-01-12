import { MdlHoldstErrlogDtbase } from "../model/mdlHoldstParams";

// Treatment function params
export function FncHoldstDetailParams(params: MdlHoldstErrlogDtbase) {
  return {
    depart: params.depart || "",
    worker: params.worker || "",
    datefl: params.datefl || "",
  } as MdlHoldstErrlogDtbase;
}
