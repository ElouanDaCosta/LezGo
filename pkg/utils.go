package pkg

import "fmt"

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
