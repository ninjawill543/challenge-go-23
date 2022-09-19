package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := []rune(os.Args[0])
	for i := 2; i < len(arr); i++ {
		z01.PrintRune(arr[i])
	}
	z01.PrintRune('\n')
}
