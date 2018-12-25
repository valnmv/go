package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-My-Header", "I am setting a header!")
	w.Write([]byte("Hello World\n"))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
