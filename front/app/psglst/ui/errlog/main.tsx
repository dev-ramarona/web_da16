import { ApiPsglstErrlogDtbase } from "../../api/apiPsglstErrlog";
import { MdlPsglstErrlogDtbase } from "../../model/mdlPsglstParams";
import UixPsglstErrlogTablex from "./uixPsglstErrlog";


export default async function UixPsglstErrlogMainpg() {
  // await new Promise((r) => setTimeout(r, 2000));
  const errlog: MdlPsglstErrlogDtbase[] = await ApiPsglstErrlogDtbase();
  return (
    <>
      {errlog.length > 0 ? (
        <UixPsglstErrlogTablex errlog={errlog} />
      ) : (
        <div className="afull flexctr text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
    </>
  );
}
