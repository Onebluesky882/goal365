import React from "react";
import { Atom } from "react-loading-indicators";

const LoadingIndicators = () => {
  return (
    <div className="fixed  w-full h-full  inset-0 z-50 flex items-center justify-center bg-black/5">
      <Atom color="#477AC6" size="small" text="" textColor="" />
    </div>
  );
};

export default LoadingIndicators;
