package generate

import (
	"fmt"
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
	file := string(input)
	file = strings.Replace(file, "_TEMPLATE_", strings.Title(domain), -1)
	file = strings.Replace(file, "_TABLE_", domain+"s", -1)
	file = strings.Replace(file, "template", "migrations", -1)
	err = ioutil.WriteFile("migrations/"+domain+"s.go", []byte(file), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}
