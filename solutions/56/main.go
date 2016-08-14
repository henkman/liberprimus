package main

import (
	"fmt"

	"github.com/henkman/euler/sieve"
	"github.com/henkman/liberprimus"
	"github.com/henkman/liberprimus/gematriaprimus"
)

func main() {
	page := liberprimus.Pages[56]
	primes := make(chan uint64)
	go func(n uint64) {
		ps := sieve.Sieve(n)
		var i uint64
		for i = 0; i < n; i++ {
			if ps.IsPrime(i) {
				primes <- i
			}
		}
	}(uint64(len(page)))
	n := 0
	for _, r := range page {
		if r == gematriaprimus.SPACE {
			fmt.Print(" ")
		} else if r == '\n' {
			fmt.Println()
		} else if !gematriaprimus.IsRune(r) {
			fmt.Printf("%c", r)
		} else {
			o := gematriaprimus.RuneIndex(r)
			if n != 57-1 { // 57th rune is unencrypted
				p := int(<-primes)
				// o + -(p + 57)
				o = (o - p + 59) % len(gematriaprimus.Letters)
				if o < 0 {
					o = len(gematriaprimus.Letters) + o
				}
			}
			fmt.Print(gematriaprimus.Letters[o])
			n++
		}
	}
	fmt.Println()
}
