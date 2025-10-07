"use client";
import {
  FncGlobalFormatFilter,
  FncGlobalFormatRoutfl,
} from "@/app/global/function/fncGlobalFormat";
import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import {
  MdlJeddahInputxAllprm,
  MdlJeddahPnrdtlSearch,
} from "../../model/mdlJeddahMainpr";
import { useEffect, useState } from "react";
import { mdlGlobalAlluserFilter } from "@/app/global/model/mdlGlobalPrimer";
import { ApiJeddahPnrdtlDownld } from "../../api/apiJeddahPnrdtl";


export default function UixJeddahPnrdtlSearch({
  trtprm,
}: {
  trtprm: MdlJeddahInputxAllprm;
}) {
  const [params, paramsSet] = useState<MdlJeddahPnrdtlSearch>({
    clssfl_pnrdtl: trtprm.clssfl_pnrdtl || "",
    airlfl_pnrdtl: trtprm.airlfl_pnrdtl || "",
    datefl_pnrdtl: trtprm.datefl_pnrdtl || "",
    pnrcde_pnrdtl: trtprm.pnrcde_pnrdtl || "",
    flnbfl_pnrdtl: trtprm.flnbfl_pnrdtl || "",
    routfl_pnrdtl: trtprm.routfl_pnrdtl || "",
    agtnme_pnrdtl: trtprm.agtnme_pnrdtl || "",
    srtspl_pnrdtl: trtprm.srtspl_pnrdtl || "",
    srtcxl_pnrdtl: trtprm.srtcxl_pnrdtl || "",
    pagenw_pnrdtl: trtprm.pagenw_pnrdtl || 1,
    limitp_pnrdtl: trtprm.limitp_pnrdtl || 15,
  });

  // Monitor PNR change
  useEffect(() => {
    const handler = setTimeout(() => {
      paramsSet((prev) => ({
        ...prev,
        clssfl_pnrdtl: trtprm.clssfl_pnrdtl || "",
        airlfl_pnrdtl: trtprm.airlfl_pnrdtl || "",
        datefl_pnrdtl: trtprm.datefl_pnrdtl || "",
        pnrcde_pnrdtl: trtprm.pnrcde_pnrdtl || "",
        flnbfl_pnrdtl: trtprm.flnbfl_pnrdtl || "",
        routfl_pnrdtl: trtprm.routfl_pnrdtl || "",
        agtnme_pnrdtl: trtprm.agtnme_pnrdtl || "",
        srtspl_pnrdtl: trtprm.srtspl_pnrdtl || "",
        srtcxl_pnrdtl: trtprm.srtcxl_pnrdtl || "",
      }));
    }, 500);
    return () => clearTimeout(handler);
  }, [trtprm.pnrclk_pnrdtl]);

  // Replace params
  const rplprm = FncGlobalParamsEdlink();
  const repprm = (e: React.ChangeEvent<HTMLInputElement>) => {
    const filter: mdlGlobalAlluserFilter[] = [
      { keywrd: "HIG", output: "Highest" },
      { keywrd: "LOW", output: "Lowest" },
    ];
    const namefl = e.currentTarget.id;
    let valuef = e.currentTarget.value;
    if (namefl == "flnbfl_pnrdtl") valuef = valuef.replace(/[^0-9]/g, "");
    else if (namefl == "routfl_pnrdtl") valuef = FncGlobalFormatRoutfl(valuef);
    else if (["srtspl_pnrdtl", "srtcxl_pnrdtl"].includes(namefl))
      valuef = FncGlobalFormatFilter(valuef, filter);
    else valuef = valuef.toUpperCase();
    paramsSet({
      ...params,
      [namefl]: valuef,
    });
    setTimeout(() => rplprm([namefl, "pagenw_pnrdtl"], [valuef, ""]), 500);
  };

  // Download csv summary pnr
  const [dwnrsp, dwnrspSet] = useState("Download");
  const dwnapi = async () => {
    dwnrspSet("Wait");
    const rspdwn = await ApiJeddahPnrdtlDownld(params);
    rspdwn ? dwnrspSet("Success") : dwnrspSet("Failed");
    setTimeout(() => dwnrspSet("Download"), 500);
  };

  // Reset function
  const resetx = () => {
    paramsSet({
      clssfl_pnrdtl: "",
      airlfl_pnrdtl: "",
      datefl_pnrdtl: "",
      pnrcde_pnrdtl: "",
      flnbfl_pnrdtl: "",
      routfl_pnrdtl: "",
      agtnme_pnrdtl: "",
      srtspl_pnrdtl: "",
      srtcxl_pnrdtl: "",
      pagenw_pnrdtl: 1,
      limitp_pnrdtl: 15,
    });
    rplprm(
      [
        "prmkey_pnrsmr",
        "prmkey_pnrdtl",
        "clssfl_pnrdtl",
        "airlfl_pnrdtl",
        "datefl_pnrdtl",
        "pnrcde_pnrdtl",
        "pnrclk_pnrdtl",
        "flnbfl_pnrdtl",
        "routfl_pnrdtl",
        "agtnme_pnrdtl",
        "srtspl_pnrdtl",
        "srtcxl_pnrdtl",
        "pagenw_pnrdtl",
        "limitp_pnrdtl",
      ],
      ""
    );
  };
  return (
    <div className="w-full h-20 min-h-fit pb-1.5 flexctr flex-wrap gap-y-3">
      <div className="w-1/3 md:w-28 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"date"}
          length={undefined}
          queryx={"datefl_pnrdtl"}
          params={params.datefl_pnrdtl}
          plchdr="Flight Date"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-28 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={2}
          queryx={"airlfl_pnrdtl"}
          params={params.airlfl_pnrdtl}
          plchdr="Airline"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-28 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={4}
          queryx={"flnbfl_pnrdtl"}
          params={params.flnbfl_pnrdtl}
          plchdr="Flight Number"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-28 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={7}
          queryx={"routfl_pnrdtl"}
          params={params.routfl_pnrdtl}
          plchdr="Route"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-28 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={100}
          queryx={"pnrcde_pnrdtl"}
          params={params.pnrcde_pnrdtl}
          plchdr="PNR Code"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-28 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"agtnme_pnrdtl"}
          params={params.agtnme_pnrdtl}
          plchdr="Agent Name"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-28 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"srtspl_pnrdtl"}
          params={params.srtspl_pnrdtl}
          plchdr="Sort Split"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-28 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"srtcxl_pnrdtl"}
          params={params.srtcxl_pnrdtl}
          plchdr="Sort Cancel"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/3 md:w-28 h-10 flexctr relative">
        <div className="afull p-1.5">
          <div className="afull btnsbm flexctr" onClick={() => dwnapi()}>
            {dwnrsp}
          </div>
        </div>
      </div>
      <div className="w-1/3 md:w-28 h-10 flexctr relative">
        <div className="afull p-1.5">
          <div className="afull btnwrn flexctr" onClick={() => resetx()}>
            Reset
          </div>
        </div>
      </div>
    </div>
  );
}
