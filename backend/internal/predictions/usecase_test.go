package predictions

import (
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
	if err := WriteFailedPredictions(failed, date); err != nil {
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
