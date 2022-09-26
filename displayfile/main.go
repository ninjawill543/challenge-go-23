package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		content, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("File name missing")
			fmt.Printf("\n")
		}
		fmt.Printf(string(content))
	} else {
		fmt.Printf("File name missing")
		fmt.Printf("\n")
	}
}
