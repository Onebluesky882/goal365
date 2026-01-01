package mytips

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFailedPredictions_CreateFile(t *testing.T) {
	// 1. สร้าง temp dir
	tmp := t.TempDir()

	// 2. เปลี่ยน working directory ให้ test คุมได้
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldWd)

	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}

	// 3. เตรียมข้อมูล
	date := "2099-01-01"
	failed := []int{1001, 1002, 1003}

	// 4. เรียก function ที่ต้องการ test
	if err := WriteFailedPredictions(failed); err != nil {
		t.Fatalf("WriteFailedPredictions error: %v", err)
	}

	// 5. ตรวจว่าไฟล์ถูกสร้างจริง
	expectedPath := filepath.Join(
		tmp,
		"bin",
		date,
		"error_query_prediction.json",
	)

	if _, err := os.Stat(expectedPath); err != nil {
		t.Fatalf("file not created: %s", expectedPath)
	}
}

func TestWriteFailedPredictions_DebugRealFile(t *testing.T) {
	// 🔴 FIXED PATH (ของจริง)
	projectRoot := "/Users/onebluesky882/local_files/myjob/mytipster/backend"

	// เปลี่ยน working directory ไปที่ backend
	if err := os.Chdir(projectRoot); err != nil {
		t.Fatal(err)
	}

	date := "2099-01-01"
	failed := []int{111, 222, 333}

	log.Println("DEBUG: writing failed predictions...")
	log.Println("DEBUG: date =", date)
	log.Println("DEBUG: failed =", failed)

	// เรียก function จริง
	if err := WriteFailedPredictions(failed); err != nil {
		t.Fatalf("WriteFailedPredictions error: %v", err)
	}

	// path ที่ควรจะมีไฟล์
	expectedPath := filepath.Join(
		projectRoot,
		"bin",
		date,
		"error_query_prediction.json",
	)

	log.Println("DEBUG: checking file:", expectedPath)

	if _, err := os.Stat(expectedPath); err != nil {
		t.Fatalf("❌ file not created: %v", err)
	}

	log.Println("✅ file created successfully")
}
