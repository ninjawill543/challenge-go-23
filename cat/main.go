package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	open := "ERROR: open "
	no := ": no such file or directory"
	er := "ERROR: "
	Nop := ": No such file or directory"
	exit := "exit status 1"
	for i := 1; i < len(os.Args); i++ {
		if len(os.Args) == 2 {
			content, err := ioutil.ReadFile(os.Args[1])
			if err != nil {
				for _, a := range open {
					z01.PrintRune(a)
				}
				for _, b := range os.Args[2] {
					z01.PrintRune(b)
				}
				for _, c := range no {
					z01.PrintRune(c)
				}
				z01.PrintRune('\n')
				for _, d := range exit {
					z01.PrintRune(d)
				}
			}
			for _, g := range string(content) {
				z01.PrintRune(g)
			}
		} else if len(os.Args) > 2 {
			content, err := ioutil.ReadFile(os.Args[i])
			if err != nil {
				for _, e := range er {
					z01.PrintRune(e)
				}
				for _, f := range os.Args[i] {
					z01.PrintRune(f)
				}
				for _, g := range Nop {
					z01.PrintRune(g)
				}
			}
			for _, h := range string(content) {
				z01.PrintRune(h)
			}
		}
	}
}
