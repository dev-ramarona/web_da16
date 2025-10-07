import { Suspense } from "react";
import { ApiGlobalCookieGetdta } from "../global/api/apiCookieParams";
import { UixGlobalIconvcSeting } from "../global/ui/server/uixGlobalIconvc";
import UixJeddahAgtnmeMainpg from "./ui/agtnme/main";
import UixGlobalLoadngAnmate from "../global/ui/server/UixGlobalLoadng";
import UixJeddahLogactMainpg from "./ui/actlog/main";
import UixJeddahPrcessMainpg from "./ui/prcess/main";
import UixJeddahPnrdtlMainpg from "./ui/pnrdtl/main";
import { MdlJeddahInputxAllprm } from "./model/mdlJeddahMainpr";
import UixJeddahPnrsmrMainpg from "./ui/pnrsmr/main";
import UixJeddahFlnsmrMainpg from "./ui/flnsmr/main";
import { FncJeddahAllpnrrMainpr } from "./function/fncJeddahMainpr";
import UixJeddahAddflnMainpg from "./ui/flnbfl/main";

export default async function Page(props: {
  searchParams: Promise<MdlJeddahInputxAllprm>;
}) {
  const cookie = await ApiGlobalCookieGetdta();
  const qryprm = await props.searchParams;
  const trtprm = FncJeddahAllpnrrMainpr(qryprm);
  // await new Promise((r) => setTimeout(r, 500));
  return (
    <div className="afull flex justify-start items-start flex-wrap p-1.5 md:p-6">

      {/* First section */}
      <div className="w-full md:min-w-1/4 md:w-[15rem] h-[24rem] md:h-[25rem] max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col shadow-md">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Log Action
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixJeddahLogactMainpg />
          </Suspense>
          <div className="w-52 h-14 md:h-20 py-3 md:py-0">
            <Suspense fallback={<UixGlobalLoadngAnmate />}>
              <UixJeddahPrcessMainpg />
            </Suspense>
          </div>
        </div>
      </div>

      {/* Second section */}
      <div className="w-full md:min-w-3/4 md:w-[25rem] h-[24rem] md:h-[25rem] max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col shadow-md">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Edit Agent Name
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixJeddahAgtnmeMainpg cookie={cookie} trtprm={trtprm} />
          </Suspense>
        </div>
      </div>

      {/* Third section */}
      <div className="w-full md:w-full h-[45rem] md:h-[37rem] max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col shadow-md">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            {trtprm.chssmr === "" ? "Summary PNR" : "Summary Flight Number"}
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            {trtprm.chssmr === "" ? (
              <UixJeddahPnrsmrMainpg trtprm={trtprm} cookie={cookie} />
            ) : (
              <UixJeddahFlnsmrMainpg trtprm={trtprm} />
            )}
          </Suspense>
        </div>
      </div>

      {/* Fourth section */}
      <div className="w-full md:w-full h-[45rem] md:h-[37rem] max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col shadow-md">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Detail PNR
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixJeddahPnrdtlMainpg trtprm={trtprm} />
          </Suspense>
        </div>
      </div>

      {/* Fifth section */}
      <div className="w-full md:w-full h-fit p-3">
        <div className="afull rounded-xl py-1.5 px-3 flexstr flex-col shadow-md">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Add New Flight Jeddah
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixJeddahAddflnMainpg cookie={cookie} />
          </Suspense>
        </div>
      </div>
    </div>
  );
}
