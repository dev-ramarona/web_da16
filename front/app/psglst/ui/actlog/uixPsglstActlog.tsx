import { FncGlobalFormatDatefm } from "@/app/global/function/fncGlobalFormat";
import { MdlPsglstActlogDtbase } from "../../model/mdlPsglstParams";

export default function UixPsglstActlogTablex({
  actlog,
}: {
  actlog: MdlPsglstActlogDtbase[];
}) {
  return (
    <>
      <div className="afull max-h-fit overflow-auto rounded-lg ring-2 ring-sky-800">
        <table className="w-full">
          <thead className="sticky top-0 z-10 text-white">
            <tr>
              {actlog && actlog.length > 0
                ? Object.entries(actlog[0]).map(([key]) =>
                  key != "prmkey" ? (
                    <th key={key} className="thhead">
                      {key}
                    </th>
                  ) : (
                    ""
                  )
                )
                : ""}
            </tr>
          </thead>
          <tbody className="text-slate-700 bg-sky-100">
            {actlog.map((log, idx) => (
              <tr className="h-8 group" key={idx}>
                {Object.entries(log).map(([key, val]) => (
                  <td className="tdbody text-center" key={key}>
                    {["dateup", "datenb", "timeup"].includes(key)
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
