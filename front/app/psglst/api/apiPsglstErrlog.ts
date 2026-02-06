import { MdlPsglstErrlogSrcprm } from "../model/mdlPsglstParams";

// Function get jeddah database Errlog
export async function ApiPsglstErrlogDtbase(prmErrlog: MdlPsglstErrlogSrcprm) {
  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_URL_AXIOSB}/psglst/errlog/getall`,
      {
        method: "POST",
        body: JSON.stringify(prmErrlog),
        headers: {
          "Content-Type": "application/json",
        },
        next: {
          revalidate: 30,
          tags: [
            `errlog-${prmErrlog.pagenw_errlog}-${prmErrlog.erdvsn_errlog}-${prmErrlog.update_global}`,
          ],
        },
      },
    );
    if (!res.ok) throw new Error("Failed fetch errlog");
    return await res.json();
  } catch (err) {
    console.error(err);
    return { arrdta: [], totdta: 0 };
  }
}
