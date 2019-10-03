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
		case "request":
			generate.Request(c)
		}

	case "db":
		postgres := instance.Postgres()
		defer postgres.Close()
		switch c.Args().Get(1) {
		case "fresh":
			db.Fresh(postgres)
		case "clear":
			db.Drop(postgres)
		case "migrate":
			db.Migrate(postgres)
		}
	case "serve":
		app.Serve()

	default:
		fmt.Println("Command Does Not Exist")
	}

	return nil
}
