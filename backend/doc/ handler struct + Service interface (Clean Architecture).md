```go
type Service interface {
InsertPicked(items []bets_models.MyBets, analyticsID uuid.UUID, ctx context.Context) error
}

type handler struct {
service Service
}

func NewHandler(service Service) \*handler {
return &handler{service: service}
}
```

1️⃣ Handler “ไม่ควรรู้” ว่าข้อมูลมาจากไหน

“ฉันเรียก InsertPicked() ได้ก็พอ
ไม่สนใจว่าข้างในทำอะไร”

2️⃣ Interface = สัญญา (Contract)
“อะไรก็ตามที่ทำตามสัญญานี้
ถือว่าเป็น Service ได้”

3️⃣ ทำไม handler ต้องเป็น struct?
ถ้าเป็น function ธรรมดา
func InsertPicked(c \*fiber.Ctx) error {
InsertPickedService(...)
}
🔴 ปัญหา:
• inject dependency ไม่ได้
• mock ไม่ได้
• global variable เต็มไปหมด

4️⃣ ทำไมต้องมี NewHandler(...)
func NewHandler(service Service) \*handler {
return &handler{service: service}
}

    เหตุผลหลัก 3 ข้อ

✅ 1. บังคับ inject dependency
✅ แบบนี้ปลอดภัย: h := NewHandler(service)
✅ 2. เปลี่ยน implementation ได้โดยไม่แตะ handler
NewHandler(&BetsService{db})
NewHandler(&MockService{})
✅ 3. อ่านแล้วรู้ทันทีว่า handler ต้องการอะไร
NewHandler(service)
