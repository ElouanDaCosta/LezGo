package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ElouanDaCosta/LezGo/src"
)

var usage = `Usage: lezgo command [options]

A simple build and dependency management tool for Go, offering multiple build profiles and enhanced workflows for developers

Options:

Commands:

Init

`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}

	flag.Parse()
	if flag.NArg() < 1 {
		usageAndExit("")
	}

	var cmd *src.Command

	switch os.Args[1] {
	case "init":
		cmd = src.NewInitCommand()

	default:
		usageAndExit(fmt.Sprintf("lezgo: '%s' is not a lezgo command.\n", os.Args[1]))
	}

	cmd.Init(os.Args[2:])
	cmd.Run()
}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n")
	}

	flag.Usage()
	os.Exit(0)
}
