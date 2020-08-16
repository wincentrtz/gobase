package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/wincentrtz/gobase/migrations"
)

func Migrate(db *sql.DB, schemaName string) {
	var wg sync.WaitGroup
	var schemas []string
	if schemaName == "all" {
		schemas = migrations.GetMigrations()
	} else {
		schemas = migrations.GetMigrationFromSchemaName(schemaName)
	}

	if len(schemas) == 0 {
		log.Panic("Table is not found !")
	}

	for _, v := range schemas {
		wg.Add(1)
		go func(wg *sync.WaitGroup, v string) {
			defer wg.Done()
			db.Exec(v)
		}(&wg, v)
	}
	wg.Wait()
	fmt.Print("Successfully migrate all tables")
}
