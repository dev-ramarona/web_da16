"use client";

import {
  FncGlobalFormatCpnfmt,
  FncGlobalFormatDatefm,
  FncGlobalFormatRoutfl,
} from "@/app/global/function/fncGlobalFormat";
import {
  MdlPsglstAcpedtDtbase,
  MdlPsglstPsgdtlFrntnd,
} from "../../model/mdlPsglstParams";
import {
  UixGlobalIconvcCancel,
  UixGlobalIconvcCeklis,
  UixGlobalIconvcEditdt,
} from "@/app/global/ui/server/uixGlobalIconvc";
import { useState } from "react";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";

export default function UixPsglstPsgdtlTablex({
  psgdtl,
  edtprm,
}: {
  psgdtl: MdlPsglstPsgdtlFrntnd[];
  edtprm: MdlPsglstAcpedtDtbase[];
}) {
  console.log(edtprm);

  const [edtobj, edtobjSet] = useState<MdlPsglstPsgdtlFrntnd>();
  const actedt = (e: React.ChangeEvent<HTMLInputElement>) => {
    const key = e.currentTarget.id;
    let val = e.currentTarget.value;
    if (key == "routvc") val = FncGlobalFormatRoutfl(val);
    if (key == "cpnbvc") val = FncGlobalFormatCpnfmt(val);
    else if (["tktnbr", "flnbvc"].includes(key))
      val = val.replace(/[^0-9]/g, "");
    edtobjSet({
      ...edtobj,
      [key]: val,
    } as MdlPsglstPsgdtlFrntnd);
  };

  return (
    <>
      <div className="afull max-h-fit overflow-auto rounded-lg ring-2 ring-sky-800">
        <table className="w-full">
          <thead className="sticky top-0 z-20 text-white">
            <tr>
              <th className="thhead sticky left-0">Action</th>
              {psgdtl && psgdtl.length > 0
                ? Object.entries(psgdtl[0]).map(([key]) => (
                  <th key={key} className="thhead">
                    {key}
                  </th>
                ))
                : ""}
            </tr>
          </thead>
          <tbody className="text-slate-700 bg-sky-100">
            {psgdtl.map((log, idx) => (
              <tr className="group" key={idx}>
                <td
                  className={`tdbody text-center sticky left-0 z-10 shadow-md ${edtobj?.prmkey === log.prmkey ? "bg-sky-200" : "bg-sky-100"
                    }`}
                >
                  <div className="afull flexctr gap-x-1.5 relative">
                    <div
                      className={`flexctr btnsbm duration-300 cursor-pointer ${edtobj?.prmkey === log.prmkey
                        ? "opacity-100"
                        : "opacity-0 select-none pointer-events-none"
                        }`}
                      onClick={() => edtobjSet({ ...log, prmkey: "" })}
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
                      onClick={() => edtobjSet(log)}
                    >
                      <UixGlobalIconvcEditdt
                        bold={2.5}
                        color="white"
                        size={1.4}
                      />
                    </div>
                  </div>
                </td>
                {Object.entries(log).map(([key, val]) => (
                  <td
                    className={`tdbody text-center z-0 h-8 ${edtobj?.prmkey === log.prmkey ? "bg-sky-200" : ""
                      }`}
                    key={key}
                  >
                    {edtobj?.prmkey === log.prmkey &&
                      edtprm.some((item) => item.params === key) ? (
                      <div className="relative flexctr">
                        <span className="invisible">
                          XXXX{String(edtobj[key as keyof typeof edtobj])}
                        </span>
                        <div className="h-8 absolute">
                          <UixGlobalInputxFormdt
                            typipt="text"
                            length={
                              edtprm.find((item) => item.params === key)?.length
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
                      <div className="flexctr h-8">
                        {["datefl", "timeup"].includes(key)
                          ? FncGlobalFormatDatefm(String(val))
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
