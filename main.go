package main

import (
	"log"
	"net/http"

	"github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge/handlers"
)

func returnStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", returnStatus)
	mux.HandleFunc("/webhook", handlers.WebhookHandler)

	log.Println("Server starting on :8080 ...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
