package db

import (
	"fmt"
	"github.com/wincentrtz/gobase/seeders"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	for _, v := range seeders.SeederList() {
		r := v.(seeders.Seeder)
		r.SeedTable(db)
	}
	fmt.Printf("Successfully seed tables\n")
}
