import { Suspense } from "react";
import { UixGlobalIconvcSeting } from "../global/ui/server/uixGlobalIconvc";
import UixGlobalLoadngAnmate from "../global/ui/server/UixGlobalLoadng";
import UixHoldstPrcessMainpg from "./ui/prcess/main";

export default async function Page() {
  return (
    <div className="afull flex justify-start items-start flex-wrap p-1.5 md:p-6">
      <div className="w-full md:w-[20rem] min-w-full h-[15rem] md:h-[15rem] max-h-fit p-3">
        <div className="afull max-h-fit rounded-xl py-1.5 px-3 flexstr flex-col shadow-md">
          <div className="w-full text-slate-800 font-semibold text-base py-1.5 flexstr">
            Process manual
            <UixGlobalIconvcSeting color="gray" size={1.3} bold={3} />
          </div>
          <Suspense fallback={<UixGlobalLoadngAnmate />}>
            <UixHoldstPrcessMainpg />
          </Suspense>
        </div>
      </div>
    </div>
  );
}
