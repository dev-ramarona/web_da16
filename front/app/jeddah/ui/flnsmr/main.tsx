import UixGlobalPagntnMainpg from "@/app/global/ui/client/uixGlobalPagntn";
import {
  MdlJeddahInputxAllprm,
  MdlJeddahFlnsmrDtbase,
} from "../../model/mdlJeddahMainpr";
import UixJeddahFlnsmrSearch from "./search";
import UixJeddahFlnsmrTablex from "./table";
import { ApiJeddahFlnsmrGetall } from "../../api/apiJeddahFlnsmr";

export default async function UixJeddahFlnsmrMainpg({
  trtprm,
}: {
  trtprm: MdlJeddahInputxAllprm;
}) {
  // await new Promise((r) => setTimeout(r, 2000));
  const flnsmr = await ApiJeddahFlnsmrGetall(trtprm);
  const arrdta: MdlJeddahFlnsmrDtbase[] = flnsmr.arrdta;
  const totdta: number = flnsmr.totdta;
  return (
    <>
      <UixJeddahFlnsmrSearch trtprm={trtprm} />
      {arrdta.length > 0 ? (
        <UixJeddahFlnsmrTablex flnsmr={arrdta} />
      ) : (
        <div className="afull flexctr text-base font-semibold text-sky-800">
          No database Summary PNR
        </div>
      )}
      <UixGlobalPagntnMainpg
        pgview={15}
        pgenbr={trtprm.pagenw_flnsmr}
        pgestr="pagenw_flnsmr"
        totdta={totdta}
      />
    </>
  );
}
