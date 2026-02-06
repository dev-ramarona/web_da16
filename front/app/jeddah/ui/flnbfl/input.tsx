'use client'
import UixGlobalInputxFormdt from "@/app/global/ui/client/uixGlobalInputx";
import { MdlJeddahFlnbflDtbase } from "../../model/mdlJeddahMainpr";
import { useState } from "react";
import { mdlGlobalAlluserFilter, mdlGlobalAllusrCookie } from "@/app/global/model/mdlGlobalPrimer";
import { FncGlobalFormatFilter, FncGlobalFormatRoutfl } from "@/app/global/function/fncGlobalFormat";
import { ApiJeddahflnbflUpdate } from "../../api/apiJeddahFlnbfl";

export default function UixJeddahFlnbflInputx({
    cookie,
}: {
    cookie: mdlGlobalAllusrCookie;
}) {

    // Input partial data
    const defipt: MdlJeddahFlnbflDtbase = {
        airlfl: "", datefl: "", flnbfl: "",
        routfl: "", fltype: "", updtby: cookie.usrnme,
    }
    const [inputx, inputxSet] = useState(defipt);
    const iptrsp: MdlJeddahFlnbflDtbase = {
        airlfl: "Airlines is Empty",
        datefl: "Flight Date is Empty",
        flnbfl: "Flight Number is Empty",
        routfl: "Route is Empty",
        fltype: "Flight Type is Empty",
        updtby: "Re Login",
    };

    // Action add Flight number
    const [inptrs, inptrsSet] = useState("Submit");
    const iptact = async () => {
        for (const [key, val] of Object.entries(inputx))
            if (val == "") return inptrsSet(iptrsp[key as keyof typeof iptrsp]);
        const updapi = await ApiJeddahflnbflUpdate(inputx);
        if (updapi) { inptrsSet("Success"); inputxSet(defipt) }
        else inptrsSet("Failed");
        setTimeout(() => inptrsSet("Submit"), 700);
    };

    // Onchange input data
    const iptcgh = (e: React.ChangeEvent<HTMLInputElement>) => {
        const filter: mdlGlobalAlluserFilter[] = [
            { keywrd: "OUT", output: "Outgoing" },
            { keywrd: "INC", output: "Incoming" },
            { keywrd: "OJD", output: "Non Jeddah" },
        ];
        const nameid = e.currentTarget.id;
        let valuex = e.currentTarget.value;
        if (nameid == "flnbfl") valuex = valuex.replace(/[^0-9]/g, "");
        else if (nameid == "routfl") valuex = FncGlobalFormatRoutfl(valuex);
        else if (nameid == "fltype")
            valuex = FncGlobalFormatFilter(valuex, filter);
        else valuex = valuex.toUpperCase();
        inputxSet((prev) => ({
            ...prev,
            [nameid]: valuex,
        }));
    };
    return (
        <>
            <div className="w-7/12 h-fit py-1 flexctr flex-wrap gap-y-3 border-r-2 border-sky-200">
                <div className="w-1/2 md:w-1/6 min-w-20 md:min-w-32 h-10 flexctr">
                    <UixGlobalInputxFormdt
                        typipt={"text"}
                        length={2}
                        queryx={"airlfl"}
                        params={inputx.airlfl}
                        plchdr="Airline"
                        repprm={iptcgh}
                        labelx=""
                    />
                </div>
                <div className="w-1/2 md:w-1/6 min-w-20 md:min-w-32 h-10 flexctr">
                    <UixGlobalInputxFormdt
                        typipt={"text"}
                        length={4}
                        queryx={"flnbfl"}
                        params={inputx.flnbfl}
                        plchdr="Flight Number"
                        repprm={iptcgh}
                        labelx=""
                    />
                </div>
                <div className="w-1/2 md:w-1/6 min-w-20 md:min-w-32 h-10 flexctr">
                    <UixGlobalInputxFormdt
                        typipt={"text"}
                        length={7}
                        queryx={"routfl"}
                        params={inputx.routfl}
                        plchdr="Route"
                        repprm={iptcgh}
                        labelx=""
                    />
                </div>
                <div className="w-1/2 md:w-1/6 min-w-20 md:min-w-32 h-10 flexctr">
                    <UixGlobalInputxFormdt
                        typipt={"text"}
                        length={undefined}
                        queryx={"fltype"}
                        params={inputx.fltype}
                        plchdr="Flight Type"
                        repprm={iptcgh}
                        labelx=""
                    />
                </div>
                <div className="w-1/2 md:w-1/6 min-w-20 md:min-w-32 h-10 flexctr">
                    <UixGlobalInputxFormdt
                        typipt={"date"}
                        length={undefined}
                        queryx={"datefl"}
                        params={inputx.datefl}
                        plchdr="Flight Date"
                        repprm={iptcgh}
                        labelx=""
                    />
                </div>
                <div className="w-1/2 md:w-1/6 min-w-20 md:min-w-32 h-10 flexctr">
                    <div className="afull flexctr p-1.5">
                        <div className="afull btnsbm flexctr" onClick={() => iptact()}>{inptrs}</div>
                    </div>
                </div>
            </div>
        </>
    );
}