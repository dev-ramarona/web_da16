"use client";
import Link from "next/link";
import { FncGlobalParamsEdlink } from "../../function/fncGlobalParams";
import {
  UixGlobalIconvcNextpg,
  UixGlobalIconvcPrevpg,
} from "../server/uixGlobalIconvc";

export default function UixGlobalPagntnMainpg({
  pgestr,
  totdta,
  pgenbr,
}: {
  pgestr: string;
  totdta: number;
  pgenbr: number;
}) {
  const rplprm = FncGlobalParamsEdlink();
  const maxpge = Math.ceil(totdta / 15);
  const totpge = Math.min(Math.ceil(totdta / 15), 10);
  return (
    <div className="w-full h-16 pt-1.5 flexbtw">
      <div
        className="w-6 h-6 flexctr cursor-pointer"
        onClick={() => rplprm(pgestr, 0 + "")}
      >
        <UixGlobalIconvcPrevpg bold={3} color="black" size={1.3} />
      </div>
      <div
        className="w-6 h-6 flexctr cursor-pointer"
        onClick={() => rplprm(pgestr, Math.max(pgenbr - 10, 1) + "")}
      >
        <UixGlobalIconvcPrevpg bold={3} color="black" size={1.3} />
      </div>
      {Array.from({ length: totpge }, (_, i) => {
        const pgenow = (Math.ceil(pgenbr / 10) - 1) * 10 + 1 + i;
        if (pgenow > maxpge) return null;
        return (
          <div
            className={`w-4 md:w-7 min-w-fit h-4 md:h-7 flexctr  cursor-pointer text-[0.45rem] md:text-xs ${
              pgenow == Number(pgenbr)
                ? "btnsbm ring-2 ring-sky-800"
                : "hover:bg-slate-200 ring-2 ring-slate-400 rounded-md duration-300"
            }`}
            key={pgenow}
            onClick={() => rplprm(pgestr, pgenow + "")}
          >
            <div>{pgenow}</div>
          </div>
        );
      })}
      <div
        className="w-6 h-6 flexctr cursor-pointer"
        onClick={() => rplprm(pgestr, Math.min(pgenbr + 10, maxpge) + "")}
      >
        <UixGlobalIconvcNextpg bold={3} color="black" size={1.3} />
      </div>
      <div
        className="w-6 h-6 flexctr cursor-pointer"
        onClick={() => rplprm(pgestr, maxpge + "")}
      >
        <UixGlobalIconvcNextpg bold={3} color="black" size={1.3} />
      </div>
    </div>
  );
}
