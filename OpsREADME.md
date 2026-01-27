# Cyborg Conductor Core - Operations Guide

## Overview

This document provides operational guidance for deploying and managing the Cyborg Conductor Core system in production environments.

## Deployment

### Docker Deployment

```bash
# Build the Docker image
docker build -t cyborg-conductor-core .

# Run the container
docker run -d \
  --name cyborg-conductor \
  -p 8080:8080 \
  -v /var/lib/cyborg/evidence:/var/lib/cyborg/evidence \
  cyborg-conductor-core
```

### Kubernetes Deployment

```bash
# Create ConfigMap with schema files
kubectl create configmap cyborg-schema \
  --from-file=proto/ \
  --from-file=cyborgs/

# Deploy the application
kubectl apply -f deployment.yaml
```

## Configuration

### Environment Variables

The system supports configuration through environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVER_HOST` | Server host address | `localhost` |
| `SERVER_PORT` | Server port | `8080` |
| `MAX_CONTEXT_BYTES` | Maximum context size in bytes | `104857600` (100MB) |
| `LOG_LEVEL` | Logging level | `info` |
| `LOG_FORMAT` | Log format (json/text) | `json` |

### Configuration File

The system can also be configured via environment variables. A `config.json` file is no longer required as configuration is now handled through environment variables.

## Health Checks

### Health Endpoint

```bash
curl http://localhost:8080/healthz
```

Returns `200 OK` when the service is healthy.

### Status Endpoint

```bash
curl http://localhost:8080/api/v1/status
```

Returns JSON with system status information.

## Monitoring & Metrics

### Prometheus Metrics

The system exposes Prometheus metrics at `/metrics` endpoint. Key metrics include:

- `cbg_cyborg_registered_total` - Total number of registered cyborgs
- `cbg_scheduler_queue_len` - Current queue length
- `cbg_llm_stream_active_sessions` - Active LLM streaming sessions

### Logging

Structured logging is implemented using logrus. Logs are output in JSON format by default.

## Alerting Thresholds

### Critical Alerts

- **Service Unavailable**: `/healthz` returns non-200 status for 3 consecutive checks
- **Memory Usage**: Context memory usage exceeds 90% of configured limit
- **Cyborg Registration Failure**: Failed to register any cyborgs during startup

### Warning Alerts

- **High Queue Load**: Scheduler queue length exceeds 1000 pending jobs
- **Low Resource**: Available memory drops below 20% of total system memory

## Rollback Procedures

### Docker Rollback

1. Stop current container:

   ```bash
   docker stop cyborg-conductor
   ```

2. Start previous version:

   ```bash
   docker run -d --name cyborg-conductor -p 8080:8080 previous-version-image
   ```

### Kubernetes Rollback

1. Scale down current deployment:

   ```bash
   kubectl scale deployment cyborg-conductor --replicas=0
   ```

2. Rollback to previous version:

   ```bash
   kubectl rollout undo deployment/cyborg-conductor
   ```

## Scaling Considerations

### Horizontal Scaling

The system supports horizontal scaling by running multiple instances behind a load balancer. Each instance maintains independent cyborg registration and scheduling.

### Resource Requirements

- **CPU**: Minimum 0.5 cores per instance
- **Memory**: Minimum 2GB RAM per instance
- **Storage**: 10GB for evidence logging and 1GB for application files

### Auto-scaling

Configure auto-scaling based on:

- CPU utilization (target 70%)
- Memory usage (target 80%)
- Scheduler queue length (threshold 500 jobs)

## Backup Strategy

### Evidence Data

Evidence data is stored in `/var/lib/cyborg/evidence` and should be backed up regularly using standard file system backup tools.

### Configuration

Configuration is managed through environment variables and Kubernetes ConfigMaps. Maintain version control of these configurations.

## Troubleshooting

### Common Issues

1. **Failed Cyborg Registration**
   - Check that all `.txtpb` files in `cyborgs/` directory are valid
   - Verify file permissions for cyborg directory
   - Check system logs for specific error messages

2. **Health Check Failures**
   - Verify the service is listening on correct port
   - Check resource constraints
   - Review system logs for error messages

3. **High Memory Usage**
   - Monitor context size limits
   - Check for memory leaks in components
   - Review evidence logging policies

### Debugging Commands

```bash
# Check container logs
docker logs cyborg-conductor

# Monitor resource usage
docker stats cyborg-conductor

# Verify cyborg registration
curl http://localhost:8080/api/v1/status
```

## Security Considerations

### Access Control

- Restrict access to administrative endpoints
- Implement proper authentication for sensitive operations
- Use network policies to limit inter-service communication

### Data Protection

- Encrypt sensitive data at rest
- Implement secure communication using TLS
- Regular security scanning of container images

## Maintenance Windows

### Scheduled Maintenance

Perform maintenance during low-traffic periods to minimize impact on cyborg operations.

### System Updates

1. Plan updates during maintenance windows
2. Test updates in staging environment first
3. Roll out updates to production in batches
4. Monitor system health post-update

## Performance Optimization

### Memory Management

- Configure appropriate `MAX_CONTEXT_BYTES` based on workload
- Monitor LRU cache behavior in memory manager
- Optimize evidence logging policies

### Scheduler Optimization

- Monitor queue length and adjust worker pool size
- Optimize cyborg selection algorithms
- Review resource quota configurations

## Compliance

### Audit Requirements

- Maintain logs of all cyborg registration and scheduling operations
- Implement retention policies for evidence data
- Ensure compliance with data protection regulations

## Contact Information

For support with this system, contact the Cyborg Conductor team.
