package commands

import (
	"fmt"
	"log"
	"os/user"
	"strings"

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
		Aliases: []string{"list"},
		Name:    "ls",
		Usage:   "list variables",
		Action:  bootstrap(List),
	},
	{
		Name:   "set",
		Usage:  "set a variable",
		Action: Set,
	},
	{
		Name:  "workspaces",
		Usage: "manage workspaces",
		Subcommands: []cli.Command{
			{
				Name:   "add",
				Usage:  "add a workspace",
				Action: bootstrap(WorkspaceAdd),
			},
			{
				Name:    "ls",
				Aliases: []string{"list"},
				Usage:   "list workspaces",
				Action:  bootstrap(WorkspaceList),
			},
			{
				Name:   "remove",
				Usage:  "remove a workspace",
				Action: bootstrap(WorkspaceRemove),
			},
			{
				Name:    "set-default",
				Aliases: []string{"default"},
				Usage:   "set the default workspace",
				Action:  bootstrap(WorkspaceSetDefault),
			},
		},
	},
}

type appCtx struct {
	cli *cli.Context
	app app.App
}

func bootstrap(command func(*appCtx) error) func(context *cli.Context) {
	return func(c *cli.Context) {
		path := c.String("config")
		if len(path) == 0 {
			currentUser, _ := user.Current()
			homeDir := currentUser.HomeDir
			path = strings.Replace(app.DEFAULT_CONFIG_PATH, "~", homeDir, 1)
		}
		appl, err := app.NewApp(path)
		if err != nil {
			log.Fatal(err)
		}

		if err := command(&appCtx{c, appl}); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}
	}
}
