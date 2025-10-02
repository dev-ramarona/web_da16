import {
  MdlPsglstActlogParams,
  MdlPsglstEdtprmParams,
  MdlPsglstErrlogParams,
  MdlPsglstDetailParams,
} from "../model/mdlPsglstParams";

// Function get jeddah Edit param accepted
export async function ApiPsglstDtbaseEdtprm() {
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
  const tempry: MdlPsglstEdtprmParams[] = [
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

// Function get jeddah database log action
export async function ApiPsglstDtbaseActlog() {
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
  const tempry: MdlPsglstActlogParams[] = [
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

// Function get jeddah database Errlog
export async function ApiPsglstDtbaseErrlog() {
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
  const tempry: MdlPsglstErrlogParams[] = [
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

// Function get jeddah database Errlog
export async function ApiPsglstDtbaseDetail() {
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
  const tempry: MdlPsglstDetailParams[] = [
    {
      prmkey: "xxx",
      mnthfl: "2508",
      datefl: "250824",
      timefl: "2508240804",
      datevc: "250824",
      timevc: "2508240804",
      airlfl: "JT",
      airlvc: "JT",
      flnbfl: "789",
      flnbvc: "789",
      flhour: 1.5,
      depart: "CGK",
      arrivl: "DPS",
      airtyp: "738",
      routfl: "CGK-DPS",
      routvc: "CGK-DPS",
      routac: "CGK-AMQ-DPS",
      routcm: "CGK-AMQ-DPS-UPG",
      lnenbr: "1",
      isflwn: "Flown",
      istrst: "Transit",
      gender: "M",
      seatpx: "17A",
      groupx: "AX7",
      pnrcde: "TNJIRO",
      tktnbr: "9901341312312",
      Tktprv: "9901341312312",
      status: "USED",
      fstnme: "BAMBANG PAMUNGKAS MR",
      lstnme: "STEPEN",
      cpnbfl: "C01",
      cpnbvc: "C02",
      clssfl: "W",
      clssvc: "A",
      cabnfl: "Y",
      cabnvc: "Y",
      frrate: 0.55,
      framntflflwn: 1139764,
      framntflvcrs: 1139764,
      frtxyqflflwn: 113976,
      frtxyqflvcrs: 113976,
      frtxyrflflwn: 11397,
      frtxyrflvcrs: 11397,
      agtdie: "BHK",
      ireglr: "",
      istedt: "",
      notedt: "",
      nottxt: "",
      enhkey: "",
      updtby: "",
    },
  ];
  for (let i = 0; i < 20; i++) {
    let temp = { ...tempry[0] };
    temp.prmkey = "xxxaaa" + i;
    tempry.push(temp);
  }
  return tempry;
}
