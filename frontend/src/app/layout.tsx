import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import { AuthProvider } from "@/GlobalContext/auth-provider";
import Headers from "@/components/Header/Header";
import { ToastProvider } from "@/GlobalContext/Toast";
import Footer from "@/components/Footer";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Goal365 แทงบอลฟรี ราคาบอลสด วิเคราห์บอล",
  description: "Goal365 แทงบอลฟรี ราคาบอลสด วิเคราห์บอล",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="dark">
      <head>
        <link rel="icon" href="/favicon.svg" type="image/svg+xml"></link>
      </head>

      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased overflow-x-hidden`}
      >
        <AuthProvider>
          <ToastProvider>
            <Headers />
            <div className=" min-h-screen bg-primary-foreground/80 h-full pt-1 max-sm:pt-1">
              {children}
            </div>
            <Footer />
          </ToastProvider>
        </AuthProvider>
      </body>
    </html>
  );
}
