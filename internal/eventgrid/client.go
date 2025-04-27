package eventgrid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fivetran-justinbeausoleil/fivetran-webhook-bridge/models"
)

type Client struct {
	TopicURL string
	SASKey   string
}

// NewClient creates a new Event Grid client
func NewClient(topicURL, sasKey string) *Client {
	return &Client{
		TopicURL: topicURL,
		SASKey:   sasKey,
	}
}

func (c *Client) Send(payload any) error {
	events, ok := payload.([]models.EventGridEvent)
	if !ok {
		return fmt.Errorf("invalid payload type for Event Grid sender")
	}
	data, err := json.Marshal(events)
	if err != nil {
		return fmt.Errorf("failed to marshal events: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.TopicURL, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("aeg-sas-key", c.SASKey)

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
