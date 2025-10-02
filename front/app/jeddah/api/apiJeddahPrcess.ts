"/jeddah/prcess";

import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";

export async function ApiJeddahPrcessManual() {
  try {
    const rspnse = await ApiGlobalAxiospParams.post("/jeddah/prcess");
    if (rspnse.status === 200) {
      const fnlobj = await rspnse.data;
      console.log(fnlobj);
    }
  } catch (error) {
    console.log(error);
  }
}

