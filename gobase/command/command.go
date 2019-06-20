package command

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/wincentrtz/gobase/gobase/command/db"
	"github.com/wincentrtz/gobase/gobase/command/generate"
)

func Command(c *cli.Context) error {
	switch c.Args().Get(0) {
	case "generate":
		switch c.Args().Get(1) {
		case "domain":
			generate.Domain(c)
		}
	case "db":
		switch c.Args().Get(1) {
		case "clear":
			db.Drop()
		}
	default:
		fmt.Println("Command Does Not Exist")
	}

	return nil
}
