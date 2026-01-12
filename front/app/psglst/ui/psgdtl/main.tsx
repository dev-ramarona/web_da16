import UixGlobalPagntnMainpg from "@/app/global/ui/client/uixGlobalPagntn";
import { ApiPsglstAcpedtDtbase } from "../../api/apiPsglstAcpedt";
import { ApiPsglstPsgdtlGetall } from "../../api/apiPsglstPsgdtl";
import {
  MdlPsglstAcpedtDtbase,
  MdlPsglstSrcprmAllprm,
  MdlPsglstPsgdtlFrntnd,
} from "../../model/mdlPsglstParams";
import UixPsglstDetailSearch from "./search";
import UixPsglstDetailTablex from "./tablex";
import { mdlGlobalAllusrCookie } from "@/app/global/model/mdlGlobalPrimer";

export default async function UixPsglstDetailMainpg({
  trtprm,
  datefl,
  cookie,
}: {
  trtprm: MdlPsglstSrcprmAllprm;
  datefl: string[];
  cookie: mdlGlobalAllusrCookie;
}) {
  // await new Promise((r) => setTimeout(r, 2000));
  const psgdtl = await ApiPsglstPsgdtlGetall(trtprm);
  const arrdta: MdlPsglstPsgdtlFrntnd[] = psgdtl.arrdta;
  const totdta: number = psgdtl.totdta;
  const acpedt: MdlPsglstAcpedtDtbase[] = await ApiPsglstAcpedtDtbase();
  return (
    <>
      <UixPsglstDetailSearch trtprm={trtprm} datefl={datefl} />
      {arrdta.length > 0 ? (
        <UixPsglstDetailTablex detail={arrdta} acpedt={acpedt} cookie={cookie} />
      ) : (
        <div className="w-full h-fit flexctr text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
      <UixGlobalPagntnMainpg
        pgenbr={trtprm.pagenw_psgdtl}
        pgestr="pagenw_psgdtl"
        totdta={totdta}
      />
    </>
  );
}
