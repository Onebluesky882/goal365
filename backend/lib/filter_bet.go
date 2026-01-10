package lib

import (
	"fmt"
	"math"
	m "mytipster/models"
	"sort"
	"strconv"
	"strings"
)

// PairScore - คะแนนความใกล้เคียง 2.0 ของคู่ Home/Away
type PairScore struct {
	HomeValue m.Value
	AwayValue m.Value
	AvgDiff   float64
	HasHome   bool
	HasAway   bool
}

// FilterBetPairsNear2 - กรองคู่ของ bet (Home/Away หรือ Over/Under) ที่มี odd ใกล้ 2.0 มากที่สุด
func FilterBetPairsNear2(values []m.Value, topPairs int) []m.Value {
	if len(values) == 0 {
		return []m.Value{}
	}

	pairs := make(map[string]*PairScore)

	for _, v := range values {
		// แปลง v.Value เป็น string
		valStr := ""
		switch val := v.Value.(type) {
		case string:
			valStr = val
		default:
			valStr = fmt.Sprintf("%v", val)
		}

		if valStr == "" {
			continue
		}

		// แยกส่วน เช่น "Home -0.25" -> ["Home", "-0.25"]
		// หรือ "Over 2.5" -> ["Over", "2.5"]
		parts := strings.Fields(valStr)
		if len(parts) < 2 {
			continue
		}

		side := parts[0]     // "Home", "Away", "Over", "Under"
		handicap := parts[1] // "-0.25", "+0", "2.5", etc.

		// ใช้ handicap เป็น key
		pairKey := handicap

		pair, exists := pairs[pairKey]
		if !exists {
			pair = &PairScore{}
			pairs[pairKey] = pair
		}

		// จัดเก็บตาม side
		if side == "Home" || side == "Over" {
			pair.HomeValue = v
			pair.HasHome = true
		} else if side == "Away" || side == "Under" {
			pair.AwayValue = v
			pair.HasAway = true
		}
	}

	// คำนวณค่าเฉลี่ยความห่างจาก 2.0 ของแต่ละคู่
	var pairScores []PairScore
	for _, pair := range pairs {
		// ต้องมีทั้ง 2 ข้าง
		if !pair.HasHome || !pair.HasAway {
			continue
		}

		homeOdd, err1 := strconv.ParseFloat(pair.HomeValue.Odd, 64)
		awayOdd, err2 := strconv.ParseFloat(pair.AwayValue.Odd, 64)

		if err1 != nil || err2 != nil {
			continue
		}

		// คำนวณค่าเฉลี่ยความห่างจาก 2.0
		homeDiff := math.Abs(homeOdd - 2.0)
		awayDiff := math.Abs(awayOdd - 2.0)
		pair.AvgDiff = (homeDiff + awayDiff) / 2.0

		pairScores = append(pairScores, *pair)
	}

	// เรียงตามความใกล้เคียง 2.0 (น้อยที่สุด = ใกล้ที่สุด)
	sort.Slice(pairScores, func(i, j int) bool {
		return pairScores[i].AvgDiff < pairScores[j].AvgDiff
	})

	// จำกัดจำนวนคู่
	if len(pairScores) > topPairs {
		pairScores = pairScores[:topPairs]
	}

	// แปลงกลับเป็น []m.Value (Home, Away, Home, Away, ...)
	var result []m.Value
	for _, pair := range pairScores {
		result = append(result, pair.HomeValue, pair.AwayValue)
	}

	return result
}

// FilterMatchWinnerNear2 - กรองสำหรับ Match Winner (Home, Draw, Away)
func FilterMatchWinnerNear2(values []m.Value) []m.Value {
	if len(values) == 0 {
		return []m.Value{}
	}

	type valWithDiff struct {
		v    m.Value
		diff float64
	}

	var temp []valWithDiff
	for _, v := range values {
		odd, err := strconv.ParseFloat(v.Odd, 64)
		if err != nil {
			continue
		}
		diff := math.Abs(odd - 2.0)
		temp = append(temp, valWithDiff{v, diff})
	}

	// เรียงจากใกล้ 2.0 มากที่สุด
	sort.Slice(temp, func(i, j int) bool {
		return temp[i].diff < temp[j].diff
	})

	// จำกัดสูงสุด 3 ตัว
	top := 5
	if len(temp) < top {
		top = len(temp)
	}

	result := make([]m.Value, top)
	for i := 0; i < top; i++ {
		result[i] = temp[i].v
	}

	return result
}

 
