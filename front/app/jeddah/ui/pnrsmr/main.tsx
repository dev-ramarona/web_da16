import { mdlGlobalAllusrCookie } from "@/app/global/model/mdlGlobalPrimer";
import {
  MdlJeddahInputxAllpnr,
  MdlJeddahPnrsmrDtbase,
} from "../../model/mdlJeddahMainpr";
import UixJeddahPnrsmrSearch from "./search";
import UixJeddahPnrsmrTablex from "./table";
import UixGlobalPagntnMainpg from "@/app/global/ui/client/uixGlobalPagntn";
import UixJeddahPnrsmrUpldwn from "./upload";
import { ApiJeddahPnrsmrGetall } from "../../api/apiJeddahPnrsmr";

export default async function UixJeddahPnrsmrMainpg({
  trtprm,
  cookie,
}: {
  trtprm: MdlJeddahInputxAllpnr;
  cookie: mdlGlobalAllusrCookie;
}) {
  // await new Promise((r) => setTimeout(r, 2000));
  const pnrsmr = await ApiJeddahPnrsmrGetall(trtprm);
  const arrdta: MdlJeddahPnrsmrDtbase[] = pnrsmr.arrdta;
  const totdta: number = pnrsmr.totdta;

  return (
    <>
      <UixJeddahPnrsmrSearch trtprm={trtprm} />
      {arrdta.length > 0 ? (
        <UixJeddahPnrsmrTablex
          pnrsmr={arrdta}
          pnrcde={trtprm.pnrcde_pnrsmr}
          cookie={cookie}
        />
      ) : (
        <div className="w-full h-fit flexctr text-base font-semibold text-sky-800">
          No database Summary PNR
        </div>
      )}
      <UixGlobalPagntnMainpg
        pgenbr={trtprm.pagenw_pnrsmr}
        pgestr="pagenw_pnrsmr"
        totdta={totdta}
      />
      <UixJeddahPnrsmrUpldwn trtprm={trtprm} cookie={cookie} />
    </>
  );
}
