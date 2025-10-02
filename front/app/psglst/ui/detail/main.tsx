import {
  ApiPsglstDtbaseEdtprm,
  ApiPsglstDtbaseDetail,
} from "../../api/apiPsglstDtbase";
import {
  MdlPsglstAllprmSrcprm,
  MdlPsglstEdtprmParams,
  MdlPsglstDetailParams,
} from "../../model/mdlPsglstParams";
import UixPsglstDetailSearch from "./search";
import UixPsglstDetailTablex from "./uixPsglstDetail";

export default async function UixPsglstDetailMainpg({
  trtprm,
}: {
  trtprm: MdlPsglstAllprmSrcprm;
}) {
  // await new Promise((r) => setTimeout(r, 2000));
  const detail: MdlPsglstDetailParams[] = await ApiPsglstDtbaseDetail();
  const edtprm: MdlPsglstEdtprmParams[] = await ApiPsglstDtbaseEdtprm();
  return (
    <>
      <UixPsglstDetailSearch trtprm={trtprm} />
      {detail.length > 0 ? (
        <UixPsglstDetailTablex detail={detail} edtprm={edtprm} />
      ) : (
        <div className="afull flexctr text-base font-semibold text-sky-800">
          No database Log Action
        </div>
      )}
    </>
  );
}
