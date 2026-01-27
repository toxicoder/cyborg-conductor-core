# User Guide

Welcome to the Cyborg Conductor Core User Guide. This guide provides comprehensive information on how to use, deploy, and manage the Cyborg Conductor Core system.

## Table of Contents

1. [Overview](overview.md)
2. [Getting Started](getting-started.md)
3. [Agent Management](agent-management.md)
4. [Use Cases](use-cases.md)
5. [Troubleshooting](troubleshooting.md)

## System Architecture

The Cyborg Conductor Core is designed to orchestrate and manage 108 cyborgs described by the Master Matrix. It provides:

- Autonomous orchestration of cyborgs with deterministic handlers and streaming LLM sessions
- Full traceability via immutable Merkle-log of evidence
- Production-grade observability (Prometheus metrics, health checks)
- Rollout-ready packaging (Docker/Helm)
- Support for the complete organizational agent framework with 97+ roles across 10+ squads

## Key Components

- **Cyborg Registry**: Manages all registered cyborgs with their capabilities and configurations
- **Job Scheduler**: Intelligent scheduler that matches jobs to appropriate cyborgs based on capabilities
- **Context Manager**: Handles memory management, evidence logging, and context tolerance policies
- **Execution Manager**: Runs deterministic handlers and manages subprocess execution
- **LLM Session Manager**: Manages streaming LLM sessions for complex tasks
- **Adapters**: Python and Node.js wrappers for external integrations
- **Observability Layer**: Prometheus metrics, health checks, and logging

## Getting Started

To get started with the Cyborg Conductor Core, please follow our [Getting Started Guide](getting-started.md).

## Agent Framework

The system supports a comprehensive organizational agent framework with roles across various squads. Learn more about [Agent Management](agent-management.md) and see examples in our [Use Cases](use-cases.md).

## Contributing

We welcome contributions from the community. Please see our [Contributing Guidelines](../../CONTRIBUTING.md) for more information on how to contribute to this project.