package pkg

import "fmt"

func LezgoAsciiArt() {
	const asciiArt = `
   __           _____   
  / /  ___ ___ / ___/__ 
 / /__/ -_)_ // (_ / _ \
/____/\__//__/\___/\___/

	`
	fmt.Printf("%+v\n", asciiArt)
}
