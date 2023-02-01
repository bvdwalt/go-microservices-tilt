package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, from the 'world' microservice!")
	})

	http.ListenAndServe(":8082", nil)
}
