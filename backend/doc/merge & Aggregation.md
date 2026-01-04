Merge
รวมข้อมูลหลายชุด
Aggregation
รวมข้อมูลย่อย → ชุดใหญ่

result := make(map[int][]odds_models.Bet)

// Merge odds data
for k, bets := range oddsMap {
result[k] = append(result[k], bets...)
}
