import { MdlPsglstAcpedtDtbase } from "../model/mdlPsglstParams";

// Function get jeddah Edit param accepted
export async function ApiPsglstAcpedtDtbase() {
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
    const tempry: MdlPsglstAcpedtDtbase[] = [
        {
            params: "tktnbr",
            length: 13,
        },
        {
            params: "airlvc",
            length: 2,
        },
        {
            params: "flnbvc",
            length: 4,
        },
        {
            params: "cpnbvc",
            length: 2,
        },
        {
            params: "routvc",
            length: 7,
        },
        {
            params: "status",
            length: 4,
        },
        {
            params: "notedt",
            length: 90,
        },
    ];

    return tempry;
}