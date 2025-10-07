"use client";
import { useActionState, useEffect, useState } from "react";
import { ApiGlobalAllusrLogin } from "../../api/apiGlobalAllusr";

export default function UixGlobalLoginxFormdt() {
  const [formac, formacSet] = useActionState(ApiGlobalAllusrLogin, null);
  const [formdt, formdtSet] = useState({ usrnme: "", psswrd: "" });
  const [rspnse, rspnseSet] = useState({
    dfault: formac?.dfault || "",
    usrnme: formac?.usrnme || "",
    psswrd: formac?.psswrd || "",
    rspnse: formac?.rspnse || "",
  });

  // Monitor
  useEffect(() => {
    rspnseSet((prev) => ({
      ...prev,
      dfault: formac?.dfault || "",
      usrnme: formac?.usrnme || "",
      psswrd: formac?.psswrd || "",
      rspnse: formac?.rspnse || "",
    }));
  }, [formac]);

  // Function onchange
  const onchng = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { id, value } = e.target;
    formdtSet((prev) => ({
      ...prev,
      [id as keyof typeof prev]: value,
    }));
    rspnseSet((prev) => ({
      ...prev,
      [id as keyof typeof prev]: "",
    }));
  };

  return (
    <div className="afull flexctr text-sky-900">
      <form
        className="w-80 md:w-96 min-w-fit h-80 min-h-fit"
        action={formacSet}
      >
        <div className="afull flexctr flex-col bg-gradient-to-br from-white-800 via-sky-200 to-white-800 rounded-xl shadow-lg">
          <div className="w-3/4 py-1.5 flexctr flex-col">
            <div className="font-black text-2xl">LOGIN</div>
            <div className="font-semibold">Data Analyst IC</div>
          </div>
          <div className="w-1/2 py-1.5 flexctr flex-col relative">
            <label
              className={`px-1.5 afull font-semibold text-sky-800/50 relative flexstr`}
              htmlFor="usrnme"
            >
              <div className="opacity-0">Username</div>
              <div
                className={`absolute opacity-100 cursor-text ${formdt.usrnme.length > 0
                  ? "translate-y-0 pt-0 pb-1"
                  : "translate-y-full pt-1 pb-0"
                  } duration-300`}
              >
                Username
              </div>
            </label>
            <input
              className="afull bg-white text-slate-800 p-1.5 rounded-md"
              defaultValue={rspnse.dfault}
              type="text"
              name="usrnme"
              id="usrnme"
              onChange={(e) => onchng(e)}
            />
            <div className="w-full absolute -bottom-2.5 flexend text-red-500 text-[0.65rem] px-1.5 pt-0.5">
              {rspnse.usrnme}
            </div>
          </div>
          <div className="w-1/2 py-1.5 flexctr flex-col relative">
            <label
              className={`px-1.5 afull font-semibold text-sky-800/50 relative flexstr`}
              htmlFor="psswrd"
            >
              <div className="opacity-0">Password</div>
              <div
                className={`absolute opacity-100 cursor-text ${formdt.psswrd.length > 0
                  ? "translate-y-0 pt-0 pb-1"
                  : "translate-y-full pt-1 pb-0"
                  } duration-300`}
              >
                Password
              </div>
            </label>
            <input
              className="afull bg-white text-slate-800 p-1.5 rounded-md"
              type="password"
              name="psswrd"
              id="psswrd"
              onChange={(e) => onchng(e)}
            />
            <div className="w-full absolute -bottom-2.5 flexend text-red-500 text-[0.65rem] px-1.5 pt-0.5">
              {rspnse.psswrd}
            </div>
          </div>
          <button
            className="w-1/2 h-16 py-3 flexctr cursor-pointer group relative"
            type="submit"
          >
            <div className="afull btnsbm flexctr">Login</div>
            <div className="w-full absolute -bottom-2.5 flexctr text-red-500 text-[0.65rem] px-1.5 pt-0.5">
              {rspnse.rspnse}
            </div>
          </button>
        </div>
      </form>
    </div>
  );
}
