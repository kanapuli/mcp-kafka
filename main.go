package main

import (
	"os"

	kafka "github.com/kanapuli/mcp-kafka/kafka"
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
	"go.uber.org/zap"
)

// version is set during build time using ldflags
var version string

// Request represents the MCP related request information
type Request struct {
	Submitter         string `json:"submitter" jsonschema:"required,description=The name of the thing calling this tool"`
	Topic             string `json:"topic" jsonschema:"required,description=The topic which the kafka client should work with"`
	NumPartitions     int32  `json:"num_partitions" jsonschema:"optional,description=The number of partitions for the topic"`
	ReplicationFactor int16  `json:"replication_factor" jsonschema:"optional,description=The replication factor for the topic"`
}

func main() {
	zap.S().Infof("Starting mcp-kafkaversion %s", version)

	done := make(chan struct{})

	kafkaClient, err := kafka.NewClient(kafka.WithBootstrapServers([]string{"localhost:9092"}))
	if err != nil {
		zap.S().Errorf("error creating kafka client: %v", err)
		os.Exit(1)
	}
	defer kafkaClient.Close()

	kafkaHandler := &KafkaHandler{
		Client: kafkaClient,
	}
	server := mcp_golang.NewServer(stdio.NewStdioServerTransport(), mcp_golang.WithName("kafka"))

	err = server.RegisterTool("create_topic", "Create a topic with the given number of partitions and replication factor", kafkaHandler.CreateTopic)
	if err != nil {
		zap.S().Errorf("error registering kafka topic resource: %v", err)
		os.Exit(1)
	}

	err = server.RegisterTool("list_topics", "List all available topics", kafkaHandler.ListTopics)
	if err != nil {
		zap.S().Errorf("error registering kafka topic resource: %v", err)
		os.Exit(1)
	}

	err = server.RegisterTool("delete_topic", "Delete a topic", kafkaHandler.DeleteTopic)
	if err != nil {
		zap.S().Errorf("error registering kafka topic resource: %v", err)
		os.Exit(1)
	}

	err = server.RegisterTool("describe_topic", "Describe a kafka topic", kafkaHandler.DescribeTopic)
	if err != nil {
		zap.S().Errorf("error registering kafka topic resource: %v", err)
		os.Exit(1)
	}

	if err := server.Serve(); err != nil {
		zap.S().Errorf("error serving kafka mcp server: %v", err)
		os.Exit(1)
	}
	<-done
}
