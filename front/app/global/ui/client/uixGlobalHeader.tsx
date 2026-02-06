"use client";

import { usePathname } from "next/navigation";
import { useEffect, useState } from "react";
import { MdlGlobalApplstDtbase } from "../../model/mdlGlobalApplst";

export default function UixGlobalHeaderClient({
  applst,
}: {
  applst: MdlGlobalApplstDtbase[];
}) {
  const pthnme = usePathname(); // misal: "/opclss"
  const [nowpth, nowpthSet] = useState("");
  const [dtlpth, dtlpthSet] = useState("Wellcome");
  useEffect(() => {
    const sgment = pthnme.split("/").filter(Boolean).pop();
    nowpthSet(sgment || "");
    applst.forEach((app) => {
      if (app.prmkey == sgment) {
        dtlpthSet(app.detail);
      }
    });
  }, [pthnme, applst]);
  if (nowpth == "" || nowpth == "global") return;
  return (
    <>
      <div className="font-semibold text-2xl text-slate-800 tracking-wide">
        {dtlpth}
      </div>
      <div className="text-lg text-slate-800 px-1.5">Page</div>
    </>
  );
}
