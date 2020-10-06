package migrations

import (
	"github.com/wincentrtz/gobase/models/entity"
)

var user entity.User

func MigrationList() []interface{} {
	return []interface{}{&user}
}
