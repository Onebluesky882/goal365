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
	err := os.MkdirAll(path, 0755)
	err = os.Chdir("/Users/onebluesky882/local_files/myjob/mytipster/backend")
	if err != nil {
		t.Fatal(err)
	}

	err = lib.WriteJSON(filepath.Join(path, "error_query_odds.json"), failed)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(filepath.Join(path, "error_query_odds.json")); err != nil {
		t.Fatal("file not created")
	}
}
