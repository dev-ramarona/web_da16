import UixGlobalPagntnMainpg from "@/app/global/ui/client/uixGlobalPagntn";
import { ApiJeddahPnrdtlGetarr } from "../../api/apiJeddahDtbase";
import {
  MdlJeddahInputxAllpnr,
  MdlJeddahParamsPnrdtl,
} from "../../model/mdlJeddahParams";

import UixJeddahPnrdtlSearch from "./search";
import UixJeddahPnrdtlTablex from "./table";

export default async function UixJeddahPnrdtlMainpg({
  trtprm,
}: {
  trtprm: MdlJeddahInputxAllpnr;
}) {
  // await new Promise((r) => setTimeout(r, 2000));
  const pnrdtl = await ApiJeddahPnrdtlGetarr(trtprm);
  const arrdta: MdlJeddahParamsPnrdtl[] = pnrdtl.arrdta;
  const totdta: number = pnrdtl.totdta;

  return (
    <>
      <UixJeddahPnrdtlSearch trtprm={trtprm} />
      {arrdta.length > 0 ? (
        <UixJeddahPnrdtlTablex
          pnrdtl={arrdta}
          pnrcde={trtprm.pnrcde_pnrdtl || ""}
        />
      ) : (
        <div className="w-full h-fit flexctr text-base font-semibold text-sky-800">
          No database Summary PNR
        </div>
      )}
      <UixGlobalPagntnMainpg
        pgenbr={trtprm.pagenw_pnrdtl}
        pgestr="pagenw_pnrdtl"
        totdta={totdta}
      />
    </>
  );
}
