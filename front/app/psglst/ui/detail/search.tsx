"use client";
import { useEffect, useState } from "react";
import {
  MdlPsglstAllprmSrcprm,
  MdlPsglstDetailSrcprm,
} from "../../model/mdlPsglstParams";
import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";
import { mdlGlobalAlluserFilter } from "@/app/global/model/mdlGlobalPrimer";
import {
  FncGlobalFormatFilter,
  FncGlobalFormatRoutfl,
} from "@/app/global/function/fncGlobalFormat";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";

export default function UixPsglstDetailSearch({
  trtprm,
}: {
  trtprm: MdlPsglstAllprmSrcprm;
}) {
  const [params, paramsSet] = useState<MdlPsglstDetailSrcprm>({
    mnthfl_detail: trtprm.mnthfl_detail || "",
    datefl_detail: trtprm.datefl_detail || "",
    airlfl_detail: trtprm.airlfl_detail || "",
    flnbfl_detail: trtprm.flnbfl_detail || "",
    depart_detail: trtprm.depart_detail || "",
    routfl_detail: trtprm.routfl_detail || "",
    isflwn_detail: trtprm.isflwn_detail || "",
    istrst_detail: trtprm.istrst_detail || "",
    pnrcde_detail: trtprm.pnrcde_detail || "",
    tktnbr_detail: trtprm.tktnbr_detail || "",
    istirg_detail: trtprm.istirg_detail || "",
  });

  // Monitor change
  useEffect(() => {
    paramsSet({
      mnthfl_detail: trtprm.mnthfl_detail || "",
      datefl_detail: trtprm.datefl_detail || "",
      airlfl_detail: trtprm.airlfl_detail || "",
      flnbfl_detail: trtprm.flnbfl_detail || "",
      depart_detail: trtprm.depart_detail || "",
      routfl_detail: trtprm.routfl_detail || "",
      isflwn_detail: trtprm.isflwn_detail || "",
      istrst_detail: trtprm.istrst_detail || "",
      pnrcde_detail: trtprm.pnrcde_detail || "",
      tktnbr_detail: trtprm.tktnbr_detail || "",
      istirg_detail: trtprm.istirg_detail || "",
    });
  }, [trtprm]);

  // Replace params
  const rplprm = FncGlobalParamsEdlink();
  const repprm = (e: React.ChangeEvent<HTMLInputElement>) => {
    const filter: mdlGlobalAlluserFilter[] = [{ keywrd: "", output: "True" }];
    const namefl = e.currentTarget.id;
    let valuef = e.currentTarget.value;
    if (["istrst_detail", "isflwn_detail", "istirg_detail"].includes(namefl))
      valuef = FncGlobalFormatFilter(valuef, filter);
    else if (namefl == "flnbfl_pnrsmr") valuef = valuef.replace(/[^0-9]/g, "");
    else if (namefl == "routfl_pnrsmr") valuef = FncGlobalFormatRoutfl(valuef);
    else if (namefl == "psdate_pnrsmr")
      valuef = valuef = FncGlobalFormatFilter(valuef, [
        { keywrd: "", output: "Show Past Date" },
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
    // const rspdwn = await ApiPsglstPnrsmrDownld(params);
    // rspdwn ? dwnrspSet("Success") : dwnrspSet("Failed");
    setTimeout(() => dwnrspSet("Download"), 500);
  };

  // Reset function
  const resetx = () => {
    rplprm(
      [
        "prmkey_smrfln",
        "prmkey_pnrsmr",
        "airlfl_pnrsmr",
        "flnbfl_pnrsmr",
        "routfl_pnrsmr",
        "datefl_pnrsmr",
        "pnrcde_pnrsmr",
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
    <div className="w-full h-20 min-h-fit pb-1.5 flexctr flex-wrap gap-y-3">
      <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"month"}
          length={undefined}
          queryx={"mnthfl_detail"}
          params={params.mnthfl_detail}
          plchdr="Flight Month"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"date"}
          length={undefined}
          queryx={"datefl_detail"}
          params={params.datefl_detail}
          plchdr="Flight Date"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"airlfl_detail"}
          params={params.airlfl_detail}
          plchdr="Airline"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"flnbfl_detail"}
          params={params.flnbfl_detail}
          plchdr="Flight Number"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"depart_detail"}
          params={params.depart_detail}
          plchdr="Departure"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"routfl_detail"}
          params={params.routfl_detail}
          plchdr="Route"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"pnrcde_detail"}
          params={params.pnrcde_detail}
          plchdr="PNR Code"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"tktnbr_detail"}
          params={params.tktnbr_detail}
          plchdr="Ticket Number"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"isflwn_detail"}
          params={params.isflwn_detail}
          plchdr="Flown Only"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"istrst_detail"}
          params={params.istrst_detail}
          plchdr="Transit Only"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"istirg_detail"}
          params={params.istirg_detail}
          plchdr="Irreg Only"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
        <div className="afull p-1.5">
          <div className="afull btnsbm flexctr" onClick={() => dwnapi()}>
            {dwnrsp}
          </div>
        </div>
      </div>
      <div className="w-1/2 md:w-[6.5rem] h-10 flexctr relative">
        <div className="afull p-1.5">
          <div className="afull btnwrn flexctr" onClick={() => resetx()}>
            Reset
          </div>
        </div>
      </div>
    </div>
  );
}
