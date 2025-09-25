import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";
import {
  MdlJeddahDefapihPnrdtl,
  MdlJeddahDefapihFlnsmr,
  MdlJeddahDefapihPnrsmr,
  MdlJeddahParamsAddlfn,
  MdlJeddahSearchPnrdtl,
  MdlJeddahSearchFlnsmr,
  MdlJeddahSearchPnrsmr,
  MdlJeddahParamsPnrsmr,
} from "../model/mdlJeddahParams";

// Function get jeddah summary PNR
export async function ApiJeddahPnrdtlGetarr(params: MdlJeddahSearchPnrdtl) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post("/jeddah/pnrdtl", params);
    if (rspnse.status === 200) {
      const fnlobj: MdlJeddahDefapihPnrdtl = await rspnse.data;
      return fnlobj;
    }
  } catch (error) {
    console.log(error);
  }
  const defdta: MdlJeddahDefapihPnrdtl = { arrdta: [], totdta: 1 };
  return defdta;
}

// // Function get jeddah summary PNR
export async function ApiJeddahPnrdtlDownld(params: MdlJeddahSearchPnrdtl) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post(
      "/jeddah/pnrdtl/downld",
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

// Function get jeddah summary PNR
export async function ApiJeddahPnrsmrGetarr(params: MdlJeddahSearchPnrsmr) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post("/jeddah/pnrsmr", params);
    if (rspnse.status === 200) {
      const fnlobj: MdlJeddahDefapihPnrsmr = await rspnse.data;
      return fnlobj;
    }
  } catch (error) {
    console.log(error);
  }
  const defdta: MdlJeddahDefapihPnrsmr = { arrdta: [], totdta: 1 };
  return defdta;
}

// // Function get jeddah summary PNR
export async function ApiJeddahPnrsmrDownld(
  params: MdlJeddahSearchPnrsmr,
  typdwn: "downld" | "rtlsrs"
) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post(
      `/jeddah/pnrsmr/${typdwn}`,
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

// Function get jeddah summary Flight Number
export async function ApiJeddahFlnsmrGetarr(params: MdlJeddahSearchFlnsmr) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post("/jeddah/flnsmr", params);
    if (rspnse.status === 200) {
      const fnlobj: MdlJeddahDefapihFlnsmr = await rspnse.data;
      return fnlobj;
    }
  } catch (error) {
    console.log(error);
  }
  const defdta: MdlJeddahDefapihFlnsmr = { arrdta: [], totdta: 1 };
  return defdta;
}

// Function get jeddah summary Flight number
export async function ApiJeddahFlnsmrDownld(params: MdlJeddahSearchFlnsmr) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post(
      "/jeddah/flnsmr/downld",
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

// Function get jeddah database log action
export async function ApiJeddahDtbaseLogact() {
  try {
    const rspnse = await ApiGlobalAxiospParams.get("/jeddah/logact");
    if (rspnse.status === 200) {
      const fnlobj = await rspnse.data;
      return fnlobj;
    }
  } catch (error) {
    console.log(error);
  }
  return [];
}

// Function add flight number to database
export async function ApiJeddahDtbaseAddfln(addfln: MdlJeddahParamsAddlfn) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post("/jeddah/addfln", addfln);
    if (rspnse.status === 200) {
      const fnlobj = await rspnse.data;
      console.log(fnlobj);
    }
  } catch (error) {
    console.log(error);
  }
}

// Function update PNR Summary Retali or series
export async function ApiJeddahRtlsrsUpdate(params: MdlJeddahParamsPnrsmr) {
  try {
    const rspnse = await ApiGlobalAxiospParams.post(
      "/jeddah/rtlsrs/update",
      params
    );
    if (rspnse.status === 200) {
      const fnlobj: MdlJeddahDefapihFlnsmr = await rspnse.data;
      return fnlobj;
    }
  } catch (error) {
    console.log(error);
  }
  const defdta: MdlJeddahDefapihFlnsmr = { arrdta: [], totdta: 1 };
  return defdta;
}

// Function upload PNR Summary Retali or series
export async function ApiJeddahRtlsrsUpload(
  fl1eup: FileList | null,
  upldby: string
) {
  if (!fl1eup) return;
  const frmdta = new FormData();
  for (let i = 0; i < fl1eup.length; i++) frmdta.append("fl1eup", fl1eup[i]);
  try {
    const rspnse = await ApiGlobalAxiospParams.post(
      `/jeddah/rtlsrs/upload/${upldby || "x"}`,
      frmdta,
      {
        headers: { "Content-Type": "multipart/form-data" },
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
    else return;

    // Trigger unduhan file
    const a = document.createElement("a");
    a.href = dwnurl;
    a.download = flname;
    document.body.appendChild(a);
    a.click();

    // Bersihkan URL blob
    window.URL.revokeObjectURL(dwnurl);
    document.body.removeChild(a);
  } catch (error) {
    console.log(error);
  }
}

// Hit status sabre api
export async function ApiJeddahRtlsrsStatus() {
  try {
    const rspnse = await ApiGlobalAxiospParams.get("/jeddah/rtlsrs/status");
    if (rspnse.status === 200) {
      const fnlstr: string = await rspnse.data;
      return fnlstr;
    }
  } catch (error) {
    console.log(error);
  }
  return "failed";
}
