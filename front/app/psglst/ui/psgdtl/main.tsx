import UixGlobalPagntnMainpg from "@/app/global/ui/client/uixGlobalPagntn";
import { ApiPsglstAcpedtDtbase } from "../../api/apiPsglstAcpedt";
import { ApiPsglstPsgdtlGetall } from "../../api/apiPsglstPsgdtl";
import {
  MdlPsglstAcpedtDtbase,
  MdlPsglstPsgdtlFrntnd,
  MdlPsglstPsgdtlSrcprm,
} from "../../model/mdlPsglstParams";
import UixPsglstDetailSearch from "./search";
import UixPsglstDetailTablex from "./tablex";
import { mdlGlobalAllusrCookie } from "@/app/global/model/mdlGlobalPrimer";

export default async function UixPsglstDetailMainpg({
  prmPsgdtl,
  datefl,
  cookie,
}: {
  prmPsgdtl: MdlPsglstPsgdtlSrcprm;
  datefl: string[];
  cookie: mdlGlobalAllusrCookie;
}) {
  const psgdtl = await ApiPsglstPsgdtlGetall(prmPsgdtl);
  const arrdta: MdlPsglstPsgdtlFrntnd[] = psgdtl.arrdta;
  const totdta: number = psgdtl.totdta;
  const acpedt: MdlPsglstAcpedtDtbase[] = await ApiPsglstAcpedtDtbase();
  return (
    <>
      <UixPsglstDetailSearch prmPsgdtl={prmPsgdtl} datefl={datefl} />
      {arrdta.length > 0 ? (
        <UixPsglstDetailTablex detail={arrdta} acpedt={acpedt} cookie={cookie} />
      ) : (
        <div className="w-full h-fit flexctr text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
      <UixGlobalPagntnMainpg
        pgview={15}
        pgenbr={prmPsgdtl.pagenw_psgdtl}
        pgestr="pagenw_psgdtl"
        totdta={totdta}
      />
    </>
  );
}
