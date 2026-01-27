# Contributing to Cyborg Conductor Core

Thank you for your interest in contributing to the Cyborg Conductor Core! This document outlines the guidelines for contributing to this project.

## How to Contribute

### Reporting Bugs

If you find a bug, please create an issue on our GitHub repository with:

- A clear description of the problem
- Steps to reproduce the issue
- Expected vs actual behavior
- Environment details (Go version, OS, etc.)

### Suggesting Enhancements

We welcome suggestions for new features or improvements:

1. Open an issue to discuss your idea
2. Provide context about the problem it solves
3. Include any relevant design considerations

### Pull Request Process

1. Fork the repository
2. Create a feature branch from `main`
3. Make your changes with clear, descriptive commit messages
4. Ensure all tests pass
5. Update documentation as needed
6. Submit a pull request to the `main` branch

## Development Workflow

### Setting Up Your Environment

1. Clone the repository:

   ```bash
   git clone https://github.com/toxicoder/cyborg-conductor-core.git
   cd cyborg-conductor-core
   ```

2. Install required dependencies:
   - Go 1.22 or higher
   - Protobuf tools
   - Docker (for integration tests)

3. Install Go modules:

   ```bash
   go mod tidy
   ```

### Code Style Guidelines

This project follows specific style guides for each language:

- **Go**: Follow the [Go Style Guide](docs/style_guides/golang/)
- **Markdown**: Follow the [Markdown Style Guide](docs/style_guides/markdown/)
- **Protobuf**: Follow the [Protobuf Style Guide](docs/style_guides/proto/)
- **Shell/Bash**: Follow the [Bash Style Guide](docs/style_guides/bash/)
- **TypeScript**: Follow the [TypeScript Style Guide](docs/style_guides/typescript/)
- **Python**: Follow the [Python Style Guide](docs/style_guides/python/)

### Testing

All code changes must include appropriate tests:

- Unit tests for new functionality
- Integration tests where applicable
- Ensure test coverage is maintained (target 85%+)

Run all tests with:

```bash
go test ./... -v
```

Run tests with coverage:

```bash
go test ./... -coverprofile=coverage.txt -covermode=atomic
```

### Documentation

All new features or significant changes must include documentation updates:

- Update the technical design document if core architecture changes
- Add usage examples where applicable
- Ensure documentation follows the project's style guides

## Commit Guidelines

- Use clear, descriptive commit messages
- Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification
- Keep commits focused on a single logical change

## Review Process

1. All pull requests require at least one review from core maintainers
2. CI checks must pass before merging
3. Code must adhere to project style guides
4. Documentation must be updated to reflect changes
