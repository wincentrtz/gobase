package migrations

import (
	"github.com/wincentrtz/gobase/models/entity"
)

func MigrationList() []interface{} {
	return []interface{}{
		&entity.User{},
		&entity.Merchant{},
		&entity.Product{},
	}
}
