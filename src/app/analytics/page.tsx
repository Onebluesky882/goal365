"use client";

import { useRouter } from "next/navigation";
import { useEffect } from "react";
import { useAuth } from "@/GlobalContext/auth-provider";
import MyAnalytics from "@/pre-load/MyAnalytices";

export default function Page() {
  const { session, isLoading } = useAuth();
  const router = useRouter();

  const userAllow = ["wansing882@gmail.com"];

  useEffect(() => {
    if (!isLoading) {
      if (!userAllow.includes(session?.user?.email || "")) {
        router.replace("/"); // 🔥 redirect
      }
    }
  }, [session, isLoading]);

  if (isLoading) return <div>loading...</div>;

  return <MyAnalytics />;
}
