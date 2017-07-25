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

	err, commands := cmds.LoadCommands(cmds.CommandDefs(ops.CommandDefs))
	if err != nil {
		fmt.Printf("Error loading commands, error: %s\n", err.Error())
	}

	app.Commands = commands

	app.Run(os.Args)
}