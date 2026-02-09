import { Suspense } from "react";
import { UixGlobalIconvcSeting } from "../global/ui/server/uixGlobalIconvc";
import UixGlobalLoadngAnmate from "../global/ui/server/UixGlobalLoadng";
import UixPsglstActlogMainpg from "../psglst/ui/actlog/main";
import { ApiPsglstActlogDtbase } from "../psglst/api/apiPsglstActlog";
import { MdlPsglstActlogDtbase, MdlPsglstGlobalSrcprm } from "../psglst/model/mdlPsglstParams";
import UixPsglstErrlogMainpg from "./ui/errlog/main";
import { FncSlsflwErrlogSrcprm, FncSlsflwPsgdtlSrcprm } from "./function/fncSlsflwParams";
import UixSlsflwDetailMainpg from "./ui/psgdtl/main";
import { ApiGlobalCookieGetdta } from "../global/api/apiCookieParams";


export default async function Page(props: {
  searchParams: Promise<MdlPsglstGlobalSrcprm>;
}) {
  const cookie = await ApiGlobalCookieGetdta();
  const qryprm = await props.searchParams;
  const actobj = await ApiPsglstActlogDtbase();
  const actlog: MdlPsglstActlogDtbase[] = actobj.actlog;
  const actdte: string[] = actobj.datefl;
  const prmErrlog = FncSlsflwErrlogSrcprm(qryprm);
  const prmPsgdtl = FncSlsflwPsgdtlSrcprm(qryprm, actdte);
  return (
    <div className="afull flex justify-start items-start flex-wrap p-1.5 md:p-6">
      <div className="w-full md:w-[10rem] min-w-1/5 h-[15rem] md:h-[20rem] max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col shadow-md">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Log Action
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixPsglstActlogMainpg actlog={actlog} />
          </Suspense>
        </div>
      </div>
      <div className="w-full md:w-[15rem] min-w-4/5 h-[30rem] md:h-[20rem] max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col shadow-md">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Log error
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixPsglstErrlogMainpg prmErrlog={prmErrlog} />
          </Suspense>
        </div>
      </div>
      <div className="w-full md:w-[20rem] min-w-full h-[45rem] md:h-[40rem] max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col shadow-md">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Passangger detail
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixSlsflwDetailMainpg prmPsgdtl={prmPsgdtl} datefl={actdte} cookie={cookie} />
          </Suspense>
        </div>
      </div>
      <div className="w-full md:w-[20rem] min-w-full h-[15rem] md:h-[15rem] max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col shadow-md">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Process manual
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <></>
          </Suspense>
        </div>
      </div>
    </div >
  );
}
