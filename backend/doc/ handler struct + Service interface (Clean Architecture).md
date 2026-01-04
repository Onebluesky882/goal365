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

Handler
↓
AnalyticService (interface)
↓
analyticsService (struct)
↓
bun.DB

หน้าที่
• บอกว่า Service นี้ทำอะไรได้บ้าง
• ไม่สนใจ implementation
• ไม่รู้จัก bun.DB
• ใช้สำหรับ:
• handler
• mock ใน unit test

```go
// ! ตัวอย่าง
// AnalyticService คือ interface (สัญญา / abstraction) abstract class
type AnalyticService interface {
	InsertManual(ctx context.Context, item *m.MyAnalytics) error
	InsertMany(ctx context.Context, items []m.MyAnalytics) error
	PredictionByDay(ctx context.Context, date string) ([]m.MyAnalytics, error)
}

// constructor / factory
// (เชื่อม interface ↔ struct) คือค่า interface
func NewAnalyticService(db *bun.DB) AnalyticService {
	return &analyticsService{
		db: db,
	}
}

// inform receiver type
// คือ concrete implementation
type analyticsService struct {
	db *bun.DB
}
```
