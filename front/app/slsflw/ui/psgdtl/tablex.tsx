"use client";

import {
  FncGlobalFormatCpnfmt,
  FncGlobalFormatDatefm,
  FncGlobalFormatRoutfl,
} from "@/app/global/function/fncGlobalFormat";
import {
  UixGlobalIconvcCancel,
  UixGlobalIconvcCeklis,
  UixGlobalIconvcEditdt,
} from "@/app/global/ui/server/uixGlobalIconvc";
import { useState } from "react";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import { mdlGlobalAllusrCookie } from "@/app/global/model/mdlGlobalPrimer";
import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";
import { MdlPsglstAcpedtDtbase, MdlPsglstPsgdtlFrntnd } from "@/app/psglst/model/mdlPsglstParams";
import { ApiPsglstPsgdtlUpdate } from "@/app/psglst/api/apiPsglstPsgdtl";

export default function UixSlsflwDetailTablex({
  detail,
  acpedt,
  cookie,
}: {
  detail: MdlPsglstPsgdtlFrntnd[];
  acpedt: MdlPsglstAcpedtDtbase[];
  cookie: mdlGlobalAllusrCookie;
}) {
  const [edtobj, edtobjSet] = useState<MdlPsglstPsgdtlFrntnd>();
  const [okeupd, okeupdSet] = useState<string>("");
  const [cxlupd, cxlupdSet] = useState<string>("");
  const [cxlrsp, cxlrspSet] = useState<string>("");
  const rplprm = FncGlobalParamsEdlink();
  const actedt = (e: React.ChangeEvent<HTMLInputElement>) => {
    const key = e.currentTarget.id;
    let val: string | number = e.currentTarget.value;
    if (key == "routvc") val = FncGlobalFormatRoutfl(val);
    else if (key == "cpnbvc") val = FncGlobalFormatCpnfmt(val);
    else if (["tktnbr", "flnbvc"].includes(key))
      val = val.replace(/[^0-9]/g, "");
    else if (["ntafvc", "ntaffl"].includes(key))
      val = Number(val);
    else val = val.toUpperCase();
    edtobjSet({
      ...edtobj,
      [key]: val,
    } as MdlPsglstPsgdtlFrntnd);
  };

  // Confirm update retail or series
  const update = async (log: MdlPsglstPsgdtlFrntnd) => {
    const rspupd: string = await ApiPsglstPsgdtlUpdate(log);
    edtobjSet({ ...log, prmkey: "" })
    if (rspupd == "success") {
      okeupdSet(log.prmkey);
    } else {
      cxlupdSet(log.prmkey);
      cxlrspSet(rspupd);
    }
    setTimeout(() => {
      okeupdSet("");
      cxlupdSet("");
      cxlrspSet("");
      rplprm(["update_psgdtl"], String(Math.random()));
    }, 1000);
  };

  return (
    <>
      <div className="afull max-h-fit overflow-auto rounded-lg ring-2 ring-sky-800">
        <table className="w-full">
          <thead className="sticky top-0 z-20 text-white">
            <tr>
              <th className="thhead sticky left-0">Action</th>
              {detail && detail.length > 0
                ? Object.entries(detail[0]).map(([key]) => (
                  <th key={key} className="thhead">
                    {key}
                  </th>
                ))
                : ""}
            </tr>
          </thead>
          <tbody className="text-slate-700 bg-sky-100">
            {detail.map((log, idx) => (
              <tr className="group" key={idx}>
                <td
                  className={`tdbody text-center sticky left-0 z-10 shadow-md drop-shadow-lg 
                    ${edtobj?.prmkey === log.prmkey ? "bg-sky-200" :
                      okeupd === log.prmkey ? "bg-green-400 shkeit" :
                        cxlupd === log.prmkey ? "bg-red-400 shkeit" : "bg-sky-100"}`}

                >
                  <div className="afull flexctr gap-x-1.5 relative">
                    <div
                      className={`flexctr btnsbm duration-300 cursor-pointer ${edtobj?.prmkey === log.prmkey
                        ? "opacity-100"
                        : "opacity-0 select-none pointer-events-none"
                        }`}
                      onClick={() => update(edtobj as MdlPsglstPsgdtlFrntnd)}
                    >
                      <UixGlobalIconvcCeklis
                        bold={2.5}
                        color="#53eafd"
                        size={1.4}
                      />
                    </div>
                    <div
                      className={`flexctr btnsbm duration-300 cursor-pointer ${edtobj?.prmkey === log.prmkey
                        ? "opacity-100"
                        : "opacity-0 select-none pointer-events-none"
                        }`}
                      onClick={() => edtobjSet({ ...log, prmkey: "" })}
                    >
                      <UixGlobalIconvcCancel
                        bold={2.5}
                        color="#fb2c36"
                        size={1.4}
                      />
                    </div>
                    <div
                      className={`absolute flexctr btnsbm duration-300 cursor-pointer ${edtobj?.prmkey === log.prmkey
                        ? "opacity-0 select-none pointer-events-none"
                        : "opacity-100"
                        }`}
                      onClick={() => edtobjSet({ ...log, updtby: cookie.usrnme, prmkey: log.prmkey })}
                    >
                      <UixGlobalIconvcEditdt
                        bold={2.5}
                        color="white"
                        size={1.4}
                      />
                    </div>
                  </div>
                  <div className={`${cxlupd === log.prmkey ? "h-8 flexctr font-semibold text-white" :
                    "h-0 opacity-0"} duration-300`}>{cxlrsp}</div>
                </td>
                {Object.entries(log).map(([key, val]) => (
                  <td
                    className={`tdbody text-center z-0 h-8 w-fit ${edtobj?.prmkey === log.prmkey ? "bg-sky-200" :
                      okeupd === log.prmkey ? "bg-green-400 shkeit" :
                        cxlupd === log.prmkey ? "bg-red-400 shkeit" : "bg-sky-100"}`}
                    key={key}
                  >
                    {edtobj?.prmkey === log.prmkey &&
                      acpedt.some((item) => item.params === key) ? (
                      <div className="relative flexctr">
                        <span className="invisible">
                          XXXXXXXXXXXXX{String(edtobj[key as keyof typeof edtobj])}
                        </span>
                        <div className="h-8 absolute">
                          <UixGlobalInputxFormdt
                            typipt={key == "timeis" ? "datetime-local" : key == "datevc" ? "date" : "text"}
                            length={
                              acpedt.find((item) => item.params === key)?.length
                            }
                            queryx={key.toString()}
                            params={String(edtobj[key as keyof typeof edtobj])}
                            plchdr=""
                            repprm={actedt}
                            labelx=""
                          />
                        </div>
                      </div>
                    ) : (
                      <div className="w-fit flexctr h-8">
                        {["datefl", "daterv", "datevc", "timefl", "timevc", "timerv", "timeis", "timecr", "mnthfl"].includes(key)
                          ? FncGlobalFormatDatefm(String(val))
                          : ["ntaffl", "ntafvc", "yqtxfl", "yqtxvc", "qsrcrw", "qsrcvc"].includes(key)
                            ? <div className="w-full text-right">{val.toLocaleString("en-US")}</div>
                            : val}
                      </div>
                    )}
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
