package eventgrid

import (
	"log"
	"os"
)

// Config holds settings needed for Event Grid
type Config struct {
	TopicURL string
	SASKey   string
}

func LoadConfig() *Config {
	topicURL := os.Getenv("EVENT_GRID_TOPIC_URL")
	sasKey := os.Getenv("EVENT_GRID_SAS_KEY")

	if topicURL == "" || sasKey == "" {
		log.Fatal("Missing Event Grid config environment variables")
	}

	return &Config{
		TopicURL: topicURL,
		SASKey:   sasKey,
	}
}
