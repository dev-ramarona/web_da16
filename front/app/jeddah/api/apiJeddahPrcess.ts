"/jeddah/prcess";

import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";
import { MdlJeddahParamsActlog } from "../model/mdlJeddahMainpr";

export async function ApiJeddahPrcessManual(params: MdlJeddahParamsActlog) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post("/jeddah/prcess", params);
    if (rspnse.status === 200) {
      const fnlobj = await rspnse.data;
      console.log(fnlobj);
    }
  } catch (error) {
    console.log(error);
  }
}

