package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initializeRouter()
}

func initializeRouter() {
	router := mux.NewRouter()
	router.Use(commonMiddleware)
	router.HandleFunc("/", WebhookHome).Methods("GET")
	router.HandleFunc("/webhook/verify", WebhookVerify).Methods("POST")

	fmt.Println("App listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
