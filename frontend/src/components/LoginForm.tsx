"use client";

import { authClient } from "@/lib/auth-client";
import React, { FormEvent, useState } from "react";

async function LoginForm() {
  // State สำหรับเก็บค่าของ input
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  // ฟังก์ชันจัดการ submit form
  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    // ตรวจสอบว่ากรอกครบหรือไม่
    if (!email || !password) {
      setError("กรุณากรอก Email และ Password");
      return;
    }

    // ตัวอย่าง simple validation ของ email
    const emailRegex = /\S+@\S+\.\S+/;
    if (!emailRegex.test(email)) {
      setError("กรุณากรอก Email ให้ถูกต้อง");
      return;
    }

    // ถ้าผ่าน validation
    setError("");
    console.log("Login with:", { email, password });
    alert(`Login success!\nEmail: ${email}`);

    // รีเซ็ต form
    setEmail("");
    setPassword("");
  };

  try {
    setError("");
    const result = await authClient.signIn.email({ email, password });
    console.log("Login success:", result);
    setEmail("");
    setPassword("");
  } catch (error) {
    console.error(error);
  }
  return (
    <div
      style={{
        maxWidth: "400px",
        margin: "50px auto",
        padding: "20px",
        border: "1px solid #ccc",
        borderRadius: "8px",
      }}
    >
      <h2>Login</h2>
      {error && <p style={{ color: "red" }}>{error}</p>}
      <form onSubmit={handleSubmit}>
        <div style={{ marginBottom: "15px" }}>
          <label>Email:</label>
          <br />
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Enter your email"
            style={{ width: "100%", padding: "8px", boxSizing: "border-box" }}
          />
        </div>
        <div style={{ marginBottom: "15px" }}>
          <label>Password:</label>
          <br />
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Enter your password"
            style={{ width: "100%", padding: "8px", boxSizing: "border-box" }}
          />
        </div>
        <button type="submit" style={{ padding: "10px 20px" }}>
          Login
        </button>
      </form>
    </div>
  );
}

export default LoginForm;
