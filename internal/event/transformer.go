package event

import (
	"time"

	"github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge/models"
)

func TransformFivetranToEventGrid(e *FivetranEvent) (models.EventGridEvent, error) {
	return models.EventGridEvent{
		ID:              e.SyncID,
		EventType:       e.Event,
		Subject:         "/fivetran/webhook",
		EventTime:       time.Now(),
		Data:            e.Data,
		DataVersion:     "1.0",
		MetadataVersion: "1",
		Topic:           "", // Event Grid fills in topic when using custom topics
	}, nil
}
