"use client";

import { useEffect, useState } from "react";
import { ApiJeddahPrcessManual } from "../../api/apiJeddahPrcess";
import {
  ApiGlobalStatusIntrvl,
  ApiGlobalStatusSbrapi,
} from "@/app/global/api/apiGlobalPrimer";

export default function UixJeddahPrcessManual() {
  // Get status first
  const apistt = ApiGlobalStatusSbrapi;
  const [statfn, statfnSet] = useState("Upload");
  const [intrvl, intrvlSet] = useState<NodeJS.Timeout | null>(null);
  useEffect(() => {
    const gtstat = async () => {
      const status = await ApiGlobalStatusSbrapi();
      statfnSet(status);
      if (status != "Done") {
        await ApiGlobalStatusIntrvl(statfnSet, intrvlSet, apistt);
      } else statfnSet("Upload");
    };
    gtstat();
  }, []);

  // Hit the database and get interval status
  const prcess = async () => {
    const status = await ApiGlobalStatusSbrapi();
    if (status == "Done") {
      statfnSet("Wait");
      ApiJeddahPrcessManual();
      await ApiGlobalStatusIntrvl(statfnSet, intrvlSet, apistt);
    } else statfnSet(status);
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
