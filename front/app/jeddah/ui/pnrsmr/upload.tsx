"use client";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import { useEffect, useState } from "react";
import {
  ApiJeddahPnrsmrDownld,
  ApiJeddahRtlsrsStatus,
  ApiJeddahRtlsrsUpload,
} from "../../api/apiJeddahDtbase";
import {
  MdlJeddahInputxAllpnr,
  MdlJeddahSearchPnrsmr,
} from "../../model/mdlJeddahParams";
import { ApiGlobalStatusIntrvl } from "@/app/global/api/apiGlobalPrimer";
import { mdlGlobalAllusrCookie } from "@/app/global/model/mdlGlobalAllusr";
import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";

export default function UixJeddahPnrsmrUpldwn({
  trtprm,
  cookie,
}: {
  trtprm: MdlJeddahInputxAllpnr;
  cookie: mdlGlobalAllusrCookie;
}) {
  // File Upload csv Variable
  const [fileup, fileupSet] = useState<FileList | null>(null);
  const [filenm, filenmSet] = useState<string>("");
  const filech = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      fileupSet(e.target.files);
      if (e.target.files[0]) {
        filenmSet(e.target.files[0].name);
        if (e.target.files.length > 1)
          filenmSet(`${e.target.files.length} files selected`);
      }
    }
  };

  // File Upload csv Function
  const apistt = ApiJeddahRtlsrsStatus;
  const [alrtup, alrtupSet] = useState<string>("Upload");
  const [intrvl, intrvlSet] = useState<NodeJS.Timeout | null>(null);
  useEffect(() => {
    const gtstat = async () => {
      const status = await ApiJeddahRtlsrsStatus();
      alrtupSet(status);
      if (status != "Done") {
        await ApiGlobalStatusIntrvl(alrtupSet, intrvlSet, apistt);
      } else alrtupSet("Upload");
    };
    gtstat();
  }, []);

  // Hit the database and get interval status
  const rplprm = FncGlobalParamsEdlink();
  const actupl = async () => {
    if (!fileup) alrtupSet("Failed File not selected");
    if (fileup)
      if (fileup.length < 1) alrtupSet("Failed file not selected");
      else if (fileup[0].type != "text/csv")
        alrtupSet("Failed, please upload CSV file");
      else {
        const status = await ApiJeddahRtlsrsStatus();
        if (status == "Done") {
          alrtupSet("Wait");
          await ApiJeddahRtlsrsUpload(fileup, cookie.stfnme);
          rplprm(["pnrclk_pnrdtl", "pnrclk_pnrsmr"], String(Math.random()));
          return await ApiGlobalStatusIntrvl(alrtupSet, intrvlSet, apistt);
        } else alrtupSet(status);
      }
    setTimeout(() => alrtupSet("Upload"), 800);
  };

  const getprm = (trtprm: MdlJeddahSearchPnrsmr) => {
    return {
      airlfl_pnrsmr: trtprm.airlfl_pnrsmr || "",
      flnbfl_pnrsmr: trtprm.flnbfl_pnrsmr || "",
      routfl_pnrsmr: trtprm.routfl_pnrsmr || "",
      datefl_pnrsmr: trtprm.datefl_pnrsmr || "",
      pnrcde_pnrsmr: trtprm.pnrcde_pnrsmr || "",
      agtnme_pnrsmr: trtprm.agtnme_pnrsmr || "",
      psdate_pnrsmr: trtprm.psdate_pnrsmr || "",
      srtspl_pnrsmr: trtprm.srtspl_pnrsmr || "",
      srtcxl_pnrsmr: trtprm.srtcxl_pnrsmr || "",
      pagenw_pnrsmr: trtprm.pagenw_pnrsmr || 1,
      limitp_pnrsmr: trtprm.limitp_pnrsmr || 15,
    } as MdlJeddahSearchPnrsmr;
  };
  const [params, paramsSet] = useState(getprm(trtprm));
  useEffect(() => {
    paramsSet(getprm(trtprm));
  }, [trtprm]);

  // Download csv summary pnr
  const [dwnrsp, dwnrspSet] = useState("Download PNR Null Retail/Series");
  const dwnapi = async () => {
    dwnrspSet("Wait");
    const rspdwn = await ApiJeddahPnrsmrDownld(params, "rtlsrs");
    rspdwn ? dwnrspSet("Success") : dwnrspSet("Failed");
    setTimeout(() => dwnrspSet("Download PNR Null Retail/Series"), 500);
  };
  return (
    <div className="w-full h-fit py-1 flexstr flex-wrap gap-y-3 border-r-2 border-sky-200">
      <div className="w-full md:w-72 h-10 flexctr">
        <UixGlobalInputxFormdt
          typipt={"file"}
          length={undefined}
          queryx={"csfile_addfln"}
          params={filenm}
          plchdr="Choose file"
          repprm={filech}
          labelx="hidden"
        />
      </div>
      <div className="w-1/2 md:w-36 h-10 flexctr">
        <div className="afull flexctr p-1.5">
          <div
            className={`afull flexctr text-center ${
              alrtup.includes("Failed") ? "shkeit btncxl" : "btnsbm"
            } duration-300`}
            onClick={() => actupl()}
          >
            <span className="leading-3">{alrtup}</span>
          </div>
        </div>
      </div>
      <div className="w-1/2 md:w-36 h-10 flexctr">
        <div className="afull flexctr p-1.5">
          <div
            className="afull btnwrn flexctr text-center"
            onClick={() => dwnapi()}
          >
            <span className="leading-3">{dwnrsp}</span>
          </div>
        </div>
      </div>
    </div>
  );
}
