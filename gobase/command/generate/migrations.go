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

func AppendToMigration(domain string) {
	fileName := "migrations/migration_list.go"
	res, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panic(err)
	}
	migrationList := migrations.GetAllMigrations()
	if len(migrationList[domain]) == 0 {
		fileContent := string(res)
		indexToAppend := strings.LastIndex(fileContent, ",") + 1
		newSchemaSet := "\"" + strings.ToLower(domain) + "\"" + ": " +
			strings.ToUpper(string(domain[0])) + domain[1:] + "Schema(),"
		finalString := fmt.Sprintf("%v\n\t\t%v%v", fileContent[:indexToAppend],
			newSchemaSet, fileContent[indexToAppend:])
		ioutil.WriteFile(fileName, []byte(finalString), 0775)
	}
}

func Migration(c *cli.Context) {
	input, err := ioutil.ReadFile("gobase/template/migrations.go")
	if err != nil {
		log.Fatal("File is Not Exist: %v", err)
	}
	domain := strings.ToLower(c.Args().Get(2))
	pluralDomain := utils.ConvertToPluralNoun(domain)
	file := string(input)
	file = strings.Replace(file, "_TEMPLATE_", strings.Title(domain), -1)
	file = strings.Replace(file, "_TABLE_", pluralDomain, -1)
	file = strings.Replace(file, "template", "migrations", -1)
	err = ioutil.WriteFile("migrations/"+pluralDomain+".go", []byte(file), 0755)
	if err != nil {
		log.Fatal("Unable to write file: %v", err)
	}
	AppendToMigration(domain)
	fmt.Printf("Successfully create %v migration", domain)
}
