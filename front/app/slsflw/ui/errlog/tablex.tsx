'use client'
import { FncGlobalFormatDatefm } from "@/app/global/function/fncGlobalFormat";
import {
  UixGlobalIconvcIgnore,
  UixGlobalIconvcRfresh,
} from "@/app/global/ui/server/uixGlobalIconvc";
import { ApiGlobalStatusIntrvl, ApiGlobalStatusPrcess } from "@/app/global/api/apiGlobalPrimer";
import { useEffect, useState } from "react";
import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";
import { MdlPsglstErrlogDtbase } from "@/app/psglst/model/mdlPsglstParams";
import { ApiPsglstPrcessManual } from "@/app/psglst/api/apiPsglstPrcess";

export default function UixPsglstErrlogTablex({
  errlog,
}: {
  errlog: MdlPsglstErrlogDtbase[];
}) {

  // Hit the database and get interval status
  const rplprm = FncGlobalParamsEdlink();
  const [statfn, statfnSet] = useState("Done");
  const [onpkey, onpkeySet] = useState("Done");
  const [intrvl, intrvlSet] = useState<NodeJS.Timeout | null>(null);
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
      statfnSet("Wait");
      onpkeySet(params.prmkey);
      ApiPsglstPrcessManual(nowParams);
      await ApiGlobalStatusIntrvl(statfnSet, intrvlSet, "sbrapi");
    } else statfnSet(`Wait ${status.sbrapi}%`);
  };

  // refresh page
  useEffect(() => {
    if (statfn == "Process Done") setTimeout(() => {
      rplprm(["update_psgdtl"], String(Math.random()));
    }, 1000);
  }, [statfn])

  return (
    <>
      <div className="afull max-h-fit overflow-auto rounded-lg ring-2 ring-sky-800">
        <table className="w-full">
          <thead className="sticky top-0 z-10 text-white">
            <tr>
              <th className="thhead">Action</th>
              {errlog && errlog.length > 0
                ? Object.entries(errlog[0]).map(([key]) => (
                  <th key={key} className="thhead">
                    {key}
                  </th>
                ))
                : ""}
            </tr>
          </thead>
          <tbody className="text-slate-700 bg-sky-100">
            {errlog.map((log, idx) => (
              <tr className="h-8 group" key={idx}>
                <td className="tdbody text-center">
                  <div className={`${onpkey == log.prmkey && statfn != "Done" ? "" :
                    "h-0 opacity-0 select-none pointer-events-none"} duration-300`}>{statfn}</div>
                  <div className="afull flexctr gap-x-1.5">
                    <div className="w-1/2 flexctr btnsbm duration-300 cursor-pointer"
                      onClick={() => prcess(log)}>
                      <UixGlobalIconvcRfresh
                        bold={3}
                        color="#53eafd"
                        size={1.4}
                      />
                    </div>
                    <div className="w-1/2 flexctr btnsbm duration-300 cursor-pointer"
                      onClick={() => prcess({ ...log, erignr: log.prmkey })}>
                      <UixGlobalIconvcIgnore
                        bold={3}
                        color="#ffd230"
                        size={1.4}
                      />
                    </div>
                  </div>
                </td>
                {Object.entries(log).map(([key, val]) => (
                  <td className="tdbody text-center" key={key}>
                    {["datefl", "timeup"].includes(key)
                      ? FncGlobalFormatDatefm(String(val))
                      : val}
                  </td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  );
}
