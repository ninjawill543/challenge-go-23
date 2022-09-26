package main

import "github.com/01-edu/z01"

func main() {
	stupid := "x = 42, y = 21"
	for _, i := range stupid {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
