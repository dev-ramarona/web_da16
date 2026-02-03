"use client";
import { useRef, useState } from "react";
export default function UixGlobalInputxFormdt({
  typipt,
  labelx,
  length,
  queryx,
  params,
  plchdr,
  repprm,
}: {
  typipt: "text" | "date" | "month" | "number" | "file" | "datetime-local";
  labelx: "hidden" | string;
  length: number | undefined | string[];
  queryx: string;
  params: string;
  plchdr: string;
  repprm: null | ((e: React.ChangeEvent<HTMLInputElement>) => void);
}) {
  const refdte = useRef<HTMLInputElement>(null);
  const pckrdt = () => refdte.current?.showPicker();
  return (
    <div className="afull py-1.5 px-1 flexstr relative text-[0.6rem] md:text-[0.66rem] z-0 hover:z-10">
      {typipt == "file" ? (
        // Input type file
        <>
          <input
            className="hidden peer"
            type={typipt}
            id={queryx}
            accept=".csv"
            multiple
            hidden={labelx == "hidden" ? true : false}
            onChange={(e) => (repprm ? repprm(e) : "")}
          />
          <label
            className={`afull bg-white cursor-pointer rounded-md p-1.5 shadow-md ${params != ""
              ? "text-slate-700 overflow-hidden whitespace-nowrap"
              : "text-white peer-focus:text-slate-500"
              } duration-300`}
            htmlFor={queryx}
          >
            {params}
          </label>
          <label
            className={`absolute left-3 select-none cursor-pointer whitespace-nowrap ${params != ""
              ? `h-1/2 -translate-y-full mb-1 text-slate-600 font-semibold text-[0.65rem] ${labelx == "hidden" ? "opacity-0" : ""
              }`
              : `text-slate-400 peer-focus:h-1/2 peer-focus:-translate-y-full peer-focus:mb-1 ${labelx == "hidden" ? "peer-focus:opacity-0" : ""
              }`
              } duration-300`}
            htmlFor={queryx}
          >
            {plchdr}
          </label>
        </>
      ) : (
        // Input type All
        <>
          <input
            className={`afull bg-white rounded-md p-1.5 peer shadow-md ${params != ""
              ? "text-slate-700"
              : "text-white focus:text-slate-500"
              } duration-300`}
            value={params}
            maxLength={length && typeof length == "number" ? length : undefined}
            min={(Array.isArray(length) ? length[0] : length)}
            max={(Array.isArray(length) ? length[length.length - 1] : length)}
            type={typipt}
            id={queryx}
            onChange={(e) => (repprm ? repprm(e) : "")}
            onClick={() =>
              typipt == "date" || typipt == "month" || typipt == "datetime-local" ? pckrdt() : null
            }
            ref={typipt == "date" || typipt == "month" || typipt == "datetime-local" ? refdte : null}
          />
          <label
            className={`absolute left-3 select-none cursor-text whitespace-nowrap ${params != ""
              ? `h-1/2 -translate-y-full mb-1 text-slate-600 font-semibold text-[0.65rem] ${labelx == "hidden" ? "opacity-0" : ""
              }`
              : `text-slate-400 peer-focus:h-1/2 peer-focus:-translate-y-full peer-focus:mb-1 ${labelx == "hidden" ? "peer-focus:opacity-0" : ""
              }`
              } group/fst duration-300`}
            htmlFor={queryx}
          >
            <div className="flexctr">
              <div className="px-1">{plchdr}</div>
              {labelx != "hidden" && labelx != "" ? (
                <div className="rounded-full flexctr ring w-3 h-3 opacity-0 group-hover/fst:opacity-100 duration-300 pl-[0.6px] cursor-pointer group/scd">
                  <div>?</div>
                  <div className="absolute -right-1 translate-x-full opacity-0 group-hover/scd:opacity-100 w-0 group-hover/scd:w-28 h-0 group-hover/scd:h-12 bg-white ring-2 duration-300 p-1.5 rounded-md overflow-auto">
                    {labelx.split("|").map((lbl, idx) => (
                      <div key={idx} className="text-[0.6rem] text-slate-700 font-mono">
                        {lbl}
                      </div>
                    ))}
                  </div>
                </div>) : ""}
            </div>
          </label>
        </>
      )}
    </div>
  );
}
