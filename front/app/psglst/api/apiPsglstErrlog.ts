import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";
import {
  MdlPsglstErrlogDtbase,
  MdlPsglstSrcprmAllprm,
} from "../model/mdlPsglstParams";

// Function get jeddah database Errlog
export async function ApiPsglstErrlogDtbase(params: MdlPsglstSrcprmAllprm) {
  var arrdta: MdlPsglstErrlogDtbase[] = [];
  var fnlrsl = { arrdta: arrdta, totdta: 0 };
  try {
    const rspnse = await ApiGlobalAxiospParams.post(
      "/psglst/errlog/getall",
      params,
    );
    if (rspnse.status === 200) {
      fnlrsl = await rspnse.data;
    }
  } catch (error) {
    console.log(error);
  }
  return fnlrsl;
}
