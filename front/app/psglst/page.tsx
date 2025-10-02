import { Suspense } from "react";
import { ApiGlobalCookieGetdta } from "../global/api/apiCookieParams";
import { UixGlobalIconvcSeting } from "../global/ui/server/uixGlobalIconvc";
import UixGlobalLoadngAnmate from "../global/ui/server/UixGlobalLoadng";
import UixPsglstActlogMainpg from "./ui/actlog/main";
import UixPsglstErrlogMainpg from "./ui/errlog/main";
import UixPsglstDetailMainpg from "./ui/detail/main";
import { MdlPsglstAllprmSrcprm } from "./model/mdlPsglstParams";
import { FncPsglstDetailParams } from "./function/fncPsglstParams";

export default async function Page(props: {
  searchParams: Promise<MdlPsglstAllprmSrcprm>;
}) {
  const cookie = await ApiGlobalCookieGetdta();
  const qryprm = await props.searchParams;
  const trtprm = FncPsglstDetailParams(qryprm);
  // await new Promise((r) => setTimeout(r, 3000));
  return (
    <div className="afull flex justify-start items-start flex-wrap p-1.5 md:p-6">
      <div className="w-full md:w-[10rem] min-w-1/5 h-[15rem] md:h-[20rem] max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col shadow-md">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Setup Parameter
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixPsglstActlogMainpg />
          </Suspense>
        </div>
      </div>
      <div className="w-full md:w-[15rem] min-w-4/5 h-[30rem] md:h-[20rem] max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col shadow-md">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Setup Parameter
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixPsglstErrlogMainpg />
          </Suspense>
        </div>
      </div>
      <div className="w-full md:w-[20rem] min-w-full h-[45rem] md:h-[30rem] max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col shadow-md">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Setup Parameter
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixPsglstDetailMainpg trtprm={trtprm} />
          </Suspense>
        </div>
      </div>
    </div>
  );
}
