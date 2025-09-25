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
  typipt: "text" | "date" | "month" | "number" | "file";
  labelx: "hidden" | "";
  length: number | undefined;
  queryx: string;
  params: string;
  plchdr: string;
  repprm: null | ((e: React.ChangeEvent<HTMLInputElement>) => void);
}) {
  const refdte = useRef<HTMLInputElement>(null);
  const pckrdt = () => refdte.current?.showPicker();
  return (
    <div className="afull py-1.5 px-1 flexstr relative text-[0.6rem] md:text-[0.66rem]">
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
            maxLength={length ? length : undefined}
            type={typipt}
            id={queryx}
            onChange={(e) => (repprm ? repprm(e) : "")}
            onClick={() =>
              typipt == "date" || typipt == "month" ? pckrdt() : null
            }
            ref={typipt == "date" || typipt == "month" ? refdte : null}
          />
          <label
            className={`absolute left-3 select-none cursor-text whitespace-nowrap ${params != ""
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
      )}
    </div>
  );
}
