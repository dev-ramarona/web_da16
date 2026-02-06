import UixGlobalPagntnMainpg from "@/app/global/ui/client/uixGlobalPagntn";
import UixPsglstErrlogTablex from "./tablex";
import { MdlPsglstErrlogDtbase, MdlPsglstErrlogSrcprm } from "@/app/psglst/model/mdlPsglstParams";
import { ApiPsglstErrlogDtbase } from "@/app/psglst/api/apiPsglstErrlog";



export default async function UixPsglstErrlogMainpg({ prmErrlog }: { prmErrlog: MdlPsglstErrlogSrcprm }) {
  const rslobj = await ApiPsglstErrlogDtbase(prmErrlog);
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
