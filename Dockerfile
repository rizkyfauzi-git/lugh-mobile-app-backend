# Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Run stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .
# Copy .env file if you want to include it (though it's better to use environment variables in hosting)
# COPY .env .

# Expose port
EXPOSE 8080

# Command to run
CMD ["./main"]
