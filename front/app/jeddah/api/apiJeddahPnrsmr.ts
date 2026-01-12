import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";
import {
  MdlJeddahPnrsmrDefapi,
  MdlJeddahPnrsmrSearch,
} from "../model/mdlJeddahMainpr";

// Function get jeddah summary PNR
export async function ApiJeddahPnrsmrGetall(params: MdlJeddahPnrsmrSearch) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post(
      "/jeddah/pnrsmr/getall",
      params
    );
    if (rspnse.status === 200) {
      const fnlobj: MdlJeddahPnrsmrDefapi = await rspnse.data;
      return fnlobj;
    }
  } catch (error) {
    console.log(error);
  }
  const defdta: MdlJeddahPnrsmrDefapi = { arrdta: [], totdta: 1 };
  return defdta;
}

// Function get jeddah summary PNR
export async function ApiJeddahPnrsmrDownld(
  params: MdlJeddahPnrsmrSearch,
  typdwn: "downld"
) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post(
      `/jeddah/pnrsmr/getall/${typdwn}`,
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
