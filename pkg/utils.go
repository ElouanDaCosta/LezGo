package pkg

import (
	"errors"
	"fmt"
	"os"
)

func LezgoAsciiArt() {
	var Reset = "\033[0m"
	var Cyan = "\033[36m"
	const asciiArt = `
   __           _____   
  / /  ___ ___ / ___/__ 
 / /__/ -_)_ // (_ / _ \
/____/\__//__/\___/\___/

	`
	fmt.Println(Cyan + asciiArt + Reset)
}

func IsFileExist(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}
