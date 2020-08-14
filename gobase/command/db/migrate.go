package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/wincentrtz/gobase/migrations"
)

func Migrate(db *sql.DB) {
	var wg sync.WaitGroup
	schemas := migrations.GetMigrations()
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
