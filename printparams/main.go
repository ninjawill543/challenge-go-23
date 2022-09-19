package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i := 1; i < len(arguments); i++ {
		for _, j := range os.Args[i] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
