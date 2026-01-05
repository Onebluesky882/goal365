"use client";
import React from "react";
import { useLogin } from "@/hooks/useSignAuth";

const LoginLine: React.FC = () => {
  const { lineLogin } = useLogin();

  return (
    <div className="flex flex-col items-center justify-center min-h-screen gap-4 bg-gray-50">
      <div className="bg-white p-8 rounded-lg shadow-lg text-center">
        <h2 className="text-3xl font-bold mb-4">Welcome!</h2>
        <p className="text-gray-600 mb-6">Please login to continue</p>
        <button
          onClick={lineLogin}
          className="px-8 py-3 bg-green-500 text-white rounded-lg hover:bg-green-600 transition font-semibold text-lg"
        >
          Login with LINE
        </button>
      </div>
    </div>
  );
};

export default LoginLine;
