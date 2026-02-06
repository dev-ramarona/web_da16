import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";
import { MdlPsglstAcpedtDtbase } from "../model/mdlPsglstParams";

// Function get jeddah Edit param accepted
export async function ApiPsglstAcpedtDtbase() {
  let fnlobj: MdlPsglstAcpedtDtbase[] = [];
  try {
    const rspnse = await ApiGlobalAxiospParams.get("/psglst/acpedt/getall");
    if (rspnse.status === 200) {
      fnlobj = await rspnse.data;
    }
  } catch (error) {
    console.log(error);
  }
  return fnlobj;
}
