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
		fmt.Println(err)
	}
	for rows.Next() {
		var table string
		err = rows.Scan(
			&table,
		)

		if err != nil {
			fmt.Println(err)
		}

		if stringBuilder.Len() == 0 {
			stringBuilder.WriteString("\"" + table + "\"")
		} else {
			stringBuilder.WriteString("," + "\"" + table + "\"")
		}
	}
	return stringBuilder.String()
}

func Drop(db *sql.DB, domainName string) {
	var tables string
	if domainName == "all" {
		tables = GetAllTables(db)
	} else {
		tables = domainName
	}
	query := "DROP TABLE " + tables
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully drop tables")
}
