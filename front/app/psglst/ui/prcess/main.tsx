import { mdlGlobalAllusrCookie } from "@/app/global/model/mdlGlobalPrimer";
import UixPsglstPrcessManual from "./process";

export default function UixPsglstPrcessMainpg({ cookie }: { cookie: mdlGlobalAllusrCookie }) {
    return <UixPsglstPrcessManual cookie={cookie} />;
}