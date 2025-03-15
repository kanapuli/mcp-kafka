package kakfa

import (
	"errors"
)

type kafkaOptions func(client *client) error

// WithBootstrapServers sets the bootstrap servers for the client.
func WithBootstrapServers(servers []string) kafkaOptions {
	return func(c *client) error {
		if len(servers) == 0 {
			return errors.New("bootstrap servers cannot be empty")
		}
		c.bootstrapServers = servers
		return nil
	}
}

// WithUsername sets the username for the client.
func WithUsername(username string) kafkaOptions {
	return func(c *client) error {
		c.username = username
		return nil
	}
}

// WithPassword sets the password for the client.
func WithPassword(password string) kafkaOptions {
	return func(c *client) error {
		c.password = password
		return nil
	}
}

// WithProducerTopic sets the topic for the producer.
func WithProducerTopic(topic string) kafkaOptions {
	return func(c *client) error {
		if topic == "" {
			return errors.New("topic cannot be empty")
		}
		c.producer.topic = topic
		return nil
	}
}

// WithVerbose sets the verbose flag for the client.
func WithVerbose(verbose bool) kafkaOptions {
	return func(c *client) error {
		c.verbose = verbose
		return nil
	}
}
