package main

import (
	"os"

	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
	"go.uber.org/zap"
)

// Request represents the MCP related request information
type Request struct {
	Submitter         string `json:"submitter" jsonschema:"required,description=The name of the thing calling this tool"`
	Topic             string `json:"topic" jsonschema:"required,description=The topic which the kafka client should work with"`
	NumPartitions     int32  `json:"num_partitions" jsonschema:"optional,description=The number of partitions for the topic"`
	ReplicationFactor int16  `json:"replication_factor" jsonschema:"optional,description=The replication factor for the topic"`
}

func main() {
	done := make(chan struct{})

	kafkaHandler := &KafkaHandler{}
	server := mcp_golang.NewServer(stdio.NewStdioServerTransport(), mcp_golang.WithName("kafka"))

	err := server.RegisterTool("create_topic", "Create a topic with the given number of partitions and replication factor", kafkaHandler.CreateTopic)
	if err != nil {
		zap.S().Errorf("error registering kafka topic resource: %v", err)
		os.Exit(1)
	}

	err = server.RegisterTool("list_topics", "List all available topics in a table format", kafkaHandler.ListTopics)
	if err != nil {
		zap.S().Errorf("error registering kafka topic resource: %v", err)
		os.Exit(1)
	}

	err = server.RegisterTool("delete_topic", "Delete a topic", kafkaHandler.DeleteTopic)
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
