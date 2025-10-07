"use client";
import {
  FncGlobalFormatFilter,
  FncGlobalFormatRoutfl,
} from "@/app/global/function/fncGlobalFormat";
import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import {
  MdlJeddahInputxAllprm,
  MdlJeddahFlnsmrSearch,
} from "../../model/mdlJeddahMainpr";
import { useState } from "react";
import { ApiJeddahFlnsmrDownld } from "../../api/apiJeddahFlnsmr";


export default function UixJeddahFlnsmrSearch({
  trtprm,
}: {
  trtprm: MdlJeddahInputxAllprm;
}) {
  const [params, paramsSet] = useState<MdlJeddahFlnsmrSearch>({
    airlfl_flnsmr: trtprm.airlfl_flnsmr || "",
    datefl_flnsmr: trtprm.datefl_flnsmr || "",
    flnbfl_flnsmr: trtprm.flnbfl_flnsmr || "",
    routfl_flnsmr: trtprm.routfl_flnsmr || "",
    psdate_flnsmr: trtprm.psdate_flnsmr || "",
    srtspl_flnsmr: trtprm.srtspl_flnsmr || "",
    srtcxl_flnsmr: trtprm.srtcxl_flnsmr || "",
    pagenw_flnsmr: trtprm.pagenw_flnsmr || 1,
    limitp_flnsmr: trtprm.limitp_flnsmr || 15,
  });

  // Replace params
  const rplprm = FncGlobalParamsEdlink();
  const repprm = (e: React.ChangeEvent<HTMLInputElement>) => {
    const namefl = e.currentTarget.id;
    let valuef = e.currentTarget.value;
    if (namefl == "flnbfl_flnsmr") valuef = valuef.replace(/[^0-9]/g, "");
    else if (namefl == "routfl_flnsmr") valuef = FncGlobalFormatRoutfl(valuef);
    else if (namefl == "psdate_flnsmr")
      valuef = valuef = FncGlobalFormatFilter(valuef, [
        { keywrd: "", output: "Hide Past Date" },
      ]);
    else if (["srtspl_flnsmr", "srtcxl_flnsmr"].includes(namefl))
      valuef = FncGlobalFormatFilter(valuef, [
        { keywrd: "HIG", output: "Highest" },
        { keywrd: "LOW", output: "Lowest" },
      ]);
    else valuef = valuef.toUpperCase();
    paramsSet({
      ...params,
      [namefl]: valuef,
    });
    rplprm([namefl, "pagenw_flnsmr"], [valuef, ""]);
  };

  // Download csv summary pnr
  const [dwnrsp, dwnrspSet] = useState("Download");
  const dwnapi = async () => {
    dwnrspSet("Wait");
    const rspdwn = await ApiJeddahFlnsmrDownld(params);
    rspdwn ? dwnrspSet("Success") : dwnrspSet("Failed");
    setTimeout(() => dwnrspSet("Download"), 500);
  };
  return (
    <div className="w-full h-20 min-h-fit pb-1.5 flexstr flex-wrap gap-y-3">
      <div className="w-1/3 md:w-28 h-10 flexctr relative">
        <div className="afull p-1.5">
          <div
            className="afull btnsbm flexctr"
            onClick={() => rplprm("chssmr", "")}
          >
            <span className="lead-3 text-center">View Smr PNR</span>
          </div>
        </div>
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"date"}
          length={undefined}
          queryx={"datefl_flnsmr"}
          params={params.datefl_flnsmr}
          plchdr="Flight Date"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={2}
          queryx={"airlfl_flnsmr"}
          params={params.airlfl_flnsmr}
          plchdr="Airline"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={4}
          queryx={"flnbfl_flnsmr"}
          params={params.flnbfl_flnsmr}
          plchdr="Flight Number"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={2}
          queryx={"routfl_flnsmr"}
          params={params.routfl_flnsmr}
          plchdr="Route"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={2}
          queryx={"psdate_flnsmr"}
          params={params.psdate_flnsmr}
          plchdr="Past Date"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"srtspl_flnsmr"}
          params={params.srtspl_flnsmr}
          plchdr="Sort Split"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"srtcxl_flnsmr"}
          params={params.srtcxl_flnsmr}
          plchdr="Sort Cancel"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <div className="afull p-1.5">
          <div className="afull btnsbm flexctr" onClick={() => dwnapi()}>
            {dwnrsp}
          </div>
        </div>
      </div>
    </div>
  );
}
