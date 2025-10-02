import { mdlGlobalAllusrCookie } from "@/app/global/model/mdlGlobalPrimer";
import UixJeddahFlnbflUpload from "./upload";
import UixJeddahFlnbflInputx from "./input";

export default async function UixJeddahAddflnMainpg({ cookie }: {
  cookie: mdlGlobalAllusrCookie
}) {
  return (
    <div className="w-full h-fit flexctr">
      <UixJeddahFlnbflUpload cookie={cookie} />
      <UixJeddahFlnbflInputx cookie={cookie} />
    </div>
  );
}
