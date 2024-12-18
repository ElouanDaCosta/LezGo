package src

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

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
	if profileName == "" {
		pkg.CustomError("Please refer the name of the profile you want to build")
	}
	f, err := os.ReadFile("lezgo.yaml")
	if err != nil {
		log.Fatal(err)
	}

	profile := pkg.Config{}

	err = yaml.Unmarshal(f, &profile)
	if err != nil {
		panic(err)
	}

	if profileName != "release" && profileName != "debug" {
		pkg.CustomError("Please refer the name of an existing profile")
	}

	var buildArgs []string
	var mkdirPath string
	switch profileName {
	case "release":
		buildArgs = profile.Profiles.Release.Flags
		buildArgs = append([]string{profile.Profiles.Release.Output}, buildArgs...)
		mkdirPath = profile.Profiles.Release.Output
		break
	case "debug":
		buildArgs = profile.Profiles.Debug.Flags
		mkdirPath = profile.Profiles.Debug.Output
	}
	fmt.Println("Starting the build process...")
	err = pkg.IsFileExist(profile.Entrypoint)
	if err != nil {
		pkg.CustomError("Entrypoint not found")
	}

	mkdirOutput := exec.Command("mkdir", mkdirPath)

	_, err = mkdirOutput.Output()
	if err != nil {
		pkg.CustomError(err.Error())
	}

	goBuild := exec.Command("go build -o", buildArgs...)
	_, err = goBuild.Output()
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
