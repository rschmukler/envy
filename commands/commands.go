package commands

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/rschmukler/envy/app"
)

var Commands = []cli.Command{
	{
		Name:   "load",
		Usage:  "load workspace for sourcing",
		Action: Load,
	},
	{
		Name:   "set",
		Usage:  "set a variable",
		Action: Set,
	},
	{
		Name:   "workspaces",
		Usage:  "manage workspaces",
		Action: Workspaces,
	},
}

type contextCommandLine struct {
	ctx *cli.Context
	app app.App
}

func bootstrap(command func(*contextCommandLine) error) func(context *cli.Context) {
	return func(c *cli.Context) {
		path := c.String("config")
		if len(path) == 0 {
			path = app.DEFAULT_CONFIG_PATH
		}
		appl := app.NewApp(path)

		if err := command(&contextCommandLine{c, appl}); err != nil {
			log.Fatal(err)
		}
	}
}
