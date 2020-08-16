package db

import "database/sql"

func Fresh(db *sql.DB, schemaName string) {
	Drop(db)
	Migrate(db, schemaName)
}
