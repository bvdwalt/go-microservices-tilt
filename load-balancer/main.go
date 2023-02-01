package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	//routes and where they should be redirected to
	routes := map[string]string{
		"/hello": "http://localhost:8081/hello",
		"/world": "http://localhost:8082/world",
	}

	for k, v := range routes {
		r.HandleFunc(k, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, v, http.StatusPermanentRedirect)
		})
	}

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "Welcome to the loadbalancer, here are the routes available:\n%s\n\n\n%v", printRoutes(routes), routes)
	})

	http.ListenAndServe(":8080", r)
}

func printRoutes(routes map[string]string) string {
	s := "<ol>"
	s += "<li><a href=\"http://localhost:8080\">/index</li>"
	for k, _ := range routes {
		s += fmt.Sprintf("<li><a href=\"http://localhost:8080%s\">%s</a></li>\n", k, k)
	}
	s += "</ol>"

	return s
}
