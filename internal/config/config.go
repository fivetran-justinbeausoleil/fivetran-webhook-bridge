package config

import (
	"log"
	"os"
)

// Config holds settings needed for all external services used by the webhook server.
type Config struct {
	EventGrid EventGridConfig
}

// EventGridConfig holds Azure Event Grid-specific configuration such as the
// topic URL and SAS authentication key.
type EventGridConfig struct {
	TopicURL string
	SASKey   string
}

// LoadConfig reads required environment variables and constructs a Config object.
// It logs a fatal error and exits the program if required values are missing.
func LoadConfig() *Config {
	topicURL := os.Getenv("EVENT_GRID_TOPIC_URL")
	sasKey := os.Getenv("EVENT_GRID_SAS_KEY")

	if topicURL == "" || sasKey == "" {
		log.Fatal("Missing Event Grid config environment variables")
	}

	return &Config{
		EventGrid: EventGridConfig{
			TopicURL: topicURL,
			SASKey:   sasKey,
		},
	}
}
