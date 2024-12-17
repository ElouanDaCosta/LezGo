package src

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ElouanDaCosta/LezGo/pkg"
	"github.com/ElouanDaCosta/LezGo/src/templates"
)

var forceFlag bool = false
var projectName string

var initUsage = `Initialize LezGo in the current directory.

Usage: lezgo init [OPTIONS]

Options:
	-f, --force    skip the confirmation prompt
	-n, --name     give the name of the application
`

var initFunc = func(cmd *Command, args []string) {
	c := make(chan bool)
	go pkg.Throbber(pkg.Spin, 50*time.Millisecond, c)

	errCreateConfig := createConfigFile(projectName)
	if errCreateConfig != nil {
		c <- true
		panic(errCreateConfig)
	}
	c <- true
	if forceFlag {
		fmt.Printf("%+v\n", "force")
	}
}

func createConfigFile(name string) error {
	// If the file doesn't exist, create it, or append to the file
	template := templates.RenderLezgoConfigTemplate(name)
	f, err := os.OpenFile("lezgo.yaml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write([]byte(template))
	if err != nil {
		log.Fatal(err)
	}

	f.Close()
	return nil
}

func NewInitCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("", flag.ExitOnError),
		Execute: initFunc,
	}

	cmd.flags.BoolVar(&forceFlag, "force", false, "Long declaration to skip the confirmation prompt")
	cmd.flags.BoolVar(&forceFlag, "f", false, "Short declaration to skip the confirmation prompt")
	cmd.flags.StringVar(&projectName, "name", "new_app", "Long declaration to give the name of the application")
	cmd.flags.StringVar(&projectName, "n", "new_app", "Short declaration to give the name of the application")

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, initUsage)
	}

	return cmd
}
