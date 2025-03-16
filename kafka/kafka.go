package kafka

import (
	"time"

	"github.com/IBM/sarama"
)

// Client is a kafka Client
type Client struct {
	saramaClient     sarama.Client
	admin            sarama.ClusterAdmin
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
func NewClient(opts ...kafkaOptions) (*Client, error) {
	client := Client{}
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

	admin, err := sarama.NewClusterAdmin(client.bootstrapServers, config)
	if err != nil {
		return nil, err
	}

	client.saramaClient = saramaClient
	client.admin = admin

	return &client, nil
}

func (c *Client) Close() error {
	return c.saramaClient.Close()
}
