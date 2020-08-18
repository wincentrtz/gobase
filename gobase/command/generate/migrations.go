package generate

import (
	"fmt"
	"github.com/wincentrtz/gobase/gobase/utils"
	"github.com/wincentrtz/gobase/migrations"
	"io/ioutil"
	"log"
	"strings"

	"github.com/urfave/cli"
)

func AppendToMigration(schemaName string) {
	fileName := "migrations/migration_list.go"
	res, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panic(err)
	}
	if len(migrations.MigrationList[schemaName]) == 0 {
		fileContent := string(res)
		indexToAppend := strings.LastIndex(fileContent, ",") + 1
		newSchemaSet := "\"" + strings.ToLower(schemaName) + "\"" + ":\t\t" +
			strings.ToUpper(string(schemaName[0])) + schemaName[1:] + "Schema(),"
		finalString := fmt.Sprintf("%v\n\t%v%v", fileContent[:indexToAppend],
			newSchemaSet, fileContent[indexToAppend:])
		ioutil.WriteFile(fileName, []byte(finalString), 0775)
	}
}

func Migration(c *cli.Context) {
	input, err := ioutil.ReadFile("gobase/template/migrations.go")
	if err != nil {
		log.Fatal("File is Not Exist: %v", err)
	}
	schemaName := strings.ToLower(c.Args().Get(2))
	pluralDomain := utils.ConvertToPluralNoun(schemaName)
	file := string(input)
	file = strings.Replace(file, "_TEMPLATE_", strings.Title(schemaName), -1)
	file = strings.Replace(file, "_TABLE_", pluralDomain, -1)
	file = strings.Replace(file, "template", "migrations", -1)
	err = ioutil.WriteFile("migrations/"+pluralDomain+".go", []byte(file), 0755)
	if err != nil {
		log.Fatal("Unable to write file: %v", err)
	}
	AppendToMigration(schemaName)
	fmt.Printf("Successfully create %v migration", schemaName)
}
