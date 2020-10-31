package migrations

import (
	"github.com/wincentrtz/gobase/models/entity"
)

var user entity.User
var product entity.Product

func MigrationList() []interface{} {
	return []interface{}{&user, &product}
}
