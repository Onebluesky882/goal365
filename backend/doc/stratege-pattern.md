จะสร้างแบบนี้ ต้องคิดถึงอะไรก่อนเป็นอันดับแรก
map[string]BetSettleFunc คือการสร้าง map เพื่อ return function ใช่หรือไม่

คำตอบสั้นก่อน
• ✅ ใช่ — map[string]BetSettleFunc คือ Function Registry
✅ สิ่งที่ต้องคิด ก่อนเขียน map คือ Domain + Contract

🥇 อันดับที่ 1 — Domain & Output ต้องชัด

“ระบบนี้ตัดสินผลอะไร และผลลัพธ์มีอะไรบ้าง”

✅ ใช่ — มันคือ
• Strategy Registry
• Dispatch Table
• Function Lookup Table

มันทำหน้าที่:
• map market name → logic
• แทน switch-case
• เพิ่ม market ใหม่โดยไม่แก้ของเก่า
