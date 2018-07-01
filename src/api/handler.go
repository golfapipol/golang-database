package api

import (
	"encoding/json"
	"net/http"
)

type Sayhi struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

type Api struct {
	ReadFunc func(string) (Sayhi, error)
}

func (api Api) ReadHandler(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	sayhi, err := api.ReadFunc(id)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	responseSayhi, err := json.Marshal(sayhi)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	response.Write(responseSayhi)
}
