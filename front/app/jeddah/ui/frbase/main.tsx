import { ApiJeddahDtbaseActlog } from "../../api/apiJeddahActlog";
import { MdlJeddahParamsActlog } from "../../model/mdlJeddahMainpr";
import UixJeddahActlogTablex from "./table";


export default async function UixJeddahFrbaseMainpg() {
  const frbase: MdlJeddahParamsActlog[] = await ApiJeddahDtbaseActlog();
  return (
    <>
      {frbase.length > 0 ? (
        <UixJeddahActlogTablex frbase={frbase} />
      ) : (
        <div className="w-full h-fit flexctr text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
    </>
  );
}
