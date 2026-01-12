import { FncGlobalFormatDatefm } from "@/app/global/function/fncGlobalFormat";
import { MdlJeddahParamsActlog } from "../../model/mdlJeddahMainpr";

export default function UixJeddahActlogTablex({
  frbase,
}: {
  frbase: MdlJeddahParamsActlog[];
}) {
  const header = {
    datenb: "Date Flown",
    dateup: "Date Upload",
    statdt: "Status Data",
  };
  return (
    <>
      <div className="afull max-h-fit overflow-auto rounded-lg ring-2 ring-sky-800">
        <table className="w-full">
          <thead className="sticky top-0 z-10 text-white">
            <tr>
              {frbase && frbase.length > 0
                ? Object.entries(frbase[0]).map(([key]) =>
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
            {frbase.map((log, idx) => (
              <tr className="h-8 group" key={idx}>
                {Object.entries(log).map(([key, val]) => (
                  <td className="tdbody text-center" key={key}>
                    {key == "datenb" || key == "dateup"
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
