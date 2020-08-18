package db

import (
	"database/sql"
	"github.com/wincentrtz/gobase/gobase/utils"
)

func Fresh(db *sql.DB, schemaName string) {
	pluralSchemaName := utils.ConvertToPluralNoun(schemaName)
	Drop(db, pluralSchemaName)
	Migrate(db, schemaName)
}
