"use client";
import { FncGlobalParamsEdlink } from "../../function/fncGlobalParams";
import {
  UixGlobalIconvcNextpg,
  UixGlobalIconvcPrevpg,
  UixGlobalIconvcRfresh,
} from "../server/uixGlobalIconvc";
import { useEffect, useState } from "react";

export default function UixGlobalPagntnMainpg({
  pgview,
  pgestr,
  totdta,
  pgenbr,
}: {
  pgview: number;
  pgestr: string;
  totdta: number;
  pgenbr: number;
}) {
  const rplprm = FncGlobalParamsEdlink();
  const maxpge = Math.ceil(totdta / pgview);
  const totpge = Math.min(Math.ceil(totdta / pgview), 10);
  const [chnged, chngedSet] = useState(false);
  const pagent = function (qry: string | string[], prm: string | string[]) {
    rplprm(qry, prm)
    if (pgenbr != Number(prm)) {
      chngedSet(true);
    }
  }
  useEffect(() => {
    chngedSet(false);
  }, [pgenbr]);
  return (
    <div className="w-full h-16 pt-1.5 flexctr">
      <div className={`${chnged ? "w-16 h-10 translate-y-0" : "w-0 h-0 opacity-0 -translate-y-10"} z-10 absolute bg-white ring-2 ring-sky-300 px-5 py-2 rounded-xl flexctr duration-300`}>
        <div>Wait</div>
        <div className="animate-spin"><UixGlobalIconvcRfresh bold={2} color="black" size={1} /></div>
      </div>
      <div className="afull flexbtw">
        <div
          className="w-6 h-6 flexctr cursor-pointer"
          onClick={() => pagent(pgestr, 0 + "")}
        >
          <UixGlobalIconvcPrevpg bold={3} color="black" size={1.3} />
        </div>
        <div
          className="w-6 h-6 flexctr cursor-pointer"
          onClick={() => pagent(pgestr, Math.max(pgenbr - 10, 1) + "")}
        >
          <UixGlobalIconvcPrevpg bold={3} color="black" size={1.3} />
        </div>
        {Array.from({ length: totpge }, (_, i) => {
          const pgenow = (Math.ceil(pgenbr / 10) - 1) * 10 + 1 + i;
          if (pgenow > maxpge) return null;
          return (
            <div
              className={`w-4 md:w-7 min-w-fit h-4 md:h-7 flexctr  cursor-pointer text-[0.45rem] md:text-xs ${pgenow == Number(pgenbr)
                ? "btnsbm ring-2 ring-sky-800"
                : "hover:bg-slate-200 ring-2 ring-slate-400 rounded-md duration-300"
                }`}
              key={pgenow}
              onClick={() => pagent(pgestr, pgenow + "")}
            >
              <div>{pgenow}</div>
            </div>
          );
        })}
        <div
          className="w-6 h-6 flexctr cursor-pointer"
          onClick={() => pagent(pgestr, Math.min(pgenbr + 10, maxpge) + "")}
        >
          <UixGlobalIconvcNextpg bold={3} color="black" size={1.3} />
        </div>
        <div
          className="w-6 h-6 flexctr cursor-pointer"
          onClick={() => pagent(pgestr, maxpge + "")}
        >
          <UixGlobalIconvcNextpg bold={3} color="black" size={1.3} />
        </div>
      </div>
    </div>
  );
}
