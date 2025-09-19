"use server";

import { cookies } from "next/headers";
import { mdlGlobalAllusrCookie } from "../model/mdlGlobalAllusr";

// Handle Cookie
export async function ApiGlobalCookieGetdta() {
  try {
    const nowusr = (await cookies()).get("nowusr")?.value || "";
    const Objusr: mdlGlobalAllusrCookie = JSON.parse(nowusr);
    return Objusr;
  } catch (error) {
    const Objusr: mdlGlobalAllusrCookie = {
      stfnme: "",
      usrnme: "",
      access: ["null"],
      keywrd: ["null"],
    };
    return Objusr;
  }
}
