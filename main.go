package main

import (
	"context"
	"os"
	"time"

	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
	"go.uber.org/zap"
)

type FuncArgs struct {
	Submitter         string `json:"submitter" jsonschema:"required,description=The name of the thing calling this tool."`
	Content           string `json:"content" jsonschema:"required,description=The content of the message"`
	Topic             string `json:"topic" jsonschema:"required,description=The topic which the kafka client should work with"`
	NumPartitions     int32  `json:"num_partitions" jsonschema:"required,description=The number of partitions for the topic"`
	ReplicationFactor int16  `json:"replication_factor" jsonschema:"required,description=The replication factor for the topic"`
}

func main() {
	done := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())
	kafkaHandler := &KafkaHandler{}
	// Todo: Add more description
	err := server.RegisterTool("kafka", "Performs client operations on Kafka brokers", kafkaHandler.Register)
	if err != nil {
		zap.S().Errorf("error registering kafka mcp server: %v", err)
		os.Exit(1)
	}

	go func(ctx context.Context) {
		err := server.RegisterTool("create_topic", "Create a topic with the given number of partitions and replication factor", kafkaHandler.CreateTopic)
		if err != nil {
			zap.S().Errorf("error registering kafka topic resource: %v", err)
		}

		time.Sleep(1 * time.Second)

		err = server.DeregisterResource("kafka://topic")
		if err != nil {
			zap.S().Errorf("error deregistering kafka topic resource: %v", err)
		}
	}(ctx)

	if err = server.Serve(); err != nil {
		zap.S().Errorf("error serving kafka mcp server: %v", err)
		os.Exit(1)
	}

	<-done
}
