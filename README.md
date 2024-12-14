# todo-service

## Overview

The Todo Service is a RESTful API that allows users to manage Todo items. It supports:

    1. Uploading files to S3.
    2. Creating Todo items stored in a PostgreSQL database.
    3. Sending Todo item details to an SQS queue.

This project is built using the ** Gin framework ** and follows ** Hexagonal Architecture ** principles for better modularity and testability.
## Features

- File Upload: Users can upload files that are stored in an S3 bucket, with support for file validation.
- Todo Management: Users can create Todo items with descriptions, due dates, and optional file associations.
- SQS Integration: Each Todo item is sent as a message to an SQS queue after creation.
- Hexagonal Architecture: Clean separation of domain, application logic, and infrastructure concerns.
- Dockerized Setup: Includes a PostgreSQL database and mock AWS services using LocalStack.
- Testing & Benchmarks: Unit tests and benchmarks for key operations using mocks.

## Technologies Used

- Programming Language: Go (Golang)
- Framework: Gin
- Database: PostgreSQL
- Object Storage: S3 (mocked via LocalStack)
- Message Queue: SQS (mocked via LocalStack)
- Containerization: Docker, Docker Compose
- Testing: Go testing framework with mocks

## Project Structure
```
todo-service/
├── cmd/
│   └── main.go                 # Main entry point for the application
├── config/
│   └── config.go               # Configuration setup for environment variables
├── db/
│   └── db.go                   # Database connection and migrations
├── internal/
│   ├── domain/                 # Domain models and interfaces
│   │   ├── interfaces.go       # Interfaces for S3, SQS, and Todo services
│   │   └── todo.go             # Todo entity
│   ├── ports/                  # Adapters for external services
│   │   ├── s3.go               # Interface implementation for S3 interactions
│   │   └── sqs.go              # Interface implementation for SQS interactions
│   └── services/               # Application services
│       ├── sqs_service.go      # SQS service implementation
│       └── todo_service.go     # Todo service implementation
├── Dockerfile                  # Dockerfile for the Go application
├── docker-compose.yml          # Docker Compose setup for PostgreSQL and LocalStack
├── go.mod                      # Go module dependencies
└── README.md                   # Project documentation

```

## Getting Started
### Prerequisites

Ensure you have the following installed on your system:

- Docker
- DockerCompose
- Go (version 1.23 or newer)
- Make

### Environment Variables

Create a .env file in the root directory with the following:
```
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=todo_db
AWS_REGION=us-east-1
S3_BUCKET_NAME=my-todo-bucket
SQS_QUEUE_URL=http://localhost:4566/000000000000/todo-queue
```
### Running the Application

1. **Clone the repository:**

```git clone https://github.com/your-repo/todo-service.git
cd todo-service
```

2. **Start services using Docker Compose:**

    ```make run```

3. **Access the application:** The API will be available at ```http://localhost:8080```.

## API Endpoints
1. **Upload File**
   - Endpoint: POST /upload
   - Description: Upload a file to S3.
   - Request: Form data with the file (file).
   -     
    Response:

    ```{
      "fileId": "generated-file-id",
      "fileUrl": "https://s3-bucket-url/file"
    }

2. **Create TodoItem**
   - Endpoint: POST /todo
   - Description: Create a new Todo item.
   - Request:
  ```
  {"description": "Complete project documentation",
  "dueDate": "2024-12-15T12:00:00Z",
  "fileId": "optional-file-id"}
```
   - Response:
       ```
      {"id": "generated-uuid",
      "description": "Complete project documentation",
      "dueDate": "2024-12-15T12:00:00Z",
      "fileId": "optional-file-id" }

## Testing
### Run Unit Tests

```make test```

### Run Benchmarks

``make benchmark``

## Makefile

The ``Makefile`` simplifies common tasks:

```run:
    docker-compose up --build

test:
    go test ./...

benchmark:
    go test -bench ./...

