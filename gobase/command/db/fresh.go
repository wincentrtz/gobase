package db

import (
	"gorm.io/gorm"
)

func Fresh(db *gorm.DB) {
	Drop(db)
	Migrate(db)
}
