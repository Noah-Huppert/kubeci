package cmds

import (
	"testing"
	"github.com/urfave/cli"
	"github.com/stretchr/testify/assert"
)

// MockCommand is a mock Command interface for testing purposes, returns all values in respective fields
type MockCommand struct {
	MockName []string
	MockUsage string
	MockArgs []Argument
	MockFlags []cli.Flag
	MockSubcommands []Command
	MockAction func(c *cli.Context) error
}

func (m MockCommand) Name() []string {
	return m.MockName
}

func (m MockCommand) Usage() string {
	return m.MockUsage
}

func (m MockCommand) Args() []Argument {
	return m.MockArgs
}

func (m MockCommand) Flags() []cli.Flag {
	return m.MockFlags
}

func (m MockCommand) Subcommands() []Command {
	return m.MockSubcommands
}

func (m MockCommand) Action(c *cli.Context) error {
	return m.MockAction(c)
}

// defaultMockSubcommands holds 2 empty mock Commands for use when one does not wish to provide some mock subcommands
var defaultMockSubcommands []Command = []Command{
	NewEmptyMockCommand(),
	NewEmptyMockCommand(),
}

// defaultMockAction is a default action method for a Command interface
func defaultMockAction (c *cli.Context) error {
	return nil
}

// NewEmptyMockCommand creates a new MockCommand with empty values
func NewEmptyMockCommand() MockCommand {
	return MockCommand{
		MockName: []string{},
		MockUsage: "",
		MockArgs: []Argument{},
		MockFlags: []cli.Flag{},
		MockSubcommands: []Command{},
		MockAction: defaultMockAction,
	}
}

// NewFullMockCommand creates a new MockCommand with full dummy values
func NewFullMockCommand(subcommands []Command, mockAction func(c *cli.Context) error) MockCommand {
	return MockCommand{
		MockName: []string{"MockName", "MN"},
		MockUsage: "How to use MockCommand",
		MockArgs: []Argument{
			{
				Name: "MockArgument1",
				Usage: "How to use MockArgument1",
			},
			{
				Name: "MockArgument2",
				Usage: "How to use MockArgument2",
			},
		},
		MockFlags: []cli.Flag{
			cli.StringFlag{
				Name: "MockFlag1",
				Usage: "How to use MockFlag1",
			},
			cli.StringFlag{
				Name: "MockFlag2",
				Usage: "How to use MockFlag2",
			},
		},
		MockSubcommands: subcommands,
		MockAction: mockAction,
	}
}

func Test_CheckCommand_NameFull_Ok(t *testing.T) {
	cmd := NewFullMockCommand(defaultMockSubcommands, defaultMockAction)


}