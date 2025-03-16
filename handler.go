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

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Topic %s is created", req.Topic))), nil
}

// DeleteTopic deletes an existing Kafka topic
func (k *KafkaHandler) DeleteTopic(ctx context.Context, req Request) (*mcp_golang.ToolResponse, error) {

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	kafkaClient, err := kafka.NewClient(kafka.WithBootstrapServers([]string{"localhost:9092"}))
	if err != nil {
		return nil, err
	}
	defer kafkaClient.Close()

	if err := kafkaClient.DeleteTopic(req.Topic); err != nil {
		return nil, err
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Topic %s is deleted", req.Topic))), nil
}

// ListTopics lists all existing Kafka topics
func (k *KafkaHandler) ListTopics(ctx context.Context, req Request) (*mcp_golang.ToolResponse, error) {

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	kafkaClient, err := kafka.NewClient(kafka.WithBootstrapServers([]string{"localhost:9092"}))
	if err != nil {
		return nil, err
	}
	defer kafkaClient.Close()

	topics, err := kafkaClient.ListTopics()
	if err != nil {
		return nil, err
	}

	i := 1
	response := "Available Topics\n"
	for k, v := range topics {
		response += fmt.Sprintf("%d - Name: %s, Replication Factor: %d, Partitions: %d\n", i, k, v.ReplicationFactor, v.NumPartitions)

	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(response)), nil
}

// DescribeTopic describes the specified topic
func (k *KafkaHandler) DescribeTopic(ctx context.Context, req Request) (*mcp_golang.ToolResponse, error) {

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	kafkaClient, err := kafka.NewClient(kafka.WithBootstrapServers([]string{"localhost:9092"}))
	if err != nil {
		return nil, err
	}
	defer kafkaClient.Close()

	details, err := kafkaClient.DescribeTopic(req.Topic)
	if err != nil {
		return nil, err
	}

	response := fmt.Sprintf("Hey %s, Format the following details in a nice table format or in a json format\n", req.Submitter)
	for _, d := range details {
		response += fmt.Sprintf("Topic Name: %s\n", d.Name)
		if d.Err != 0 {
			response += fmt.Sprintf("Topic Error: %s\n", d.Err)
		}

		switch d.IsInternal {
		case true:
			response += fmt.Sprintf("Topic is an internal\n")
		case false:
			response += fmt.Sprintf("Topic is not internal\n")
		}

		response += fmt.Sprintf("## Topic Partition Details")
		for i, partition := range d.Partitions {
			response += fmt.Sprintf("%d) Partition ID: %d\n", i+1, partition.ID)
			response += fmt.Sprintf("Partition Leader Broker ID: %d\n", partition.Leader)
			response += fmt.Sprintf("Partition Leader Epoch: %d\n", partition.LeaderEpoch)
			response += fmt.Sprintf("Partition Replicas: %v\n", partition.Replicas)
			response += fmt.Sprintf("Partition Errors: %v\n", partition.Err)
			response += fmt.Sprintf("Insync replicas: %v\n", partition.Isr)
			response += fmt.Sprintf("Offline Replicas: %v\n", partition.OfflineReplicas)
		}
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(response)), nil
}
