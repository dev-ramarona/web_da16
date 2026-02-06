import { AxiosError } from "axios";
import { ApiGlobalAxiospParams } from "./apiGlobalPrimer";
import { MdlGlobalApplstDtbase } from "../model/mdlGlobalApplst";
import { MdlGlobalAllusrParams } from "../model/mdlGlobalPrimer";

// API Login
export async function ApiGlobalAllusrLogin(
  prvstt: object | null | void,
  formdt: FormData,
) {
  const usrnme = formdt.get("usrnme") as string;
  const psswrd = formdt.get("psswrd") as string;
  const dataInput = { usrnme: usrnme, psswrd: psswrd };
  const errorObject = {
    usrnme: "",
    psswrd: "",
    rspnse: "",
    dfault: usrnme,
  };
  if (!usrnme) errorObject.usrnme = "Username Empty";
  if (!psswrd) errorObject.psswrd = "psswrd Empty";
  if (!usrnme || !psswrd) return errorObject;
  else
    try {
      await ApiGlobalAxiospParams.post("/allusr/loginx", dataInput, {
        withCredentials: true,
      });
      window.location.href = "/home";
    } catch (error) {
      if (error instanceof AxiosError) {
        if (error.response?.data.error === "usrnme") {
          errorObject.usrnme = "Username not found";
        } else if (error.response?.data.error === "psswrd") {
          errorObject.psswrd = "Password invalid";
        } else {
          errorObject.rspnse = "Invalid credential";
        }
      } else if (error instanceof Error) {
        errorObject.rspnse = error.message;
      } else {
        errorObject.rspnse = "An unknown error occurred";
      }
      return errorObject;
    }
}

// API Logout
export async function ApiGlobalAllusrLogout() {
  try {
    ApiGlobalAxiospParams.get("/allusr/logout", {
      withCredentials: true,
    });
    window.location.href = "/log";
  } catch (error) {
    console.log("Error logout", error);
  }
}

// API Applist
export async function ApiGlobalAllusrApplst() {
  const fnlObj: MdlGlobalApplstDtbase[] = [];
  try {
    const rspnse = await ApiGlobalAxiospParams.get("/allusr/applst");
    if (rspnse.status === 200) {
      const fnlObj: MdlGlobalApplstDtbase[] = await rspnse.data;
      return fnlObj;
    }
  } catch (error) {
    console.log(error);
  }
  return fnlObj;
}

// API Register
export async function ApiGlobalAllusrRegist(params: MdlGlobalAllusrParams) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post("/allusr/regist", params);
    if (rspnse.status === 200) return true;
  } catch (error) {
    console.log(error);
  }
  return false;
}
