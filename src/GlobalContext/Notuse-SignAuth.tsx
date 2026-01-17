// // hooks/useSignAuth.tsx
// "use client";
// import { liff } from "@line/liff";
// import { authClient } from "@/lib/auth-client";
// import { useEffect, useState, useRef } from "react";

// export const useLogin = () => {
//   const [idToken, setIdToken] = useState<string | null>(null);
//   const [accessToken, setAccessToken] = useState<string | null>(null);
//   const hasRun = useRef(false);

//   // Init LIFF
//   useEffect(() => {
//     if (!process.env.LIFF_ID) return;

//     const init = async () => {
//       await liff.init({ liffId: process.env.LIFF_ID! });

//       // ✅ เช็คว่ามี code ใน URL หรือเปล่า (คือพึ่ง redirect กลับมาจาก LINE)
//       const urlParams = new URLSearchParams(window.location.search);
//       const hasCode = urlParams.has("code");

//       if (liff.isLoggedIn() && hasCode) {
//         setIdToken(liff.getIDToken());
//         setAccessToken(liff.getAccessToken());
//       }
//     };
//     init();
//   }, []);

//   // Authenticate - ทำงานแค่ครั้งเดียว
//   useEffect(() => {
//     if (!idToken || !accessToken || hasRun.current) return;
//     hasRun.current = true;
//     console.log("process.env.API_URL ", process.env.API_URL);

//     const auth = async () => {
//       try {
//         // ส่ง idToken ไป backend
//         const resp = await fetch(
//           `${process.env.API_URL}/api/auth/verify-liff`,
//           {
//             method: "POST",
//             headers: { "Content-Type": "application/json" },
//             body: JSON.stringify({ idToken }),
//           }
//         );

//         await authClient.signIn.social({
//           provider: "line",
//           idToken: { token: idToken, accessToken: accessToken },
//         });

//         await new Promise((r) => setTimeout(r, 1000));

//         // ✅ Redirect และ clear query params
//         window.location.href = "/";
//       } catch (err) {
//         console.error(err);
//         hasRun.current = false;
//       }
//     };
//     auth();
//   }, [idToken]);

//   const lineLogin = () => {
//     if (!liff.isLoggedIn()) {
//       liff.login();
//     }
//   };

//   const lineLogOut = async () => {
//     await authClient.signOut();
//     if (liff.isLoggedIn()) liff.logout();
//     window.location.href = "/line-login";
//   };

//   return { lineLogin, lineLogOut };
// };
