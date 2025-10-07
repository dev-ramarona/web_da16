import { ApiJeddahDtbaseActlog } from "../../api/apiJeddahActlog";
import { MdlJeddahParamsActlog } from "../../model/mdlJeddahMainpr";
import UixJeddahActlogTablex from "./table";


export default async function UixJeddahLogactMainpg() {
  const logact: MdlJeddahParamsActlog[] = await ApiJeddahDtbaseActlog();
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
