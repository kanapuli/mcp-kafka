package kafka

import (
	"github.com/IBM/sarama"
)

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
