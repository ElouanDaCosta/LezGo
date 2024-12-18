package src

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ElouanDaCosta/LezGo/pkg"
	"github.com/ElouanDaCosta/LezGo/src/templates"
)

var forceFlag bool = false
var projectName string

var initUsage = `Initialize LezGo in the current directory.

Usage: lezgo init [OPTIONS]

Options:
	-n, --name     give the name of the application
`

/*
Initialize LezGo to the current project,
create the config file with default configuration
*/
var initFunc = func(cmd *Command, args []string) {
	// ascii art
	pkg.LezgoAsciiArt()
	fmt.Println("Creating the config file...")

	// create config file
	errCreateConfig := createConfigFile(projectName)
	if errCreateConfig != nil {
		panic(errCreateConfig)
	}
	if forceFlag {
		fmt.Printf("%+v\n", "force")
	}
	fmt.Println("LezGo initialize successfully!")
}

func createConfigFile(name string) error {
	// If the file doesn't exist, create it, or append to the file
	template := templates.RenderLezgoConfigTemplate(name)
	f, err := os.OpenFile("lezgo.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
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
