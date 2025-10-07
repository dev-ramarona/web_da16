import { MdlPsglstActlogDtbase } from "../model/mdlPsglstParams";

// Function get jeddah database log action
export async function ApiPsglstActlogDtbase() {
    //   try {
    //     const rspnse = await ApiGlobalAxiospParams.get("/opclss/logact");
    //     if (rspnse.status === 200) {
    //       const fnlobj = await rspnse.data;
    //       return fnlobj;
    //     }
    //   } catch (error) {
    //     console.log(error);
    //   }
    //   return [];
    const tempry: MdlPsglstActlogDtbase[] = [
        {
            dateup: 2508221647,
            statdt: "Final",
            timeup: 2508221647,
        },
    ];
    for (let i = 0; i < 20; i++) {
        tempry.push(tempry[0]);
    }
    return tempry;
}