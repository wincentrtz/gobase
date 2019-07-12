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
	schema := migrations.Schema()

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully migrate user table")
}
