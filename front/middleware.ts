import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { ApiGlobalAxiospParams } from "./app/global/api/apiGlobalPrimer";
import { mdlGlobalAllusrCookie } from "./app/global/model/mdlGlobalPrimer";

export async function middleware(req: NextRequest) {
  const tknnme = process.env.NEXT_PUBLIC_TKN_COOKIE || "x"
  const tokenx = req.cookies.get(tknnme)?.value || "";
  const pathnm = req.nextUrl.pathname.split("/")[1];

  // Jika belum login, arahkan ke "/"
  if (tokenx == "" || !tokenx) {
    return NextResponse.redirect(new URL("/", req.url));
  }

  // Try hit API
  try {

    // Validasi token ke backend
    const rspnse = await ApiGlobalAxiospParams.get("/allusr/tokenx", {
      headers: {
        Authorization: tokenx,
      },
    });

    // Jika token tidak valid, kembalikan ke halaman login "/"
    if (rspnse.status !== 200)
      return NextResponse.redirect(new URL("/", req.url));

    // Jika user tidak memiliki akses ke path tsb, arahkan ke halaman default (misal /home atau /dashboard)
    const fnlobj: mdlGlobalAllusrCookie = await rspnse.data;
    if (pathnm && !fnlobj.access.includes(pathnm))
      return NextResponse.redirect(new URL("/global", req.url));
    return NextResponse.next();
  }

  // Cath error
  catch (error) {
    console.log(error);
    return NextResponse.redirect(new URL("/", req.url));
  }
}

// Menerapkan middleware ke semua route di bawah '/global'
export const config = {
  matcher: '/((?!$|_next|favicon.ico|.*\\.(?:png|jpg|jpeg|gif|svg|ico|webp|css|js)).*)',
};
