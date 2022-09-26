package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file := os.Args[1]
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("File name missing")
	}

	fmt.Println(string(content))
}
