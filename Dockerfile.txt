# Minimal Dockerfile for testing GitHub Actions + Docker Hub

# Stage 1: Build Go binary
FROM golang:1.24 AS builder
WORKDIR /app

# Copy code
COPY . .

# Build binary
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o agent_ai main.go

# Stage 2: Minimal runtime
FROM debian:bookworm-slim
WORKDIR /app

# Copy binary
COPY --from=builder /app/agent_ai ./

# Expose port
EXPOSE 8080

# Run app
CMD ["./agent_ai"]
