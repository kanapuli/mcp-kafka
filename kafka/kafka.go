package kakfa

// client is a kafka client
type client struct {
	bootstrapServers []string
	username         string
	password         string
	topic            []string
}

type kafkaOptions func(client *client)

func WithBootstrapServers(servers []string) kafkaOptions {
	return func(c *client) {
		c.bootstrapServers = servers
	}
}

func WithUsername(username string) kafkaOptions {
	return func(c *client) {
		c.username = username
	}
}

func WithPassword(password string) kafkaOptions {
	return func(c *client) {
		c.password = password
	}
}

func WithTopic(topic []string) kafkaOptions {
	return func(c *client) {
		c.topic = topic
	}
}
