package main

import (
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

type FuncArgs struct {
	Submitter string `json:"submitter" jsonschema:"required,description=The name of the thing calling this tool."`
	Content   string `json:"content" jsonschema:"required,description=The content of the message"`
	Topic     string `json:"topic" jsonschema:"required,description=The topic which the kafka client will publish the message to"`
}

func main() {
	done := make(chan struct{})
	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())
	kafkaHandler := &KafkaHandler{}
	// Todo: Add more description
	err := server.RegisterTool("kafka", "Performs client operations on Kafka brokers", kafkaHandler.Register)

	if err != nil {
		panic(err)
	}

	err = server.Serve()
	if err != nil {
		panic(err)
	}

	<-done
}
