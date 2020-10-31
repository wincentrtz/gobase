package seeders

import (
	"github.com/wincentrtz/gobase/models/entity"
	"gorm.io/gorm"
)

type product entity.Product

func (*product) SeedTable(db *gorm.DB) {
	products := []product{}
	db.Create(&products)
}
