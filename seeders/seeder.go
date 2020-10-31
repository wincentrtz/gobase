package seeders

import "gorm.io/gorm"

type Seeder interface {
	SeedTable(db *gorm.DB)
}
