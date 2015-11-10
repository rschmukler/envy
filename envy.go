package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/rschmukler/envy/app"
	"github.com/rschmukler/envy/commands"
	"os"
)

var AppHelpTemplate = `Usage: {{.Name}} {{if .Flags}}[OPTIONS] {{end}}COMMAND [arg...]

{{.Usage}}

Version: {{.Version}}{{if or .Author .Email}}
Author:{{if .Author}}
  {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
  {{.Email}}{{end}}{{end}}
{{if .Flags}}
Options:
  {{range .Flags}}{{.}}
  {{end}}{{end}}
Commands:
  {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}
Run '{{.Name}} COMMAND --help' for more information on a command.
`

var CommandHelpTemplate = `Usage: docker-machine {{.Name}}{{if .Flags}} [OPTIONS]{{end}} [arg...]
{{.Usage}}{{if .Description}}
Description:
   {{.Description}}{{end}}{{if .Flags}}
Options:
   {{range .Flags}}
   {{.}}{{end}}{{ end }}
`

func main() {
	cli.AppHelpTemplate = AppHelpTemplate
	cli.CommandHelpTemplate = CommandHelpTemplate

	prog := cli.NewApp()
	prog.Name = "envy"
	prog.Usage = "Manage environment variables with style"

	prog.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Usage: fmt.Sprintf("Path to envy config (default: %s)", app.DEFAULT_CONFIG_PATH),
		},
	}

	prog.Commands = commands.Commands
	prog.Run(os.Args)
}
