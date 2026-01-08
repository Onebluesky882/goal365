"use client";

import { useAuth } from "@/GlobalContext/auth-provider";
import Link from "next/link";
import { useRouter } from "next/navigation";

const Headers = () => {
  const router = useRouter();
  const { session, isLoading } = useAuth();
  if (isLoading) {
    return <div>Loading...</div>;
  }
  console.log("session :", session);
  const user = session?.user;
  const credit = 1000;
  return (
    <>
      <div className="bg-web-bg-dark border-b border-web-primary-dark shadow-md">
        <div className="container mx-auto px-6 py-4 flex items-center justify-between">
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

          <div className="flex items-center justify-end space-x-4   rounded-xl shadow-md w-full max-w-md">
            {/* Balance Section */}
            <div className="flex flex-col items-end">
              {user ? (
                <>
                  <span className="text-sm  text-custom-gray ">
                    ยอดเงินคงเหลือ
                  </span>
                  <p className="text-xl font-bold  text-custom-gray ">{`${credit} บาท`}</p>
                </>
              ) : (
                <span
                  onClick={() => router.push("/free-bet")}
                  className="italic text-gray-400 text-sm cursor-pointer"
                >
                  รับเครดิตฟรี
                </span>
              )}
            </div>

            {/* Divider */}
            <div className="w-px h-10 bg-gray-300" />

            {/* User Info or Join Button */}
            <div className="flex flex-col items-start">
              {user ? (
                <p className="text-sm text-custom-gray font-medium">
                  {user.name}
                </p>
              ) : (
                <button
                  onClick={() => router.push("/sign-in")}
                  className="bg-yellow-400 hover:bg-yellow-500 text-gray-800 font-semibold px-5 py-2 rounded-full shadow transition-all"
                >
                  Join Now
                </button>
              )}
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Headers;
