import {
  ApiPsglstDtbaseEdtprm,
  ApiPsglstDtbasePsgdtl,
} from "../../api/apiPsglstDtbase";
import {
  MdlPsglstAllprmSrcprm,
  MdlPsglstEdtprmParams,
  MdlPsglstPsgdtlParams,
} from "../../model/mdlPsglstParams";
import UixPsglstPsgdtlSearch from "./search";
import UixPsglstPsgdtlTablex from "./uixPsglstPsglst";

export default async function UixPsglstPsgdtlMainpg({
  trtprm,
}: {
  trtprm: MdlPsglstAllprmSrcprm;
}) {
  // await new Promise((r) => setTimeout(r, 2000));
  const psgdtl: MdlPsglstPsgdtlParams[] = await ApiPsglstDtbasePsgdtl();
  const edtprm: MdlPsglstEdtprmParams[] = await ApiPsglstDtbaseEdtprm();
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
