package template

import "gorm.io/gorm"

type templateRepository struct {
	db *gorm.DB
}

func NewTemplateRepository(db *gorm.DB) TemplateRepository {
	return &templateRepository{
		db: db,
	}
}
