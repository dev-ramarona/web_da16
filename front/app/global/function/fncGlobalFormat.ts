import { mdlGlobalAlluserFilter } from "../model/mdlGlobalAllusr";

// Fucntion change format data yymmdd/hhmm to dd-MMM-yyyy hh:mm
export function FncGlobalFormatDatefm(inputd: string): string {
  if (inputd.length !== 6 && inputd.length !== 10) return "Invalid format";
  const yearnw = inputd.slice(0, 2);
  const monthn = inputd.slice(2, 4);
  const daynow = inputd.slice(4, 6);
  const yearfl = parseInt(yearnw) < 50 ? `20${yearnw}` : `19${yearnw}`;
  const hournw = inputd.length === 10 ? inputd.slice(6, 8) : null;
  const minute = inputd.length === 10 ? inputd.slice(8, 10) : null;

  // Buat Date object
  const datenw = new Date(
    `${yearfl}-${monthn}-${daynow}T${hournw ?? "00"}:${minute ?? "00"}`
  );
  const optons: Intl.DateTimeFormatOptions = {
    day: "2-digit",
    month: "short",
    year: "numeric",
  };
  const datetx = datenw.toLocaleDateString("en-GB", optons).replace(/ /g, "-");
  if (hournw && minute) return `${datetx} ${hournw}:${minute}`;
  return datetx;
}

// Fucntion change format data yymmdd/hhmm to dd-MMM-yyyy hh:mm
export function FncGlobalFormatDateip(inputd: string): string {
  if (inputd.length !== 6) return "Format harus YYMMDD";
  const year = parseInt(inputd.slice(0, 2), 10) + 2000; // "25" → 2025
  const month = inputd.slice(2, 4); // "07"
  const day = inputd.slice(4, 6); // "30"
  return `${year}-${month}-${day}`;
}

// Fucntion change format data dd-MMM-yyyy hh:mm to yymmdd
export function FncGlobalFormatIpdate(inputd: string): string {
  if (inputd.length !== 10) return "Invalid format";

  const yearnw = inputd.slice(0, 4);
  const monthn = inputd.slice(5, 7);
  const daynow = inputd.slice(8, 10);

  // Buat Date object
  const datetx = `${yearnw}-${monthn}-${daynow}`;
  return datetx;
}

// Function change format routef to 3-3 characters
export function FncGlobalFormatRoutfl(routef: string) {
  let raw = routef.toUpperCase().replace(/[^A-Z]/g, "");
  if (raw.length > 6) raw = raw.slice(0, 6);
  let formatted = raw;
  if (raw.length > 3) formatted = raw.slice(0, 3) + "-" + raw.slice(3);
  return formatted;
}

// Function change format routef to 3-3 characters
export function FncGlobalFormatPercnt(percent: string, prvprc: string) {
  if (!percent.includes("%")) {
    if (!prvprc.includes("%")) return percent + "%";
    else if (percent.length == 1) return "";
    return percent.substring(0, percent.length - 1) + "%";
  }
  let raw = percent.replace(/[^0-9]/g, "");
  return raw + "%";
}

// Function change format routef to 3-3 characters
export function FncGlobalFormatCpnfmt(cpnnbr: string) {
  if (cpnnbr === "") return cpnnbr;
  let raw = cpnnbr.toUpperCase().replace(/[^A-Z]/g, "");
  let nbr = parseInt(raw);
  if (isNaN(nbr)) return "";
  if (nbr < 10) {
    return `C0${nbr}`;
  }
  return `C${nbr}`;
}

// Function change format routef to 3-3 characters
export function FncGlobalFormatSorthl(params: string) {
  const raw = params.trim().toUpperCase();
  if (raw === "") return "";
  if (raw.length < 3 && /^[LOW]/i.test(raw)) return "Lowest";
  if (raw.length < 3 && /^[HIG]/i.test(raw)) return "Highest";
  else if (raw.length === 1) return "Highest";
  return ""; // default aman
}

// Function change format routef to 3-3 characters
export function FncGlobalFormatFilter(
  params: string,
  arrays: mdlGlobalAlluserFilter[]
) {
  const raw = params.trim().toUpperCase();
  if (raw === "") return "";
  for (let i = 0; i < arrays.length; i++) {
    const arr = arrays[i];
    const reg = new RegExp(`^[${arr.keywrd}]`, "i");
    if (raw.length <= arr.keywrd.length && reg.test(raw)) return arr.output;
  }
  if (raw.length === 1) return arrays[0].output;
  return ""; // default aman
}

// Function format arr split and cancel jeddah
export function FncGlobalFormatArrcpn(str: string) {
  if (!str.includes(":") && !str.includes("-")) return str;
  const val = str.split("|");
  const arr = [];
  for (let i = 0; i < val.length; i++) {
    const sep = val[i].includes(":") ? ":" : "-";
    const tmp = val[i].split(sep);
    const arx = [];
    for (let i = 0; i < tmp.length; i++) {
      const elm = tmp[i];
      if (Number.isInteger(Number(elm)) && elm.length >= 6)
        arx.push(FncGlobalFormatDatefm(elm));
      else arx.push(elm);
    }
    arr.push(arx.join("-"));
  }
  return arr.join(" | ");
}
