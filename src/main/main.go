package main

import (
	"api"
	"net/http"
)

func main() {
	http.HandleFunc("/helloworld/read", api.ReadHandler)

	http.ListenAndServe(":3000", nil)
}
