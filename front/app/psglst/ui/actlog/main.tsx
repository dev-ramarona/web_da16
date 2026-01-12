import { MdlPsglstActlogDtbase } from "../../model/mdlPsglstParams";
import UixPsglstActlogTablex from "./tablex";

export default async function UixPsglstActlogMainpg({ actlog }: { actlog: MdlPsglstActlogDtbase[] }) {

  return (
    <>
      {actlog.length > 0 ? (
        <UixPsglstActlogTablex actlog={actlog} />
      ) : (
        <div className="w-full h-fit flexctr text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
    </>
  );
}
