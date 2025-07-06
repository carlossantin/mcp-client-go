# MCP Client Go

A Go-based client for interacting with Model Context Protocol (MCP) servers, providing AI agents with tool capabilities through streaming and non-streaming interfaces.

## Overview

This project demonstrates how to build a Go client that connects to MCP servers and creates AI agents with access to external tools. The client supports both Azure OpenAI and other LLM providers, with configurable MCP server connections and tool access control.

## Features

- **Multi-Provider LLM Support**: Currently supports Azure OpenAI with extensible provider architecture
- **MCP Server Integration**: Connect to MCP servers via SSE (Server-Sent Events) or local processes
- **Tool Access Control**: Configure which tools each agent can access from connected MCP servers
- **Streaming Support**: Real-time streaming responses with tool execution feedback
- **Interactive CLI**: Command-line interface for conversational interactions
- **Configuration-Driven**: YAML-based configuration for easy setup and deployment

## Prerequisites

- Go 1.24.1 or higher
- Access to an Azure OpenAI instance (or other supported LLM provider)
- Running MCP server(s) to connect to

## Installation

1. Clone the repository:
```bash
git clone https://github.com/carlossantin/mcp-client-go.git
cd mcp-client-go
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up environment variables:
```bash
export AZURE_OPENAI_API_KEY="your-azure-openai-api-key"
export AZURE_URL="your-azure-openai-endpoint"
export MY_MCP_SERVER_URL="http://localhost:8080/mcp/events"
```

## Configuration

The application uses a `config.yaml` file to define providers, servers, and agents:

```yaml
providers:
  - name: my-azure-provider
    type: AZURE
    token: ${AZURE_OPENAI_API_KEY}
    baseUrl: ${AZURE_URL}
    model: gpt-4o-mini
    version: 2025-01-01-preview

servers:
  - name: my-mcp-server
    type: sse
    url: ${MY_MCP_SERVER_URL|http://localhost:8080/mcp/events}

agents:
  - name: my-agent
    servers:
      - Name: my-mcp-server
        AllowedTools:
          - component_documentation
          - component_egress_dependencies
    provider: my-azure-provider
```

### Configuration Options

#### Providers
- `name`: Unique identifier for the provider
- `type`: Provider type (currently supports "AZURE")
- `token`: API key for the provider
- `baseUrl`: Base URL for the provider's API
- `model`: Model name to use
- `version`: API version

#### Servers
- `name`: Unique identifier for the MCP server
- `type`: Connection type ("sse" for Server-Sent Events, "local" for local processes)
- `url`: Server URL (for SSE connections)
- `command`: Command to execute (for local processes)

#### Agents
- `name`: Unique identifier for the agent
- `servers`: List of MCP servers this agent can access
  - `Name`: MCP server name
  - `AllowedTools`: List of tools the agent can use (empty list allows all tools)
- `provider`: LLM provider to use for this agent

## Usage

### Running the Application

```bash
go run main.go
```

The application will start an interactive CLI where you can:
- Enter questions or prompts
- See real-time streaming responses
- View tool usage and responses
- Type "quit" or "exit" to terminate

### Example Interaction

```
Enter your question: What components are available in the system?

Generating content...
[tool_usage] component_documentation

[tool_response] my-mcp-server__component_documentation: {"components": [...]}

Based on the available components, here are the systems currently documented:
...
```

## Code Structure

- `main.go`: Main application entry point with CLI interface
- `config.yaml`: Configuration file for providers, servers, and agents
- External dependencies:
  - `github.com/carlossantin/mcp-agents-go`: Core MCP agent functionality
  - `github.com/tmc/langchaingo`: LLM integration library

## Key Components

### Agent Configuration
Agents are configured with:
- Access to specific MCP servers
- Tool-level permissions
- LLM provider settings

### Streaming vs Non-Streaming
The client supports both modes:
- **Streaming**: Real-time response generation with tool execution visibility
- **Non-Streaming**: Complete response after processing

### Tool Execution
Tools are executed with:
- Automatic JSON argument parsing
- Error handling and response formatting
- Tool response integration back to the LLM

## Development

### Adding New Providers
To add support for new LLM providers, extend the provider configuration in the `mcp-agents-go` library.

### Adding New MCP Server Types
New connection types can be added by implementing the appropriate interfaces in the server package.

### Customizing Tool Access
Fine-tune tool access by modifying the `AllowedTools` list in the agent configuration.

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `AZURE_OPENAI_API_KEY` | Azure OpenAI API key | Required |
| `AZURE_URL` | Azure OpenAI endpoint URL | Required |
| `MY_MCP_SERVER_URL` | MCP server URL | `http://localhost:8080/mcp/events` |

## Dependencies

- **github.com/carlossantin/mcp-agents-go**: Core MCP agent functionality
- **github.com/tmc/langchaingo**: LLM integration and tool calling
- **github.com/mark3labs/mcp-go**: MCP protocol implementation

## License

This project is licensed under the terms found in the LICENSE file.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## Support

For issues and questions:
- Check the existing issues on GitHub
- Create a new issue with detailed information about your problem
- Include relevant configuration and error messages

## Roadmap

- [ ] Support for additional LLM providers
- [ ] Enhanced error handling and recovery
- [ ] Metrics and logging improvements
- [ ] Web UI interface
- [ ] Docker deployment support
