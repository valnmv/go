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

func writeFile() {
	s := "Hello World"
	err := ioutil.WriteFile("example03.txt", []byte(s), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func readDir() {
	files, err := ioutil.ReadDir("c:/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Mode(), file.Name())
	}
}

func main() {
	readFile()
	writeFile()
	readDir()
}
