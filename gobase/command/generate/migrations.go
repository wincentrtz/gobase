package generate

import (
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli"
)

func Migration(c *cli.Context) {
	domain := c.Args().Get(2)

	err := ioutil.WriteFile("migrations/"+domain+"s.go", []byte("package migrations"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	fmt.Printf("Generate migration")
}
