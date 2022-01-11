package main

import (
	"encoding/json"
	"net/http"
)

func WebhookHome(writer http.ResponseWriter, request *http.Request) {
	response := map[string]interface{}{
		"status":  "Successful",
		"message": "Request Succesful",
	}

	json.NewEncoder(writer).Encode(response)
}

func Execute(writer http.ResponseWriter, request *http.Request) {
	var payload map[string]interface{}
	json.NewDecoder(request.Body).Decode(&payload)
	response := map[string]interface{}{
		"status":  "Successful",
		"message": "Request Successful",
		"data": map[string]interface{}{
			"challenge": payload["challenge"],
		},
	}
	json.NewEncoder(writer).Encode(response)
}
