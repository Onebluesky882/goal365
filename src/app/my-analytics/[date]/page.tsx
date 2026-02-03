import MyAnalytics from "@/pre-load/MyAnalytices";
import React from "react";
type Props = {
  params: {
    date: string;
  };
};

const page = ({ params }: Props) => {
  const { date } = params;
  return (
    <div>
      <MyAnalytics params={{date}} />
    </div>
  );
};

export default page;
