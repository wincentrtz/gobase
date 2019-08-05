package template

import (
	"database/sql"

	"github.com/wincentrtz/gobase/gobase/config"
)

type templateRepository struct {
	db *sql.DB
}

func NewTemplateRepository() TemplateRepository {
	return &templateRepository{
		db: config.InitDb(),
	}
}
