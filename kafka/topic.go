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
	return c.admin.DeleteTopic(topic)
}

// ListTopics lists all topics and their details.
func (c *client) ListTopics() (map[string]sarama.TopicDetail, error) {
	return c.admin.ListTopics()
}

// DescribeTopics describes the specified topics.
func (c *client) DescribeTopics(topics []string) ([]*sarama.TopicMetadata, error) {
	return c.admin.DescribeTopics(topics)
}
