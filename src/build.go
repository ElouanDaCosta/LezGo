package src

import (
	"flag"
	"fmt"
	"os"
)

var profileName string

var buildUsage = `Build LezGo given profile.

Usage: lezgo build [OPTIONS]

Options:
	-p, --profile     give the name of the profile
`

var buildFunc = func(cmd *Command, args []string) {
	fmt.Println("Starting the build process...")
}

func NewBuildCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("", flag.ExitOnError),
		Execute: buildFunc,
	}

	cmd.flags.StringVar(&profileName, "profile", "", "Long declaration to give the name of the profile")
	cmd.flags.StringVar(&profileName, "p", "", "Short declaration to give the name of the profile")

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, initUsage)
	}

	return cmd
}
