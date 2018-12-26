package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User is a struct with GitHub user personal data
type User struct {
	Name string `json:"name"`
	Blog string `json:"blog"`
}

func encodeJSON() {
	type Person struct {
		Name    string   `json:"name,omitempty"`
		Age     int      `json:"age,omitempty"`
		Hobbies []string `json:"hobbies,omitempty"`
		Married bool     `json:"married,omitempty"`
	}

	hobbies := []string{"Cycling", "Cheese", "Techno"}
	p := Person{
		Name:    "George",
		Age:     40,
		Hobbies: hobbies,
	}
	fmt.Printf("%+v\n", p)
	jsonByteData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	jsonStringData := string(jsonByteData)
	fmt.Println(jsonStringData)
}

func decodeJSON() {
	var u User
	res, err := http.Get("https://api.github.com/users/shapeshed")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", u)
}

func main() {
	encodeJSON()
	decodeJSON()
}
