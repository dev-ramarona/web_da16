"use client";
import {
  FncGlobalFormatDatefm,
  FncGlobalFormatDateip,
} from "@/app/global/function/fncGlobalFormat";
import { MdlJeddahParamsFlnsmr } from "../../model/mdlJeddahParams";
import { FncGlobalParamsEdlink } from "@/app/global/function/fncGlobalParams";

export default function UixJeddahFlnsmrTablex({
  flnsmr,
  onpkey,
}: {
  flnsmr: MdlJeddahParamsFlnsmr[];
  onpkey: string;
}) {
  const rplprm = FncGlobalParamsEdlink();

  // Focus on this Flight
  const flnnow = (pnr: MdlJeddahParamsFlnsmr) => {
    rplprm(
      [
        "flnbfl_pnrdtl",
        "airlfl_pnrdtl",
        "datefl_pnrdtl",
        "routfl_pnrdtl",
        "pagenw_pnrdtl",
        "pnrclk_pnrdtl",
      ],
      [
        pnr.flnbfl,
        pnr.airlfl,
        FncGlobalFormatDateip(String(pnr.datefl)),
        pnr.routfl,
        "1",
        String(Math.random()),
      ]
    );
  };

  return (
    <>
      <div className="afull max-h-fit overflow-auto rounded-lg ring-2 ring-sky-800">
        <table className="w-full">
          <thead className="sticky top-0 z-10 text-white">
            <tr>
              {flnsmr && flnsmr.length > 0 ? (
                Object.entries(flnsmr[0]).map(([key]) =>
                  key != "prmkey" ? (
                    <th key={key} className="thhead">
                      {key}
                    </th>
                  ) : (
                    <th key={key} className="thhead">
                      Show PNR Detail
                    </th>
                  )
                )
              ) : (
                <th className="thhead">Empty</th>
              )}
            </tr>
          </thead>
          <tbody className="text-slate-700 bg-sky-100">
            {flnsmr && flnsmr.length > 0 ? (
              flnsmr.map((log, idx) => (
                <tr className="h-8 group" key={idx}>
                  {Object.entries(log).map(([key, val]) =>
                    key != "prmkey" ? (
                      <td
                        className={`tdbody text-center ${
                          log.prmkey == onpkey ? "bg-white font-semibold" : ""
                        } duration-300`}
                        key={key}
                      >
                        {["datefl", "dateup", "timeup"].includes(key)
                          ? FncGlobalFormatDatefm(String(val))
                          : val}
                      </td>
                    ) : (
                      <td
                        className={`tdbody text-center ${
                          log.prmkey == onpkey ? "bg-white font-semibold" : ""
                        } duration-300`}
                        key={key}
                      >
                        <div
                          onClick={() => flnnow(log)}
                          className="font-semibold text-sky-700 italic hover:underline cursor-pointer"
                        >
                          {val}
                        </div>
                      </td>
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
