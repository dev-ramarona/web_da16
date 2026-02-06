import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";
import { MdlPsglstActlogDtbase } from "../model/mdlPsglstParams";

// Function get jeddah database log action
export async function ApiPsglstActlogDtbase() {
  const actlog: MdlPsglstActlogDtbase[] = [];
  const datefl: string[] = [];
  let fnlrsl = { actlog: actlog, datefl: datefl };
  try {
    const rspnse = await ApiGlobalAxiospParams.get("/psglst/actlog/getall");
    if (rspnse.status === 200) {
      const fnlobj = await rspnse.data;
      fnlrsl = fnlobj;
    }
  } catch (error) {
    console.log(error);
  }
  return fnlrsl;
}
