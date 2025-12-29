package oddstoday

import (
	"encoding/json"
	"fmt"
	"log"
	oddsmap "mytipster/lib/odds_map"
	odds_models "mytipster/models/odds"
	"os"
)

func WriteJSON(path string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// อ่าน JSON ที่เป็น map structure
func ReadOddsMap(path string) (odds_models.OddsMap, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var oddsMap odds_models.OddsMap
	if err := json.Unmarshal(data, &oddsMap); err != nil {
		return nil, err
	}

	return oddsMap, nil
}

// ฟังก์ชันหลักสำหรับ process odds file
func ProcessOddsFile(inputPath, outputPath string) error {
	// อ่านข้อมูล
	oddsMap, err := ReadOddsMap(inputPath)
	if err != nil {
		return fmt.Errorf("error reading odds map: %w", err)
	}

	log.Printf("✅ โหลดข้อมูลสำเร็จ: %d fixtures\n", len(oddsMap))

	// กรองข้อมูลตามเงื่อนไข
	filtered := oddsmap.FilterOddsMap(oddsMap)

	log.Printf("\n🎯 ผลการกรอง:ที่ตรงเงื่อนไข  %d fixtures \n", len(filtered))

	// เขียนผลลัพธ์
	if err := WriteJSON(outputPath, filtered); err != nil {
		return fmt.Errorf("error writing output: %w", err)
	}

	log.Printf("✅ บันทึกผลลัพธ์ที่: %s\n", outputPath)
	return nil
}
