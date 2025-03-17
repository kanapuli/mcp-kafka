package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Produce sends a message to the Kafka topic.
func (c *Client) Produce(topic string, key []byte, message []byte) (string, error) {
	if key == nil {
		key = []byte(uuid.New().String())
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
		Key:   sarama.StringEncoder(key),
	}

	partition, offset, err := c.producer.SendMessage(msg)
	if err != nil {
		return "", err
	}
	responseMsg := fmt.Sprintf("message successfully produced to the Topic: %s, Partition: %d, Offset: %d", topic, partition, offset)
	zap.S().Infof(responseMsg)
	return responseMsg, nil
}
