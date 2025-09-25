"use client";
import {
  FncGlobalFormatFilter,
  FncGlobalFormatRoutfl,
} from "@/app/global/function/fncGlobalFormat";
import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import {
  MdlJeddahInputxAllpnr,
  MdlJeddahSearchAgtnme,
} from "../../model/mdlJeddahParams";
import { useEffect, useState } from "react";

export default function UixJeddahAgtnmerSearch({
  trtprm,
}: {
  trtprm: MdlJeddahInputxAllpnr;
}) {
  const [params, paramsSet] = useState<MdlJeddahSearchAgtnme>({
    airlfl_agtnme: trtprm.airlfl_agtnme || "",
    agtnme_agtnme: trtprm.agtnme_agtnme || "",
    srtnul_agtnme: trtprm.srtnul_agtnme || "Null",
    pagenw_agtnme: trtprm.pagenw_agtnme || 1,
    limitp_agtnme: trtprm.limitp_agtnme || 15,
  });

  // Monitor change
  useEffect(() => {
    paramsSet({
      airlfl_agtnme: trtprm.airlfl_agtnme,
      agtnme_agtnme: trtprm.agtnme_agtnme,
      srtnul_agtnme: trtprm.srtnul_agtnme,
      pagenw_agtnme: trtprm.pagenw_agtnme,
      limitp_agtnme: trtprm.limitp_agtnme,
    });
  }, []);

  // Replace params
  const rplprm = FncGlobalParamsEdlink();
  const repprm = (e: React.ChangeEvent<HTMLInputElement>) => {
    const namefl = e.currentTarget.id;
    let valuef = e.currentTarget.value;
    if (namefl == "srtnul_agtnme")
      valuef = valuef = FncGlobalFormatFilter(valuef, [
        { keywrd: "", output: "All" },
      ]);
    else valuef = valuef.toUpperCase();
    paramsSet({
      ...params,
      [namefl]: valuef,
    });
    rplprm([namefl, "pagenw_agtnme"], [valuef, ""]);
  };

  // Reset function
  const resetx = () => {
    paramsSet({
      airlfl_agtnme: "",
      agtnme_agtnme: "",
      srtnul_agtnme: "",
      pagenw_agtnme: 1,
      limitp_agtnme: 15,
    });
    rplprm(
      [
        "airlfl_agtnme",
        "agtnme_agtnme",
        "srtnul_agtnme",
        "pagenw_agtnme",
        "limitp_agtnme",
      ],
      ""
    );
  };
  return (
    <div className="w-full h-10 min-h-fit pb-1.5 flexstr flex-wrap gap-y-3">
      <div className="w-1/2 md:w-36 h-9 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={20}
          queryx={"airlfl_agtnme"}
          params={params.airlfl_agtnme}
          plchdr="Airline"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-36 h-9 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={20}
          queryx={"agtnme_agtnme"}
          params={params.agtnme_agtnme}
          plchdr="Agent Name"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-36 h-9 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={20}
          queryx={"srtnul_agtnme"}
          params={params.srtnul_agtnme}
          plchdr="Sort All Agent"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-36 h-9 flexctr relative">
        <div className="afull p-1.5">
          <div className="afull btnwrn flexctr" onClick={() => resetx()}>
            Reset
          </div>
        </div>
      </div>
    </div>
  );
}
