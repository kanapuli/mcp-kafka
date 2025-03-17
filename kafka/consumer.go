package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type ConsumerHandler struct {
	ready chan bool
	// MessageHandler is a function that handles a message from the Kafka consumer.
	MessageHandler   func(message *sarama.ConsumerMessage) error
	consumedMessages []string
}

// Setup is called at the beginning of a new session, before ConsumeClaim
func (consumer *ConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

// Cleanup is called at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *ConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if err := consumer.MessageHandler(message); err != nil {
			return err
		}
		session.MarkMessage(message, "")
	}
	return nil
}

func (c *Client) ConsumeWithoutTimeout(topics []string, messageHandler func(*sarama.ConsumerMessage) error, timeout time.Duration) error {
	handler := &ConsumerHandler{
		ready:          make(chan bool),
		MessageHandler: messageHandler,
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	go func() {
		for {
			if err := c.consumer.Consume(ctx, topics, handler); err != nil {
				c.logger.Errorf("Error consuming messages: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			handler.ready = make(chan bool)
		}
	}()

	<-handler.ready
	zap.S().Infof("Kafka consumer started for topics: %v", topics)

	select {
	case <-ctx.Done():
		zap.S().Infof("Kafka consumer stopped due to context cancellation")
	}
	return nil
}

func (c *Client) SimpleConsume(topics []string, timeout time.Duration) ([]string, error) {
	defaultHandler := func(message *sarama.ConsumerMessage) error {
		msg := fmt.Sprintf("Message received: topic=%s, partition=%d, offset=%d, key=%s, value=%s", message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
		select {
		case c.messagesChan <- msg:
		// message added to the channel
		default:
			c.logger.Debugf("Message dropped because channel is full: %s", msg)
		}
		return nil
	}

	zap.S().Infof("SimpleConsume is called for topics: %v and timeout: %v ", topics, timeout)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	go c.ConsumeWithoutTimeout(topics, defaultHandler, timeout)
	var messages []string
	for {
		select {
		case msg := <-c.messagesChan:
			messages = append(messages, msg)
			c.logger.Debugf("Message received: %s", msg)
		case <-ctx.Done():
			return messages, nil
		}
	}
}
