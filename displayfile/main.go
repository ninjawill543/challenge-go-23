package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("File name missing")
	}

	fmt.Println(string(content))
}
