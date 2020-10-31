package seeders

import (
	"github.com/wincentrtz/gobase/models/entity"
	"gorm.io/gorm"
)

type user entity.User

func (*user) SeedTable(db *gorm.DB) {
	users := []entity.User{
		{
			gorm.Model{},
			"name",
			"email",
		},
		{
			gorm.Model{},
			"name1",
			"email1",
		},
	}

	if len(users) != 0 {
		db.Create(&users)
	}
}
