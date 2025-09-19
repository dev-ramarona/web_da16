import Link from "next/link";
import { ApiGlobalCookieGetdta } from "./api/apiCookieParams";
import { UixGlobalIconvcTolink } from "./ui/server/uixGlobalIconvc";

export default async function Page() {
  const cookie = await ApiGlobalCookieGetdta();
  return (
    <div className="afull flexctr flex-col text-sky-900 fixed top-0">
      <div className="text-3xl">
        Wellcome <span className="font-semibold">{cookie.stfnme}</span>
      </div>
      <div>You're only accepted on Page</div>
      <div className="w-2/3 md:w-1/3 flexctr flex-wrap">
        {cookie.access.map((item, index) => (
          <Link
            href={"/" + item}
            className="flexctr text-base group"
            key={index}
          >
            <div className="group-hover:scale-110 group-hover:rotate-45 duration-300">
              <UixGlobalIconvcTolink color="#024a70" size={1.1} bold={2} />
            </div>
            <div className="font-semibold">{item.toUpperCase()}</div>
            <div className="pl-0 pr-2">,</div>
          </Link>
        ))}
      </div>
      <div className="w-full flexctr flex-col text-center py-5">
        <div>for request Access or new User please confirm to email :</div>
        <div className="font-semibold">rama.rona@lionair.com</div>
      </div>
    </div>
  );
}
