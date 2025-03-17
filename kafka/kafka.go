package kafka

import (
	"time"

	"github.com/IBM/sarama"
)

// Client is a kafka Client
type Client struct {
	admin            sarama.ClusterAdmin
	producer         sarama.SyncProducer
	bootstrapServers []string
	username         string
	password         string
	verbose          bool
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
	config := sarama.NewConfig()
	// TODO: These configs should be really configurable by the LLM user
	config.Admin.Timeout = 3 * time.Second
	config.ClientID = "mcp-kafka"
	config.Producer.Compression = sarama.CompressionSnappy
	// Return.Successes is specific to SyncProducer in order to work.
	config.Producer.Return.Successes = true

	for _, opt := range opts {
		err := opt(&client)
		if err != nil {
			return nil, err
		}
	}

	admin, err := sarama.NewClusterAdmin(client.bootstrapServers, config)
	if err != nil {
		return nil, err
	}
	client.admin = admin

	producer, err := sarama.NewSyncProducer(client.bootstrapServers, config)
	if err != nil {
		return nil, err
	}
	client.producer = producer

	return &client, nil
}

func (c *Client) Close() error {
	// The close calls must be called in the correct order
	if c.producer != nil {
		err := c.producer.Close()
		if err != nil {
			return err
		}
	}
	if c.admin != nil {
		err := c.admin.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
