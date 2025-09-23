import axios from "axios";

// Global Axios
export const ApiGlobalAxiospParams = axios.create({
  baseURL: `${process.env.NEXT_PUBLIC_IPV_ADRESS}:${process.env.NEXT_PUBLIC_PRT_GOLANG}`,
  timeout: 1000000,
  headers: { "Content-Type": "application/json" },
});

// Hit status sabre api
export async function ApiGlobalStatusSbrapi() {
  try {
    const rspnse = await ApiGlobalAxiospParams.get("/global/status");
    if (rspnse.status === 200) {
      const fnlstr: string = await rspnse.data;
      return fnlstr;
    }
  } catch (error) {
    console.log(error);
  }
  return "failed";
}

// Hit status api with interval time
export async function ApiGlobalStatusIntrvl(
  statfnSet: (v: string) => void,
  intrvlSet: (v: NodeJS.Timeout | null) => void,
  statusApi: () => Promise<string>
) {
  const strtiv = setInterval(async () => {
    const status = await statusApi();
    statfnSet(status);
    if (status === "Done") {
      clearInterval(strtiv);
      intrvlSet(null);
      statfnSet("Process Done");
      setTimeout(() => statfnSet("Upload"), 1000);
    }
  }, 3000);
  intrvlSet(strtiv);
}
