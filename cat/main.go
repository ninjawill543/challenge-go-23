package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	open := "ERROR: open "
	no := ": no such file or directory"
	er := "ERROR: "
	Nop := ": No such file or directory"
	exit := "exit status 1"
	for i := 1; i <= len(os.Args); i++ {
		if len(os.Args) == 2 {
			content, err := ioutil.ReadFile(os.Args[2])
			if err != nil {
				fmt.Printf(open)
				fmt.Printf(os.Args[2])
				fmt.Printf(no)
				fmt.Printf("\n")
				fmt.Print(exit)
			}
			fmt.Printf(string(content))
		} else if len(os.Args) > 2 {
			content, err := ioutil.ReadFile(os.Args[i])
			if err != nil {
				fmt.Printf(er)
				fmt.Printf(os.Args[i])
				fmt.Printf(Nop)
			}
			fmt.Printf(string(content))
		}
	}
}
