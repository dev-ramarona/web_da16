import { mdlGlobalAllusrCookie } from "@/app/global/model/mdlGlobalPrimer";
import UixJeddahFlnbflUpload from "./upload";
import UixJeddahFlnbflInputx from "./input";

export default async function UixJeddahAddflnMainpg({ cookie }: {
  cookie: mdlGlobalAllusrCookie
}) {
  return (
    <div className="w-full h-fit flexctr">
      <UixJeddahFlnbflUpload cookie={cookie} />
      <div className="h-36 md:h-16 w-1 flexctr bg-sky-800 rounded-full"></div>
      <UixJeddahFlnbflInputx cookie={cookie} />
    </div>
  );
}
