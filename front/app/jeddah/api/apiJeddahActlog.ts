import { ApiGlobalAxiospParams } from "@/app/global/api/apiGlobalPrimer";

// Function get jeddah database log action
export async function ApiJeddahDtbaseActlog() {
    try {
        const rspnse = await ApiGlobalAxiospParams.get("/jeddah/actlog/getall");
        if (rspnse.status === 200) {
            const fnlobj = await rspnse.data;
            return fnlobj;
        }
    } catch (error) {
        console.log(error);
    }
    return [];
}