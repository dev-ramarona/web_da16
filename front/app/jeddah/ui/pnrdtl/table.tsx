"use client";
import {
  FncGlobalFormatArrcpn,
  FncGlobalFormatDatefm,
} from "@/app/global/function/fncGlobalFormat";
import { MdlJeddahPnrdtlDtbase } from "../../model/mdlJeddahMainpr";
import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";
export default function UixJeddahPnrdtlTablex({
  pnrdtl,
  pnrcde,
}: {
  pnrdtl: MdlJeddahPnrdtlDtbase[];
  pnrcde: string;
}) {
  // Filter all PNR arr split
  const rplprm = FncGlobalParamsEdlink();
  const pnrspl = (val: string, str: string) => {
    const arr = val.split("|");
    const pnr = [str];
    if (pnrcde != "" && !pnrcde.includes(str)) pnr.push(pnrcde);
    for (let i = 0; i < arr.length; i++) pnr.push(arr[i].substring(0, 6));
    rplprm(
      ["pnrcde_pnrdtl", "pnrclk_pnrdtl", "pagenw_pnrdtl"],
      [pnr.join("|"), String(Math.random()), ""]
    );
  };
  return (
    <>
      <div className="afull max-h-fit overflow-auto rounded-lg ring-2 ring-sky-800">
        <table className="w-full">
          <thead className="sticky top-0 z-10 text-white">
            <tr>
              {pnrdtl && pnrdtl.length > 0 ? (
                Object.entries(pnrdtl[0]).map(([key]) =>
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
            {pnrdtl && pnrdtl.length > 0 ? (
              pnrdtl.map((log, idx) => (
                <tr className="h-8 group" key={idx}>
                  {Object.entries(log).map(([key, val]) =>
                    key != "prmkey" ? (
                      <td className="tdbody text-center" key={key}>
                        {[
                          "dateup",
                          "timeup",
                          "timecr",
                          "datefl",
                          "",
                          "",
                        ].includes(key) ? (
                          FncGlobalFormatDatefm(String(val))
                        ) : key == "arrspl" ? (
                          <div
                            onClick={() => pnrspl(String(val), log.pnrcde)}
                            className="font-semibold text-sky-700 italic hover:underline cursor-pointer"
                          >
                            {FncGlobalFormatArrcpn(String(val))}
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
