package src

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ElouanDaCosta/LezGo/pkg"
	"gopkg.in/yaml.v3"
)

var checkUsage = `Health check the config file.

Usage: lezgo check [OPTIONS]

Options:
	
`

var checkFunc = func(cmd *Command, args []string) {
	fmt.Println("Checking config file...")
	f, err := os.ReadFile("lezgo.yaml")
	if err != nil {
		log.Fatal(err)
	}

	profile := pkg.Config{}

	err = yaml.Unmarshal(f, &profile)
	if err != nil {
		panic(err)
	}

	fmt.Print(profile.Project)

	return
}

func NewCheckCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("", flag.ExitOnError),
		Execute: checkFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, checkUsage)
	}

	return cmd
}
