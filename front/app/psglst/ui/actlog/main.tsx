
import { ApiPsglstActlogDtbase } from "../../api/apiPsglstActlog";
import { MdlPsglstActlogDtbase } from "../../model/mdlPsglstParams";
import UixPsglstActlogTablex from "./uixPsglstActlog";

export default async function UixPsglstActlogMainpg() {
  // await new Promise((r) => setTimeout(r, 2000));
  const actlog: MdlPsglstActlogDtbase[] = await ApiPsglstActlogDtbase();
  return (
    <>
      {actlog.length > 0 ? (
        <UixPsglstActlogTablex actlog={actlog} />
      ) : (
        <div className="afull flexctr text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
    </>
  );
}
