# Troubleshooting Guide

This guide provides solutions for common issues and problems that users may encounter when working with the Cyborg Conductor Core system.

## System Startup Issues

### Server Fails to Start

**Symptoms**:
- Server crashes immediately on startup
- Error messages about port binding
- Configuration file parsing errors

**Solutions**:
1. Check if port 8080 is already in use:
   ```bash
   lsof -i :8080
   ```

2. Verify configuration files are properly formatted:
   ```bash
   cat config/app.yaml | yaml-lint
   ```

3. Ensure required environment variables are set:
   ```bash
   echo $PORT
   echo $LOG_LEVEL
   ```

4. Check system resource limits:
   ```bash
   ulimit -a
   ```

### Memory Allocation Errors

**Symptoms**:
- "out of memory" errors during startup
- Memory usage exceeds system limits
- System becomes unresponsive

**Solutions**:
1. Reduce `MAX_CONTEXT_BYTES` in configuration:
   ```bash
   export MAX_CONTEXT_BYTES=536870912  # 512MB
   ```

2. Monitor memory usage:
   ```bash
   top -p $(pgrep cyborg-conductor-core)
   ```

3. Check for memory leaks with profiling:
   ```bash
   go tool pprof -http=:8081 cpu.pprof
   ```

## Agent Registration Problems

### Registration Failures

**Symptoms**:
- Registration requests return "already registered" errors
- "Invalid descriptor" validation errors
- gRPC connection timeouts

**Solutions**:
1. Verify unique `cyborg_id` in descriptor:
   ```bash
   grep "cyborg_id" your_descriptor.pb.txt
   ```

2. Validate descriptor against schema:
   ```bash
   protoc --decode=cyborg.v1.CyborgDescriptor proto/cyborg.proto < your_descriptor.pb.txt
   ```

3. Check gRPC service availability:
   ```bash
   grpc_health_probe -addr=:8080
   ```

### Agent Not Appearing in Registry

**Symptoms**:
- Registered agents don't appear in `/api/v1/status`
- Agents are missing from scheduled jobs
- Context management fails for registered agents

**Solutions**:
1. Check registration logs for errors:
   ```bash
   tail -f logs/system.log
   ```

2. Verify agent descriptor contains all required fields:
   - `cyborg_id`
   - `display_name`
   - `category`
   - `capabilities_list`

3. Restart the server to refresh the registry:
   ```bash
   systemctl restart cyborg-conductor-core
   ```

## Performance Issues

### Slow Job Processing

**Symptoms**:
- Jobs take longer than expected to complete
- Scheduler appears to be blocking
- High CPU usage on the server

**Solutions**:
1. Check system resource utilization:
   ```bash
   htop
   ```

2. Review job scheduling logs:
   ```bash
   grep "scheduler" logs/system.log
   ```

3. Optimize agent capabilities:
   - Ensure agents have appropriate resource allocation
   - Review capabilities_list for relevant match requirements

4. Check for back-pressure issues:
   ```bash
   curl http://localhost:8080/metrics | grep backpressure
   ```

### High Memory Usage

**Symptoms**:
- Memory consumption consistently high
- Evidence logging causing storage issues
- System becomes sluggish

**Solutions**:
1. Monitor context usage:
   ```bash
   curl http://localhost:8080/api/v1/status
   ```

2. Adjust evidence logging retention:
   ```bash
   export EVIDENCE_RETENTION_DAYS=7
   ```

3. Implement memory cache size limits:
   ```bash
   export MAX_CONTEXT_BYTES=1073741824  # 1GB
   ```

## Network and Communication Issues

### gRPC Connection Problems

**Symptoms**:
- "connection refused" errors
- "deadline exceeded" messages
- Timeout during agent communication

**Solutions**:
1. Verify gRPC service is running:
   ```bash
   systemctl status cyborg-conductor-core
   ```

2. Test gRPC connectivity:
   ```bash
   grpcurl -plaintext localhost:8080 list
   ```

3. Check firewall rules:
   ```bash
   ufw status
   ```

4. Ensure correct TLS configuration if enabled:
   ```bash
   openssl s_client -connect localhost:8080 -servername localhost
   ```

## Monitoring and Logging Issues

### Missing Metrics

**Symptoms**:
- Prometheus endpoint returns 404
- Metrics not appearing in Grafana
- Health checks fail

**Solutions**:
1. Verify metrics endpoint is enabled:
   ```bash
   curl http://localhost:8080/metrics
   ```

2. Check logging configuration:
   ```bash
   cat config/logging.yaml
   ```

3. Restart monitoring service if needed:
   ```bash
   systemctl restart prometheus
   ```

### Log File Corruption

