import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";
import { MdlHoldstErrlogDtbase } from "../model/mdlHoldstParams";

export async function ApiHoldstPrcessManual(params: MdlHoldstErrlogDtbase) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post("/holdst/prcess", params);
    if (rspnse.status === 200) {
      const fnlobj = await rspnse.data;
      console.log(fnlobj);
    }
  } catch (error) {
    console.log(error);
  }
}
