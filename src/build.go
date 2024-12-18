package src

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ElouanDaCosta/LezGo/pkg"
	"gopkg.in/yaml.v3"
)

var profileName string

var buildUsage = `Build LezGo given profile.

Usage: lezgo build [OPTIONS]

Options:
	-p, --profile     give the name of the profile
`

var buildFunc = func(cmd *Command, args []string) {
	fmt.Println("Starting the build process...")
	f, err := os.ReadFile("lezgo.yaml")
	if err != nil {
		log.Fatal(err)
	}

	profile := pkg.Config{}

	err = yaml.Unmarshal(f, &profile)
	if err != nil {
		panic(err)
	}

	// var config map[string]interface{}
	// if err := yaml.Unmarshal(f, &config); err != nil {
	// 	fmt.Println("YAML syntax error:", err)
	// } else {
	// 	fmt.Println("YAML is well-formed!")
	// }
	fmt.Println(profile.Profiles.Release)
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
