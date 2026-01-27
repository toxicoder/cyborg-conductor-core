# Integration Tests

This directory contains all integration tests for the Cyborg Conductor Core system.

## Test Structure
- Integration tests verify end-to-end functionality
- Tests use Docker Compose for environment setup
- Tests cover the complete orchestration workflow

## Running Tests
```bash
cd test/integration
docker-compose up --build