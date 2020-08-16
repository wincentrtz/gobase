package generate

import (
	"fmt"
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
	fileContent := string(res)
	indexToAppend := strings.LastIndex(fileContent, ",") + 1
	newSchemaSet := "\"" + strings.ToLower(domain) + "\"" + ": " +
		strings.ToUpper(string(domain[0])) + domain[1:] + "Schema(),"
	new := fmt.Sprintf("%v\n\t\t%v%v", fileContent[:indexToAppend], newSchemaSet, fileContent[indexToAppend:])
	ioutil.WriteFile(fileName, []byte(new), 0775)
}

func Migration(c *cli.Context) {
	input, err := ioutil.ReadFile("gobase/template/migrations.go")
	if err != nil {
		log.Fatal("File is Not Exist: %v", err)
	}
	domain := c.Args().Get(2)
	file := string(input)
	file = strings.Replace(file, "_TEMPLATE_", strings.Title(domain), -1)
	file = strings.Replace(file, "_TABLE_", domain+"s", -1)
	file = strings.Replace(file, "template", "migrations", -1)
	err = ioutil.WriteFile("migrations/"+domain+"s.go", []byte(file), 0755)
	if err != nil {
		log.Fatal("Unable to write file: %v", err)
	}
	AppendToMigration(domain)
	fmt.Println("Successfully create migration")
}
