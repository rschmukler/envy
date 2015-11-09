package commands

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func List(ctx *appCtx) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Variable", "Workspace", "Value"})
	table.SetBorder(false)

	for _, val := range ctx.app.Vals {
		table.Append([]string{val.Name, val.Workspace, val.Value})
	}

	table.Render()
	return nil
}
