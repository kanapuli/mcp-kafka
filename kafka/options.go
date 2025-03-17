package kafka

import (
	"errors"
)

type kafkaOptions func(client *Client) error

// WithBootstrapServers sets the bootstrap servers for the client.
func WithBootstrapServers(servers []string) kafkaOptions {
	return func(c *Client) error {
		if len(servers) == 0 {
			return errors.New("bootstrap servers cannot be empty")
		}
		c.bootstrapServers = servers
		return nil
	}
}

// WithUsername sets the username for the client.
func WithUsername(username string) kafkaOptions {
	return func(c *Client) error {
		c.username = username
		return nil
	}
}

// WithPassword sets the password for the client.
func WithPassword(password string) kafkaOptions {
	return func(c *Client) error {
		c.password = password
		return nil
	}
}

// WithVerbose sets the verbose flag for the client.
func WithVerbose(verbose bool) kafkaOptions {
	return func(c *Client) error {
		c.verbose = verbose
		return nil
	}
}

func WithConsumerGroupID(id string) kafkaOptions {
	return func(c *Client) error {
		if id == "" {
			id = "mcp-kafka-consumer"
		}
		c.consumerGroupID = id
		return nil
	}
}
