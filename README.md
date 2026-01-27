# Cyborg Conductor Core

[![Cyborg Conductor Core banner](static/images/cyborg-conductor-core-banner-wide.png)]()

A distributed system for efficiently managing *Cyborgs* and the physical resources they run on.

## Definitions

### Cyborg

A *Cyborg* is a logically isolated entity identified by a UUID. It bundles immutable capability specifications, either *Deterministic executable handlers* or *Streamable LLM inference streams*.

Cyborgs are described in a `CyborgDescriptor` protobuf with runtime state (current phase, active streams, and latency budgets) tracked separately in the `CyborgFusionEnvelope` protobuf to guarantee predictable execution and strong isolation across deployments.

## Project Structure

```
.
├── cmd/                    # Command-line tools
│   └── server/             # Main server application
├── pkg/                    # Core packages
│   ├── core/               # Core data structures and types
│   │   ├── pb/             # Protobuf-generated types and registry
│   │   └── types/          # Core types
│   ├── config/             # Configuration management
│   ├── context/            # Context utilities
│   └── memory/             # Memory management and evidence logging
├── internal/               # Internal packages
│   ├── context/            # Context management components
│   ├── runner/             # Cyborg runner implementations
│   └── conductor/          # Job scheduling and orchestration
├── cyborgs/                # Cyborg management
│   └── [cyborg-id]/        # Individual cyborg directories
├── proto/                  # Protocol buffer definitions
├── adapters/               # Language adapters (Python, Node.js)
├── test/                   # Test suites
│   ├── unit/               # Unit tests
│   └── integration/        # Integration tests
├── docs/                   # Documentation
├── .github/workflows/      # CI/CD pipelines
└── go.mod                  # Go module file
```

## Features

- **Distributed Cyborg Management**: Manage multiple Cyborgs across different environments.
- **Adaptive Scheduling**: Dynamic task assignment based on Cyborg capabilities.
- **Back-pressure Handling**: Intelligent flow control to prevent system overload.
- **LLM Integration**: Support for large language model streaming sessions.
- **Observability**: Comprehensive monitoring and admin interfaces.
- **Production-Ready**: Full CI/CD pipeline, security checks, and test coverage.

## Deployment

### Docker Deployment

```bash
# Build Docker image
docker build -t cyborg-conductor-core .

# Run container
docker run -p 8080:8080 cyborg-conductor-core
```

### Kubernetes Deployment

```bash
# Create ConfigMap with schema
kubectl create configmap cyborg-schema --from-file=proto/

# Deploy with Helm chart
helm install cyborg-conductor ./charts/cyborg-conductor-core
```

## Documentation

For detailed information about the Cyborg Conductor Core system, please refer to the following documentation:

- [User Guide] - Complete user documentation including getting started, Cyborg management, and use cases
- [Technical Design Document] - In-depth architectural and implementation details
- [Contributing Guidelines] - How to contribute to the project

---

## Support My Projects

If you find this repository helpful and would like to support its development, consider making a donation:

### GitHub Sponsors

[![Sponsor](https://img.shields.io/badge/Sponsor-%23EA4AAA?style=for-the-badge&logo=github)](https://github.com/sponsors/toxicoder)

### Buy Me a Coffee

<a href="https://www.buymeacoffee.com/toxicoder" target="_blank">
    <img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" height="41" width="174">
</a>

### PayPal

[![PayPal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate/?hosted_button_id=LSHNL8YLSU3W6)

### Ko-fi

<a href="https://ko-fi.com/toxicoder" target="_blank">
    <img src="https://storage.ko-fi.com/cdn/kofi3.png" alt="Ko-fi" height="41" width="174">
</a>

### Coinbase

[![Donate via Coinbase](https://img.shields.io/badge/Donate%20via-Coinbase-0052FF?style=for-the-badge&logo=coinbase&logoColor=white)](https://commerce.coinbase.com/checkout/e07dc140-d9f7-4818-b999-fdb4f894bab7)

Your support helps maintain and improve this collection of development tools and templates. Thank you for contributing to open source!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
