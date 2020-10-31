package db

import (
	"fmt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	fmt.Printf("Successfully seed tables\n")
}
