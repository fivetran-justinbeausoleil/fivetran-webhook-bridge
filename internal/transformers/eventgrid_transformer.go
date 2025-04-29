package transformers

import (
	"github.com/google/uuid"
	"time"

	"github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge/models"
)

// TransformFivetranToEventGrid maps a FivetranEvent to the Event Grid transformers schema.
// It validates required fields and constructs a compatible EventGridEvent for publishing.
// Returns an error if validation fails.
func TransformFivetranToEventGrid(e *models.FivetranEvent) (models.EventGridEvent, error) {
	id := e.SyncID
	if id == "" {
		id = uuid.New().String() // Generate a random UUID if missing
	}

	data := e.Data
	if data == nil {
		data = []byte("{}") // Empty JSON object if no data
	}

	return models.EventGridEvent{
		ID:              id,
		EventType:       e.Event,
		Subject:         "/fivetran/webhook",
		EventTime:       time.Now().UTC(),
		Data:            data,
		DataVersion:     "1.0",
		MetadataVersion: "1",
		Topic:           "",
	}, nil
}
