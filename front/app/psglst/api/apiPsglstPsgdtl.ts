import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";
import {
  MdlPsglstPsgdtlFrntnd,
  MdlPsglstSrcprmAllprm,
} from "../model/mdlPsglstParams";

// Function get psglst database
export async function ApiPsglstPsgdtlGetall(params: MdlPsglstSrcprmAllprm) {
  var arrdta: MdlPsglstPsgdtlFrntnd[] = [];
  var fnlrsl = { arrdta: arrdta, totdta: 0 };
  try {
    const rspnse = await ApiGlobalAxiospParams.post(
      "/psglst/psgdtl/getall",
      params
    );
    if (rspnse.status === 200) {
      fnlrsl = await rspnse.data;
    }
  } catch (error) {
    console.log(error);
  }
  return fnlrsl;
}

// Function get psglst database
export async function ApiPsglstPsgdtlUpdate(params: MdlPsglstPsgdtlFrntnd) {
  var fnlrsl: string = "";
  if (params.tktnvc == "") {
    return "tktnvc empty";
  }
  if (params.tktnvc.length !== 13) {
    return "tktnvc invalid";
  }
  if (params.airlvc == "") {
    return "airlvc empty";
  }
  if (params.flnbvc == "") {
    return "flnbvc empty";
  }
  if (params.cpnbvc == 0) {
    return "cpnbvc empty";
  }
  if (params.routvc == "") {
    return "routvc empty";
  }
  if (params.statvc == "") {
    return "statvc empty";
  }
  if (
    params.slsrpt == "NOT CLEAR" &&
    (params.curncy == "" || params.ntafvc == 0)
  )
    return "curncy empty";
  try {
    const rspnse = await ApiGlobalAxiospParams.post(
      "/psglst/psgdtl/update",
      params
    );
    if (rspnse.status === 200) {
      fnlrsl = await rspnse.data;
    }
  } catch (error) {
    console.log(error);
  }
  return fnlrsl;
}
