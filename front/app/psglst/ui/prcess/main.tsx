import { ApiPsglstAcpedtDtbase } from "../../api/apiPsglstAcpedt";
import { ApiPsglstPsgdtlGetall } from "../../api/apiPsglstPsgdtl";
import { MdlPsglstAcpedtDtbase, MdlPsglstPsgdtlFrntnd, MdlPsglstSrcprmAllprm } from "../../model/mdlPsglstParams";
import UixPsglstPsgdtlSearch from "./search";
import UixPsglstPsgdtlTablex from "./uixPsglstPsglst";

export default async function UixPsglstPsgdtlMainpg({
  trtprm,
}: {
  trtprm: MdlPsglstSrcprmAllprm;
}) {
  // await new Promise((r) => setTimeout(r, 2000));
  const psgdtl: MdlPsglstPsgdtlFrntnd[] = await ApiPsglstPsgdtlGetall();
  const edtprm: MdlPsglstAcpedtDtbase[] = await ApiPsglstAcpedtDtbase();
  return (
    <>
      <UixPsglstPsgdtlSearch trtprm={trtprm} />
      {psgdtl.length > 0 ? (
        <UixPsglstPsgdtlTablex psgdtl={psgdtl} edtprm={edtprm} />
      ) : (
        <div className="afull flexctr text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
    </>
  );
}
