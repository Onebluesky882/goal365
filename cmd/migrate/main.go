package main

import (
	"log"
	oddstoday "mytipster/internal/odds-today"
)

func main() {
	// ctx := context.Background()

	// database, err := db.NewDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	in := "/Users/onebluesky882/local_files/myjob/mytipster/bin/data.json"
	out := "/Users/onebluesky882/local_files/myjob/mytipster/bin/output.json"

	if err := oddstoday.ProcessOddsFile(in, out); err != nil {
		log.Fatal(err)
	}

	// if err := mytips.InsertData(ctx, database ,); err != nil {
	// 	// 	log.Fatal(err)
	// 	// }
	// }
	log.Println("migration finished ✅")
}
