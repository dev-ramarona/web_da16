import { mdlGlobalAllusrCookie } from "@/app/global/model/mdlGlobalPrimer";
import { ApiJeddahAgtnmeNullnm } from "../../api/apiJeddahAgtnme";
import {
  MdlJeddahInputxAllprm,
  MdlJeddahAgtnmeDtbase,
} from "../../model/mdlJeddahMainpr";
import UixJeddahAgtnmeTablex from "./table";
import UixJeddahAgtnmerSearch from "./search";
import UixGlobalPagntnMainpg from "@/app/global/ui/client/uixGlobalPagntn";

export default async function UixJeddahAgtnmeMainpg({
  cookie,
  trtprm,
}: {
  cookie: mdlGlobalAllusrCookie;
  trtprm: MdlJeddahInputxAllprm;
}) {
  // await new Promise((r) => setTimeout(r, 5000));
  const agtnul = await ApiJeddahAgtnmeNullnm(trtprm);
  const arrdta: MdlJeddahAgtnmeDtbase[] = agtnul.arrdta;
  const totdta: number = agtnul.totdta;
  return (
    <>
      <UixJeddahAgtnmerSearch trtprm={trtprm} />
      {arrdta.length > 0 ? (
        <>
          <UixJeddahAgtnmeTablex agtnul={arrdta} cookie={cookie} />
          <UixGlobalPagntnMainpg
            pgview={15}
            pgenbr={trtprm.pagenw_agtnme}
            pgestr="pagenw_agtnme"
            totdta={totdta}
          />
        </>
      ) : (
        <div className="w-full h-fit flexctr text-base font-semibold text-sky-800">
          All Agent Name Updated
        </div>
      )}
    </>
  );
}
