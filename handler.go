package main

import (
	"context"
	"fmt"

	kafka "github.com/kanapuli/mcp-kafka/kafka"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

// KafkaHandler is a struct that handles Kafka operations for the mcp-kafka tool
type KafkaHandler struct{}

// CreateTopic creates a new Kafka topic
// Optional parameters that can be passed via FuncArgs are:
// - NumPartitions: number of partitions for the topic
// - ReplicationFactor: replication factor for the topic
func (k *KafkaHandler) CreateTopic(ctx context.Context, req Request) (*mcp_golang.ToolResponse, error) {

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	kafkaClient, err := kafka.NewClient(kafka.WithBootstrapServers([]string{"localhost:9092"}))
	if err != nil {
		return nil, err
	}
	defer kafkaClient.Close()

	if err := kafkaClient.CreateTopic(req.Topic, req.NumPartitions, req.ReplicationFactor); err != nil {
		return nil, err
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Creating topic %s", req.Topic))), nil
}
