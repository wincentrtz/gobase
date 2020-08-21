package generate

import (
	"fmt"
	"github.com/wincentrtz/gobase/gobase/utils"
	"io/ioutil"
	"strings"

	"github.com/urfave/cli"
)

func Migration(c *cli.Context) {
	input, err := ioutil.ReadFile("gobase/template/migrations.go")
	if err != nil {
		fmt.Printf("File is Not Exist: %v", err)
	}
	domain := strings.ToLower(c.Args().Get(2))
	pluralDomain := utils.ConvertToPluralNoun(domain)
	file := string(input)
	file = strings.Replace(file, "_TEMPLATE_", strings.Title(domain), -1)
	file = strings.Replace(file, "_TABLE_", pluralDomain, -1)
	file = strings.Replace(file, "template", "migrations", -1)
	err = ioutil.WriteFile("migrations/"+pluralDomain+".go", []byte(file), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}
