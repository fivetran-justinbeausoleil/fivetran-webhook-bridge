package event

import (
	"encoding/json"
	"time"
)

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
