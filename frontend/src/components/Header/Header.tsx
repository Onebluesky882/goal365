"use client";

import { useAuth } from "@/GlobalContext/auth-provider";
import Link from "next/link";
import { useRouter } from "next/navigation";
import ProfileMenu from "./profileMenu";
import LoadingIndicators from "../Loading_indicators";
import { useEffect, useRef, useState } from "react";
const Headers = () => {
  const router = useRouter();
  const { session, isLoading } = useAuth();

  const user = session?.user;

  const lastScrollY = useRef(0);
  const [hideHeader, setHideHeader] = useState(false);

  useEffect(() => {
    const onScroll = () => {
      const currentY = window.scrollY;

      if (currentY > lastScrollY.current) {
        // scroll down
        setHideHeader(true);
      } else {
        // scroll up
        setHideHeader(false);
      }

      lastScrollY.current = currentY;
    };

    window.addEventListener("scroll", onScroll, { passive: true });
    return () => window.removeEventListener("scroll", onScroll);
  }, []);

  if (isLoading) {
    return <LoadingIndicators />;
  }

  return (
    <div
      className="sticky top-0 z-50  "
      style={{
        transform: hideHeader ? "translateY(-100%)" : "translateY(0)",
      }}
    >
      <div className="  mx-auto px-6 py-4 flex items-center justify-between   ">
        <div className="flex items-center space-x-6">
          <span className="text-web-primary font-extrabold text-3xl tracking-wide cursor-pointer hover:text-web-primary-dark transition-colors">
            <Link href={"/"}>Goal365</Link>
          </span>

          <nav className="hidden md:flex space-x-8  text-custom-gray  text-sm font-medium">
            {["Sports", "In-Play", "Casino", "Odds"].map((item) => (
              <Link key={item} href={`/${item.toLowerCase()}`}>
                <span
                  key={item}
                  className="cursor-pointer hover:text-web-primary transition-colors"
                >
                  {item}
                </span>
              </Link>
            ))}
          </nav>
        </div>

        <div className="flex items-center justify-end space-x-4 w-full max-w-md">
          <div className="flex flex-col items-start">
            {user ? (
              <>
                <ProfileMenu name={user.name} />
              </>
            ) : (
              <button
                onClick={() => router.push("sign-in")}
                className="bg-yellow-400 hover:bg-yellow-500 text-gray-800 font-semibold px-5 py-2 rounded-full shadow transition-all"
              >
                Join Now
              </button>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Headers;
