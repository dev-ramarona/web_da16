import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";
import {
  MdlPsglstPsgdtlFrntnd,
  MdlPsglstPsgdtlSearch,
} from "../model/mdlPsglstParams";

// Function get psglst database
export async function ApiPsglstPsgdtlGetall(params: MdlPsglstPsgdtlSearch) {
  var arrdta: MdlPsglstPsgdtlFrntnd[] = [];
  var fnlrsl = { arrdta: arrdta, totdta: 0 };
  try {
    const rspnse = await ApiGlobalAxiospParams.post(
      "/psglst/psgdtl/getall",
      params
    );
    if (rspnse.status === 200) {
      fnlrsl = await rspnse.data;
    }
  } catch (error) {
    console.log(error);
  }
  return fnlrsl;
}

// Function get psglst database
export async function ApiPsglstPsgdtlUpdate(params: MdlPsglstPsgdtlFrntnd) {
  var fnlrsl: string = "";
  if (params.tktnvc == "") {
    return "tktnvc empty";
  }
  if (params.tktnvc.length !== 13) {
    return "tktnvc invalid";
  }
  if (params.airlvc == "") {
    return "airlvc empty";
  }
  if (params.flnbvc == "") {
    return "flnbvc empty";
  }
  if (params.cpnbvc == 0) {
    return "cpnbvc empty";
  }
  if (params.routvc == "") {
    return "routvc empty";
  }
  if (params.statvc == "") {
    return "statvc empty";
  }
  if (
    params.slsrpt == "NOT CLEAR" &&
    (params.curncy == "" || params.ntafvc == 0)
  )
    return "curncy empty";
  try {
    const rspnse = await ApiGlobalAxiospParams.post(
      "/psglst/psgdtl/update",
      params
    );
    if (rspnse.status === 200) {
      fnlrsl = await rspnse.data;
    }
  } catch (error) {
    console.log(error);
  }
  return fnlrsl;
}

// Function get psglst summary PNR
export async function ApiPsglstPnrsmrDownld(params: MdlPsglstPsgdtlSearch) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post(
      `/psglst/psgdtl/getall/downld`,
      params,
      {
        timeout: 70000,
        responseType: "blob",
      }
    );

    // Buat Blob URL untuk file besar
    const blobfl = new Blob([rspnse.data], { type: "text/csv" });
    const dwnurl = window.URL.createObjectURL(blobfl);

    // Ambil nama file dari header
    const rawnme = rspnse.headers["content-disposition"];
    let flname = "download.csv";
    if (rawnme && rawnme.includes("filename="))
      flname = rawnme.split("filename=")[1].replace(/['"]/g, "");

    // Trigger unduhan file
    const a = document.createElement("a");
    a.href = dwnurl;
    a.download = flname;
    document.body.appendChild(a);
    a.click();

    // Bersihkan URL blob
    window.URL.revokeObjectURL(dwnurl);
    document.body.removeChild(a);

    return true;
  } catch (error) {
    console.error("Error downloading CSV file:", error);
  }
  return false;
}
