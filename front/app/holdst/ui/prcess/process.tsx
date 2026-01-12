"use client";

import { ApiGlobalStatusIntrvl, ApiGlobalStatusPrcess } from "@/app/global/api/apiGlobalPrimer";
import { useEffect, useState } from "react";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import { FncGlobalParamsHminfr } from "@/app/global/function/fncGlobalParams";
import { FncGlobalFormatDatefm } from "@/app/global/function/fncGlobalFormat";
import { MdlHoldstErrlogDtbase } from "../../model/mdlHoldstParams";
import { ApiHoldstPrcessManual } from "../../api/apiHoldstPrcess";


export default function UixHoldstPrcessManual() {

  // Get status first
  const hminfr = FncGlobalParamsHminfr(2)
  const dfault: MdlHoldstErrlogDtbase = {
    depart: "", worker: 3, datefl: 1
  }
  const [nwhour, nwhourSet] = useState((Number(new Date().getHours().toString().padStart(2, '0'))))
  const [params, paramsSet] = useState<MdlHoldstErrlogDtbase>(dfault)
  const [statfn, statfnSet] = useState("Done");
  const [intrvl, intrvlSet] = useState<NodeJS.Timeout | null>(null);
  useEffect(() => {
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
  const prcess = async (params: MdlHoldstErrlogDtbase) => {
    const status = await ApiGlobalStatusPrcess();
    let nowParams = { ...params };
    if (status.sbrapi == 0) {
      statfnSet("Wait");
      ApiHoldstPrcessManual(nowParams);
      await ApiGlobalStatusIntrvl(statfnSet, intrvlSet, "sbrapi");
    } else statfnSet(`Wait ${status.sbrapi}%`);
  };

  return (
    <div className="afull flexctr flex-col">
      <div className="h-1/3 w-full flexstr py-1.5">
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
      </div>
      <div className="h-1/2 md:h-2/3 w-full flexstr">
        {hminfr.map((val, idx) => (
          <div className="afull max-w-64 px-1.5 flexctr" key={idx}>
            <button
              className={`w-full h-full md:h-1/2 flexctr ${nwhour > 11 && idx == hminfr.length - 1 ? "btnoff select-none pointer-events-none" : "btnsbm"}`}
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
