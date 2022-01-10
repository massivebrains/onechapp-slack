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
	router.HandleFunc("/webhook/verify", WebhookVerify).Methods("POST")

	fmt.Println("App listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
