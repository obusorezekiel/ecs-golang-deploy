# Stage 1: Build the Go application
FROM golang:1.22.2-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o /gin-gorm-crud

# Stage 2: Run the Go application
FROM alpine:latest

# Set environment variable for the application
# Set environment variables for the application
ENV GIN_MODE=release \
    DATABASE_HOST=localhost \
    DATABASE_PORT=5432 \
    DATABASE_USERNAME=obusor \
    DATABASE_PASSWORD=password \
    DATABASE_NAME=mydatabase

# Expose the port the app runs on
EXPOSE 8888

# Copy the pre-built binary from the builder stage
COPY --from=builder /gin-gorm-crud /gin-gorm-crud

# Run the Go app
CMD ["/gin-gorm-crud"]
