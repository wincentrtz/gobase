package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/wincentrtz/gobase/gobase/utils"
)

func GetAllTables(db *sql.DB) string {
	stringBuilder := strings.Builder{}
	query := utils.NewQueryBuilder().
		Table("information_schema.tables").
		Select("table_name").
		Where("table_schema", "=", "public").
		AndWhere("table_type", "=", "BASE TABLE").
		Build()
	rows, err := db.Query(query)
	if err != nil || rows == nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var table string
		err = rows.Scan(
			&table,
		)

		if err != nil {
			log.Fatal(err)
		}

		if stringBuilder.Len() == 0 {
			stringBuilder.WriteString("\"" + table + "\"")
		} else {
			stringBuilder.WriteString("," + "\"" + table + "\"")
		}
	}
	return stringBuilder.String()
}

func Drop(db *sql.DB, schemaName string) {
	var tables string
	var tableIdentifier string
	if schemaName == "all" {
		tables = GetAllTables(db)
		tableIdentifier = "tables"
	} else {
		tables = schemaName
		tableIdentifier = "table"
	}
	query := "DROP TABLE " + tables
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully drop %v %v\n", schemaName, tableIdentifier)
}
