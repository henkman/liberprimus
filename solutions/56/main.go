package main

import (
	"fmt"

	"github.com/henkman/liberprimus"
	"github.com/henkman/liberprimus/gematriaprimus"
	"github.com/henkman/liberprimus/utils"
)

func main() {
	page := liberprimus.Pages[56]
	pg := utils.MakePrimeGenerator()
	pi := 0
	ri := 0
	for _, r := range page {
		if r == gematriaprimus.SPACE {
			fmt.Print(" ")
		} else if r == '\n' {
			fmt.Println()
		} else if !gematriaprimus.IsRune(r) {
			fmt.Printf("%c", r)
		} else {
			o := gematriaprimus.RuneIndex(r)
			if ri != 56 { // 57th rune is unencrypted
				p := int(pg.GetNth(pi))
				o = (1 + o - p) % len(gematriaprimus.Letters)
				if o < 0 {
					o = len(gematriaprimus.Letters) + o
				}
				pi++
			}
			fmt.Print(gematriaprimus.Letters[o])
			ri++
		}
	}
	fmt.Println()
}
