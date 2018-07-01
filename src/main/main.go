package main

import (
	apiLibrary "api"
	"net/http"
)

func main() {
	stubReadFunc := func(string) (apiLibrary.Sayhi, error) {
		return apiLibrary.Sayhi{Id: 1, Description: "hello world"}, nil
	}
	api := apiLibrary.Api{
		ReadFunc: stubReadFunc,
	}
	http.HandleFunc("/helloworld/read", api.ReadHandler)

	http.ListenAndServe(":3000", nil)
}
