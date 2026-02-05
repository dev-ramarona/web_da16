import UixGlobalPagntnMainpg from "@/app/global/ui/client/uixGlobalPagntn";
import UixPsglstErrlogTablex from "./tablex";
import { MdlPsglstErrlogDtbase, MdlPsglstSrcprmAllprm } from "@/app/psglst/model/mdlPsglstParams";
import { ApiPsglstErrlogDtbase } from "@/app/psglst/api/apiPsglstErrlog";



export default async function UixPsglstErrlogMainpg({ trtprm }: { trtprm: MdlPsglstSrcprmAllprm }) {
  const rslobj = await ApiPsglstErrlogDtbase(trtprm);
  const errlog: MdlPsglstErrlogDtbase[] = rslobj.arrdta
  const totdta: number = rslobj.totdta
  return (
    <>
      {errlog.length > 0 ? (
        <>
          <UixPsglstErrlogTablex errlog={errlog} />
          <UixGlobalPagntnMainpg
            pgview={5}
            pgenbr={trtprm.pagenw_errlog}
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
