# Build stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Install necessary build tools
RUN apk add --no-cache git gcc musl-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd/server

# Final stage
FROM alpine:latest

WORKDIR /app

# Install necessary runtime dependencies
RUN apk --no-cache add ca-certificates tzdata sqlite

# Copy the binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Create necessary directories
RUN mkdir -p /app/uploads/images /app/data

# Set environment variables
ENV GIN_MODE=release
ENV PORT=8080

# Create volumes for persistent data
VOLUME ["/app/data", "/app/uploads"]

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"] 