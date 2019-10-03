package db

import "database/sql"

func Fresh(db *sql.DB) {
	Drop(db)
	Migrate(db)
}
