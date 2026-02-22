"use client";

import { useAuth } from "@/GlobalContext/auth-provider";
import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import ProfileMenu from "./profileMenu";
import LoadingIndicators from "../Loading_indicators";
import { useEffect, useRef, useState } from "react";
import { CgMenuGridO } from "react-icons/cg";

export const HeaderComponent = () => {
  const pathname = usePathname();
  const [closeHeader, setCloseHeader] = useState(false);

  useEffect(() => {
    if (pathname.startsWith("/sportbook")) {
      setCloseHeader(true);
    } else {
      setCloseHeader(false);
    }
  }, [pathname]);

  return <>{!closeHeader && <Headers />}</>;
};

const Headers = () => {
  const router = useRouter();

  const { session, isLoading } = useAuth();

  const user = session?.user;

  const lastScrollY = useRef(0);
  const [hideHeader, setHideHeader] = useState(false);
  const SCROLL_THRESHOLD = 280;
  useEffect(() => {
    const onScroll = () => {
      const currentY = window.scrollY;
      const diff = currentY - lastScrollY.current;
      if (diff > SCROLL_THRESHOLD) {
        // scroll down
        setHideHeader(true);
        lastScrollY.current = currentY;
      }
      if (diff < -SCROLL_THRESHOLD) {
        setHideHeader(false);
        lastScrollY.current = currentY;
      }
    };

    window.addEventListener("scroll", onScroll, { passive: true });
    return () => window.removeEventListener("scroll", onScroll);
  }, []);

  if (isLoading) {
    return <LoadingIndicators />;
  }
  const format = (d: Date) => d.toLocaleDateString("en-CA");

  const todayDate = new Date();

  const today = format(todayDate);

  const yesterdayDate = new Date(todayDate);
  yesterdayDate.setDate(todayDate.getDate() - 1);

  const topMenuBar = [
    {
      name: "Today",
      path: "sportbook",
    },
    {
      name: "Coming Soon",
      path: "",
    },

    {
      name: "Live",
      path: "",
    },
    {
      name: "Analytics",
      path: `my-analytics`,
    },
    {
      name: "Mytips",
      path: `my-reviews?date=${today}&picked=true`,
    },
  ];
  return (
    <div
      className="sticky top-0 z-50 bg-background  "
      style={{
        transform: hideHeader ? "translateY(-100%)" : "translateY(0)",
      }}
    >
      <div className="  mx-auto px-2 py-2 flex items-center justify-between   ">
        <div className="flex items-center space-x-6">
          <span className="text-web-primary font-extrabold text-2xl tracking-wide cursor-pointer hover:text-web-primary-dark transition-colors">
            <Link href={"/"}>Goal365</Link>
          </span>

          <nav className="hidden md:flex space-x-8  text-custom-gray  text-sm font-medium">
            {topMenuBar.map((item) => (
              <Link key={item.name} href={`/${item.path}`}>
                <span
                  key={item.name}
                  className="cursor-pointer hover:text-web-primary transition-colors"
                >
                  {item.name}
                </span>
              </Link>
            ))}
          </nav>
        </div>

        <div className="flex items-center justify-end space-x-4 w-full max-w-md">
          <div className="flex gap-2 items-center ">
            {user ? (
              <>
                <ProfileMenu name={user.name} />
              </>
            ) : (
              <>
                <button
                  onClick={() => router.push("sign-in")}
                  className="text-sm  bg-yellow-400 hover:bg-yellow-500 text-gray-800 font-semibold px-4 py-1 rounded-md shadow transition-all"
                >
                  เข้าสู่ระบบ
                </button>
                <CgMenuGridO
                  className=" hover:border-blue-500 border rounded-sm  "
                  size={32}
                  color="white"
                />
              </>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Headers;
