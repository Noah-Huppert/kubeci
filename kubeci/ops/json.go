package ops

import (
	"github.com/urfave/cli"
	"github.com/Noah-Huppert/kubeci/kubeci/cmds"
)

type JsonCommand struct {}

func (j JsonCommand) Name() []string {
	return []string{"json"}
}

func (j JsonCommand) Usage() string {
	return "Runs a Go template against the specified Json data"
}

func (j JsonCommand) Args() []cmds.Argument {
	return []cmds.Argument{
		{
			Name: "template",
			Usage: "Go template to run against Json data",
		},
	}
}

func (j JsonCommand) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name: "file",
			Usage: "File to load Json data from, if specified can not provide env-var argument",
		},
		cli.StringFlag{
			Name: "env-var",
			Usage: "Environment variable to load Json data from, assumes data is base64 encoded in env var, if specified can not provide file argument",
		},
	}
}

func (j JsonCommand) Subcommands() []cmds.Command {
	return []cmds.Command{}
}

func (j JsonCommand) Action(c *cli.Context) error {
	return nil
}