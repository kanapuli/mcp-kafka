# mcp-kafka

A Model Context Protocol (MCP) server for performing Kafka client operations from AI assistants.

## Overview

mcp-kafka provides a bridge between AI assistants and Apache Kafka, allowing them to interact with Kafka clusters through the Model Context Protocol. This tool enables AI assistants to create, manage, and interact with Kafka topics and messages directly.

## Features

The mcp-kafka server provides the following Kafka operations:

- **Create Topic**: Create a new Kafka topic with configurable partitions and replication factor
- **List Topics**: Get a list of all available Kafka topics in the cluster
- **Delete Topic**: Remove an existing Kafka topic
- **Describe Topic**: Get detailed information about a specific topic, including partition details
- **Produce Message**: Send messages to a Kafka topic with support for message keys and headers
- **Consume Messages**: Read messages from a Kafka topic with configurable timeout

## Installation

### Prerequisites

- Go 1.24 or higher
- A running Kafka cluster (default connection: localhost:9092)

### Download Prebuilt Binaries

Prebuilt binaries for multiple platforms are available on the [releases page](https://github.com/kanapuli/mcp-kafka/releases).

Supported platforms:
- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

### Building from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/kanapuli/mcp-kafka.git
   cd mcp-kafka
   ```

2. Build the application:
   ```bash
   make build
   ```

3. Optionally, build for a specific platform:
   ```bash
   make build GOOS=darwin GOARCH=arm64
   ```

### Installing as a Claude Desktop Tool

To use mcp-kafka with Claude Desktop:

1. Download the latest release from the [releases page](https://github.com/kanapuli/mcp-kafka/releases) or build from source.

2. Place the executable in a location included in your system PATH or in a dedicated tools directory.

3. Configure Claude Desktop to recognize the tool:
   - Open Claude Desktop settings
   - Navigate to the Tools section
   - Add a new tool pointing to the mcp-kafka executable
   - Save your changes

## Sample Claude Desktop config json

Please look at the [Model Context Protocol for Claude Desktop](https://modelcontextprotocol.io/quickstart/user) documentation to get a general idea of how to setup the MCP tool.

```json
{
  "mcpServers": {
    "kafka": {
	    "command": "/Path-to-your-mcp-kafka-binary/mcp-kafka-darwin-arm64",
      "args": [
          "--bootstrap-servers=localhost:9092",
          "--consumer-group-id=mcp-kafka-consumer-group",
          "--username=",
          "--password="
        ],
      "env": {}
    }
  }
}
```


## Configuration

The mcp-kafka tool accepts the following configuration parameters:

| Parameter | Description | Default |
|-----------|-------------|---------|
| topic | Topic to interact with | (required) |
| num_partitions | Number of partitions for topic creation | (optional) |
| replication_factor | Replication factor for topic creation | (optional) |
| produce_message_key | Key for produced messages | (optional) |
| produce_message_value | Value for produced messages | (optional) |
| produce_message_headers | Headers for produced messages | (optional) |
| consumer_timeout | Timeout in seconds for message consumption | 10 |

These parameters will be automatically derived from your Natural language message to the LLM.


### CLI flags for the mcp-kafka tool

The following flags should be used to configure the Kafka client:

```
--bootstrap-servers=localhost:9092
--consumer-group-id=mcp-kafka-consumer-group
--username='your_sasl_username'
--password='your_sasl_password'

```

NOTE: Currently, SASL_PLAINTEXT is supported along with PLAINTEXT authentication. SASL_SSL is not supported.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)

## Acknowledgments

- Built using the [MCP Golang library](https://github.com/metoro-io/mcp-golang)
