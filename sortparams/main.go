package main

import (
	//"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	list := []rune{}
	arguments := os.Args
	for i := 1; i < len(arguments); i++ {
		for _, j := range os.Args[i] {
			list = append(list, j)
		}
	}
	for e := 0; e < len(list); e++ {
		for k := 32; k < 127; k++ {
			if rune(list[e]) == rune(k) {
				z01.PrintRune(rune(k))
			}
		}
	}
}
