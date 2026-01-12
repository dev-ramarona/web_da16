import UixGlobalPagntnMainpg from "@/app/global/ui/client/uixGlobalPagntn";
import { ApiPsglstErrlogDtbase } from "../../api/apiPsglstErrlog";
import { MdlPsglstErrlogDtbase, MdlPsglstSrcprmAllprm } from "../../model/mdlPsglstParams";
import UixPsglstErrlogTablex from "./tablex";



export default async function UixPsglstErrlogMainpg({ trtprm }: { trtprm: MdlPsglstSrcprmAllprm }) {
  // await new Promise((r) => setTimeout(r, 2000));
  const rslobj = await ApiPsglstErrlogDtbase(trtprm);
  const errlog: MdlPsglstErrlogDtbase[] = rslobj.arrdta
  const totdta: number = rslobj.totdta
  return (
    <>
      {errlog.length > 0 ? (
        <>
          <UixPsglstErrlogTablex errlog={errlog} />
          <UixGlobalPagntnMainpg
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
