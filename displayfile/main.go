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
			fmt.Println("File name missing")
		}
		fmt.Println(string(content))
	} else {
		fmt.Println("File name missing")
	}
}
