package db

import (
	"fmt"
	"github.com/wincentrtz/gobase/models/entity"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
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
	db.Create(&users)

	fmt.Printf("Successfully seed tables\n")
}
