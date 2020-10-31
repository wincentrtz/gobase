package template

import (
	"gorm.io/gorm"
)

type template struct{}

func (*template) SeedTable(db *gorm.DB) {
	templates := []template{}
	db.Create(&templates)
}
