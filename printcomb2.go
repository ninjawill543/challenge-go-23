package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	a := '0'
	b := '0'
	c := '0'
	d := '0'
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			if d == '9' {
				if d > b && c > a {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(' ')
					z01.PrintRune(c)
					z01.PrintRune(d)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				d = '0'
				c++
			}
			d++
			if d > b && c >= a {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(' ')
				z01.PrintRune(c)
				z01.PrintRune(d)
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			if d == '9' && c == '9' {
				b++
				d = '0'
				c = '0'
			}
			if b == '9' {
				b = '0'
				a++
			}
		}
	}
}
