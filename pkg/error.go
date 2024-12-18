package pkg

import "fmt"

func Warning(message string) {
	var Yellow = "\033[33m"
	var Reset = "\033[0m"
	fmt.Println(Yellow + "Warning: " + message + Reset)
}
