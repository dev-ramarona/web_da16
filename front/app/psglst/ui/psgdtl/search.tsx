"use client";
import { useEffect, useState } from "react";

import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";
import {
  FncGlobalFormatFilter,
  FncGlobalFormatRoutfl,
} from "@/app/global/function/fncGlobalFormat";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import { MdlPsglstPsgdtlSearch, MdlPsglstSrcprmAllprm } from "../../model/mdlPsglstParams";
import { UixGlobalIconvcRfresh } from "@/app/global/ui/server/uixGlobalIconvc";

export default function UixPsglstDetailSearch({
  trtprm,
  datefl,
}: {
  trtprm: MdlPsglstSrcprmAllprm;
  datefl: string[];
}) {
  const [params, paramsSet] = useState<MdlPsglstPsgdtlSearch>({
    update_psgdtl: trtprm.update_psgdtl || "",
    mnthfl_psgdtl: trtprm.mnthfl_psgdtl || "",
    datefl_psgdtl: trtprm.datefl_psgdtl || "",
    airlfl_psgdtl: trtprm.airlfl_psgdtl || "",
    flnbfl_psgdtl: trtprm.flnbfl_psgdtl || "",
    depart_psgdtl: trtprm.depart_psgdtl || "",
    routfl_psgdtl: trtprm.routfl_psgdtl || "",
    pnrcde_psgdtl: trtprm.pnrcde_psgdtl || "",
    tktnfl_psgdtl: trtprm.tktnfl_psgdtl || "",
    isitfl_psgdtl: trtprm.isitfl_psgdtl || "",
    isittx_psgdtl: trtprm.isittx_psgdtl || "",
    isitir_psgdtl: trtprm.isitir_psgdtl || "",
    nclear_psgdtl: trtprm.nclear_psgdtl || "",
    pagenw_psgdtl: trtprm.pagenw_psgdtl || 1,
    limitp_psgdtl: trtprm.limitp_psgdtl || 15,
  });

  // Monitor change
  const [chnged, chngedSet] = useState<boolean>(false);
  useEffect(() => {
    chngedSet(false);
    paramsSet({
      update_psgdtl: trtprm.update_psgdtl || "",
      mnthfl_psgdtl: trtprm.mnthfl_psgdtl || "",
      datefl_psgdtl: trtprm.datefl_psgdtl || "",
      airlfl_psgdtl: trtprm.airlfl_psgdtl || "",
      flnbfl_psgdtl: trtprm.flnbfl_psgdtl || "",
      depart_psgdtl: trtprm.depart_psgdtl || "",
      routfl_psgdtl: trtprm.routfl_psgdtl || "",
      pnrcde_psgdtl: trtprm.pnrcde_psgdtl || "",
      tktnfl_psgdtl: trtprm.tktnfl_psgdtl || "",
      isitfl_psgdtl: trtprm.isitfl_psgdtl || "",
      isittx_psgdtl: trtprm.isittx_psgdtl || "",
      isitir_psgdtl: trtprm.isitir_psgdtl || "",
      nclear_psgdtl: trtprm.nclear_psgdtl || "",
      pagenw_psgdtl: trtprm.pagenw_psgdtl || 1,
      limitp_psgdtl: trtprm.limitp_psgdtl || 15,
    });
  }, [trtprm]);

  // Replace params
  const rplprm = FncGlobalParamsEdlink();
  const repprm = (e: React.ChangeEvent<HTMLInputElement>) => {
    chngedSet(true);
    const namefl = e.currentTarget.id;
    let valuef = e.currentTarget.value;
    if (["isittx_psgdtl", "isitfl_psgdtl", "isitir_psgdtl"].includes(namefl))
      valuef = FncGlobalFormatFilter(valuef,
        [{ keywrd: "fl", output: "Flown" }, { keywrd: "no", output: "Not flown" }]);
    else if (namefl == "nclear_psgdtl") valuef = FncGlobalFormatFilter(valuef,
      [{ keywrd: "a", output: "ALL" },
      { keywrd: "spt", output: "SLSRPT" },
      { keywrd: "mnf", output: "MNFEST" }]);
    else if (["flnbfl_psgdtl", "tktnfl_psgdtl"].includes(namefl))
      valuef = valuef.replace(/[^0-9]/g, "");
    else if (namefl == "routfl_psgdtl") valuef = FncGlobalFormatRoutfl(valuef);
    else valuef = valuef.toUpperCase();
    paramsSet({
      ...params,
      [namefl]: valuef,
    });
    rplprm([namefl, "pagenw_psgdtl"], [valuef, ""]);
  };

  // Reset function
  const resetx = () => {
    rplprm(
      [
        "prmkey_psgdtl",
        "nclear_psgdtl",
        "mnthfl_psgdtl",
        "datefl_psgdtl",
        "airlfl_psgdtl",
        "flnbfl_psgdtl",
        "depart_psgdtl",
        "routfl_psgdtl",
        "pnrcde_psgdtl",
        "tktnfl_psgdtl",
        "isitfl_psgdtl",
        "isittx_psgdtl",
        "isitir_psgdtl",
        "pagenw_psgdtl",
      ],
      ""
    );
  };
  return (
    <div className="w-full h-20 min-h-fit pb-1.5 flexctr relative">
      <div className={`${chnged ? "w-16 h-10 translate-y-0" : "w-0 h-0 opacity-0 -translate-y-10"} z-10 absolute bg-white ring-2 ring-sky-300 px-5 py-2 rounded-xl flexctr duration-300`}>
        <div>Wait</div>
        <div className="animate-spin"><UixGlobalIconvcRfresh bold={2} color="black" size={1} /></div>
      </div>
      <div className={`afull flexctr flex-wrap gap-y-3 ${chnged ? "animate-pulse select-none" : ""} duration-300`}>
        <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"nclear_psgdtl"}
            params={params.nclear_psgdtl}
            plchdr="Not Clear"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"month"}
            length={undefined}
            queryx={"mnthfl_psgdtl"}
            params={params.mnthfl_psgdtl}
            plchdr="Flight Month"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"date"}
            length={datefl}
            queryx={"datefl_psgdtl"}
            params={params.datefl_psgdtl}
            plchdr="Flight Date"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"airlfl_psgdtl"}
            params={params.airlfl_psgdtl}
            plchdr="Airline"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"flnbfl_psgdtl"}
            params={params.flnbfl_psgdtl}
            plchdr="Flight Number"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"depart_psgdtl"}
            params={params.depart_psgdtl}
            plchdr="Departure"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"routfl_psgdtl"}
            params={params.routfl_psgdtl}
            plchdr="Route"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"pnrcde_psgdtl"}
            params={params.pnrcde_psgdtl}
            plchdr="PNR Code"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"tktnfl_psgdtl"}
            params={params.tktnfl_psgdtl}
            plchdr="Ticket Number"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"isitfl_psgdtl"}
            params={params.isitfl_psgdtl}
            plchdr="Flown Only"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"isittx_psgdtl"}
            params={params.isittx_psgdtl}
            plchdr="Transit Only"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={undefined}
            queryx={"isitir_psgdtl"}
            params={params.isitir_psgdtl}
            plchdr="Irreg Only"
            repprm={repprm}
            labelx=""
          />
        </div>
        <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
          <form className="afull p-1.5"
            method="POST"
            action={`${process.env.NEXT_PUBLIC_URL_AXIOSB}/psglst/psgdtl/getall/downld`}>
            <input type="hidden" name="data" value={JSON.stringify(params)} />
            <button type="submit" className="afull btnsbm flexctr">
              Download
            </button>
          </form>
        </div>
        <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
          <div className="afull p-1.5">
            <div className="afull btnwrn flexctr" onClick={() => resetx()}>
              Reset
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
