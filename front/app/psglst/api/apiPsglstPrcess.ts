

import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";
import { MdlPsglstErrlogDtbase } from "../model/mdlPsglstParams";

export async function ApiPsglstPrcessManual(params: MdlPsglstErrlogDtbase) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post("/psglst/prcess", params);
    if (rspnse.status === 200) {
      const fnlobj = await rspnse.data;
      console.log(fnlobj);
    }
  } catch (error) {
    console.log(error);
  }
}

