package src

import (
	"flag"
	"fmt"
	"os"
)

var forceFlag bool = false

var initUsage = `Initialize LezGo in the current directory.

Usage: lezgo init [OPTIONS]

Options:
	-f, --force    skip the confirmation prompt
`

var initFunc = func(cmd *Command, args []string) {
	fmt.Printf("%+v\n", "here")
	if forceFlag {
		fmt.Printf("%+v\n", "force")
	}
}

func NewInitCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("", flag.ExitOnError),
		Execute: initFunc,
	}

	cmd.flags.BoolVar(&forceFlag, "force", false, "Long declaration to skip the confirmation prompt")
	cmd.flags.BoolVar(&forceFlag, "f", false, "short declaration to skip the confirmation prompt")

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, initUsage)
	}

	return cmd
}
