
'use client'
import React, { useEffect, useState } from "react";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import { mdlGlobalAllusrCookie } from "@/app/global/model/mdlGlobalPrimer";
import { ApiJeddahFlnbflTmplte, ApiJeddahFlnbflUpload } from "../../api/apiJeddahFlnbfl";
import { ApiGlobalStatusIntrvl, ApiGlobalStatusPrcess } from "@/app/global/api/apiGlobalPrimer";
import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";

export default function UixJeddahFlnbflUpload({
  cookie,
}: {
  cookie: mdlGlobalAllusrCookie;
}) {

  // Upload data
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

  // Download template
  const dwntmp = () => {
    ApiJeddahFlnbflTmplte();
  }

  // File Upload csv Function
  const [alrtup, alrtupSet] = useState<string>("Upload");
  useEffect(() => {
    const gtstat = async () => {
      const status = await ApiGlobalStatusPrcess();
      alrtupSet(status.sbrapi == 0 ? "Done" : `Wait ${status.sbrapi}%`);
      if (status.sbrapi != 0) {
        await ApiGlobalStatusIntrvl(alrtupSet, "action");
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
          await ApiJeddahFlnbflUpload(fileup, cookie.stfnme);
          rplprm(["pnrclk_pnrdtl", "pnrclk_pnrsmr"], String(Math.random()));
          return await ApiGlobalStatusIntrvl(alrtupSet, "action");
        } else alrtupSet(`Wait ${status.action}%`);
      }
    setTimeout(() => alrtupSet("Upload"), 800);
  };

  return (
    <>
      <div className="w-5/12 h-fit py-1 flexctr flex-wrap gap-y-3">
        <div className="w-full md:w-2/4 min-w-20 md:min-w-32 h-10 flexctr">
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
        <div className="w-full md:w-1/4 min-w-20 md:min-w-32 h-10 flexctr">
          <div className="afull flexctr p-1.5">
            <div className="afull btnsbm flexctr" onClick={() => actupl()}>{alrtup}</div>
          </div>
        </div>
        <div className="w-full md:w-1/4 min-w-20 md:min-w-32 h-10 flexctr">
          <div className="afull flexctr p-1.5">
            <div className="afull btnsbm flexctr text-center" onClick={() => dwntmp()}>
              <span className="leading-3">Download Template</span>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
