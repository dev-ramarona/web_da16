import {
  MdlPsglstPsgdtlFrntnd,
  MdlPsglstPsgdtlSrcprm,
} from "../model/mdlPsglstParams";

// Function get psglst database
export async function ApiPsglstPsgdtlGetall(prmPsgdtl: MdlPsglstPsgdtlSrcprm) {
  const tag = [
    "psgdtl",
    prmPsgdtl.mnthfl_psgdtl,
    prmPsgdtl.datefl_psgdtl,
    prmPsgdtl.airlfl_psgdtl,
    prmPsgdtl.flnbfl_psgdtl,
    prmPsgdtl.depart_psgdtl,
    prmPsgdtl.routfl_psgdtl,
    prmPsgdtl.pnrcde_psgdtl,
    prmPsgdtl.tktnfl_psgdtl,
    prmPsgdtl.isitfl_psgdtl,
    prmPsgdtl.isittx_psgdtl,
    prmPsgdtl.isitir_psgdtl,
    prmPsgdtl.nclear_psgdtl,
    prmPsgdtl.pagenw_psgdtl,
  ]
    .filter(Boolean)
    .join(":");
  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_URL_AXIOSB}/psglst/psgdtl/getall`,
      {
        method: "POST",
        body: JSON.stringify(prmPsgdtl),
        headers: { "Content-Type": "application/json" },
        next: { revalidate: 30, tags: [tag] },
      },
    );
    if (!res.ok) throw new Error("Failed fetch psgdtl");
    return await res.json();
  } catch (err) {
    console.error(err);
    return { arrdta: [], totdta: 0 };
  }
}

// Function get psglst database
export async function ApiPsglstPsgdtlUpdate(
  params: MdlPsglstPsgdtlFrntnd,
): Promise<string> {
  // Validation
  if (params.tktnvc === "") return "tktnvc empty";
  if (params.tktnvc.length !== 13) return "tktnvc invalid";
  if (params.airlvc === "") return "airlvc empty";
  if (params.flnbvc === "") return "flnbvc empty";
  if (params.cpnbvc === 0) return "cpnbvc empty";
  if (params.routvc === "") return "routvc empty";
  if (params.statvc === "") return "statvc empty";
  if (
    params.slsrpt === "NOT CLEAR" &&
    (params.curncy === "" || params.ntafvc === 0)
  ) {
    return "curncy empty";
  }

  // Call API
  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_URL_AXIOSB}/psglst/psgdtl/update`,
      {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(params),
        cache: "no-store",
      },
    );
    if (!res.ok) {
      throw new Error("Failed update psgdtl");
    }
    const data = await res.json();
    return await data;
  } catch (error) {
    console.error(error);
    return "update failed";
  }
}
