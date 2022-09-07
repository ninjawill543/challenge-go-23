package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	a := '0'
	b := '0'
	c := '0'
	for c < '9' {
		c++
		if c == '9' && b != '9' {
			if c > b && b > a {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)
				z01.PrintRune(44)
				z01.PrintRune(32)
			}
			if c == '9' && b == '8' && a == '7' {
				goto Exit
			}
			b++
			c = 0
		}
		if b == '9' {
			a++
			b = 0
			if c == '9' && b == '8' && a == '7' {
				goto Exit
			}
		}
		if b == '9' && c == '9' {
			b = 0
			c = 0
			a++
			if c == '9' && b == '8' && a == '7' {
				goto Exit
			}
		}
		if c > b && b > a {
			z01.PrintRune(a)
			z01.PrintRune(b)
			z01.PrintRune(c)
			z01.PrintRune(44)
			z01.PrintRune(32)
			if c == '9' && b == '8' && a == '7' {
				goto Exit
			}
		}
	}
Exit:
}
