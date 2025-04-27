package main

import (
	"log"
	"net/http"

	"github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge/handlers"
	"github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge/internal/eventgrid"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	eventGridCfg := eventgrid.LoadConfig()
	eventGridClient := eventgrid.NewClient(eventGridCfg.TopicURL, eventGridCfg.SASKey)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /", healthCheck)
	mux.HandleFunc("POST /webhook/eventgrid", handlers.EventGrid(eventGridClient))

	log.Println("Server starting on :8080 ...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
