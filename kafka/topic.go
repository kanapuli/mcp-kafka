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

// DescribeTopic describes the specified topic.
// Note: The method accepts currently just a single topic though the underlying client can handle a slice of topics
// Let's leave the job to the LLM to call this method iteratively to handle multiple topics
func (c *client) DescribeTopic(topics string) ([]*sarama.TopicMetadata, error) {
	return c.admin.DescribeTopics([]string{topics})
}
