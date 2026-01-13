"use client";

import { authClient } from "@/lib/auth-client";
import { useRouter } from "next/navigation";
import React, { FormEvent, useEffect, useState } from "react";
import { toast } from "sonner";

function AuthForm() {
  const [isLogin, setIsLogin] = useState(true);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [name, setName] = useState("");
  const [loading, setLoading] = useState(false);

  const router = useRouter();

  const { data } = authClient.useSession();

  useEffect(() => {
    if (data?.session) {
      router.replace("/players");
    }
  }, [data, router]);

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!email || !password) {
      toast.error("กรุณากรอกข้อมูลให้ครบถ้วน");
      return;
    }

    if (password.length < 8) {
      toast.error("รหัสผ่านต้องมีความยาวอย่างน้อย 8 ตัวอักษร");
      return;
    }

    setLoading(true);

    try {
      if (isLogin) {
        const { error: signInError } = await authClient.signIn.email({
          email,
          password,
        });

        if (signInError) {
          toast.error(signInError.message || "Email หรือ Password ไม่ถูกต้อง");
        } else {
          toast.success("เข้าสู่ระบบสำเร็จ");
          router.push("/players");
        }
      } else {
        const { error: signUpError } = await authClient.signUp.email({
          email,
          password,
          name: name || email.split("@")[0],
        });

        if (signUpError) {
          toast.error(signUpError.message || "ไม่สามารถสมัครสมาชิกได้");
        } else {
          toast.success("สมัครสมาชิกสำเร็จ! กรุณาเข้าสู่ระบบ");
          setIsLogin(true);
        }
      }
    } catch {
      toast.error("เกิดข้อผิดพลาดบางอย่าง กรุณาลองใหม่");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={containerStyle}>
      <div style={cardStyle}>
        <h2 style={headerStyle}>{isLogin ? "เข้าสู่ระบบ" : "สมัครสมาชิก"}</h2>

        <form onSubmit={handleSubmit} style={formStyle}>
          {!isLogin && (
            <input
              type="text"
              placeholder="ชื่อของคุณ"
              value={name}
              onChange={(e) => setName(e.target.value)}
              style={inputStyle}
            />
          )}

          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            style={inputStyle}
            required
          />

          <input
            type="password"
            placeholder="Password (อย่างน้อย 8 ตัว)"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            style={inputStyle}
            required
          />

          <button
            type="submit"
            disabled={loading}
            style={isLogin ? loginButtonStyle : signUpButtonStyle}
          >
            {loading ? "กำลังประมวลผล..." : isLogin ? "Login" : "Sign Up"}
          </button>
        </form>

        <div style={toggleContainerStyle}>
          <p style={toggleTextStyle}>
            {isLogin ? "ยังไม่มีบัญชีใช่ไหม?" : "มีบัญชีอยู่แล้วใช่ไหม?"}{" "}
            <span onClick={() => setIsLogin(!isLogin)} style={toggleLinkStyle}>
              {isLogin ? "สมัครสมาชิกที่นี่" : "เข้าสู่ระบบที่นี่"}
            </span>
          </p>
        </div>
      </div>
    </div>
  );
}

/* ===== styles (เหมือนเดิม) ===== */

// --- Styles ---
const containerStyle: React.CSSProperties = {
  display: "flex",
  justifyContent: "center",
  alignItems: "center",
  minHeight: "100vh",
  background: "linear-gradient(135deg, #6b73ff 0%, #000dff 100%)",
  padding: "1rem",
};

const cardStyle: React.CSSProperties = {
  background: "#fff",
  padding: "2.5rem 2rem",
  borderRadius: "12px",
  boxShadow: "0 8px 24px rgba(0,0,0,0.15)",
  width: "100%",
  maxWidth: "420px",
};

const headerStyle: React.CSSProperties = {
  textAlign: "center",
  marginBottom: "1.5rem",
  fontSize: "1.8rem",
  color: "#333",
  fontWeight: "bold",
};

const errorBoxStyle: React.CSSProperties = {
  color: "#b00020",
  backgroundColor: "#fddede",
  padding: "12px",
  borderRadius: "6px",
  fontSize: "14px",
  marginBottom: "1rem",
  textAlign: "center",
};

const formStyle: React.CSSProperties = {
  display: "flex",
  flexDirection: "column",
  gap: "1rem",
};

const inputStyle: React.CSSProperties = {
  padding: "12px",
  borderRadius: "6px",
  border: "1px solid #ccc",
  fontSize: "16px",
};

const loginButtonStyle: React.CSSProperties = {
  padding: "12px",
  backgroundColor: "#0070f3",
  color: "#fff",
  border: "none",
  borderRadius: "6px",
  cursor: "pointer",
  fontWeight: "bold",
  fontSize: "16px",
};

const signUpButtonStyle: React.CSSProperties = {
  ...loginButtonStyle,
  backgroundColor: "#00c853",
};

const toggleContainerStyle: React.CSSProperties = {
  marginTop: "1.5rem",
  textAlign: "center",
};

const toggleTextStyle: React.CSSProperties = {
  fontSize: "14px",
  color: "#555",
};

const toggleLinkStyle: React.CSSProperties = {
  color: "#0070f3",
  cursor: "pointer",
  textDecoration: "underline",
};

export default AuthForm;
