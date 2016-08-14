package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"

	"github.com/henkman/liberprimus"
	"github.com/henkman/liberprimus/gematriaprimus"
)

func xor(dst, src []uint) {
	for i, _ := range dst {
		dst[i] ^= src[i%len(src)]
	}
}

func frombase(s string, base uint, digits string) (uint, error) {
	var e, n uint
	for i := len(s) - 1; i >= 0; i-- {
		v := strings.Index(digits[:base], string(s[i]))
		if v == -1 {
			return 0, fmt.Errorf("unknown digit %c at position %d", s[i], i)
		}
		n = n + uint(v)*uint(math.Pow(float64(base), float64(e)))
		e++
	}
	return n, nil
}

func main() {
	if false {
		pairs := make([]string, 0, 100)
		rePair := regexp.MustCompile("[a-zA-Z0-9]{2}")
		for _, page := range liberprimus.Pages[49:51] {
			pairs = append(pairs, rePair.FindAllString(page, -1)...)
		}
		d := map[byte]uint{}
		for _, pair := range pairs {
			if _, ok := d[pair[0]]; ok {
				d[pair[0]]++
			} else {
				d[pair[0]] = 1
			}
			if _, ok := d[pair[1]]; ok {
				d[pair[1]]++
			} else {
				d[pair[1]] = 1
			}
		}
		for k, v := range d {
			fmt.Printf("%c:%d\n", k, v)
		}
		return
	}

	if true {
		pairs := make([]string, 0, 100)
		rePair := regexp.MustCompile("[a-zA-Z0-9]{2}")
		for _, page := range liberprimus.Pages[49:51] {
			pairs = append(pairs, rePair.FindAllString(page, -1)...)
		}
		// const digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		const digits = "0123456789abcdeghijklmnopqrstuvwxABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for _, pair := range pairs {
			x, err := frombase(pair, uint(len(digits)), digits)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(x)
		}
		return
	}

	if false {
		rect := [][]uint{
			{3258, 3222, 3152, 3038},
			{3278, 3299, 3298, 2838},
			{3288, 3294, 3296, 2472},
			{4516, 1206, 708, 1820},
		}
		for nl, line := range rect {
			for nc, col := range line {
				if nl == 3 && nc <= 1 {
					fmt.Printf("%d ", 3301+col)
				} else {
					fmt.Printf("%d ", 3301-col)
				}

			}
			fmt.Println()
		}
		fmt.Println(4516 - 1206)
		return
	}

	if false {
		values := make([][]uint, len(liberprimus.Pages))
		for i, page := range liberprimus.Pages {
			values[i] = make([]uint, 0, 100)
			for _, r := range page {
				if gematriaprimus.IsRune(r) {
					values[i] = append(values[i], gematriaprimus.RuneToValue(r))
				}
			}
		}
		for i, pvs := range values {
			fmt.Println("---", i, "---")
			for _, pv := range pvs {
				fmt.Printf("%d ", pv)
			}
			fmt.Println()
		}
	}
}
