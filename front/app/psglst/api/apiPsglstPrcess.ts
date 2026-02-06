import { MdlPsglstErrlogDtbase } from "../model/mdlPsglstParams";

export async function ApiPsglstPrcessManual(params: MdlPsglstErrlogDtbase) {
  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_URL_AXIOSB}/psglst/prcess`,
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
    console.log(data);

    return await data;
  } catch (error) {
    console.error(error);
    return "update failed";
  }
}
