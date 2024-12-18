package pkg

import (
	"fmt"
	"os"
)

func Warning(message string) {
	var Yellow = "\033[33m"
	var Reset = "\033[0m"
	fmt.Println(Yellow + "Warning: " + message + Reset)
}

func CustomError(message string) {
	var Red = "\033[31m"
	var Reset = "\033[0m"
	fmt.Println(Red + "Error: " + message + Reset)
	os.Exit(1)
}
