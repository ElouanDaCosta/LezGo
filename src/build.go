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

	if crossCompilation {
		crossCompilationProfileBuild(profile)
		return
	}

	profileBuildArgs, profilePath := profileBuild(profile)

	fmt.Println("Starting the build process...")
	err = pkg.IsFileExist(profile.Entrypoint)
	if err != nil {
		pkg.CustomError("Entrypoint not found")
	}

	if err := pkg.IsFileExist(profilePath); err != nil {
		mkdirOutput := exec.Command("mkdir", profilePath)

		_, err = mkdirOutput.Output()
		if err != nil {
			pkg.CustomError(err.Error())
		}
	}
	fmt.Println(profileBuildArgs)
	goBuild := exec.Command("go", profileBuildArgs...)
	_, err = goBuild.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("Build process complete successfully")
}

func crossCompilationProfileBuild(profile pkg.Config) {
	var crossCompilation string
	var binaryNameOutput string
	var buildArgs []string
	switch profileName {
	case "release":
		crossCompilation = fmt.Sprintf(`GOOS=%v GOARCH=%v`, profile.Build.Profiles.Release.OsTarget, profile.Build.Profiles.Release.Arch)
		binaryNameOutput = fmt.Sprintf(`%v_%v`, profile.Name, profile.Build.Profiles.Release.Arch)
		buildArgs = append(buildArgs, profile.Build.Profiles.Release.Output+binaryNameOutput)
		buildArgs = append(buildArgs, profile.Entrypoint)
	case "debug":
		crossCompilation = fmt.Sprintf(`GOOS=%v GOARCH=%v`, profile.Build.Profiles.Debug.OsTarget, profile.Build.Profiles.Release.Arch)
		binaryNameOutput = fmt.Sprintf(`%v_%v`, profile.Name, profile.Build.Profiles.Debug.Arch)
		buildArgs = append(buildArgs, profile.Build.Profiles.Debug.Output+binaryNameOutput)
		buildArgs = append(buildArgs, profile.Entrypoint)
	}
	buildInput := []string{crossCompilation, "go build -o", buildArgs[0], buildArgs[1]}
	goBuild := exec.Command("env", buildInput...)
	fmt.Println(goBuild)
	if err := goBuild.Run(); err != nil {
		fmt.Println("Error running command:", err)
		return
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

func profileBuild(profile pkg.Config) ([]string, string) {
	var binaryNameOutput string
	var buildArgs []string
	switch profileName {
	case "release":
		binaryNameOutput = profile.Build.Profiles.Release.Output + profile.Name
		buildArgs = profile.Build.Profiles.Release.Flags
		buildArgs = buildArgsArray(buildArgs, binaryNameOutput, profile.Entrypoint)
		mkdirPath := profile.Build.Profiles.Release.Output
		return buildArgs, mkdirPath
	case "debug":
		binaryNameOutput = profile.Build.Profiles.Debug.Output + profile.Name
		buildArgs = profile.Build.Profiles.Debug.Flags
		buildArgs = buildArgsArray(buildArgs, binaryNameOutput, profile.Entrypoint)
		mkdirPath := profile.Build.Profiles.Debug.Output
		return buildArgs, mkdirPath
	}
	return nil, ""
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
