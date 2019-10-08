package template

import (
	"database/sql"
)

type templateRepository struct {
	db *sql.DB
}

func NewTemplateRepository(db *sql.DB) TemplateRepository {
	return &templateRepository{
		db: db,
	}
}
