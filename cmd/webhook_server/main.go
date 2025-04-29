package main

import (
	"log"
	"net/http"

	"github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge/internal/config"
	"github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge/internal/handlers"
	"github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge/pkg/eventgrid"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	eventGridCfg := config.LoadConfig()
	eventGridClient := eventgrid.NewClient(eventGridCfg.EventGrid.TopicURL, eventGridCfg.EventGrid.SASKey)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /", healthCheck)
	mux.HandleFunc("POST /webhook/eventgrid", handlers.EventGrid(eventGridClient))

	log.Println("Server starting on :8080 ...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
