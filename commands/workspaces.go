package commands

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func WorkspaceAdd(ctx *appCtx) error {
	name := ctx.cli.Args().First()
	if err := ctx.app.AddWorkspace(name); err != nil {
		return err
	}

	if err := ctx.app.Save(); err != nil {
		return err
	}
	return nil
}

func WorkspaceList(ctx *appCtx) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name"})
	table.SetBorder(false)

	for _, workspace := range ctx.app.Workspaces {
		table.Append([]string{workspace})
	}

	table.Render()
	return nil
}

func WorkspaceRemove(ctx *appCtx) error {
	return nil
}

func WorkspaceSetDefault(ctx *appCtx) error {
	return nil
}
