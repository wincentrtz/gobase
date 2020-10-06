package command

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/wincentrtz/gobase/gobase/command/app"
	"github.com/wincentrtz/gobase/gobase/command/db"
	"github.com/wincentrtz/gobase/gobase/command/generate"
	instance "github.com/wincentrtz/gobase/gobase/infrastructures/db"
)

func Command(c *cli.Context) error {
	firstArgs := c.Args().Get(0)
	secondArgs := c.Args().Get(1)
	switch firstArgs {
	case "generate":
		switch secondArgs {
		case "domain":
			generate.Domain(c)
			break
		case "response":
			generate.Response(c)
			break
		case "request":
			generate.Request(c)
		default:
			fmt.Println("Command Does Not Exist")
		}

	case "db":
		postgres := instance.Postgres()
		switch secondArgs {
		case "fresh":
			db.Fresh(postgres)
		case "clear":
			db.Drop(postgres)
		case "migrate":
			db.Migrate(postgres)
		default:
			fmt.Println("Command Does Not Exist")
		}
	case "serve":
		app.Serve()

	default:
		fmt.Println("Command Does Not Exist")
	}

	return nil
}
