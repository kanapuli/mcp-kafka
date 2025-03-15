package kafka

import (
	"time"

	"github.com/IBM/sarama"
)

// client is a kafka client
type client struct {
	saramaClient     sarama.Client
	bootstrapServers []string
	username         string
	password         string
	verbose          bool
	producer
	consumer
}

type producer struct {
	topic string
}

type consumer struct {
	topic   []string
	groupID string
}

// NewClient creates a new kafka client
func NewClient(opts ...kafkaOptions) (*client, error) {
	client := client{}
	for _, opt := range opts {
		err := opt(&client)
		if err != nil {
			return nil, err
		}
	}

	config := sarama.NewConfig()
	config.Admin.Timeout = 3 * time.Second
	saramaClient, err := sarama.NewClient(client.bootstrapServers, config)
	if err != nil {
		return nil, err
	}

	client.saramaClient = saramaClient

	return &client, nil
}

func (c *client) Close() error {
	return c.saramaClient.Close()
}
