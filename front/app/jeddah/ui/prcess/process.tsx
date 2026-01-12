"use client";

import { useEffect, useState } from "react";
import { ApiJeddahPrcessManual } from "../../api/apiJeddahPrcess";
import {
  ApiGlobalStatusIntrvl,
  ApiGlobalStatusPrcess,
} from "@/app/global/api/apiGlobalPrimer";
import { MdlJeddahParamsActlog } from "../../model/mdlJeddahMainpr";

export default function UixJeddahPrcessManual() {
  // Get status first
  const [statfn, statfnSet] = useState("Upload");
  const [intrvl, intrvlSet] = useState<NodeJS.Timeout | null>(null);
  useEffect(() => {
    const gtstat = async () => {
      const status = await ApiGlobalStatusPrcess();
      statfnSet(status.sbrapi == 0 ? "Done" : `Wait ${status.sbrapi}%`);
      if (status.sbrapi != 0) {
        await ApiGlobalStatusIntrvl(statfnSet, intrvlSet, "sbrapi");
      } else statfnSet("Upload");
    };
    gtstat();
  }, []);

  // Hit the database and get interval status
  const prcess = async () => {
    const status = await ApiGlobalStatusPrcess();
    console.log(status.sbrapi);
    if (status.sbrapi == 0) {
      statfnSet("Wait");
      const params: MdlJeddahParamsActlog = {
        airlfl: "",
        timeup: 0,
        dateup: 0,
        statdt: "",
      }
      ApiJeddahPrcessManual(params);
      await ApiGlobalStatusIntrvl(statfnSet, intrvlSet, "sbrapi");
    } else statfnSet(`Wait ${status.sbrapi}%`);
  };

  return (
    <div className="afull flexctr flex-wrap">
      <button
        className="w-full h-full md:h-1/2 btnsbm flexctr"
        onClick={() => prcess()}
      >
        {statfn == "Upload" ? "Process Manual today" : statfn}
      </button>
    </div>
  );
}
