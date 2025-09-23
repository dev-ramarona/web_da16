import { redirect } from "next/navigation";
import { ApiGlobalCookieGetdta } from "./global/api/apiCookieParams";
import UixGlobalLoginxFormdt from "./global/ui/client/uixGlobalLoginx";

export default async function Page() {
  const cookie = await ApiGlobalCookieGetdta();
  if (cookie.usrnme != "") redirect("/global");

  return (
    <div className="w-screen h-screen fixed top-0 flexctr bg-gradient-to-br from-sky-200 via-white to-emerald-100">
      <UixGlobalLoginxFormdt />
    </div>
  );
}
