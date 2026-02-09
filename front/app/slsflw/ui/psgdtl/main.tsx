import UixGlobalPagntnMainpg from "@/app/global/ui/client/uixGlobalPagntn";
import { mdlGlobalAllusrCookie } from "@/app/global/model/mdlGlobalPrimer";
import { MdlPsglstAcpedtDtbase, MdlPsglstPsgdtlFrntnd, MdlPsglstPsgdtlSrcprm } from "@/app/psglst/model/mdlPsglstParams";
import { ApiPsglstPsgdtlGetall } from "@/app/psglst/api/apiPsglstPsgdtl";
import { ApiPsglstAcpedtDtbase } from "@/app/psglst/api/apiPsglstAcpedt";
import UixSlsflwDetailSearch from "./search";
import UixSlsflwDetailTablex from "./tablex";

export default async function UixSlsflwDetailMainpg({
  prmPsgdtl,
  datefl,
  cookie,
}: {
  prmPsgdtl: MdlPsglstPsgdtlSrcprm;
  datefl: string[];
  cookie: mdlGlobalAllusrCookie;
}) {
  const psgdtl = await ApiPsglstPsgdtlGetall({
    ...prmPsgdtl, nclear_psgdtl:
      (prmPsgdtl.nclear_psgdtl == "") ? "SLSRPT" : prmPsgdtl.nclear_psgdtl
  });
  const arrdta: MdlPsglstPsgdtlFrntnd[] = psgdtl.arrdta;
  const totdta: number = psgdtl.totdta;
  const acpedt: MdlPsglstAcpedtDtbase[] = await ApiPsglstAcpedtDtbase();
  return (
    <>
      <UixSlsflwDetailSearch prmPsgdtl={prmPsgdtl} datefl={datefl} />
      {arrdta.length > 0 ? (
        <UixSlsflwDetailTablex detail={arrdta} acpedt={acpedt} cookie={cookie} />
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
