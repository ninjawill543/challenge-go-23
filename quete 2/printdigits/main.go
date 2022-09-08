package main

import "github.com/01-edu/z01"

func main() {
	alpha := "0123456789"
	for _, c := range alpha {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
