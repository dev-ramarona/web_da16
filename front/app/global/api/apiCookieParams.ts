"use server";


import { cookies } from "next/headers";
import { mdlGlobalAllusrCookie } from "../model/mdlGlobalPrimer";
import { ApiGlobalAxiospParams } from "./apiGlobalPrimer";

// Handle Cookie
export async function ApiGlobalCookieGetdta() {
  const fnccok = cookies();
  const tknnme = process.env.NEXT_PUBLIC_TKN_COOKIE || "x"
  const tokenx = (await fnccok).get(tknnme)?.value || "";
  const Objusr: mdlGlobalAllusrCookie
    = { stfnme: "", usrnme: "", access: ["null"], keywrd: ["null"] };
  if (tokenx == "" || !tokenx) return Objusr;

  // Try hit API
  try {
    const rspnse = await ApiGlobalAxiospParams.get("/allusr/tokenx", {
      headers: {
        Authorization: tokenx,
      },
    });

    // If status success
    if (rspnse.status === 200) {
      const fnlobj: mdlGlobalAllusrCookie = await rspnse.data;
      return fnlobj;
    }
  }

  // Cath error
  catch (error) {
    console.log(error);
  }

  // Return empty object
  return Objusr;
}
