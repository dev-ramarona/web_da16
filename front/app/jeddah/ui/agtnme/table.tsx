"use client";
import {
  UixGlobalIconvcCancel,
  UixGlobalIconvcCeklis,
  UixGlobalIconvcCopydt,
} from "@/app/global/ui/server/uixGlobalIconvc";
import React, { useEffect, useRef, useState } from "react";
import { MdlJeddahParamsAgtedt } from "../../model/mdlJeddahParams";
import {
  ApiJeddahAgtnmeAgtupd,
  ApiJeddahAgtnmeNulsrc,
} from "../../api/apiJeddahAgtnme";
import {
  mdlGlobalAllusrCookie,
  mdlGlobalAlluserFilter,
} from "@/app/global/model/mdlGlobalAllusr";
import { FncGlobalFormatFilter } from "@/app/global/function/fncGlobalFormat";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";

export default function UixJeddahAgtnmeTablex({
  cookie,
  agtnul,
}: {
  agtnul: MdlJeddahParamsAgtedt[];
  cookie: mdlGlobalAllusrCookie;
}) {
  // Monitoring agent detail and search
  const defagt: MdlJeddahParamsAgtedt = {
    prmkey: "",
    airlfl: "",
    agtnme: "",
    agtdtl: "",
    newdtl: "",
    agtidn: "",
    newidn: "",
    rtlsrs: "",
    updtby: cookie.stfnme || "Unknown User",
  };
  const [agtnme, agtnmeSet] = useState(defagt);

  // Action Change agent name
  const chgagt = (e: React.ChangeEvent<HTMLInputElement>) => {
    const filter: mdlGlobalAlluserFilter[] = [
      { keywrd: "RTL", output: "Retail" },
      { keywrd: "SRS", output: "Series" },
      { keywrd: "NJD", output: "Non Jeddah" },
    ];
    const nameid = e.currentTarget.id.split("|");
    const prmkey = nameid[1];
    let valuef = e.currentTarget.value;
    if (nameid[0] == "rtlsrs") valuef = FncGlobalFormatFilter(valuef, filter);
    else valuef = valuef.toUpperCase();
    prmkey != agtnme.prmkey ? agtnmeSet(defagt) : "";
    agtnmeSet((prv) => ({ ...prv, prmkey: prmkey, [nameid[0]]: valuef }));
  };

  // Monitoring agent detail and search
  useEffect(() => {
    const timeout = setTimeout(() => {
      const newidn = agtnme.newidn.trim();
      const newdtl = agtnme.newdtl.trim();
      if (newidn.length === 0 && newdtl.length === 0)
        return agtnmeSet((prev) => ({
          ...prev,
          agtdtl: "",
          agtidn: "",
          rtlsrs: "",
        }));
      (async () => {
        const agtsrc = await ApiJeddahAgtnmeNulsrc(newidn, newdtl);
        agtnmeSet((prev) => ({
          ...prev,
          agtidn: agtsrc.agtidn,
          agtdtl: agtsrc.agtdtl,
          rtlsrs: agtsrc.rtlsrs,
        }));
      })();
    }, 500);
    return () => clearTimeout(timeout);
  }, [agtnme.newidn, agtnme.newdtl]);

  useEffect(() => {
    console.log(agtnme);
  }, [agtnme]);

  // Focus confirm
  const [confrm, confrmSet] = useState(false);
  const btnrf1 = useRef<HTMLButtonElement>(null);
  const btncfm = (nowagt: string) => {
    if (agtnme.newidn != "" && agtnme.newdtl != "") {
      agtnmeSet((prv) => ({ ...prv, agtnme: nowagt }));
      confrmSet(true);
      if (btnrf1.current) btnrf1.current.focus();
    }
  };

  // Submit update data
  const btnrf2 = useRef<HTMLButtonElement>(null);
  const [rspupd, rspupdSet] = useState<string>("");
  const rplprm = FncGlobalParamsEdlink();
  const update = async (action: boolean) => {
    if (btnrf2.current) btnrf2.current.focus();
    confrmSet(false);
    if (!action) return;
    const rspnow = await ApiJeddahAgtnmeAgtupd(agtnme);
    if (rspnow != "failed") rplprm("", "")
    rspupdSet(rspnow);
    setTimeout(() => rspupdSet(""), 1000);
  };

  return (
    <>
      {/* Confirm update data */}
      <div
        className={`flexctr absolute z-20 pt-10 pb-2 ${confrm
          ? "w-72 md:w-96 h-56 opacity-100"
          : "w-72 h-0 opacity-0 pointer-events-none"
          } text-slate-600 duration-300`}
      >
        <div className="afull flexctr flex-col bg-white rounded-lg shadow-md ring-2 ring-sky-800 overflow-hidden font-semibold">
          <div className="text-base whitespace-nowrap">
            Are you sure to Update this data?
          </div>
          <div className="whitespace-nowrap text-sky-800">{agtnme.agtnme}</div>
          <div className="font-normal">as</div>
          <div className="whitespace-nowrap text-sky-800">{agtnme.newidn}</div>
          <div className="whitespace-nowrap text-sky-800">{agtnme.newdtl}</div>
          <div className="whitespace-nowrap text-sky-800">{agtnme.rtlsrs}</div>
          <div
            className={`whitespace-nowrap text-red-800 pt-1.5 ${agtnme.newidn == agtnme.agtidn || !agtnme.agtidn ? "hidden" : ""
              }`}
          >
            *Data input Different from Suggest
          </div>
          <div className="flexctr gap-5 pt-3">
            <button
              className="w-10 h-6 btnsbm flexctr"
              onClick={() => update(true)}
              ref={btnrf1}
            >
              <UixGlobalIconvcCeklis color="white" size={1.3} bold={4} />
            </button>
            <button
              className="w-10 h-6 btncxl flexctr"
              onClick={() => {
                update(false);
              }}
            >
              <UixGlobalIconvcCancel color="white" size={1.3} bold={4} />
            </button>
          </div>
        </div>
      </div>
      <div className="afull max-h-fit overflow-auto rounded-lg ring-2 ring-sky-800">
        <table className="w-full">
          <thead className="sticky top-0 z-10 text-white">
            <tr>
              <th className="thhead">Airline</th>
              <th className="thhead">Agent Name</th>
              <th className="thhead">Agent ID</th>
              <th className="thhead">Agent Name Detail</th>
              <th className="thhead">Retail or Series</th>
              <th className="thhead">Suggest Agent</th>
              <th className="thhead">Action</th>
              <th className="thhead">Updated By</th>
            </tr>
          </thead>
          <tbody className="text-slate-700 bg-sky-100">
            {agtnul.map((log, idx) => (
              <tr
                className={`h-8 group ${agtnme.agtnme == log.agtnme && rspupd != ""
                  ? rspupd == "failed"
                    ? "bg-red-300 shkeit"
                    : "bg-green-300 shkeit"
                  : ""
                  } duration-300`}
                key={idx}
              >
                <td className="tdbody text-center">{log.airlfl}</td>
                <td className="tdbody text-left">{log.agtnme}</td>
                <td className="tdbody text-center">
                  <div
                    style={{
                      height: `2rem`,
                      minWidth: `${"New ID Agent".length * 0.53}rem`,
                      width: `${(agtnme.newidn || "New ID Agent").length * 0.58
                        }rem`,
                    }}
                  >
                    <UixGlobalInputxFormdt
                      typipt={"text"}
                      length={undefined}
                      queryx={`newidn|${log.prmkey}`}
                      params={
                        log.prmkey == agtnme.prmkey
                          ? agtnme.newidn || ""
                          : log.agtidn
                      }
                      plchdr="New ID Agent"
                      repprm={chgagt}
                      labelx="hidden"
                    />
                  </div>
                </td>
                <td className="tdbody text-center">
                  <div
                    style={{
                      height: `2rem`,
                      minWidth: `${"New Detail Agent Name".length * 0.53}rem`,
                      width: `${(agtnme.newdtl || "New Detail Agent Name").length * 0.58
                        }rem`,
                    }}
                  >
                    <UixGlobalInputxFormdt
                      typipt={"text"}
                      length={undefined}
                      queryx={`newdtl|${log.prmkey}`}
                      params={
                        log.prmkey == agtnme.prmkey
                          ? agtnme.newdtl || ""
                          : log.agtdtl
                      }
                      plchdr="New Detail Agent Name"
                      repprm={chgagt}
                      labelx="hidden"
                    />
                  </div>
                </td>
                <td className="tdbody text-center">
                  <div
                    style={{
                      height: `2rem`,
                      minWidth: `${"RTL/SRS/NON".length * 0.53}rem`,
                      width: `${(agtnme.rtlsrs || "RTL/SRS/NON").length * 0.58
                        }rem`,
                    }}
                  >
                    <UixGlobalInputxFormdt
                      typipt={"text"}
                      length={undefined}
                      queryx={`rtlsrs|${log.prmkey}`}
                      params={log.prmkey == agtnme.prmkey
                        ? agtnme.rtlsrs || log.rtlsrs : log.rtlsrs}
                      plchdr="RTL/SRS/NON"
                      repprm={chgagt}
                      labelx="hidden"
                    />
                  </div>
                </td>

                <td className="tdbody">
                  {<UixJeddahAgtnmeCopydt log={log} agtnme={agtnme} agtnmeSet={agtnmeSet} />}
                </td>

                <td className="tdbody">
                  <UixJeddahAgtnmeCeklis log={log} agtnme={agtnme} btncfm={btncfm} />
                </td>
                <td className="tdbody">
                  <div className={`afull flexctr rounded-md `}>
                    {log.prmkey == agtnme.prmkey
                      ? cookie.stfnme || log.updtby
                      : ""}
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  );
}

export function UixJeddahAgtnmeCopydt({ log, agtnme, agtnmeSet }: {
  log: MdlJeddahParamsAgtedt,
  agtnme: MdlJeddahParamsAgtedt,
  agtnmeSet: React.Dispatch<React.SetStateAction<MdlJeddahParamsAgtedt>>
}
) {
  const mtcPrmkey = log.prmkey == agtnme.prmkey
  const nulAgtidn = agtnme.agtidn == ""
  const ntfAgtidn = agtnme.agtidn == "Not Found"
  return (
    <div className="flexbtw">
      {mtcPrmkey ? (
        <div className="text-[0.6rem] pr-1.5">
          <div>{agtnme.agtidn || ""}</div>
          <div>{agtnme.agtdtl || ""}</div>
        </div>
      ) : (
        <div></div>
      )}
      <button
        className={`py-0.5 ${mtcPrmkey && !nulAgtidn && !ntfAgtidn
          ? "btnsbm" : "btnoff"}`}
        disabled={!mtcPrmkey || nulAgtidn || ntfAgtidn}
        onClick={() =>
          agtnmeSet((prev) => ({
            ...prev,
            newidn: agtnme.agtidn,
            newdtl: agtnme.agtdtl,
          }))
        }
      >
        <UixGlobalIconvcCopydt
          color="white"
          size={1.3}
          bold={2.5}
        />
      </button>
    </div>
  )
}

export function UixJeddahAgtnmeCeklis({ log, agtnme, btncfm }: {
  log: MdlJeddahParamsAgtedt,
  agtnme: MdlJeddahParamsAgtedt,
  btncfm: (agtnme: string) => void
}) {
  const mtcPrmkey = log.prmkey == agtnme.prmkey
  const nulAgtidn = agtnme.newidn == ""
  const nulAgtdtl = agtnme.newdtl == ""
  const nulRtlsrs = agtnme.rtlsrs == ""
  return (
    <div className="flexctr">
      <button
        className={`py-0.5 ${mtcPrmkey && !nulAgtidn
          && !nulAgtdtl && !nulRtlsrs ? "btnsbm" : "btnoff"}`}
        disabled={!mtcPrmkey || nulAgtidn
          || nulAgtdtl || nulRtlsrs}
        onClick={() => btncfm(log.agtnme)}
      >
        <UixGlobalIconvcCeklis
          color="white"
          size={1.3}
          bold={4}
        />
      </button>
    </div>
  )
}