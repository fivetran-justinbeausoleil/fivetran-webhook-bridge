package models

import (
	"encoding/json"
	"time"
)

// EventGridEvent defines the structure of an event expected by Azure Event Grid.
// It includes metadata such as ID, event type, subject, and timestamps.
// The Data field contains the event payload and is left as raw JSON to support flexible schemas.
type EventGridEvent struct {
	ID              string          `json:"id"`
	EventType       string          `json:"eventType"`
	Subject         string          `json:"subject"`
	EventTime       time.Time       `json:"eventTime"`
	Data            json.RawMessage `json:"data"`
	DataVersion     string          `json:"dataVersion"`
	MetadataVersion string          `json:"metadataVersion"`
	Topic           string          `json:"topic"`
}
