import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";
import { MdlJeddahFlnsmrDefapi, MdlJeddahFlnsmrSearch } from "../model/mdlJeddahMainpr";

// Function get jeddah summary Flight number
export async function ApiJeddahFlnsmrDownld(params: MdlJeddahFlnsmrSearch) {
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


// Function get jeddah summary Flight Number
export async function ApiJeddahFlnsmrGetall(params: MdlJeddahFlnsmrSearch) {
    try {
        const rspnse = await ApiGlobalAxiospParams.post("/jeddah/flnsmr/getall", params);
        if (rspnse.status === 200) {
            const fnlobj: MdlJeddahFlnsmrDefapi = await rspnse.data;
            return fnlobj;
        }
    } catch (error) {
        console.log(error);
    }
    const defdta: MdlJeddahFlnsmrDefapi = { arrdta: [], totdta: 1 };
    return defdta;
}