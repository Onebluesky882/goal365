import MyAnalytics from "@/pre-load/MyAnalytices";
import React from "react";

type props = {
  params: Promise<{ date: string }>;
};

export default async function Page({ params }: props) {
  const { date } = await params;

  return <MyAnalytics date={date} />;
}
