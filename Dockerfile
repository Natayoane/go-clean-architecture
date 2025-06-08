# Build stage
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Install wire for dependency injection
RUN go install github.com/google/wire/cmd/wire@latest

# Generate wire dependencies
RUN cd cmd/ordersystem && wire

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/ordersystem

# Final stage
FROM alpine:latest

WORKDIR /app

# Install necessary packages
RUN apk add --no-cache bash curl

# Copy the binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/cmd/ordersystem/.env .

# Download wait-for-it script
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh

# Expose ports
EXPOSE 8000 8080 50051

# Wait for services to be ready before starting the application
CMD ["wait-for-it.sh", "mysql:3306", "--", "wait-for-it.sh", "rabbitmq:5672", "--", "./main"] 
