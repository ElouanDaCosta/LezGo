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
var crossCompilation bool

var buildUsage = `Build LezGo with given profile.

Usage: lezgo build [OPTIONS]

Options:
	-p, --profile     give the name of the profile
	-c, --cross-compilation     tell if the build should be using the cross-compilation option of the profile 
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

	var binaryNameOutput string

	var buildArgs []string
	var mkdirPath string
	switch profileName {
	case "release":
		binaryNameOutput = profile.Build.Profiles.Release.Output + profile.Name
		buildArgs = profile.Build.Profiles.Release.Flags
		buildArgs = buildArgsArray(buildArgs, binaryNameOutput, profile.Entrypoint)
		mkdirPath = profile.Build.Profiles.Release.Output
		break
	case "debug":
		binaryNameOutput = profile.Build.Profiles.Debug.Output + profile.Name
		buildArgs = profile.Build.Profiles.Debug.Flags
		buildArgs = buildArgsArray(buildArgs, binaryNameOutput, profile.Entrypoint)
		mkdirPath = profile.Build.Profiles.Debug.Output
		break
	}
	fmt.Println("Starting the build process...")
	err = pkg.IsFileExist(profile.Entrypoint)
	if err != nil {
		pkg.CustomError("Entrypoint not found")
	}

	if err := pkg.IsFileExist(mkdirPath); err != nil {
		mkdirOutput := exec.Command("mkdir", mkdirPath)

		_, err = mkdirOutput.Output()
		if err != nil {
			pkg.CustomError(err.Error())
		}
	}
	fmt.Println(buildArgs)
	goBuild := exec.Command("go", buildArgs...)
	_, err = goBuild.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("Build process complete successfully")
}

func buildArgsArray(array []string, binaryName string, output string) []string {
	array = append([]string{"build"}, array...)
	array = append(array, "-o")
	array = append(array, binaryName)
	array = append(array, output)
	return array
}

func NewBuildCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("", flag.ExitOnError),
		Execute: buildFunc,
	}

	cmd.flags.StringVar(&profileName, "profile", "", "Long declaration to give the name of the profile")
	cmd.flags.StringVar(&profileName, "p", "", "Short declaration to give the name of the profile")
	cmd.flags.BoolVar(&crossCompilation, "cross-compilation", false, "Long declaration to activate the cross-compilation option of the given profile")
	cmd.flags.BoolVar(&crossCompilation, "c", false, "Short declaration to activate the cross-compilation option of the given profile")

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, initUsage)
	}

	return cmd
}
