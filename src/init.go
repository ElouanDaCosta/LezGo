package src

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ElouanDaCosta/LezGo/pkg"
)

var forceFlag bool = false

var initUsage = `Initialize LezGo in the current directory.

Usage: lezgo init [OPTIONS]

Options:
	-f, --force    skip the confirmation prompt
`

var initFunc = func(cmd *Command, args []string) {
	c := make(chan bool)

	go pkg.Throbber(pkg.Spin, 50*time.Millisecond, c)
	time.Sleep(1 * time.Second)
	c <- true
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
