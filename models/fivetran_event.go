package models

import (
	"encoding/json"
	"time"
)

// FivetranEvent represents the structure of a webhook transformers received from Fivetran.
// It includes metadata about the transformers, connector details, and a flexible data payload.
// The Data field is left as raw JSON to accommodate different transformers types.
type FivetranEvent struct {
	Event              string          `json:"event"`
	Created            time.Time       `json:"created"`
	ConnectorType      string          `json:"connector_type"`
	ConnectorID        string          `json:"connector_id"`
	ConnectorName      string          `json:"connector_name"`
	SyncID             string          `json:"sync_id"`
	DestinationGroupID string          `json:"destination_group_id"`
	Data               json.RawMessage `json:"data"`
}
