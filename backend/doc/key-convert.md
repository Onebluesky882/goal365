แปลง key
สมมุตว่า odds := map[int]OddValue

ผลลัพธ์คือ สร้าง map ใหม่ ที่ key เป็น string

stringOdds := make(odds_models.OddsMap, len(odds))
for k, v := range odds {
stringOdds[strconv.Itoa(k)] = v
}
