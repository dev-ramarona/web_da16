import axios from "axios";
import { MdlGlobalStatusPrcess } from "../model/mdlGlobalPrimer";

// Global Axios
export const ApiGlobalAxiospParams = axios.create({
  baseURL: `${process.env.NEXT_PUBLIC_IPV_ADRESS}:${process.env.NEXT_PUBLIC_PRT_GOLANG}`,
  timeout: 1000000,
  headers: { "Content-Type": "application/json" },
});

// Hit status sabre api
export async function ApiGlobalStatusPrcess() {
  try {
    const rspnse = await ApiGlobalAxiospParams.get("/global/status");
    if (rspnse.status === 200) {
      const rawstr: MdlGlobalStatusPrcess = await rspnse.data;
      const fnlstr: MdlGlobalStatusPrcess = {
        action: Number(rawstr.action.toFixed(2)),
        sbrapi: Number(rawstr.sbrapi.toFixed(2))
      }
      return fnlstr;
    }
  } catch (error) {
    console.log(error);
  }
  const tmpstr: MdlGlobalStatusPrcess = { action: 0, sbrapi: 0, }
  return tmpstr
}

// Hit status api with interval time
export async function ApiGlobalStatusIntrvl(
  statfnSet: (v: string) => void,
  intrvlSet: (v: NodeJS.Timeout | null) => void,
  strVarble: "action" | "sbrapi"
) {
  const strtiv = setInterval(async () => {
    const status = await ApiGlobalStatusPrcess();
    const rawval = (strVarble == "action") ? status.action : status.sbrapi
    const nowval = Number(rawval.toFixed(2));
    const nowstr = (nowval == 0) ? "Done" : `Wait ${nowval}%`
    statfnSet(nowstr);
    if (nowstr === "Done") {
      clearInterval(strtiv);
      intrvlSet(null);
      statfnSet("Process Done");
      setTimeout(() => statfnSet("Done"), 1000);
    }
  }, 3000);
  intrvlSet(strtiv);
}
