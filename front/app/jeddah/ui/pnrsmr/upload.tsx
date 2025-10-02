"use client";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import { useEffect, useState } from "react";
import {
  MdlJeddahInputxAllpnr,
  MdlJeddahPnrsmrSearch,
} from "../../model/mdlJeddahMainpr";
import { ApiGlobalStatusIntrvl, ApiGlobalStatusPrcess } from "@/app/global/api/apiGlobalPrimer";
import { mdlGlobalAllusrCookie } from "@/app/global/model/mdlGlobalPrimer";
import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";
import { ApiJeddahPnrsmrDownld } from "../../api/apiJeddahPnrsmr";
import { ApiJeddahRtlsrsTmplte, ApiJeddahRtlsrsUpload } from "../../api/apiJeddahRtlsrs";

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
  const [alrtup, alrtupSet] = useState<string>("Upload");
  const [intrvl, intrvlSet] = useState<NodeJS.Timeout | null>(null);
  useEffect(() => {
    const gtstat = async () => {
      const status = await ApiGlobalStatusPrcess();
      alrtupSet(status.sbrapi == 0 ? "Done" : `Wait ${status.sbrapi}%`);
      if (status.sbrapi != 0) {
        await ApiGlobalStatusIntrvl(alrtupSet, intrvlSet, "action");
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
        const status = await ApiGlobalStatusPrcess();
        if (status.action == 0) {
          alrtupSet("Wait");
          await ApiJeddahRtlsrsUpload(fileup, cookie.stfnme);
          rplprm(["pnrclk_pnrdtl", "pnrclk_pnrsmr"], String(Math.random()));
          return await ApiGlobalStatusIntrvl(alrtupSet, intrvlSet, "action");
        } else alrtupSet(`Wait ${status.action}%`);
      }
    setTimeout(() => alrtupSet("Upload"), 800);
  };


  // Download csv summary pnr
  const [dwnrsp, dwnrspSet] = useState("Download PNR Null Retail/Series");
  const dwnapi = async () => {
    dwnrspSet("Wait");
    const rspdwn = await ApiJeddahRtlsrsTmplte();
    rspdwn ? dwnrspSet("Success") : dwnrspSet("Failed");
    setTimeout(() => dwnrspSet("Download PNR Null Retail/Series"), 500);
  };
  return (
    <div className="w-full h-fit py-1 flexstr flex-wrap gap-y-3 border-r-2 border-sky-200">
      <div className="w-full md:w-72 h-10 flexctr">
        <UixGlobalInputxFormdt
          typipt={"file"}
          length={undefined}
          queryx={"csfile_rtlsrs"}
          params={filenm}
          plchdr="Choose file"
          repprm={filech}
          labelx="hidden"
        />
      </div>
      <div className="w-1/2 md:w-36 h-10 flexctr">
        <div className="afull flexctr p-1.5">
          <div
            className={`afull flexctr text-center ${alrtup.includes("Failed") ? "shkeit btncxl" : "btnsbm"
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
