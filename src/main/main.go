package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", get)

	http.ListenAndServe(":3000", nil)
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
