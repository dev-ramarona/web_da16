"use client";
import {
  FncGlobalFormatArrcpn,
  FncGlobalFormatDatefm,
  FncGlobalFormatFilter,
} from "@/app/global/function/fncGlobalFormat";
import { MdlJeddahPnrsmrDtbase } from "../../model/mdlJeddahMainpr";
import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";
import {
  UixGlobalIconvcCancel,
  UixGlobalIconvcCeklis,
  UixGlobalIconvcEditdt,
} from "@/app/global/ui/server/uixGlobalIconvc";
import { useState } from "react";
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import { mdlGlobalAllusrCookie } from "@/app/global/model/mdlGlobalPrimer";
import { ApiJeddahRtlsrsUpdate } from "../../api/apiJeddahRtlsrs";

export default function UixJeddahPnrsmrTablex({
  pnrsmr,
  pnrcde,
  cookie,
}: {
  pnrsmr: MdlJeddahPnrsmrDtbase[];
  pnrcde: string;
  cookie: mdlGlobalAllusrCookie;
}) {
  const rplprm = FncGlobalParamsEdlink();

  // Focus on this PNR
  const pnrnow = (pnr: string) => {
    rplprm(
      ["pnrcde_pnrdtl", "pagenw_pnrdtl", "pnrclk_pnrdtl"],
      [pnr, "1", String(Math.random())]
    );
  };

  // Filter all PNR arr split
  const pnrspl = (val: string, str: string) => {
    const arr = val.split("|");
    const pnr = [str];
    console.log(arr, pnr);

    if (pnrcde != "" && !pnrcde.includes(str)) pnr.push(pnrcde);
    for (let i = 0; i < arr.length; i++) pnr.push(arr[i].substring(0, 6));
    rplprm(
      ["pnrcde_pnrsmr", "pnrclk_pnrsmr", "pagenw_pnrsmr"],
      [pnr.join("|"), String(Math.random()), ""]
    );
  };

  // Action edit PNR Summary Retail or series
  const [rjcedt, rjcedtSet] = useState<string>("");
  const actedt = (prm: MdlJeddahPnrsmrDtbase) => {
    if (prm.rtlsrs == "") {
      rjcedtSet(prm.prmkey);
      return setTimeout(() => rjcedtSet(""), 500);
    } else iptobjSet(prm);
  };

  // Edit Input PNR Summary Retail or series
  const [iptobj, iptobjSet] = useState<MdlJeddahPnrsmrDtbase>();
  const actipt = (
    e: React.ChangeEvent<HTMLInputElement>,
    pnrsmr: MdlJeddahPnrsmrDtbase
  ) => {
    let val = e.currentTarget.value;
    const valuef = FncGlobalFormatFilter(val, [
      { keywrd: "RTL", output: "Retail" },
      { keywrd: "SRS", output: "Series" },
      { keywrd: "NOJD", output: "-" },
    ]);
    iptobjSet({
      ...pnrsmr,
      rtlsrs: valuef,
      notedt: `Updated ${valuef} by ${cookie.stfnme}`,
    } as MdlJeddahPnrsmrDtbase);
  };

  // Confirm update retail or series
  const update = async (log: MdlJeddahPnrsmrDtbase) => {
    if (iptobj) {
      await ApiJeddahRtlsrsUpdate(iptobj);
      iptobjSet({ ...log, prmkey: "" });
      rplprm(["pnrclk_pnrsmr", "pnrclk_pnrdtl"], String(Math.random()));
    }
  };

  return (
    <>
      <div className="afull max-h-fit overflow-auto rounded-lg ring-2 ring-sky-800">
        <table className="w-full">
          <thead className="sticky top-0 z-10 text-white">
            <tr>
              {pnrsmr && pnrsmr.length > 0 ? (
                Object.entries(pnrsmr[0]).map(([key]) =>
                  key != "prmkey" ? (
                    <th key={key} className="thhead">
                      {key}
                    </th>
                  ) : (
                    ""
                  )
                )
              ) : (
                <th className="thhead">Empty</th>
              )}
            </tr>
          </thead>
          <tbody className="text-slate-700 bg-sky-100">
            {pnrsmr && pnrsmr.length > 0 ? (
              pnrsmr.map((log, idx) => (
                <tr className="h-8 group" key={idx}>
                  {Object.entries(log).map(([key, val]) =>
                    key != "prmkey" ? (
                      <td className="tdbody text-center" key={key}>
                        {[
                          "timest",
                          "timend",
                          "dateup",
                          "timeup",
                          "timecr",
                        ].includes(key) ? (
                          FncGlobalFormatDatefm(String(val))
                        ) : key == "arrspl" ? (
                          <div
                            onClick={() => pnrspl(String(val), log.pnrcde)}
                            className="font-semibold text-sky-700 italic hover:underline cursor-pointer"
                          >
                            {FncGlobalFormatArrcpn(String(val))}
                          </div>
                        ) : key == "arrcpn" ? (
                          FncGlobalFormatArrcpn(String(val))
                        ) : key == "pnrcde" ? (
                          <div
                            onClick={() => pnrnow(String(val))}
                            className="font-semibold text-sky-700 italic hover:underline cursor-pointer"
                          >
                            {val}
                          </div>
                        ) : key == "rtlsrs" ? (
                          <div className="duration-300 flexbtw">
                            <div>
                              {iptobj?.prmkey === log.prmkey ? (
                                <div className="duration-300 relative flexctr">
                                  <span className="duration-300 invisible">
                                    xxxx{iptobj.rtlsrs}
                                  </span>
                                  <div className="duration-300 h-8 absolute">
                                    <UixGlobalInputxFormdt
                                      typipt="text"
                                      length={undefined}
                                      queryx={key.toString()}
                                      params={iptobj.rtlsrs}
                                      plchdr=""
                                      repprm={(e) => actipt(e, log)}
                                      labelx=""
                                    />
                                  </div>
                                </div>
                              ) : (
                                val
                              )}
                            </div>
                            <div className="afull max-w-fit flexctr gap-x-1.5 relative">
                              <div
                                className={`flexctr btnsbm duration-300 cursor-pointer ${iptobj?.prmkey === log.prmkey
                                  ? "w-8 opacity-100"
                                  : "w-5 opacity-0 select-none pointer-events-none"
                                  }`}
                                onClick={() => update(log)}
                              >
                                <UixGlobalIconvcCeklis
                                  bold={2.5}
                                  color="#53eafd"
                                  size={1.4}
                                />
                              </div>
                              <div
                                className={`flexctr btnsbm duration-300 cursor-pointer ${iptobj?.prmkey === log.prmkey
                                  ? "w-8 opacity-100"
                                  : "w-5 opacity-0 select-none pointer-events-none"
                                  }`}
                                onClick={() =>
                                  iptobjSet({ ...log, prmkey: "" })
                                }
                              >
                                <UixGlobalIconvcCancel
                                  bold={2.5}
                                  color="#fb2c36"
                                  size={1.4}
                                />
                              </div>
                              <div
                                className={`absolute flexctr btnsbm duration-300 cursor-pointer ${iptobj?.prmkey === log.prmkey
                                  ? "w-8 opacity-0 select-none pointer-events-none"
                                  : "w-8 opacity-100"
                                  } ${rjcedt == log.prmkey ? "shkeit btncxl" : ""
                                  }`}
                                onClick={() => actedt(log)}
                              >
                                <UixGlobalIconvcEditdt
                                  bold={2.5}
                                  color="white"
                                  size={1.2}
                                />
                              </div>
                            </div>
                          </div>
                        ) : (
                          val
                        )}
                      </td>
                    ) : (
                      ""
                    )
                  )}
                </tr>
              ))
            ) : (
              <tr>
                <td className="tdbody text-center py-2">Empty</td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </>
  );
}
