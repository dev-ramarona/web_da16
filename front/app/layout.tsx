import type { Metadata } from "next";
import { Geist, Geist_Mono, Poppins } from "next/font/google";
import "./globals.css";
import UixGlobalAppbarClient from "./global/ui/client/uixGlobalAppbar";
import { ApiGlobalCookieGetdta } from "./global/api/apiCookieParams";
import { MdlGlobalApplstDtbase } from "./global/model/mdlGlobalApplst";
import { ApiGlobalAllusrApplst } from "./global/api/apiGlobalLoginx";
import UixGlobalHeaderClient from "./global/ui/client/uixGlobalHeader";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

const poppins = Poppins({
  weight: ["100", "200", "300", "400", "500", "600", "700", "800", "900"],
  variable: "--font-ubuntu",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Data Analyst Web",
  description: "Created by Data Analyst Lion Tower Internal Control",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const applst: MdlGlobalApplstDtbase[] = await ApiGlobalAllusrApplst();
  const cookie = await ApiGlobalCookieGetdta();
  return (
    <html lang="en">
      <body
        className={`${poppins.className} ${geistSans.className} ${geistMono.className} antialiased`}
      >
        <div className="w-screen max-w-full h-full max-h-full text-xs pb-16">
          <div className="afull fixed bg-gradient-to-br from-sky-200 via-white to-emerald-100 -z-50"></div>
          <div className="w-full h-16 flex justify-start items-end px-3">
            <UixGlobalHeaderClient applst={applst} />
          </div>
          {children}
          <UixGlobalAppbarClient cookie={cookie} applst={applst} />
        </div>
      </body>
    </html>
  );
}
