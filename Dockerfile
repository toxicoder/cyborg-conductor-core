FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cyborg-conductor-core-server ./cmd/server

# Final stage
FROM alpine:latest

# Install curl for health checks
RUN apk --no-cache add curl

# Create non-root user
RUN adduser -D -u 1000 cyborg

WORKDIR /root/

# Copy the binary
COPY --from=builder /app/cyborg-conductor-core-server .

# Create directories
RUN mkdir -p /var/lib/cyborg/evidence

# Change ownership
RUN chown -R cyborg:cyborg /var/lib/cyborg

# Switch to non-root user
USER cyborg

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8080/healthz || exit 1

# Run the binary
CMD ["./cyborg-conductor-core-server"]