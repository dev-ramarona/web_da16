"use client";
import { useEffect, useState } from "react";
import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";
import { mdlGlobalAlluserFilter } from "@/app/global/model/mdlGlobalPrimer";
import {
  FncGlobalFormatFilter,
  FncGlobalFormatRoutfl,
} from "@/app/global/function/fncGlobalFormat";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import { MdlJeddahPsgdtlSearch, MdlPsglstSrcprmAllprm } from "../../model/mdlPsglstParams";

export default function UixPsglstPsgdtlSearch({
  trtprm,
}: {
  trtprm: MdlPsglstSrcprmAllprm;
}) {
  const [params, paramsSet] = useState<MdlJeddahPsgdtlSearch>({
    mnthfl_psgdtl: trtprm.mnthfl_psgdtl || "",
    datefl_psgdtl: trtprm.datefl_psgdtl || "",
    airlfl_psgdtl: trtprm.airlfl_psgdtl || "",
    flnbfl_psgdtl: trtprm.flnbfl_psgdtl || "",
    depart_psgdtl: trtprm.depart_psgdtl || "",
    routfl_psgdtl: trtprm.routfl_psgdtl || "",
    isflwn_psgdtl: trtprm.isflwn_psgdtl || "",
    istrst_psgdtl: trtprm.istrst_psgdtl || "",
    pnrcde_psgdtl: trtprm.pnrcde_psgdtl || "",
    tktnbr_psgdtl: trtprm.tktnbr_psgdtl || "",
    istirg_psgdtl: trtprm.istirg_psgdtl || "",
  });

  // Monitor change
  useEffect(() => {
    paramsSet({
      mnthfl_psgdtl: trtprm.mnthfl_psgdtl || "",
      datefl_psgdtl: trtprm.datefl_psgdtl || "",
      airlfl_psgdtl: trtprm.airlfl_psgdtl || "",
      flnbfl_psgdtl: trtprm.flnbfl_psgdtl || "",
      depart_psgdtl: trtprm.depart_psgdtl || "",
      routfl_psgdtl: trtprm.routfl_psgdtl || "",
      isflwn_psgdtl: trtprm.isflwn_psgdtl || "",
      istrst_psgdtl: trtprm.istrst_psgdtl || "",
      pnrcde_psgdtl: trtprm.pnrcde_psgdtl || "",
      tktnbr_psgdtl: trtprm.tktnbr_psgdtl || "",
      istirg_psgdtl: trtprm.istirg_psgdtl || "",
    });
  }, [trtprm]);

  // Replace params
  const rplprm = FncGlobalParamsEdlink();
  const repprm = (e: React.ChangeEvent<HTMLInputElement>) => {
    const filter: mdlGlobalAlluserFilter[] = [{ keywrd: "", output: "True" }];
    const namefl = e.currentTarget.id;
    let valuef = e.currentTarget.value;
    if (["istrst_psgdtl", "isflwn_psgdtl", "istirg_psgdtl"].includes(namefl))
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
      <div className="w-1/2 md:w-32 h-10 flexctr relative">
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
      <div className="w-1/2 md:w-32 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"datefl_psgdtl"}
          params={params.datefl_psgdtl}
          plchdr="Flight Date"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-32 h-10 flexctr relative">
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
      <div className="w-1/2 md:w-32 h-10 flexctr relative">
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
      <div className="w-1/2 md:w-32 h-10 flexctr relative">
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
      <div className="w-1/2 md:w-32 h-10 flexctr relative">
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
      <div className="w-1/2 md:w-32 h-10 flexctr relative">
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
      <div className="w-1/2 md:w-32 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"tktnbr_psgdtl"}
          params={params.tktnbr_psgdtl}
          plchdr="Ticket Number"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-32 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"isflwn_psgdtl"}
          params={params.isflwn_psgdtl}
          plchdr="Flown Only"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-32 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"istrst_psgdtl"}
          params={params.istrst_psgdtl}
          plchdr="Transit Only"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-32 h-10 flexctr relative">
        <UixGlobalInputxFormdt
          typipt={"text"}
          length={undefined}
          queryx={"istirg_psgdtl"}
          params={params.istirg_psgdtl}
          plchdr="Irreg Only"
          repprm={repprm}
          labelx=""
        />
      </div>
      <div className="w-1/2 md:w-32 h-10 flexctr relative">
        <div className="afull p-1.5">
          <div className="afull btnsbm flexctr" onClick={() => dwnapi()}>
            {dwnrsp}
          </div>
        </div>
      </div>
      <div className="w-1/2 md:w-32 h-10 flexctr relative">
        <div className="afull p-1.5">
          <div className="afull btnwrn flexctr" onClick={() => resetx()}>
            Reset
          </div>
        </div>
      </div>
    </div>
  );
}
