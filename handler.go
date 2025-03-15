package main

import (
	"context"
	"fmt"

	kafka "github.com/kanapuli/mcp-kafka/kafka"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

// KafkaHandler is a struct that handles Kafka operations for the mcp-kafka tool
type KafkaHandler struct{}

func (k *KafkaHandler) Register(ctx context.Context, args FuncArgs) (*mcp_golang.ToolResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	kafkaClient, err := kafka.NewClient(kafka.WithBootstrapServers([]string{"localhost:9092"}))
	if err != nil {
		return nil, err
	}
	defer kafkaClient.Close()

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Vanakam, %s. Content: %s. Topic: %s", args.Submitter, args.Content, args.Topic))), nil
}
