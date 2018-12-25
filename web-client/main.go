package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func getExample() {
	response, err := http.Get("https://ifconfig.co/")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", body)
}

func postExample() {
	postData := strings.NewReader(`{"some": "json"}`)
	response, err := http.Post("https://httpbin.org/post", "application/json", postData)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", body)
}

func customClientExample() {
	//	debug := os.Getenv("DEBUG")
	client := &http.Client{
		Timeout: 1000 * time.Millisecond,
	}
	request, err := http.NewRequest("GET", "https://ifconfig.co", nil)
	request.Header.Add("Accept", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	// if debug == "1" {
	// 	debugRequest, err := httputil.DumpRequestOut(request, true)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Printf("%s", debugRequest)
	// }

	response, err := client.Do(request)

	// if debug == "1" {
	// 	debugResponse, err := httputil.DumpResponse(response, true)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Printf("%s", debugResponse)
	// }

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	fmt.Printf("%s", body)
}

func main() {
	getExample()
	postExample()
	customClientExample()
}
