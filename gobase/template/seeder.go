package template

import (
	"gorm.io/gorm"
)

type template struct{}

func (*template) SeedTable(db *gorm.DB) {
	templates := []template{}
	if len(templates) != 0 {
		db.Create(&templates)
	}
}
