"use client";

import { ApiGlobalStatusIntrvl, ApiGlobalStatusPrcess } from "@/app/global/api/apiGlobalPrimer";
import { useEffect, useState } from "react";
import { MdlPsglstErrlogDtbase } from "../../model/mdlPsglstParams";
import { ApiPsglstPrcessManual } from "../../api/apiPsglstPrcess";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import { FncGlobalParamsEdlink, FncGlobalParamsHminfr } from "@/app/global/function/fncGlobalParams";
import { FncGlobalFormatDatefm } from "@/app/global/function/fncGlobalFormat";
import { mdlGlobalAllusrCookie } from "@/app/global/model/mdlGlobalPrimer";


export default function UixPsglstPrcessManual({ cookie }: { cookie: mdlGlobalAllusrCookie }) {

  // Get status first
  const rplprm = FncGlobalParamsEdlink();
  const hminfr = FncGlobalParamsHminfr(4)
  const dfault: MdlPsglstErrlogDtbase = {
    prmkey: "", erstat: "", erpart: "",
    ersrce: "", erdtil: "", erdvsn: "",
    erignr: "", dateup: 0, timeup: 0,
    datefl: 0, airlfl: "", depart: "",
    flnbfl: "", Paxdif: "", flstat: "",
    flhour: 0, routfl: "", updtby: "", worker: 1,
  }
  const [nwhour, nwhourSet] = useState((Number(new Date().getHours().toString().padStart(2, '0'))))
  const [params, paramsSet] = useState<MdlPsglstErrlogDtbase>(dfault)
  const [statfn, statfnSet] = useState("Done");
  const [intrvl, intrvlSet] = useState<NodeJS.Timeout | null>(null);
  useEffect(() => {
    console.log(cookie);
    const gtstat = async () => {
      const status = await ApiGlobalStatusPrcess();
      statfnSet(status.sbrapi == 0 ? "Done" : `Wait ${status.sbrapi}%`);
      if (status.sbrapi != 0) {
        await ApiGlobalStatusIntrvl(statfnSet, intrvlSet, "sbrapi");
      } else statfnSet("Done");
    };
    gtstat();
  }, []);

  // Edit parameter
  const onchge = (e: React.ChangeEvent<HTMLInputElement>) => {
    const namefl = e.currentTarget.id;
    let valuef = e.currentTarget.value.toUpperCase();
    if (namefl == "flnbfl") valuef = valuef.replace(/[^0-9]/g, "");
    paramsSet({
      ...params,
      [namefl]: valuef,
    });
  }

  // Hit the database and get interval status
  const prcess = async (params: MdlPsglstErrlogDtbase) => {
    const status = await ApiGlobalStatusPrcess();
    let nowParams = { ...params };
    if (status.sbrapi == 0) {
      if (params.flnbfl == "") {
        nowParams.worker = 3;
        if (params.depart == "") {
          nowParams.worker = 5;
          if (params.airlfl == "")
            nowParams.worker = 8;
        }
      }

      // Cek is admin or not
      if ((cookie.keywrd && (cookie.keywrd).includes("psglst")) || nowParams.worker == 1) {
        statfnSet("Wait");
        ApiPsglstPrcessManual(nowParams);
        await ApiGlobalStatusIntrvl(statfnSet, intrvlSet, "sbrapi");
      } else {
        statfnSet("Only admin can process ALL, Please process spesific flight only");
        return setTimeout(() => statfnSet("Done"), 2000);
      }
    } else statfnSet(`Wait ${status.sbrapi}%`);
  };

  // refresh page
  useEffect(() => {
    if (statfn == "Process Done") setTimeout(() => {
      rplprm(["update_psgdtl"], String(Math.random()));
    }, 1000);
  }, [statfn])

  return (
    <div className="afull flexctr flex-col">
      <div className="h-1/3 w-full flexstr py-1.5">
        <div className="afull max-w-60 px-1.5 flexctr">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={2}
            queryx={"airlfl"}
            params={params.airlfl}
            plchdr="Airline"
            repprm={onchge}
            labelx=""
          />
        </div>
        <div className="afull max-w-60 px-1.5 flexctr">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={3}
            queryx={"depart"}
            params={params.depart}
            plchdr="Departure"
            repprm={onchge}
            labelx=""
          />
        </div>
        <div className="afull max-w-60 px-1.5 flexctr">
          <UixGlobalInputxFormdt
            typipt={"text"}
            length={4}
            queryx={"flnbfl"}
            params={params.flnbfl}
            plchdr="Flight Number"
            repprm={onchge}
            labelx=""
          />
        </div>
      </div>
      <div className="h-1/2 md:h-2/3 w-full flexstr">
        {hminfr.map((val, idx) => (
          <div className="afull max-w-64 px-1.5 flexctr" key={idx}>
            <button
              className={`w-full h-full md:h-1/2 flexctr 
                ${nwhour > 11 && idx == hminfr.length - 1 ? "btnoff select-none pointer-events-none" : "btnsbm"} 
                ${statfn.includes("admin") ? "shkeit btncxl" : ""}`}
              onClick={() => prcess({ ...params, datefl: val })}
            >
              {statfn == "Done" ? `Process Manual ${FncGlobalFormatDatefm(String(val))}` : statfn}
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}
