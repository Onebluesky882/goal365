"use client";

import { Toaster, toast } from "sonner";
import React, { createContext, useContext } from "react";

// Context สำหรับเรียก toast จากทุก component
const ToastContext = createContext<typeof toast | null>(null);

export const useToast = () => {
  const ctx = useContext(ToastContext);
  if (!ctx) throw new Error("useToast must be used within ToastProvider");
  return ctx;
};

export const ToastProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  return (
    <ToastContext.Provider value={toast}>
      {children}
      <Toaster position="bottom-right" richColors />
    </ToastContext.Provider>
  );
};
