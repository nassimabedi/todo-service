# Step 1: Build the Go application
FROM golang:1.23-alpine AS builder


# Use a lightweight base image with the migrate tool
FROM migrate/migrate:latest


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Copy migrations to the container
COPY ./migrations /migrations

# Download the Go modules
RUN go mod tidy

# Copy the entire project into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o todo-service .

# Step 2: Create a minimal image to run the Go application
FROM alpine:latest

# Install the necessary dependencies to run the Go binary (if needed)
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Go binary from the builder image
COPY --from=builder /app/todo-service .

# Expose the port the app will run on
EXPOSE 8080

# Default command to apply migrations
CMD ["-path", "/migrations", "-database", "postgres://postgres:password@db:5432/todo_db?sslmode=disable", "up"]


# Command to run the application
CMD ["./todo-service"]

