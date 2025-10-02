import { FncGlobalFormatDatefm } from "@/app/global/function/fncGlobalFormat";
import { MdlPsglstErrlogParams } from "../../model/mdlPsglstParams";
import {
  UixGlobalIconvcIgnore,
  UixGlobalIconvcRfresh,
} from "@/app/global/ui/server/uixGlobalIconvc";

export default function UixPsglstErrlogTablex({
  errlog,
}: {
  errlog: MdlPsglstErrlogParams[];
}) {
  return (
    <>
      <div className="afull max-h-fit overflow-auto rounded-lg ring-2 ring-sky-800">
        <table className="w-full">
          <thead className="sticky top-0 z-10 text-white">
            <tr>
              <th className="thhead">Action</th>
              {errlog && errlog.length > 0
                ? Object.entries(errlog[0]).map(([key]) => (
                    <th key={key} className="thhead">
                      {key}
                    </th>
                  ))
                : ""}
            </tr>
          </thead>
          <tbody className="text-slate-700 bg-sky-100">
            {errlog.map((log, idx) => (
              <tr className="h-8 group" key={idx}>
                <td className="tdbody text-center">
                  <div className="afull flexctr gap-x-1.5">
                    <div className="w-1/2 flexctr btnsbm duration-300 cursor-pointer">
                      <UixGlobalIconvcRfresh
                        bold={3}
                        color="#53eafd"
                        size={1.4}
                      />
                    </div>
                    <div className="w-1/2 flexctr btnsbm duration-300 cursor-pointer">
                      <UixGlobalIconvcIgnore
                        bold={3}
                        color="#ffd230"
                        size={1.4}
                      />
                    </div>
                  </div>
                </td>
                {Object.entries(log).map(([key, val]) => (
                  <td className="tdbody text-center" key={key}>
                    {["datefl", "timeup"].includes(key)
                      ? FncGlobalFormatDatefm(String(val))
                      : val}
                  </td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  );
}
