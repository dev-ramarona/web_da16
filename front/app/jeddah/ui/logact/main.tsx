import { ApiJeddahDtbaseLogact } from "../../api/apiJeddahDtbase";
import { MdlJeddahParamsLogact } from "../../model/mdlJeddahParams";
import UixJeddahLogactTablex from "./table";

export default async function UixJeddahLogactMainpg() {
  // await new Promise((r) => setTimeout(r, 2000));
  const logact: MdlJeddahParamsLogact[] = await ApiJeddahDtbaseLogact();
  return (
    <>
      {logact.length > 0 ? (
        <UixJeddahLogactTablex logact={logact} />
      ) : (
        <div className="w-full h-fit flexctr text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
    </>
  );
}
