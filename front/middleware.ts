import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { ApiGlobalAxiospParams } from "./app/global/api/apiGlobalPrimer";

export async function middleware(req: NextRequest) {
  const tokenx = req.cookies.get("tokenx")?.value || "";
  const nowusr = req.cookies.get("nowusr")?.value || "";

  // Jika belum login, arahkan ke "/"
  if (!tokenx || !nowusr) {
    return NextResponse.redirect(new URL("/", req.url));
  }

  try {
    // Parse userx cookie
    const Objusr: {
      stfnme: string;
      access: [string];
      keywrd: [string];
    } = JSON.parse(nowusr);

    // Validasi token ke backend
    const response = await ApiGlobalAxiospParams.get("/allusr/tokenx", {
      headers: {
        Authorization: tokenx,
      },
    });

    // Jika token tidak valid, kembalikan ke halaman login "/"
    if (response.status !== 200)
      return NextResponse.redirect(new URL("/", req.url));

    // Ambil segment pertama dari URL (misal: /opclss → "opclss")
    const path = req.nextUrl.pathname.split("/")[1];

    // Jika user tidak memiliki akses ke path tsb, arahkan ke halaman default (misal /home atau /dashboard)
    if (path && !Objusr.access.includes(path))
      return NextResponse.redirect(new URL("/global", req.url));
    return NextResponse.next();
  } catch (error) {
    console.log(error);
    return NextResponse.redirect(new URL("/", req.url));
  }
}

// Menerapkan middleware ke semua route di bawah '/global'
export const config = {
  matcher: ["/((?!_next|favicon.ico|api|$).*)"],
};
