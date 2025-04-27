package event

import (
	"fmt"
	"time"

	"github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge/models"
)

func TransformFivetranToEventGrid(e *FivetranEvent) (models.EventGridEvent, error) {
	if e.ConnectorID == "" {
		return models.EventGridEvent{}, fmt.Errorf("missing required field: connector_id")
	}

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
