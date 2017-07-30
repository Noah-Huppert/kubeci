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
	NewFullMockCommand([]Command{}, defaultMockAction),
    NewFullMockCommand([]Command{}, defaultMockAction),
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

// NewDefaultFullMockCommand constructs a new full MockCommand with default mock values
func NewDefaultFullMockCommand() MockCommand {
    return NewFullMockCommand(defaultMockSubcommands, defaultMockAction)
}

// CheckCommand
func Test_CheckCommand_NameFull_Ok(t *testing.T) {
	cmd := NewFullMockCommand(defaultMockSubcommands, defaultMockAction)

    err := CheckCommand(cmd)
    assert.Nil(t, err)
}

func Test_CheckCommand_NameEmpty_NotOk(t *testing.T) {
    cmd := NewEmptyMockCommand()

    err := CheckCommand(cmd)
    assert.NotNil(t, err)
}

func Test_CheckCommand_NameElementEmpty_NotOk(t *testing.T) {
    cmd := NewFullMockCommand(defaultMockSubcommands, defaultMockAction)
    cmd.MockName = []string{"Name",""}

    err := CheckCommand(cmd)
    assert.NotNil(t, err)
}

func Test_CheckCommand_UsageFull_Ok(t *testing.T) {
    cmd := NewDefaultFullMockCommand()

    err := CheckCommand(cmd)
    assert.Nil(t, err)
}

func Test_CheckCommand_UsageEmpty_NotOk(t *testing.T) {
    cmd := NewDefaultFullMockCommand()
    cmd.MockUsage = ""

    err := CheckCommand(cmd)
    assert.NotNil(t, err)
}

func Test_CheckCommand_ArgsFull_Ok(t *testing.T) {
    cmd := NewDefaultFullMockCommand()

    err := CheckCommand(cmd)
    assert.Nil(t, err)
}

func Test_CheckCommand_ArgsEmpty_Ok(t *testing.T) {
    cmd := NewDefaultFullMockCommand()
    cmd.MockArgs = []Argument{}

    err := CheckCommand(cmd)
    assert.Nil(t, err)
}

func Test_CheckCommand_ArgsEmptyContentsName_NotOk(t *testing.T) {
    cmd := NewDefaultFullMockCommand()
    cmd.MockArgs[0].Name = ""

    err := CheckCommand(cmd)
    assert.NotNil(t, err)
}

func Test_CheckCommand_ArgsEmptyContentsUsage_NotOk(t *testing.T) {
    cmd := NewDefaultFullMockCommand()
    cmd.MockArgs[0].Usage = ""

    err := CheckCommand(cmd)
    assert.NotNil(t, err)
}

// AssembleCommandName
func Test_AssembleCommandName_EmptyCommand_NotOk(t *testing.T) {
    cmd := NewEmptyMockCommand()

    err, _ := AssembleCommandName(cmd)
    assert.NotNil(t, err)
}

func Test_AssembleCommandName_OneName(t *testing.T) {
    cmd := NewDefaultFullMockCommand()
    cmd.MockName = []string{"Name"}

    err, name := AssembleCommandName(cmd)
    assert.Nil(t, err)
    assert.Equal(t, name, "Name")
}

func Test_AssembleCommandName_MultipleNames(t *testing.T) {
    cmd := NewDefaultFullMockCommand()
    cmd.MockName = []string{"Name", "n1", "n2"}

    err, name := AssembleCommandName(cmd)
    assert.Nil(t, err)
    assert.Equal(t, name, "{Name,n1,n2}")
}

// AssembleCommand
func Test_AssembleCommand_EmptyCommand_NotOk(t *testing.T) {
    cmd := NewEmptyMockCommand()

    err, out := AssembleCommand(cmd)
    assert.NotNil(t, err)
    assert.Nil(t, out)
}

func Test_AssembleCommand_Ok(t *testing.T) {
    cmd := NewEmptyMockCommand()
    cmd.MockName = []string{"name", "n1", "n2"}
    cmd.MockUsage = "How to use name"
    cmd.MockArgs = []Argument{
        {
            Name: "Arg1",
            Usage: "How to use Arg1",
        },
        {
            Name: "Arg2",
            Usage: "How to use Arg2",
        },
    }
    cmd.MockFlags = []cli.Flag{
        cli.StringFlag{
            Name: "Flag1",
            Usage: "How to use Flag1",
        },
        cli.StringFlag{
            Name: "Flag2",
            Usage: "How to use Flag2",
        },
    }
    cmd.MockSubcommands = defaultMockSubcommands

    err, out := AssembleCommand(cmd)
    assert.Nil(t, err)
    assert.Equal(t, out.Name, "name")
    assert.Equal(t, out.Aliases, []string{"n1", "n2"})
    assert.Equal(t, out.Usage, "How to use name")
    assert.Equal(t, out.UsageText, "kubeci {name,n1,n2} --Flag1 Flag1 --Flag2 Flag2 [Arg1] [Arg2]\n\nPOSITIONAL OPTIONS:\n   [Arg1] - How to use Arg1\n   [Arg2] - How to use Arg2")
    assert.Equal(t, out.ArgsUsage, "[Arg1] [Arg2]")
    assert.Equal(t, out.Flags, cmd.MockFlags)
    assert.NotNil(t, out.Subcommands)
    assert.NotNil(t, out.Before)
    assert.NotNil(t, out.Action)
}

// LoadCommands
func Test_LoadCommands_FullCommands_Ok(t *testing.T) {
    err, cmds := LoadCommands(defaultMockSubcommands)

    assert.Nil(t, err)
    assert.NotNil(t, cmds)
    assert.Equal(t, len(cmds), 2)
}

func Test_LoadCommands_BadCommands_NotOk(t *testing.T) {
    err, cmds := LoadCommands([]Command{
        NewEmptyMockCommand(),
        NewEmptyMockCommand(),
    })

    assert.NotNil(t, err)
    assert.Equal(t, len(cmds), 0)
}
