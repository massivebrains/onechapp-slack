package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	initializeRouter()
}

func initializeRouter() {
	router := mux.NewRouter()
	router.Use(commonMiddleware)
	router.HandleFunc("/", WebhookHome).Methods("GET")
	router.HandleFunc("/webhook", Execute).Methods("POST")

	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "3000"
	}

	message := fmt.Sprintf("App listening on port %s", port)
	fmt.Println(message)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
