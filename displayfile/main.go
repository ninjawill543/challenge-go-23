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
			fmt.Print("File name missing")
		}
		fmt.Println(string(content))
	} else {
		fmt.Print("File name missing")
	}
}
