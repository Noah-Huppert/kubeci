package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
	"github.com/Noah-Huppert/kubeci/kubeci/cmds"
	"github.com/Noah-Huppert/kubeci/kubeci/ops"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Description = "Command line tool which performs Kubernetes operations commonly required during continuous integration"
	app.Usage = "Kubernetes continuous integration command line interface"

	err, commands := cmds.LoadCommands(cmds.CommandDefs(ops.CommandDefs))
	if err != nil {
		fmt.Printf("Error loading commands, error: %s\n", err.Error())
	}

	app.Commands = commands

	app.Run(os.Args)
}
