"use client";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useEffect, useState } from "react";
import {
  UixGlobalIconvcGtotop,
  UixGlobalIconvcLogout,
  UixGlobalIconvcNextpg,
  UixGlobalIconvcPrevpg,
  UixGlobalIconvcUsrdtl,
} from "../server/uixGlobalIconvc";
import { MdlGlobalApplstDtbase } from "../../model/mdlGlobalApplst";
import { ApiGlobalAllusrLogout } from "../../api/apiGlobalLoginx";
import { mdlGlobalAllusrCookie } from "../../model/mdlGlobalAllusr";

export default function UixGlobalAppbarClient({
  cookie,
  applst,
}: {
  cookie: mdlGlobalAllusrCookie;
  applst: MdlGlobalApplstDtbase[];
}) {
  const pthnme = usePathname();
  const [lstpth, lstpthSet] = useState("xxxxxx");
  const [onclik, onclikSet] = useState(false);
  const [onhide, onhideSet] = useState(false);
  useEffect(() => {
    const segment = pthnme.split("/").filter(Boolean).pop();
    lstpthSet(segment || "");
  }, [pthnme]);

  return (
    <div className="fixed w-full h-12 bottom-4 flexctr px-3 z-30">
      <div
        className={`absolute right-0 w-10 h-full bg-sky-800 rounded-l-lg flexctr group cursor-pointer ${onhide ? "opacity-100" : "opacity-0"
          } duration-300`}
        onClick={() => onhideSet(false)}
      >
        <div className="group-hover:scale-125 duration-300 group-hover:-translate-x-1.5">
          <UixGlobalIconvcPrevpg color="white" size={2} bold={3} />
        </div>
      </div>
      <div
        className={`afull max-w-fit flexctr shadow-md shadow-slate-400 rounded-xl ${onhide ? "translate-x-[100rem]" : "translate-x-0"
          } duration-500 ease-in-out`}
      >
        <div
          className={`w-20 h-full flexctr group cursor-pointer p-1 bg-sky-800 relative ${cookie.usrnme == "" ? "rounded-xl" : "rounded-l-xl"
            }`}
        >
          <div
            className={`absolute bg-gradient-to-b from-sky-300 to-sky-800 rounded-b-xl ${onclik && cookie.usrnme != ""
              ? "w-full h-full opacity-100"
              : "w-full h-0 opacity-0"
              } duration-300`}
          ></div>
          <div
            className={`w-10 h-10 p-2 group-hover:bg-sky-500 rounded-full z-10 ${onclik ? "ring-4 ring-white" : ""
              } select-none duration-300`}
            onClick={() => onclikSet(!onclik)}
          >
            <Image
              className="invert"
              src="/lionairblack.png"
              width={1000}
              height={1000}
              alt=""
            />
          </div>
          <div
            className={`z-0 absolute left-0 bottom-12 bg-sky-300 rounded-t-xl p-1.5 ${onclik && cookie.usrnme != ""
              ? "w-[200%] h-[400%] opacity-100"
              : "w-0 h-0 opacity-0"
              } duration-300`}
          >
            <div className="w-full h-1/3 p-0.5 text-white font-semibold group/action">
              <div className="afull flexbtw btnsbm">
                <div className="whitespace-nowrap w-2/3 overflow-hidden px-0.5">
                  {cookie.stfnme}
                </div>
                <div>
                  <UixGlobalIconvcUsrdtl color="white" size={1.4} bold={3} />
                </div>
              </div>
            </div>
            <div className="w-full h-1/3 p-0.5 text-white font-semibold group/action">
              <div className="afull flexbtw btnsbm">
                <div className="whitespace-nowrap w-2/3 overflow-hidden px-0.5">
                  Go To Top
                </div>
                <div>
                  <UixGlobalIconvcGtotop color="white" size={1.4} bold={3} />
                </div>
              </div>
            </div>
            <div
              className="w-full h-1/3 p-0.5 text-white font-semibold group/action"
              onClick={() => ApiGlobalAllusrLogout()}
            >
              <div className="afull flexbtw btnsbm">
                <div className="whitespace-nowrap w-2/3 overflow-hidden px-0.5">
                  Logout
                </div>
                <div>
                  <UixGlobalIconvcLogout color="white" size={1.4} bold={3} />
                </div>
              </div>
            </div>
          </div>
        </div>
        {cookie.usrnme == "" ? (
          ""
        ) : (
          <div className="afull bg-gradient-to-r from-sky-300 to-emerald-300 overflow-x-auto rounded-r-xl">
            <div className="w-fit h-full flexctr text-white">
              {applst.map((item, idx) =>
                cookie.access.includes(item.prmkey) ? (
                  <Link
                    className="w-24 md:w-28 md:text-base md:font-semibold h-full flexctr p-1.5 group"
                    href={"/" + item.prmkey}
                    key={idx}
                  >
                    <div
                      className={`afull group-hover:bg-sky-800 rounded-lg flexctr ${lstpth == item.prmkey
                        ? "ring-2 ring-white shadow-lg shadow-slate-500"
                        : ""
                        } relative text-center duration-300`}
                    >
                      <span className="absolute duration-300 group-hover:opacity-0">
                        {item.prmkey.toUpperCase()}
                      </span>
                      <span className="absolute duration-300 text-xs opacity-0 group-hover:opacity-100 px-3">
                        {item.detail}
                      </span>
                    </div>
                  </Link>
                ) : (
                  <div
                    className="w-24 md:w-32 md:text-base md:font-semibold h-full flexctr p-3 group"
                    key={idx}
                  >
                    <div className="text-slate-300 select-none">
                      {item.prmkey.toUpperCase()}
                    </div>
                  </div>
                )
              )}
              <div
                className="w-10 md:w-16 md:text-base md:font-semibold h-full flexctr p-3 group cursor-pointer"
                onClick={() => onhideSet(!onhide)}
              >
                <div className="group-hover:scale-125 duration-300 group-hover:translate-x-1.5">
                  <UixGlobalIconvcNextpg color="white" size={2} bold={3} />
                </div>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
