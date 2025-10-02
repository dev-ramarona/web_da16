import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";
import { MdlJeddahPnrdtlDefapi, MdlJeddahPnrdtlSearch } from "../model/mdlJeddahMainpr";

// Function get jeddah summary PNR
export async function ApiJeddahPnrdtlDownld(params: MdlJeddahPnrdtlSearch) {
    try {
        const rspnse = await ApiGlobalAxiospParams.post(
            "/jeddah/pnrdtl/getall/downld",
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
export async function ApiJeddahPnrdtlGetall(params: MdlJeddahPnrdtlSearch) {
    try {
        const rspnse = await ApiGlobalAxiospParams.post("/jeddah/pnrdtl/getall", params);
        if (rspnse.status === 200) {
            const fnlobj: MdlJeddahPnrdtlDefapi = await rspnse.data;
            return fnlobj;
        }
    } catch (error) {
        console.log(error);
    }
    const defdta: MdlJeddahPnrdtlDefapi = { arrdta: [], totdta: 1 };
    return defdta;
}