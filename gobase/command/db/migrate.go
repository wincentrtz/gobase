package db

import (
	"fmt"
	"log"

	"github.com/wincentrtz/gobase/gobase/config"
	"github.com/wincentrtz/gobase/migrations"
)

func Migrate() {
	db := config.InitDb()
	defer db.Close()
	schemas := migrations.GetMigrations()
	for _, v := range schemas {
		_, err := db.Exec(v)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Successfully migrate all table")
}
