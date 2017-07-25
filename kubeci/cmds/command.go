package cmds

import (
	"fmt"
	"errors"
	"strings"

	"github.com/urfave/cli"
)

// BaseCommandName is the Name of the base executable from which all commands are called
var BaseCommandName string = "kubeci"

// Argument holds information about a command argument
type Argument struct {
	// Name is the string which the argument will be identified by in help text
	Name string

	// Usage is the help text which will be shown to the user in the command help
	Usage string
}

// Command defines a simple interface which all cli sub commands implement. It allows the loader of these sub commands
// to easily retrieve useful command information
type Command interface {
	// Name returns an array of names the command should be registered under. The first element in the returned string
	// array should be the command's full Name. All elements following the first element are considered command aliases.
	Name() []string

	// Usage return Usage help text to display the user
	Usage() string

	// Args returns an array of Arguments which represent command arguments.
	Args() []Argument

	// Flags returns an array of cli.Flag interfaces to describe valid command Flags
	Flags() []cli.Flag

	// Subcommands returns an array of Command interfaces which will be registered as sub commands to this command
	Subcommands() []Command

	// Action is the method which should run when the command is invoked. Argument values passed by the user are
	// provided in the `c` argument.
	Action(c *cli.Context) error
}

// CheckCommand inspects the command Name(), Usage(), and Args() method return values for any incorrect values. It will
// return an error if any value is invalid, and nil if successful
func CheckCommand(c Command) error {
	// Name
	name := c.Name()
	if len(name) == 0 {
		return errors.New("Command Name() method must return an array with 1 or more elements")
	}

	for key, n := range name {
		// Check not empty
		if len(n) == 0 {
			return fmt.Errorf("Command Name() must return an array of non empty strings, element %d was empty", key)
		}
	}

	// Usage
	usage := c.Usage()
	if len(usage) == 0 {
		return errors.New("Command Usage() must return a non empty string")
	}

	// Args
	args := c.Args()

	for i, arg := range args {
		// Check Name and Usage not empty
		if len(arg.Name) == 0 {
			return fmt.Errorf("Command Args() must return an array of Arguments with non empty Name fields, element %d has an empty Name field", i)
		}

		if len(arg.Usage) == 0 {
			return fmt.Errorf("Command Args() must return an array of Arguments with non empty Usage fields, element %d has an empty Usage field", i)
		}
	}

	// All good
	return nil
}

func AssembleCommandName(c Command) (error, string) {
	// Check
	err := CheckCommand(c)
	if err != nil {
		return fmt.Errorf("Error while checking command: %s", err.Error()), ""
	}

	// Assemble
	name := c.Name()

	cmdNames := name[0]
	if len(name) > 1 {
		cmdNames = fmt.Sprintf("{%s}", strings.Join(name, ","))
	}

	return nil, cmdNames
}

// AssembleCommand creates a cli.Command for the provided Command interface. If no errors occurred `nil` is returned
func AssembleCommand(c Command) (error, *cli.Command) {
	// Check command
	err := CheckCommand(c)
	if err != nil {
		return fmt.Errorf("Error while checking command: %s", err.Error()), nil
	}

	// Generate UsageText
	// A combination of the general command Usage and the defined Args
	err, cmdNames := AssembleCommandName(c)
	if err != nil {
		return err, nil
	}

	usageTxt := fmt.Sprintf("%s %s", BaseCommandName, cmdNames)
	args := c.Args()
	flags := c.Flags()

	for _, flag := range flags {
		usageTxt += fmt.Sprintf(" --%s %s", flag.GetName(), flag.GetName())
	}

	for _, arg := range args {
		usageTxt += fmt.Sprintf(" [%s]", arg.Name)
	}

	if len(args) > 0 {
		usageTxt += "\n\nPOSITIONAL OPTIONS:"

		for _, arg := range args {
			usageTxt += fmt.Sprintf("\n   [%s] - %s", arg.Name, arg.Usage)
		}
	}

	// Generate ArgsUsage text
	// String with argument names in order surrounded by square brackets
	argsUsageTxt := ""

	for i, arg := range args {
		argsUsageTxt += fmt.Sprintf("[%s]", arg.Name)

		if i != len(args) - 1 {
			argsUsageTxt += " "
		}
	}

	// Generate SubCommands array
	// Run AssembleCommand() function recursively on commands in c.Subcommands() array
	subCommands := cli.Commands{}

	for i, cmd := range c.Subcommands() {
		err, cliCmd := AssembleCommand(cmd)
		if err != nil {
			return fmt.Errorf("Error while processing element %d, error: %s", i, err.Error()), nil
		}
		subCommands = append(subCommands, *cliCmd)
	}

	name := c.Name()

	outCmd := cli.Command{
		Name:        name[0],
		Aliases:     name[1:],
		Usage: c.Usage(),
		UsageText:   usageTxt,
		ArgsUsage:   argsUsageTxt,
		Flags:       flags,
		Subcommands: subCommands,
		Before: func(c *cli.Context) error {
			// Check correct number of arguments
			if (c.NArg() < len(args)) {
				errMsg := "All positional arguments must be provided, %d required, %d provided, missing:"

				for _, arg := range args[c.NArg()-1:] {
					errMsg += fmt.Sprintf("\n    - [%s] - %s", arg.Name, arg.Usage)
				}

				return errors.New(errMsg)
			}

			return nil
		},
		Action: c.Action,
	}

	return nil, &outCmd
}

// LoadCommands will assemble all commands in the provided array.
func LoadCommands(cmdDefs []Command) (error, []cli.Command) {
	commands := cli.Commands{}

	for i, def := range cmdDefs {
		err, command := AssembleCommand(def)
		if err != nil {
			return fmt.Errorf("Error assembling command %d, error: %s", i, err.Error()), []cli.Command{}
		}

		commands = append(commands, *command)
	}

	return nil, commands
}