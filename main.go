package main

import (
	"github.com/mashup-cms/mashup-api/commands"
	"github.com/mashup-cms/mashup-api/commands/db"
	"github.com/mashup-cms/mashup-api/commands/add"
	"github.com/mashup-cms/mashup-api/commands/server"
	"github.com/mashup-cms/mashup-api/commands/worker"

	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

var Commands = []*commands.Command{
	db.UpCmd,
	db.DownCmd,
	db.RedoCmd,
	db.StatusCmd,
	db.CreateCmd,
	worker.SyncCmd,
	server.RunCmd,
	add.AdminCmd,
}

var usagePrefix = `
mashup-api is a database migration management system for Go projects.

Usage:
    mashup-api [options] <subcommand> [subcommand options]

Options:
`
var usageTmpl = template.Must(template.New("usage").Parse(
	`
Commands:{{range .}}
    {{.Name | printf "%-14s"}} {{.Summary}}{{end}}
`))

func usage() {
	fmt.Print(usagePrefix)
	flag.PrintDefaults()
	usageTmpl.Execute(os.Stdout, Commands)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 || args[0] == "-h" {
		flag.Usage()
		return
	}

	var cmd *commands.Command
	name := args[0]
	for _, c := range Commands {
		if strings.HasPrefix(c.Name, name) {
			cmd = c
			break
		}
	}

	if cmd == nil {
		fmt.Printf("error: unknown command %q\n", name)
		flag.Usage()
		os.Exit(1)
	}

	cmd.Exec(args[1:])
}