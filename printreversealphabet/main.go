package main

import "github.com/01-edu/z01"

func main() {
	alpha := "zyxwvutsrqponmlkjihgfedcba"
	for _, c := range alpha {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')

}
