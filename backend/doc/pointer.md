ใช้ pointer เมื่อ “ต้องการอ้างถึงของเดียวกัน”
ใช้ value เมื่อ “อยากได้สำเนาที่ปลอดภัย”

## ตารางตัดสินใจเร็ว (เซฟไว้ใช้)

สถานการณ์
ใช้ pointer 1. ต้องแก้ค่า *T - ค่าใหญ่ *T - optional / nil

ไม่ใช้ pointer - slice / mapT - อ่านอย่างเดียว - int/string/bool - constructor

ใช้กับ method (เรื่องสำคัญ)

❗ กฎทอง:

ถ้า method ไหนใช้ pointer → ทุก method ใช้ pointer

```go

type User struct {
	Name string
}

func (u *User) SetName(n string) {
	u.Name = n
}

func (u *User) Print() { // ✔ ใช้ pointer เหมือนกัน
	fmt.Println(u.Name)
}
```


 mental model ที่อยากให้คุณจำ

pointer = shared reality
value = private copy

ถามตัวเองทุกครั้ง:
	•	เราอยาก share หรือ isolate?
	•	เราจะแก้มั้ย?
	•	ของใหญ่หรือเล็ก?

คำตอบจะชี้ pointer เอง