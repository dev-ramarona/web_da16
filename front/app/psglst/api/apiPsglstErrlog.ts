import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";
import {
  MdlPsglstErrlogDtbase,
  MdlPsglstErrlogSrcprm,
} from "../model/mdlPsglstParams";

// Function get jeddah database Errlog
export async function ApiPsglstErrlogDtbase(prmErrlog: MdlPsglstErrlogSrcprm) {
  var arrdta: MdlPsglstErrlogDtbase[] = [];
  var fnlrsl = { arrdta: arrdta, totdta: 0 };
  try {
    const rspnse = await ApiGlobalAxiospParams.post(
      "/psglst/errlog/getall",
      prmErrlog,
    );
    if (rspnse.status === 200) {
      fnlrsl = await rspnse.data;
    }
  } catch (error) {
    console.log(error);
  }
  return fnlrsl;
}
