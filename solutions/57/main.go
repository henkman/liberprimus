package main

import (
	"fmt"

	"github.com/henkman/liberprimus"
	"github.com/henkman/liberprimus/gematriaprimus"
)

func main() {
	for _, r := range liberprimus.Pages[57] {
		if r == gematriaprimus.SPACE {
			fmt.Print(" ")
		} else if r == '\n' {
			fmt.Println()
		} else if gematriaprimus.IsRune(r) {
			fmt.Print(gematriaprimus.RuneToLetter(r))
		}
	}
	fmt.Println()
}
