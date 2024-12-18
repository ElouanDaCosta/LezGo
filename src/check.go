package src

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"unicode"

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

	checkConfigName(profile)

	fmt.Println("Config file check complete")

	return
}

func checkConfigName(yamlStruct pkg.Config) error {
	if yamlStruct.Name == "" {
		return errors.New("name field shouldn't be nil")
	} else {
		for _, r := range yamlStruct.Name {
			if unicode.IsUpper(r) {
				pkg.Warning("Name should be in lowercase only")
				break
			}
		}
	}
	return nil
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
