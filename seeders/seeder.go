package seeders

import (
	"gorm.io/gorm"
)

type Seeder interface {
	SeedTable(db *gorm.DB)
}

func SeederList() []interface{} {
	return []interface{}{
		&user{},
	}
}
