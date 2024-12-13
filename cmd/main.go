package main

import (
	"flag"
	"fmt"
	"os"
)

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n")
	}

	flag.Usage()
	os.Exit(0)
}

var usage = `Usage: lezgo command [options]

A simple build and dependency management tool for Go, offering multiple build profiles and enhanced workflows for developers

Options:

Commands:
  
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}
	usageAndExit("")
}
