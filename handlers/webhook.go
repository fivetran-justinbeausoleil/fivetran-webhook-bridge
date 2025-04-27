package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge/internal/event"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var incomingFivetranEvent event.FivetranEvent
	if err := json.NewDecoder(r.Body).Decode(&incomingFivetranEvent); err != nil {
		log.Printf("Error decoding incoming webhook: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Printf("Received webhook: EventType=%s, ConnectorID=%s", incomingFivetranEvent.Event, incomingFivetranEvent.ConnectorID)

	w.WriteHeader(http.StatusAccepted)
}
