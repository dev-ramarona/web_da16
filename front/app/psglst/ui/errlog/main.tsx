import UixGlobalPagntnMainpg from "@/app/global/ui/client/uixGlobalPagntn";
import { ApiPsglstErrlogDtbase } from "../../api/apiPsglstErrlog";
import { MdlPsglstErrlogDtbase, MdlPsglstErrlogSrcprm } from "../../model/mdlPsglstParams";
import UixPsglstErrlogTablex from "./tablex";



export default async function UixPsglstErrlogMainpg({ prmErrlog }: { prmErrlog: MdlPsglstErrlogSrcprm }) {
  const rslobj = await ApiPsglstErrlogDtbase({
    ...prmErrlog, erdvsn_errlog:
      (prmErrlog.erdvsn_errlog == "") ? "SLSRPT" : prmErrlog.erdvsn_errlog
  });
  const errlog: MdlPsglstErrlogDtbase[] = rslobj.arrdta
  const totdta: number = rslobj.totdta
  return (
    <>
      {errlog.length > 0 ? (
        <>
          <UixPsglstErrlogTablex errlog={errlog} />
          <UixGlobalPagntnMainpg
            pgview={5}
            pgenbr={prmErrlog.pagenw_errlog}
            pgestr="pagenw_errlog"
            totdta={totdta}
          />
        </>
      ) : (
        <div className="w-full h-fit flexctr text-base font-semibold text-sky-800">
          No database Log Error
        </div>
      )}

    </>
  );
}
