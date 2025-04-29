package handlers

import (
	"encoding/json"
	"github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge/pkg/sender"
	"log"
	"net/http"

	"github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge/internal/transformers"
	"github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge/models"
)

// EventGrid returns an HTTP handler for processing incoming Fivetran webhook events.
// It parses the incoming JSON payload, transforms it into Azure Event Grid format,
// and sends the resulting transformers to an Event Grid custom topic.
// Expects a POST request and returns 202 Accepted on success.
func EventGrid(eventSender sender.Sender) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		var incomingFivetranEvent models.FivetranEvent
		if err := json.NewDecoder(r.Body).Decode(&incomingFivetranEvent); err != nil {
			log.Printf("Error decoding incoming webhook: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		log.Printf("Received webhook: EventType=%s, ConnectorID=%s", incomingFivetranEvent.Event, incomingFivetranEvent.ConnectorID)

		outgoingEvent, err := transformers.EventGridTransformer(&incomingFivetranEvent)
		if err != nil {
			log.Printf("Failed to transform transformers: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		err = eventSender.Send([]models.EventGridEvent{outgoingEvent})
		if err != nil {
			log.Printf("Failed to send transformers: %v", err)
			http.Error(w, "Failed to deliver transformers", http.StatusInternalServerError)
			return
		}

		log.Printf("Successfully sent transformers to Event Grid: ID=%s", outgoingEvent.ID)

		w.WriteHeader(http.StatusAccepted)
	}
}
