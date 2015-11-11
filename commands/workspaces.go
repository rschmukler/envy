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
	table.SetHeader([]string{"Name", "Default"})
	table.SetBorder(false)

	for _, workspace := range ctx.app.Workspaces {
		var isDefault string
		if ctx.app.DefaultWorkspace == workspace {
			isDefault = "*"
		} else {
			isDefault = ""
		}
		table.Append([]string{workspace, isDefault})
	}

	table.Render()
	return nil
}

func WorkspaceRemove(ctx *appCtx) error {
	name := ctx.cli.Args().First()
	if err := ctx.app.RemoveWorkspace(name); err != nil {
		return err
	}
	if err := ctx.app.Save(); err != nil {
		return err
	}
	return nil
}

func WorkspaceSetDefault(ctx *appCtx) error {
	return nil
}
