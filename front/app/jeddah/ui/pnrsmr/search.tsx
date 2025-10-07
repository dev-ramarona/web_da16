"use client";
import {
  FncGlobalFormatFilter,
  FncGlobalFormatRoutfl,
} from "@/app/global/function/fncGlobalFormat";
import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import {
  MdlJeddahInputxAllprm,
  MdlJeddahPnrsmrSearch,
} from "../../model/mdlJeddahMainpr";
import { useEffect, useState } from "react";
import { ApiJeddahPnrsmrDownld } from "../../api/apiJeddahPnrsmr";


export default function UixJeddahPnrsmrSearch({
  trtprm,
}: {
  trtprm: MdlJeddahInputxAllprm;
}) {
  const [params, paramsSet] = useState<MdlJeddahPnrsmrSearch>({
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
  });

  // Monitor PNR change
  useEffect(() => {
    const handler = setTimeout(() => {
      paramsSet((prev) => ({
        ...prev,
        pnrcde_pnrsmr: trtprm.pnrcde_pnrsmr || "",
      }));
    }, 500);
    return () => clearTimeout(handler);
  }, [trtprm.pnrclk_pnrsmr]);

  // Replace params
  const rplprm = FncGlobalParamsEdlink();
  const repprm = (e: React.ChangeEvent<HTMLInputElement>) => {
    const namefl = e.currentTarget.id;
    let valuef = e.currentTarget.value;
    if (namefl == "flnbfl_pnrsmr") valuef = valuef.replace(/[^0-9]/g, "");
    else if (namefl == "routfl_pnrsmr") valuef = FncGlobalFormatRoutfl(valuef);
    else if (["srtspl_pnrsmr", "srtcxl_pnrsmr"].includes(namefl))
      valuef = FncGlobalFormatFilter(valuef, [
        { keywrd: "HIG", output: "Highest" },
        { keywrd: "LOW", output: "Lowest" },
      ]);
    else if (namefl == "psdate_pnrsmr")
      valuef = FncGlobalFormatFilter(valuef, [
        { keywrd: "", output: "Hide Past Date" },
      ]);
    else valuef = valuef.toUpperCase();
    paramsSet({
      ...params,
      [namefl]: valuef,
    });
    rplprm([namefl, "pagenw_pnrsmr"], [valuef, ""]);
  };

  // Download csv summary pnr
  const [dwnrsp, dwnrspSet] = useState("Download");
  const dwnapi = async () => {
    dwnrspSet("Wait");
    const rspdwn = await ApiJeddahPnrsmrDownld(params, "downld");
    rspdwn ? dwnrspSet("Success") : dwnrspSet("Failed");
    setTimeout(() => dwnrspSet("Download"), 500);
  };

  // Reset function
  const resetx = () => {
    paramsSet({
      airlfl_pnrsmr: "",
      flnbfl_pnrsmr: "",
      routfl_pnrsmr: "",
      datefl_pnrsmr: "",
      pnrcde_pnrsmr: "",
      agtnme_pnrsmr: "",
      psdate_pnrsmr: "",
      srtspl_pnrsmr: "",
      srtcxl_pnrsmr: "",
      pagenw_pnrsmr: 1,
      limitp_pnrsmr: 15,
    });
    rplprm(
      [
        "prmkey_smrfln",
        "prmkey_pnrsmr",
        "airlfl_pnrsmr",
        "flnbfl_pnrsmr",
        "routfl_pnrsmr",
        "datefl_pnrsmr",
        "pnrcde_pnrsmr",
        "pnrclk_pnrsmr",
        "agtnme_pnrsmr",
        "srtspl_pnrsmr",
        "srtcxl_pnrsmr",
        "pagenw_pnrsmr",
        "limitp_pnrsmr",
      ],
      ""
    );
  };
  return (
    <div className="w-full h-20 min-h-fit pb-1.5 flexstr flex-wrap gap-y-3">
      <div className="w-1/3 md:w-28 h-10 flexctr relative">
        <div className="afull p-1.5">
          <div
            className="afull btnsbm flexctr"
            onClick={() => rplprm("chssmr", "flnsmr")}
          >
            <span className="leading-3 text-center">View Smr Flno</span>
          </div>
        </div>
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"date"}
          length={undefined}
          queryx={"datefl_pnrsmr"}
          params={params.datefl_pnrsmr}
          plchdr="Flight Date"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={2}
          queryx={"airlfl_pnrsmr"}
          params={params.airlfl_pnrsmr}
          plchdr="Airline"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={4}
          queryx={"flnbfl_pnrsmr"}
          params={params.flnbfl_pnrsmr}
          plchdr="Flight Number"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={7}
          queryx={"routfl_pnrsmr"}
          params={params.routfl_pnrsmr}
          plchdr="Route"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={6}
          queryx={"pnrcde_pnrsmr"}
          params={params.pnrcde_pnrsmr}
          plchdr="PNR Code"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"agtnme_pnrsmr"}
          params={params.agtnme_pnrsmr}
          plchdr="Agent Name"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"psdate_pnrsmr"}
          params={params.psdate_pnrsmr}
          plchdr="Past Date"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"srtspl_pnrsmr"}
          params={params.srtspl_pnrsmr}
          plchdr="Sort Split"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"srtcxl_pnrsmr"}
          params={params.srtcxl_pnrsmr}
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
      <div className="w-1/3 md:w-24 h-10 flexctr relative">
        <div className="afull p-1.5">
          <div className="afull btnwrn flexctr" onClick={() => resetx()}>
            Reset
          </div>
        </div>
      </div>
    </div>
  );
}
