package kafka

import (
	"github.com/IBM/sarama"
)

// CreateTopic creates a new topic with the given name, number of partitions, and replication factor.
func (c *client) CreateTopic(topic string, numPartitions int32, replicationFactor int16) error {
	topicDetail := &sarama.TopicDetail{
		NumPartitions:     numPartitions,
		ReplicationFactor: replicationFactor,
		ConfigEntries:     map[string]*string{},
	}
	if err := c.admin.CreateTopic(topic, topicDetail, false); err != nil {
		return err
	}
	return nil
}

// DeleteTopic deletes the specified topic.
func (c *client) DeleteTopic(topic string) error {
	if err := c.admin.DeleteTopic(topic); err != nil {
		return err
	}
	return nil
}

// ListTopics lists all topics and their details.
func (c *client) ListTopics() (map[string]sarama.TopicDetail, error) {
	topics, err := c.admin.ListTopics()
	if err != nil {
		return nil, err
	}
	return topics, nil
}
