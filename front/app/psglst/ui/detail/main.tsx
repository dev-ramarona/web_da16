import { ApiPsglstAcpedtDtbase } from "../../api/apiPsglstAcpedt";
import { ApiPsglstPsgdtlGetall } from "../../api/apiPsglstPsgdtl";
import {
  MdlPsglstAcpedtDtbase,
  MdlPsglstSrcprmAllprm,
  MdlPsglstPsgdtlFrntnd,
} from "../../model/mdlPsglstParams";
import UixPsglstDetailSearch from "./search";
import UixPsglstDetailTablex from "./uixPsglstDetail";

export default async function UixPsglstDetailMainpg({
  trtprm,
}: {
  trtprm: MdlPsglstSrcprmAllprm;
}) {
  // await new Promise((r) => setTimeout(r, 2000));
  const detail: MdlPsglstPsgdtlFrntnd[] = await ApiPsglstPsgdtlGetall();
  const acpedt: MdlPsglstAcpedtDtbase[] = await ApiPsglstAcpedtDtbase();
  return (
    <>
      <UixPsglstDetailSearch trtprm={trtprm} />
      {detail.length > 0 ? (
        <UixPsglstDetailTablex detail={detail} acpedt={acpedt} />
      ) : (
        <div className="afull flexctr text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
    </>
  );
}
