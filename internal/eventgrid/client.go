package eventgrid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge/models"
	"net/http"
)

type Config struct {
	TopicURL string `json:"topic_url"`
	SASKey   string `json:"sas_key"`
}

// SendEvent posts the Event Grid events to Azure
func SendEvent(cfg *Config, events []models.EventGridEvent) error {
	payload, err := json.Marshal(events)
	if err != nil {
		return fmt.Errorf("failed to marshal events: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, cfg.TopicURL, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("aeg-sas-key", cfg.SASKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("received non-2xx response from Event Grid: %s", resp.Status)
	}

	return nil
}
