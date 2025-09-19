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
        prm[n] == "" ? fltprm.delete(q) : fltprm.set(q, prm[n])
      );
    } else if (!Array.isArray(qry) && !Array.isArray(prm)) {
      prm == "" ? fltprm.delete(qry) : fltprm.set(qry, prm);
    } else if (Array.isArray(qry) && !Array.isArray(prm))
      qry.forEach((q) => (prm == "" ? fltprm.delete(q) : fltprm.set(q, prm)));

    router.push(`${pathname}?${fltprm.toString()}`, { scroll: false });
  };
}
