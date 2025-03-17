package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Produce sends a message to the Kafka topic.
func (c *Client) Produce(topic string, key []byte, message []byte, headers map[string]any) (string, error) {
	if key == nil {
		key = []byte(uuid.New().String())
	}
	var (
		msgHeaders   []sarama.RecordHeader
		keyBytes     []byte
		messageBytes []byte
	)
	for k, v := range headers {
		// reset keyBytes and messageBytes
		keyBytes = keyBytes[:0]
		messageBytes = messageBytes[:0]

		keyBytes = fmt.Appendf(keyBytes, "%s", k)
		messageBytes = fmt.Appendf(messageBytes, "%v", v)

		msgHeaders = append(msgHeaders, sarama.RecordHeader{
			Key:   keyBytes,
			Value: messageBytes,
		})
	}

	msg := &sarama.ProducerMessage{
		Topic:   topic,
		Value:   sarama.StringEncoder(message),
		Key:     sarama.StringEncoder(key),
		Headers: msgHeaders,
	}

	partition, offset, err := c.producer.SendMessage(msg)
	if err != nil {
		return "", err
	}
	responseMsg := fmt.Sprintf("message successfully produced to the Topic: %s, Partition: %d, Offset: %d", topic, partition, offset)
	zap.S().Infof(responseMsg)
	return responseMsg, nil
}
