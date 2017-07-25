package cmds

// CommandDefs returns an array of Command interfaces which represent cli command definitions
// The opCommandDefs argument is the list of commands which the op package contains. This argument is necessary to
// prevent an import loop.
func CommandDefs(opCommandDefs []Command) []Command {
	return []Command{
		ListCommand{ opCommandDefs: opCommandDefs },
	}
}
