import UixGlobalPagntnMainpg from "@/app/global/ui/client/uixGlobalPagntn";
import { ApiJeddahFlnsmrGetarr } from "../../api/apiJeddahDtbase";
import {
  MdlJeddahInputxAllpnr,
  MdlJeddahParamsFlnsmr,
} from "../../model/mdlJeddahParams";
import UixJeddahFlnsmrSearch from "./search";
import UixJeddahFlnsmrTablex from "./table";

export default async function UixJeddahFlnsmrMainpg({
  trtprm,
}: {
  trtprm: MdlJeddahInputxAllpnr;
}) {
  // await new Promise((r) => setTimeout(r, 2000));
  const flnsmr = await ApiJeddahFlnsmrGetarr(trtprm);
  const arrdta: MdlJeddahParamsFlnsmr[] = flnsmr.arrdta;
  const totdta: number = flnsmr.totdta;
  return (
    <>
      <UixJeddahFlnsmrSearch trtprm={trtprm} />
      {arrdta.length > 0 ? (
        <UixJeddahFlnsmrTablex flnsmr={arrdta} onpkey={trtprm.prmkey_flnsmr} />
      ) : (
        <div className="afull flexctr text-base font-semibold text-sky-800">
          No database Summary PNR
        </div>
      )}
      <UixGlobalPagntnMainpg
        pgenbr={trtprm.pagenw_flnsmr}
        pgestr="pagenw_flnsmr"
        totdta={totdta}
      />
    </>
  );
}
