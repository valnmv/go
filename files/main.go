package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readFile() {
	bytes, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(bytes))
}

func main() {
	readFile()
}
