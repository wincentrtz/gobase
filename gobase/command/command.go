package command

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/wincentrtz/gobase/gobase/command/app"
	"github.com/wincentrtz/gobase/gobase/command/db"
	"github.com/wincentrtz/gobase/gobase/command/generate"
)

func Command(c *cli.Context) error {
	switch c.Args().Get(0) {
	case "generate":
		switch c.Args().Get(1) {
		case "domain":
			generate.Domain(c)
			break
		case "migration":
			generate.Migration(c)
			break
		case "response":
			generate.Response(c)
			break
		}

	case "db":
		switch c.Args().Get(1) {
		case "fresh":
			db.Fresh()
		case "clear":
			db.Drop()
		case "migrate":
			db.Migrate()
		}
	case "serve":
		app.Serve()

	default:
		fmt.Println("Command Does Not Exist")
	}

	return nil
}
