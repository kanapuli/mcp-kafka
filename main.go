package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	kafka "github.com/kanapuli/mcp-kafka/kafka"
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
	"go.uber.org/zap"
)

// version is set during build time using ldflags
var version string

// Request represents the MCP related request information
type Request struct {
	Submitter             string         `json:"submitter" jsonschema:"required,description=The name of the thing calling this tool"`
	Topic                 string         `json:"topic" jsonschema:"required,description=The topic which the kafka client should work with"`
	NumPartitions         int32          `json:"num_partitions" jsonschema:"optional,description=The number of partitions for the topic"`
	ReplicationFactor     int16          `json:"replication_factor" jsonschema:"optional,description=The replication factor for the topic"`
	ProduceMessageKey     string         `json:"produce_message_key" jsonschema:"optional,description=The key to use when producing messages"`
	ProduceMessageValue   string         `json:"produce_message_value" jsonschema:"optional,description=The message content to use when producing messages"`
	ProduceMessageHeaders map[string]any `json:"produce_message_headers" jsonschema:"optional,description=The message headers to use when producing messages"`
	ConsumerTimeout       int            `json:"consumer_timeout" jsonschema:"optional,description=The timeout in seconds for consuming messages.Defaults to 10 seconds" default:"10"`
}

// EnvArgs represents the environment variables that can be used to configure the Kafka client
type EnvArgs struct {
	BootstrapServers []string
	ConsumerGroupID  string
	Username         string
	Password         string
}

func main() {

	logger, err := initLogger()
	if err != nil {
		fmt.Printf("error creating logger: %v", err)
		os.Exit(1)
	}

	logger.Infof("Starting mcp-kafkaversion %s", version)

	done := make(chan struct{})

	bootstrapServersFlag := flag.String("bootstrap-servers", "localhost:9092", "Comma-separated list of Kafka bootstrap servers")
	consumerGroupID := flag.String("consumer-group-id", "mcp-kafka-consumer-group", "Consumer group ID")
	username := flag.String("username", "", "Kafka username")
	password := flag.String("password", "", "Kafka password")
	flag.Parse()

	bootstrapServers := strings.Split(*bootstrapServersFlag, ",")

	kafkaClient, err := kafka.NewClient(
		kafka.WithBootstrapServers(bootstrapServers),
		kafka.WithConsumerGroupID(*consumerGroupID),
		kafka.WithUsername(*username),
		kafka.WithPassword(*password),
		kafka.WithLogger(logger),
	)
	if err != nil {
		zap.S().Errorf("error creating kafka client: %v", err)
		os.Exit(1)
	}
	defer kafkaClient.Close()

	kafkaHandler := &KafkaHandler{
		Client: kafkaClient,
		Logger: logger,
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

	err = server.RegisterTool("produce_message", "Produce a message to a topic", kafkaHandler.Produce)
	if err != nil {
		zap.S().Errorf("error registering kafka topic resource: %v", err)
		os.Exit(1)
	}

	err = server.RegisterTool("consume_messages", "Consume messages from a topic", kafkaHandler.Consume)
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

func initLogger() (*zap.SugaredLogger, error) {
	logConfig := zap.NewProductionConfig()
	// MCP currently requires logs to be written to stderr in order to be picked up in the logs
	logConfig.OutputPaths = []string{"stderr"}
	zapLogger, err := logConfig.Build()
	if err != nil {
		return nil, err
	}

	logger := zapLogger.Sugar()

	return logger, nil
}
