# Build stage
FROM golang:1.22-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=0 ensures the binary is statically linked
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Run stage
FROM alpine:latest

# Install certificates for secure connections
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose port (default from .env)
EXPOSE 8080

# Command to run
CMD ["./main"]
