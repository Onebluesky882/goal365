package oddstoday

import (
	"mytipster/lib"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFailedFixtures(t *testing.T) {
	tmp := t.TempDir()

	date := "2099-01-01"
	failed := []int{1, 2, 3}

	path := filepath.Join(tmp, "bin", date)
	if err := os.MkdirAll(path, 0755); err != nil {
		t.Fatal(err)
	}

	file := filepath.Join(path, "error_query_odds.json")
	if err := lib.WriteJSON(file, failed); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(file); err != nil {
		t.Fatal("file not created")
	}
}
