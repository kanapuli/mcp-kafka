package kafka

import (
	"time"

	"github.com/IBM/sarama"
)

// Client is a kafka Client
type Client struct {
	admin            sarama.ClusterAdmin
	producer         sarama.SyncProducer
	consumer         sarama.ConsumerGroup
	bootstrapServers []string
	username         string
	password         string
	verbose          bool
	consumerGroupID  string
	messagesChan     chan string
}

// NewClient creates a new kafka client
func NewClient(opts ...kafkaOptions) (*Client, error) {
	client := Client{}
	config := sarama.NewConfig()
	// TODO: These configs should be really configurable by the LLM user
	config.Admin.Timeout = 3 * time.Second
	config.ClientID = "mcp-kafka"

	// Producer configs
	config.Producer.Compression = sarama.CompressionSnappy
	// Return.Successes is specific to SyncProducer in order to work.
	config.Producer.Return.Successes = true

	// Consumer configs
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

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

	if client.consumerGroupID == "" {
		client.consumerGroupID = "mcp-kafka-consumer"
	}

	consumerGroup, err := sarama.NewConsumerGroup(client.bootstrapServers, client.consumerGroupID, config)
	if err != nil {
		return nil, err
	}
	client.consumer = consumerGroup
	client.messagesChan = make(chan string, 1000)

	return &client, nil
}

func (c *Client) Close() error {
	if c.producer != nil {
		err := c.producer.Close()
		if err != nil {
			return err
		}
	}

	if c.consumer != nil {
		err := c.consumer.Close()
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
