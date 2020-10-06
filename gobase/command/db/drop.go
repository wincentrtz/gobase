package db

import (
	"fmt"

	"github.com/wincentrtz/gobase/migrations"
	"gorm.io/gorm"
)

func Drop(db *gorm.DB) {

	for _, v := range migrations.MigrationList() {
		db.Migrator().DropTable(v)
	}

	fmt.Printf("Successfully drop tables\n")
}
