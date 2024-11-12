package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, this page is deployed on Railway\n")
	fmt.Fprintf(w, "By: Mateo Pillajo :D\n")
	fmt.Fprintf(w, "Made with Golang\n")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
