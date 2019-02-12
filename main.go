package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type information struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type listFiles struct {
	Name string `json:"name"`
	File string `json:"file"`
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
		_, err := writer.Write(response)

		if err != nil {
			writer.Header().Set("Content-Type", "text/plain")
			writer.Write([]byte(err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
		} else {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
		}

	}

}

func handleListFiles(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(request.Body)

		if err != nil {
			sendError(writer, err)
			return
		}

		var payload listFiles
		err = json.Unmarshal(body, &payload)

		if err != nil {
			sendError(writer, err)
			return
		}

		resp, err := json.Marshal(payload)

		if err != nil {
			sendError(writer, err)
			return
		}

		err = createFile(payload.Name, payload.File)
		if err != nil {
			sendError(writer, err)
			return
		}

		_, err = writer.Write(resp)

		if err != nil {
			sendError(writer, err)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		return
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Header().Set("Content-Type", "text/plain")
		writer.Write([]byte(fmt.Sprintf("Http Method %s is not supported", request.Method)))
	}
}

func sendError(writer http.ResponseWriter, err error) {
	writer.WriteHeader(http.StatusInternalServerError)
	writer.Header().Set("Content-Type", "text/plain")
	writer.Write([]byte(err.Error()))
}

func main() {
	fmt.Println("Starting the server on port 6050")

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/listFiles", handleListFiles)
	http.ListenAndServe(":6050", nil)
}
