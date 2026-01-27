# Inkorporated - Eazy Homelab & Corp IaC

[![Inkorporated banner](static/images/inkorporated-globe-banner.jpg)]()

## Overview

This repository contains the complete infrastructure-as-code for the Inkorporated homelab, a comprehensive self-hosted, open-source internal work environment designed for homelab deployments. It provides a complete, integrated solution that combines infrastructure automation, GitOps deployment, and a rich set of productivity and collaboration tools.

## Architecture

Inkorporated now supports a **Hybrid Cloud** architecture, enabling workloads to run on both on-premises Proxmox clusters and public cloud providers (AWS, GCP).

* **Proxmox**: Hosts the core control plane and persistent workloads.
* **Public Cloud (AWS/GCP)**: Used for burst capacity and temporary workloads.
* **Networking**: Hybrid nodes are connected via VPN (Tailscale/Wireguard - *implementation detail*).
* **Storage**:
  * On-Prem: Longhorn / NFS
  * Cloud: Cloud-native storage (e.g., gp2/standard)

## Project Structure

The Inkorporated project is organized into several key areas:

* **Infrastructure as Code**:
  * `infrastructure/terraform`: Modular Terraform for Hybrid Cloud provisioning (Proxmox, AWS, GCP).
  * `infrastructure/ansible`: Ansible roles for hybrid k3s setup (LXC, VMs, Cloud).
* **GitOps Deployment**: ArgoCD for declarative infrastructure and application management using ApplicationSets.
* **Zero-trust Access**: Cloudflare Tunnel for secure external access
* **Centralized Authentication**: Authentik SSO/OIDC integration
* **Observability**: Prometheus, Grafana, and Loki stack
* **Storage Solutions**: Longhorn and NFS CSI driver
* **Backup & Recovery**: Velero with MinIO
* **Service Mesh**: Traefik with forward-auth middleware
* **Version Control**: Gitea instance for GitOps workflow
* **CI/CD**: Gitea runners for automated deployments

## Deployment Patterns

All services follow a consistent deployment pattern:

1. Create Namespace
2. Create Core Deployment with appropriate resources
3. Create Service
4. Create PodDisruptionBudget (where applicable)
5. Create ConfigMap (where applicable)
6. Create Ingress (where applicable)
7. Create Secrets (where applicable)

## Configuration Management

The project implements a centralized configuration management system that separates sensitive credentials from main settings:

* **Main Settings File**: `cline_mcp_settings.json` - Contains server configurations and settings (no sensitive data)
* **Configuration File**: `.devcontainer/.env` - Contains sensitive credentials and environment-specific settings (not version controlled)
* **Environment-Specific Configurations**: Different files for different environments (`.devcontainer/.env.dev`, `.devcontainer/.env.staging`, `.devcontainer/.env.prod`)

## Environment Structure

The repository now supports multiple environments with dedicated configuration directories:

* `config/environments/` - Environment-specific configurations
* `apps/environments/` - Environment-specific application overrides

### Supported Environments

1. **dev** - Development environment
2. **staging** - Staging environment
3. **autopush** - Autopush environment
4. **uat** - User Acceptance Testing environment
5. **canary** - Canary environment for feature testing
6. **prod** - Production environment
7. **priv** - Private environment for sensitive services

## Security Approach

The configuration system follows security best practices:

1. **Separation of Concerns**: Sensitive credentials are separated from settings
2. **Version Control Safety**: No sensitive data is committed to version control
3. **File Permissions**: Config files are secured with restrictive permissions (600)
4. **Environment Variables**: Configuration loaded at runtime
5. **Zero-trust Architecture**: Cloudflare Tunnel for external access
6. **Centralized Authentication**: Authentik for SSO
7. **Network Segmentation**: pfSense firewall with VLANs

## Infrastructure Components

The project implements a multi-zone network architecture with the following components:

### Core Infrastructure Services

* **Cloudflare Tunnel (cloudflared)**: Zero-trust access to services

* **Traefik**: Ingress controller with TLS termination
* **Authentik**: Central identity provider with OIDC + 2FA enforcement
* **Longhorn**: Distributed block storage
* **NFS CSI Driver**: Synology NAS integration
* **MinIO**: S3 object storage for backups
* **Velero**: Backup and restore solution
* **cert-manager**: TLS automation
* **MetalLB**: LoadBalancer provider

### Collaboration Services

* **Rocket.Chat**: Team chat platform

* **WorkAdventure**: 2D virtual office
* **Jitsi Meet**: Video conferencing
* **Coturn**: TURN/STUN server for NAT traversal

### Productivity Services

* **AppFlowy**: Collaborative knowledge base

* **LinkWarden**: Bookmark manager
* **Homebox**: Inventory tracker
* **Home Assistant**: Smart home hub

### Remote Work Services

* **Kasm Workspaces**: Browser-based workspaces

* **Coder**: Cloud IDE workspaces

### AI Services

* **Ollama**: Local LLM runner

* **Open WebUI**: Ollama web interface
* **Langflow**: Visual LangChain builder
* **Kokoro TTS**: Local TTS server
* **Docling**: Document parsing server
* **SearXNG**: Metasearch engine
* **SurfSense**: AI research agent
* **Perplexica**: AI search engine

