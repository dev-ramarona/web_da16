import { ApiPsglstActlogDtbase } from "@/app/psglst/api/apiPsglstActlog";
import { MdlPsglstActlogDtbase } from "@/app/psglst/model/mdlPsglstParams";
import UixJeddahActlogTablex from "./table";


export default async function UixJeddahLogactMainpg() {
  // await new Promise((r) => setTimeout(r, 2000));
  const logact: MdlPsglstActlogDtbase[] = await ApiPsglstActlogDtbase();
  return (
    <>
      {logact.length > 0 ? (
        <UixJeddahActlogTablex logact={logact} />
      ) : (
        <div className="w-full h-fit flexctr text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
    </>
  );
}
