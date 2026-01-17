"use client";

import { authClient } from "@/lib/auth-client";
import { useRouter } from "next/navigation";
import React, { FormEvent, useEffect, useState } from "react";
import { toast } from "sonner";
import { Eye, EyeOff, ChevronRight, HelpCircle } from "lucide-react";
import { FloatingInput } from "../components/FloatingInput";

function AuthForm() {
  const [isLogin, setIsLogin] = useState(true);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [name, setName] = useState("");
  const [loading, setLoading] = useState(false);
  const [showPassword, setShowPassword] = useState(false);

  const router = useRouter();
  const { data: session } = authClient.useSession();

  useEffect(() => {
    if (session) {
      router.replace("/players");
    }
  }, [session, router]);

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!email || !password) {
      toast.error("กรุณากรอกข้อมูลให้ครบถ้วน");
      return;
    }
    setLoading(true);

    try {
      if (isLogin) {
        const { error } = await authClient.signIn.email({ email, password });
        if (error)
          toast.error(error.message || "Email หรือ Password ไม่ถูกต้อง");
        else {
          toast.success("เข้าสู่ระบบสำเร็จ");
          router.push("/players");
        }
      } else {
        const { error } = await authClient.signUp.email({
          email,
          password,
          name: name || email.split("@")[0],
        });
        if (error) toast.error(error.message || "ไม่สามารถสมัครสมาชิกได้");
        else {
          toast.success("สมัครสมาชิกสำเร็จ! กรุณาเข้าสู่ระบบ");
          setIsLogin(true);
        }
      }
    } catch {
      toast.error("เกิดข้อผิดพลาดบางอย่าง");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className=" flex items-center justify-center p-4 font-sans">
      <div className="bg-white w-full max-w-112.5 rounded-lg shadow-2xl overflow-hidden animate-in fade-in zoom-in duration-300">
        <div className="flex justify-between items-center px-4 h-14 border-b">
          <button
            onClick={() => router.back()}
            className="p-2 hover:bg-gray-100 rounded-full"
          >
            <svg
              className="w-6 h-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>

        <div className="p-6 md:p-8">
          <div className="flex items-center justify-between mb-8">
            <h1 className="text-2xl font-bold text-gray-900">
              {isLogin ? "เข้าสู่ระบบ" : "สมัครสมาชิก"}
            </h1>
            <div className="text-right">
              <p className="text-xs text-gray-500">
                {isLogin ? "ยังไม่ได้เป็นสมาชิก?" : "มีบัญชีอยู่แล้ว?"}
              </p>
              <button
                type="button"
                onClick={() => setIsLogin(!isLogin)}
                className="text-sm font-bold text-[#00acec] hover:underline"
              >
                {isLogin ? "สมัครสมาชิก" : "เข้าสู่ระบบ"}
              </button>
            </div>
          </div>

          <form onSubmit={handleSubmit}>
            {!isLogin && (
              <FloatingInput
                id="name"
                label="ชื่อของคุณ"
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
            )}

            <FloatingInput
              id="email"
              label="ยูสเซอร์เนม/อีเมล"
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />

            <div className="relative">
              <FloatingInput
                id="password"
                label="รหัสผ่าน"
                type={showPassword ? "text" : "password"}
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
              <button
                type="button"
                onClick={() => setShowPassword(!showPassword)}
                className="absolute right-3 top-4 text-gray-500 hover:text-black z-10"
              >
                {showPassword ? <EyeOff size={20} /> : <Eye size={20} />}
              </button>
            </div>

            <div className="flex justify-between items-center mt-2 mb-8">
              <button
                type="button"
                className="text-sm font-bold text-[#00acec] hover:underline"
              >
                ลืมบัญชี ?
              </button>
              <button
                type="submit"
                disabled={loading}
                className="bg-[#1a1a1a] text-white px-6 py-2 rounded shadow-md hover:bg-black transition-all flex items-center gap-2 font-bold"
              >
                {loading
                  ? "กำลังโหลด..."
                  : isLogin
                  ? "เข้าสู่ระบบ"
                  : "สมัครสมาชิก"}
                <ChevronRight size={18} />
              </button>
            </div>
          </form>

          <div className="mt-10 p-4 bg-gray-50 rounded-lg border border-gray-200">
            <div className="flex items-center gap-2 mb-2 text-gray-700 font-bold text-sm">
              <HelpCircle size={18} className="text-gray-400" />
              ดูวิธีเข้าสู่ระบบอื่น ๆ
            </div>
            <div className="flex flex-col gap-1 pl-7">
              <button className="text-left text-sm text-[#00acec] hover:underline font-semibold">
                วิธีการเชื่อมต่อบัญชี
              </button>
              <button className="text-left text-sm text-[#00acec] hover:underline font-semibold">
                เปิดใช้งาน Passkey
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default AuthForm;
