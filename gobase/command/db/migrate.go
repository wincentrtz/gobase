package db

import (
	"fmt"
	"sync"

	"github.com/wincentrtz/gobase/gobase/config"
	"github.com/wincentrtz/gobase/migrations"
)

func Migrate() {
	var wg sync.WaitGroup
	db := config.InitDb()
	defer db.Close()
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
