package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/wincentrtz/gobase/gobase/config"
	"github.com/wincentrtz/gobase/utils"
)

func GetAllTables(db *sql.DB) string {

	var tables string

	query := utils.NewQueryBuilder().
		Table("information_schema.tables").
		Select("table_name").
		Where("table_schema", "=", "public").
		AndWhere("table_type", "=", "BASE TABLE").
		Build()

	rows, err := db.Query(query)

	if err != nil || rows == nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var table string

		err = rows.Scan(
			&table,
		)

		if len(tables) == 0 {
			tables = table
		} else {
			tables = tables + "," + table
		}
	}

	return tables
}

func Drop() {
	db := config.InitDb()
	defer db.Close()
	tables := GetAllTables(db)
	query := "DROP TABLE " + tables
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
