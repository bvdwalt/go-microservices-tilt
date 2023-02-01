package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, from the 'hello' microservice!")
	})

	http.ListenAndServe(":8081", nil)
}
