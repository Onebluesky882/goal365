package lib

import (
	"encoding/json"
	"fmt"
	"log"
	oddsmap "mytipster/lib/odds_map"
	odds_models "mytipster/models/odds"
	"os"
	"path/filepath"
	"time"
)

func WriteJSON(path string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	return os.WriteFile(path, data, 0644)
}

// WriteJSONWithDate เขียนไฟล์ JSON โดยใช้ date เป็น folder
// เช่น bin/2024-12-30/data.json
func WriteJSONWithDate(filename string, v any) error {
	// สร้าง date string (YYYY-MM-DD)
	date := time.Now().Format("2006-01-02")

	// สร้าง path: bin/{date}/{filename}
	path := filepath.Join("bin", date, filename)

	return WriteJSON(path, v)
}

func WriteJSONWithCustomDate(date string, filename string, v any) error {
	// สร้าง path: bin/{date}/{filename}
	dir := fmt.Sprintf("bin/%s", date)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("cannot create dir %s: %w", dir, err)
	}

	filePath := fmt.Sprintf("%s/%s", dir, filename)
	return WriteJSON(filePath, v)
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
func ReadJson[T any](path string) (T, error) {
	var result T
	data, err := os.ReadFile(path)
	if err != nil {
		return result, err
	}

 
	if err := json.Unmarshal(data, &result); err != nil {
		return result, err
	}

	return result, nil
}

// Process single fixture odds with retry
func ProcessOddsFile(inputPath string) error {
	// อ่านข้อมูล
	oddsMap, err := ReadOddsMap(inputPath)
	if err != nil {
		return fmt.Errorf("error reading odds map: %w", err)
	}

	log.Printf("✅ โหลดข้อมูลสำเร็จ: %d fixtures\n", len(oddsMap))

	// กรองข้อมูลตามเงื่อนไข
	filtered := oddsmap.FilterOddsMap(oddsMap)

	log.Printf("\n🎯 ผลการกรอง:ที่ตรงเงื่อนไข  %d fixtures \n", len(filtered))

	// แยก date จาก path (bin/2024-12-30/odds_data.json -> 2024-12-30)
	dir := filepath.Dir(inputPath)
	date := filepath.Base(dir)
	// เขียนผลลัพธ์
	if err := WriteJSONWithCustomDate(date, "filtered_odds.json", filtered); err != nil {
		return fmt.Errorf("error writing output: %w", err)
	}

	return nil
}