### Security Services

* **Vaultwarden**: Bitwarden-compatible password manager

* **HashiCorp Vault**: Secrets management

## Implementation Phases

The project follows a structured implementation approach:

### Phase 1: Foundation & Preparation

* Validate physical infrastructure readiness

* Create Proxmox Cloud-Init template
* Workstation setup with required tools

### Phase 2: Repository & Code Setup

* Create bootstrap repository with Terraform/Ansible

* Create apps repository with GitOps manifests

### Phase 3: Bootstrap Infrastructure

* Provision VMs with Terraform

* Install k3s with Ansible
* Bootstrap ArgoCD

### Phase 4: Core Infrastructure Deployment

* Deploy storage solutions (Longhorn, NFS CSI)

* Deploy database services (CloudNativePG, MongoDB)
* Deploy backup and monitoring stack
* Deploy core services (Authentik, Traefik, Cloudflared)

### Phase 5: Post-Deployment

* Configure authentication providers and groups

* Deploy user dashboard (Homepage)
* Enable backups and testing
* Performance and security hardening
* Documentation runbook creation

## Testing Framework

A comprehensive testing framework has been implemented to ensure reliability and security:

* Configuration validation tests
* Security permission checks
* Kubernetes manifest validation
* Integration testing for services
* Automated test suite with CI/CD integration
* Infrastructure health monitoring tests
* Backup and restore validation tests

## Development Workflow

The project follows these development workflow patterns:

* GitOps with ArgoCD for infrastructure and application deployment
* Environment variable configuration for domain flexibility
* Centralized configuration management system
* Kubernetes manifest pattern for custom resources
* Helm chart pattern for application deployment
* Service-oriented architecture with proper namespace separation
* MCP-based infrastructure automation workflows
* Automated security and compliance validation

## Tooling Preferences

The project uses the following tooling preferences:

* **Container Orchestration**: Kubernetes (k3s)
* **Deployment Management**: ArgoCD (GitOps)
* **Authentication**: Authentik (OIDC + SSO)
* **Ingress Controller**: Traefik with forward-auth
* **Zero-trust Access**: Cloudflare Tunnel (cloudflared)
* **Network Security**: pfSense VM
* **Storage**: Longhorn (block storage), NFS CSI driver (Synology NAS)
* **Backup & DR**: Velero with MinIO
* **Monitoring**: kube-prometheus-stack (Prometheus, Grafana, Loki)
* **Infrastructure as Code**: Terraform, Ansible
* **Development Tools**: kubectl, helm, terraform, ansible, cloudflared CLI
* **MCP Integration**: context7 MCP server for documentation and patterns
* **Security Tools**: Automated vulnerability scanning with MCP

## Performance Requirements

* Resource quotas and limits for all namespaces
* High availability with PodDisruptionBudgets
* Efficient resource utilization for all services
* Monitoring and alerting for performance metrics
* Backup and restore testing for disaster recovery
* Automated performance optimization with MCP tools

## Documentation Standards

All documentation follows these standards:

* Comprehensive technical documentation
* Service-specific documentation with configurations
* Implementation status tracking
* Configuration management guidelines
* Security and best practices documentation
* Troubleshooting and maintenance guides
* MCP integration documentation
* Infrastructure automation workflows documentation

## Accessibility Guidelines

While the project is primarily focused on infrastructure and backend services, accessibility considerations include:

* User-friendly dashboard interfaces
* Clear navigation and organization of services
* Consistent design patterns across applications
* Support for keyboard navigation where applicable

## Internationalization Needs

The project supports internationalization through:

* Configurable domain names and URLs
* Environment variable configuration for localization
* Multi-language support in user-facing applications
* Flexible configuration management for different regions

## Security Considerations

* All sensitive data stored as Kubernetes secrets
* Sealed secrets for Git-safe storage
* Proper RBAC and access controls
* Network policies for security segmentation
* Regular security updates for container images
* Centralized authentication with SSO
* Zero-trust network architecture
* Backup and disaster recovery procedures
* Automated security scanning with MCP tools
* Infrastructure hardening and compliance validation

## Integration Patterns

### Authentication Flow

1. User accesses any service via subdomain (e.g., `https://chat.example.com`)
2. Traefik forwards request to Authentik for authentication
3. Authentik validates credentials and issues JWT
4. User is redirected back to service with authenticated session
5. Service validates JWT and grants access

### Data Flow

1. Services communicate through internal Kubernetes networks
2. External services use Cloudflare Tunnel for secure access
3. All data is encrypted in-transit using TLS
4. Persistent data stored on Longhorn volumes or MinIO

## Maintenance and Operations

* Regular security audits and updates
* Automated container image updates
* Monitoring and alerting review
* Backup testing and verification
* Performance optimization
* Documentation updates
* Automated infrastructure health checks
* MCP-based operational workflows

## Future Enhancements

Planned features and improvements:

* Enhanced AI integration capabilities
* Improved monitoring and alerting
* Additional security hardening
* Performance optimization
* Automated testing and CI/CD improvements
* MCP-based infrastructure automation
* Advanced backup and disaster recovery
* Enhanced observability with AI insights
* More comprehensive security scanning
* Automated compliance validation

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
