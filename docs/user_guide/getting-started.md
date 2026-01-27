# Getting Started

This guide will help you set up and start using the Cyborg Conductor Core system. Whether you're deploying it for the first time or integrating with existing systems, this guide provides the necessary steps and best practices.

## Prerequisites

Before you begin, ensure you have the following:

- Go 1.22 or higher installed
- Docker and Docker Compose (for containerized deployment)
- Protobuf tools (protoc)
- Git for version control

## Installation

### From Source

1. Clone the repository:
   ```bash
   git clone https://github.com/toxicoder/cyborg-conductor-core.git
   cd cyborg-conductor-core
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the project:
   ```bash
   go build -o cyborg-conductor-core cmd/server/main.go
   ```

### Using Docker

The system can be easily deployed using Docker:

1. Build the Docker image:
   ```bash
   docker build -t cyborg-conductor-core .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 cyborg-conductor-core
   ```

### Kubernetes Deployment

For production deployments, the system supports Kubernetes:

1. Create the ConfigMap with schema:
   ```bash
   kubectl create configmap cyborg-schema --from-file=proto/
   ```

2. Deploy with Helm chart:
   ```bash
   helm install cyborg-conductor ./charts/cyborg-conductor-core
   ```

## Configuration

The system uses a configuration management approach that supports multiple environments. Configuration files are typically stored in the `config/` directory and can be overridden by environment variables.

Key configuration options include:
- `PORT`: Server port (default: 8080)
- `LOG_LEVEL`: Logging verbosity (debug, info, warn, error)
- `MAX_CONTEXT_BYTES`: Maximum memory allowed for context
- `EVIDENCE_ROOT`: Directory for evidence logging

## Running the Server

### Local Development

To run the server locally for development:

```bash
go run cmd/server/main.go
```

The server will start on port 8080 by default and expose:
- `/healthz` for health checks
- `/api/v1/status` for system status
- Prometheus metrics at `/metrics`

### Production Deployment

For production deployments, ensure:
- Proper security configurations (firewall rules, TLS)
- Sufficient resources for 108 concurrent cyborgs
- Monitoring and alerting integration
- Backup and recovery procedures

## Basic Usage

### Registering Cyborgs

Cyborgs can be registered through the gRPC API. The registration process involves:
1. Preparing a `CyborgDescriptor` protobuf message
2. Sending it to the `/RegisterCyborg` endpoint
3. Validating uniqueness and configuration

### Scheduling Jobs

Jobs are scheduled using the system's intelligent scheduler:
1. Create a job request with required capabilities
2. Submit to the job queue
3. The system matches the job to an appropriate cyborg
4. Execute the job and return results

## Monitoring and Observability

The system provides comprehensive observability features:
- Prometheus metrics endpoint at `/metrics`
- Health check endpoint at `/healthz`
- Detailed logging with structured output
- System status endpoint at `/api/v1/status`

## Next Steps

Once you've completed the basic setup, consider:
1. Exploring the agent framework and available cyborg types
2. Configuring custom cyborgs for your specific use cases
3. Setting up integration with your existing tools and workflows
4. Implementing advanced features like LLM streaming sessions

## Troubleshooting

If you encounter issues during setup or operation:
1. Check the logs for error messages
2. Verify that all prerequisites are properly installed
3. Ensure configuration files are correctly formatted
4. Consult the [Troubleshooting Guide](troubleshooting.md) for common issues