**Symptoms**:
- Log files are unreadable
- Error messages about file corruption
- Missing log entries

**Solutions**:
1. Rotate log files:
   ```bash
   logrotate -f /etc/logrotate.d/cyborg-conductor-core
   ```

2. Check disk space:
   ```bash
   df -h
   ```

3. Verify log permissions:
   ```bash
   ls -la logs/
   ```

## Security and Access Issues

### Authentication Failures

**Symptoms**:
- "unauthorized" errors for protected endpoints
- Access denied for registered agents
- Token validation failures

**Solutions**:
1. Verify authentication tokens:
   ```bash
   curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/api/v1/status
   ```

2. Check key management configuration:
   ```bash
   cat config/auth.yaml
   ```

3. Regenerate tokens if expired:
   ```bash
   ./generate-token.sh
   ```

### Permission Denied Errors

**Symptoms**:
- "permission denied" when accessing evidence files
- Failure to write to EVIDENCE_ROOT directory
- Access control violations

**Solutions**:
1. Check directory permissions:
   ```bash
   ls -ld $EVIDENCE_ROOT
   ```

2. Verify user/group ownership:
   ```bash
   stat $EVIDENCE_ROOT
   ```

3. Set proper permissions:
   ```bash
   chmod 755 $EVIDENCE_ROOT
   chown -R cyborg-user:cyborg-group $EVIDENCE_ROOT
   ```

## Deployment and Container Issues

### Docker Deployment Failures

**Symptoms**:
- Container fails to start
- "No such file or directory" errors
- Port mapping issues

**Solutions**:
1. Check Docker logs:
   ```bash
   docker logs cyborg-conductor-core
   ```

2. Verify volume mounts:
   ```bash
   docker inspect cyborg-conductor-core | grep -A 10 Mounts
   ```

3. Ensure correct image build:
   ```bash
   docker build -t cyborg-conductor-core .
   ```

### Kubernetes Deployment Problems

**Symptoms**:
- Pods stuck in "Pending" state
- Container crashes with "CrashLoopBackOff"
- ConfigMap not mounting properly

**Solutions**:
1. Check pod status:
   ```bash
   kubectl describe pod cyborg-conductor-core-xxxxx
   ```

2. Verify ConfigMap:
   ```bash
   kubectl get configmap cyborg-schema -o yaml
   ```

3. Check node resources:
   ```bash
   kubectl describe nodes
   ```

## Common Configuration Issues

### Invalid Configuration Files

**Symptoms**:
- "invalid configuration" errors
- Failed to parse YAML/JSON
- Missing required fields

**Solutions**:
1. Validate configuration with schema:
   ```bash
   cat config/app.yaml | yamllint
   ```

2. Check for syntax errors:
   ```bash
   python -m json.tool config/app.json
   ```

3. Use default configuration as reference:
   ```bash
   cp config/app.yaml.dist config/app.yaml
   ```

### Environment Variable Issues

**Symptoms**:
- "environment variable not set" errors
- Incorrect values in runtime
- Configuration not being overridden

**Solutions**:
1. List all environment variables:
   ```bash
   env | grep CYBORG
   ```

2. Set missing variables:
   ```bash
   export CYBORG_PORT=8080
   ```

3. Verify variable precedence:
   ```bash
   echo $CYBORG_PORT
   ```

## Advanced Troubleshooting

### Debugging with Logs

To enable debug logging:
```bash
export LOG_LEVEL=debug
./cyborg-conductor-core
```

To capture detailed traces:
```bash
go run cmd/server/main.go -log-level=debug -trace
```

### Profiling Performance

For performance profiling:
```bash
go build -o cyborg-conductor-core cmd/server/main.go
./cyborg-conductor-core -cpuprofile=cpu.prof
go tool pprof cpu.prof
```

### Memory Leak Detection

To detect memory leaks:
```bash
go build -o cyborg-conductor-core cmd/server/main.go
./cyborg-conductor-core -memprofile=mem.prof
go tool pprof mem.prof
```

## When to Seek Help

If you're experiencing issues that aren't resolved by this guide:

1. **Check existing issues** in the GitHub repository
2. **Review system logs** for detailed error messages
3. **Consult the documentation** for configuration details
4. **Reach out to support** with detailed error logs and reproduction steps

## Prevention Best Practices

- Regularly monitor system metrics and logs
- Keep configuration files in version control
- Test deployments in staging environments
- Implement proper backup and recovery procedures
- Schedule regular system maintenance and updates

## Additional Resources

- [GitHub Issues](https://github.com/toxicoder/cyborg-conductor-core/issues)
- [Technical Documentation](../technical_design_doc.md)
- [User Guide Index](index.md)
- [Contribution Guidelines](../../CONTRIBUTING.md)