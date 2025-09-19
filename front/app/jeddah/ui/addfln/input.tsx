"use client";

import React, { useState } from "react";
import { MdlJeddahParamsAddlfn } from "../../model/mdlJeddahParams";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import {
  FncGlobalFormatFilter,
  FncGlobalFormatRoutfl,
} from "@/app/global/function/fncGlobalFormat";
import { mdlGlobalAlluserFilter } from "@/app/global/model/mdlGlobalAllusr";

export default function UixJeddahAddflnInpupl() {
  // Input partial data
  const [inputx, inputxSet] = useState<MdlJeddahParamsAddlfn>({
    airlfl_addfln: "",
    fldtef_addfln: "",
    flnbfl_addfln: "",
    routfl_addfln: "",
    fltype_addfln: "",
  });
  const iptrsp: MdlJeddahParamsAddlfn = {
    airlfl_addfln: "Airlines is Empty",
    fldtef_addfln: "Flight Date is Empty",
    flnbfl_addfln: "Flight Number is Empty",
    routfl_addfln: "Route is Empty",
    fltype_addfln: "Flight Type is Empty",
  };

  // Action add Flight number
  const [inptrs, inptrsSet] = useState("");
  const iptact = async () => {
    inptrsSet("");
    for (const [key, val] of Object.entries(inputx))
      if (val == "") return inptrsSet(iptrsp[key as keyof typeof iptrsp]);
    const updapi: string = "Failed to connect Database";
    await new Promise((r) => setTimeout(r, 2000));
    if (updapi != "Success") return inptrsSet(updapi);
    inptrsSet("");
  };

  // Onchange input data
  const iptcgh = (e: React.ChangeEvent<HTMLInputElement>) => {
    const filter: mdlGlobalAlluserFilter[] = [
      { keywrd: "OUT", output: "Outgoing" },
      { keywrd: "INC", output: "Incoming" },
      { keywrd: "OJD", output: "Non Jeddah" },
    ];
    let nameid = e.currentTarget.id;
    let valuex = e.currentTarget.value;
    if (nameid == "flnbfl_addfln") valuex = valuex.replace(/[^0-9]/g, "");
    else if (nameid == "routfl_addfln") valuex = FncGlobalFormatRoutfl(valuex);
    else if (nameid == "fltype_addfln")
      valuex = FncGlobalFormatFilter(valuex, filter);
    else valuex = valuex.toUpperCase();
    inputxSet((prev) => ({
      ...prev,
      [nameid]: valuex,
    }));
    inptrsSet("");
  };

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

  return (
    <div className="w-full h-fit flexctr">
      <div className="w-7/12 h-fit py-1 flexctr flex-wrap gap-y-3 border-r-2 border-sky-200">
        <div className="w-1/2 md:w-1/6 min-w-20 md:min-w-32 h-10 flexctr">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={2}
            queryx={"airlfl_addfln"}
            params={inputx.airlfl_addfln}
            plchdr="Airline"
            repprm={iptcgh}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-1/6 min-w-20 md:min-w-32 h-10 flexctr">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={4}
            queryx={"flnbfl_addfln"}
            params={inputx.flnbfl_addfln}
            plchdr="Flight Number"
            repprm={iptcgh}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-1/6 min-w-20 md:min-w-32 h-10 flexctr">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={7}
            queryx={"routfl_addfln"}
            params={inputx.routfl_addfln}
            plchdr="Route"
            repprm={iptcgh}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-1/6 min-w-20 md:min-w-32 h-10 flexctr">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"fltype_addfln"}
            params={inputx.fltype_addfln}
            plchdr="Flight Type"
            repprm={iptcgh}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-1/6 min-w-20 md:min-w-32 h-10 flexctr">
          <UixGlobalInputxFormdt
            typipt={"date"}
            length={undefined}
            queryx={"fldtef_addfln"}
            params={inputx.fldtef_addfln}
            plchdr="Flight Date"
            repprm={iptcgh}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-1/6 min-w-20 md:min-w-32 h-10 flexctr">
          <div className="afull flexctr p-1.5">
            <div className="afull btnsbm flexctr">Submit</div>
          </div>
        </div>
      </div>
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
            <div className="afull btnsbm flexctr">Upload</div>
          </div>
        </div>
        <div className="w-full md:w-1/4 min-w-20 md:min-w-32 h-10 flexctr">
          <div className="afull flexctr p-1.5">
            <div className="afull btnsbm flexctr text-center">
              <span className="leading-3">Download Template</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
