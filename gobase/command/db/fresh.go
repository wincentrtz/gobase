package db

import "database/sql"

func Fresh(db *sql.DB, domainName string) {
	Drop(db, domainName)
	Migrate(db, domainName)
}
