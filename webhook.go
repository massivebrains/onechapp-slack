package main

import (
	"encoding/json"
	"net/http"
)

type VerifyPayload struct {
	Type      string `json:"type"`
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
}

func WebhookHome(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode("I'm Up.")
}

func WebhookVerify(writer http.ResponseWriter, request *http.Request) {
	var verifyPayload VerifyPayload
	json.NewDecoder(request.Body).Decode(&verifyPayload)
	responseData := map[string]interface{}{
		"challenge": verifyPayload.Challenge,
	}
	json.NewEncoder(writer).Encode(responseData)
}
