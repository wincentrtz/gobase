package db

import (
	"fmt"

	"github.com/wincentrtz/gobase/migrations"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	for _, v := range migrations.MigrationList() {
		db.AutoMigrate(v)
	}

	fmt.Printf("Successfully migrate tables\n")
}
