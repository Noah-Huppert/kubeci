package cmds

import (
	"github.com/urfave/cli"
	"fmt"
)

type ListCommand struct{
	opCommandDefs []Command
}

func (l ListCommand) Name() []string {
	return []string{"list", "l"}
}

func (l ListCommand) Usage() string {
	return "List all available kubeci operations"
}

func (l ListCommand) Flags() []cli.Flag {
	return []cli.Flag{}
}

func (l ListCommand) Args() []Argument {
	return []Argument{}
}

func (l ListCommand) Subcommands() []Command {
	return []Command{}
}

func (l ListCommand) Action(c *cli.Context) error {
	fmt.Println("Avaiable operations:\n")

	for _, op := range l.opCommandDefs {
		// Check
		err := CheckCommand(op)
		if err != nil {
			return fmt.Errorf("Error checking operation: %s", err.Error())
		}

		err, name := AssembleCommandName(op)
		if err != nil {
			return err
		}

		fmt.Printf("    %s - %s\n", name, op.Usage())
	}

	fmt.Println("")

	return nil
}