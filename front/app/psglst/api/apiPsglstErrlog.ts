import { MdlPsglstErrlogDtbase } from "../model/mdlPsglstParams";

// Function get jeddah database Errlog
export async function ApiPsglstErrlogDtbase() {
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
    const tempry: MdlPsglstErrlogDtbase[] = [
        {
            prmkey: "test",
            status: "Cancel",
            errprt: "frtaxs",
            errtxt: "Fare Taxes Cannot found on Sabre Web API	",
            datefl: 250716,
            timeup: 2507160210,
            airlnf: "JT",
            depart: "DPS",
            flnumf: "789",
            flstat: "OPENCI",
            flhour: 1,
            routef: "DPS-CGK",
            dvsion: "mnfest",
            worker: 1,
            ignore: "",
            updtby: "",
        },
    ];
    for (let i = 0; i < 20; i++) {
        tempry.push(tempry[0]);
    }
    return tempry;
}