package models

import (
	"encoding/json"
	"time"
)

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
