🔹 สำหรับ LINE LIFF + BetterAuth

ถ้าคุณอยากใช้ BetterAuth แบบมาตรฐานกับ LIFF:
	1.	Frontend ส่ง user ไป login ผ่าน /api/auth/signin/line (BetterAuth handler)
	2.	LINE login → redirect กลับ /api/auth/callback/line?code=XYZ
	3.	BetterAuth จะ handle:
	•	exchange code → id_token + access_token
	•	verify LINE JWT
	•	lookup user / create user
	•	generate JWT/session → เซ็ต cookie และ response ไป frontend
	4.	Frontend ใช้ JWT เรียก API

คุณไม่ต้องทำ verifyLineIDToken() ด้วยตัวเอง เพราะ BetterAuth handle อยู่แล้ว

⸻
