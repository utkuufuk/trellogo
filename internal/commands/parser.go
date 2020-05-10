package commands

import (
	"fmt"
)

// Parse parses program arguments (excluding executable name, i.e. os.Args[0])
// and returns the command to be executed.
func Parse(args []string) Command {
	if len(args) == 0 {
		return UndefinedCommand{Err: fmt.Errorf("too few arguments")}
	}

	if len(args) == 1 && args[0] == "list" {
		return FetchAllListsCommand{}
	}

	if len(args) > 1 && args[0] == "list" {
		return FetchListsCommand{args[1:]}
	}
	return UndefinedCommand{Err: fmt.Errorf("undefined command: %v", args)}
}
