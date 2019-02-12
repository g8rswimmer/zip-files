package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type information struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

var appInformation = information{
	Name:    "ZIP Files Application",
	Version: "0.1.0",
}

func handleIndex(writer http.ResponseWriter, request *http.Request) {

	response, err := json.Marshal(appInformation)

	if err != nil {
		writer.Header().Set("Content-Type", "text/plain")
		writer.Write([]byte(err.Error()))
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, err := writer.Write(response)

		if err != nil {
			writer.Header().Set("Content-Type", "text/plain")
			writer.Write([]byte(err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}

}

func main() {
	fmt.Println("Starting the server on port 6050")
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":6050", nil)
}
