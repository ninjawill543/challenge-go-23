package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file := os.Args[1]

	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

}
