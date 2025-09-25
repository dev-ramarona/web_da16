"use server";

import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";
import {
  MdlJeddahDefapihAgtnme,
  MdlJeddahInputxAllpnr,
  MdlJeddahParamsAgtedt,
  MdlJeddahParamsAgtnme,
  MdlJeddahSearchAgtnme,
} from "../model/mdlJeddahParams";
import { revalidatePath } from "next/cache";

// Get Agent name not complete format slice data from database
export async function ApiJeddahAgtnmeNullnm(trtprm: MdlJeddahInputxAllpnr) {
  const nowprm: MdlJeddahSearchAgtnme = {
    airlfl_agtnme: trtprm.airlfl_agtnme,
    agtnme_agtnme: trtprm.agtnme_agtnme,
    srtnul_agtnme: trtprm.srtnul_agtnme,
    pagenw_agtnme: trtprm.pagenw_agtnme,
    limitp_agtnme: trtprm.limitp_agtnme,
  };
  const fnlobj: MdlJeddahParamsAgtedt[] = [];
  try {
    const rspnse = await ApiGlobalAxiospParams.post("/jeddah/agtnme/nullvl", nowprm);
    if (rspnse.status === 200) {
      const fnlobj: MdlJeddahDefapihAgtnme = await rspnse.data;
      return fnlobj;
    }
  } catch (error) {
    console.log(error);
  }
  const defdta: MdlJeddahDefapihAgtnme = { arrdta: [], totdta: 1 };
  return defdta;
}

// Get Agent name match search params from database
export async function ApiJeddahAgtnmeNulsrc(newidn: string, newdtl: string) {
  try {
    const rspnse = await ApiGlobalAxiospParams.get(
      `/jeddah/agtnme/search/${newidn || "x"}/${newdtl || "x"}`
    );
    if (rspnse.status === 200) {
      const fnlobj: MdlJeddahParamsAgtnme = await rspnse.data;
      return fnlobj;
    }
  } catch (error) {
    console.log(error);
  }
  return {
    prmkey: "",
    airlnf: "",
    agtnme: "",
    agtdtl: "",
    newdtl: "",
    agtidn: "",
    newidn: "",
    rtlsrs: "",
  };
}

// Update Agent name detail to database
export async function ApiJeddahAgtnmeAgtupd(update: MdlJeddahParamsAgtedt) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post("/jeddah/agtnme/update", update);
    if (rspnse.status === 200) {
      const fnlstr: string = await rspnse.data;
      return fnlstr;
    }
  } catch (error) {
    console.log(error);
  }
  return "failed";
}
