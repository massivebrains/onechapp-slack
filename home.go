package main

import (
	"encoding/json"
	"net/http"
)

func HomeDefault(writer http.ResponseWriter, request *http.Request) {
	response := map[string]interface{}{
		"status":  "Successful",
		"message": "Request Succesful",
	}

	json.NewEncoder(writer).Encode(response)
}
