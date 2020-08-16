package main

import (
	"github.com/urfave/cli"
	"github.com/wincentrtz/gobase/gobase/command"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gobase"
	app.Usage = "gobase helper cli"
	app.Action = command.Command
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
