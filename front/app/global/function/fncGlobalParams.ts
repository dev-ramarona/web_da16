"use client";

import { usePathname, useRouter, useSearchParams } from "next/navigation";

// Function edit params
export function FncGlobalParamsEdlink() {
  const searchParams = useSearchParams();
  const pathname = usePathname();
  const router = useRouter();

  return function (qry: string | string[], prm: string | string[]) {
    const fltprm = new URLSearchParams(searchParams);
    if (Array.isArray(qry) && Array.isArray(prm)) {
      qry.forEach((q, n) =>
        prm[n] == "" ? fltprm.delete(q) : fltprm.set(q, prm[n]),
      );
    } else if (!Array.isArray(qry) && !Array.isArray(prm)) {
      if (prm === "") fltprm.delete(qry);
      else fltprm.set(qry, prm);
    } else if (Array.isArray(qry) && !Array.isArray(prm))
      qry.forEach((q) => (prm == "" ? fltprm.delete(q) : fltprm.set(q, prm)));

    router.push(`${pathname}?${fltprm.toString()}`, { scroll: false });
  };
}

// Function create params date now until h-4
export function FncGlobalParamsHminfr(day: number): number[] {
  const dates: number[] = [];
  const today = new Date();
  for (let i = 0; i <= day; i++) {
    const d = new Date(today);
    d.setDate(today.getDate() - i);

    // Format: YYMMDD
    const yy = d.getFullYear().toString().slice(-2);
    const mm = (d.getMonth() + 1).toString().padStart(2, "0");
    const dd = d.getDate().toString().padStart(2, "0");
    dates.push(Number(`${yy}${mm}${dd}`));
  }

  return dates;
}
