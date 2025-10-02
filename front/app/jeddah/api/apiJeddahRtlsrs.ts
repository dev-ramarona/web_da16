import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";
import { MdlJeddahFlnsmrDefapi, MdlJeddahPnrsmrDtbase } from "../model/mdlJeddahMainpr";

// Function download template flight number upload
export async function ApiJeddahRtlsrsTmplte() {
    try {
        const rspnse = await ApiGlobalAxiospParams.get(
            "/jeddah/rtlsrs/tmplte", { responseType: "blob" });

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


// Function update PNR Summary Retali or series
export async function ApiJeddahRtlsrsUpdate(params: MdlJeddahPnrsmrDtbase) {
    try {
        const rspnse = await ApiGlobalAxiospParams.post(
            "/jeddah/rtlsrs/update",
            params
        );
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