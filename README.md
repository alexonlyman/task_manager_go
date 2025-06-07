# Task Manager API

A RESTful API service for task management built with Go, using clean architecture principles.

## Features

- Create, read, update, and delete tasks
- RESTful API endpoints
- PostgreSQL database integration
- Docker support for testing
- Comprehensive test coverage
- Clean architecture implementation

## Project Structure

```
task_manager_go/
├── config/         # Configuration and database setup
├── controller/     # HTTP request handlers
├── model/         # Data models
├── repository/    # Data access layer
├── service/       # Business logic
└── test/          # Integration tests
```

## Prerequisites

- Go 1.16 or higher
- PostgreSQL 13 or higher
- Docker (for running tests)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/task_manager_go.git
cd task_manager_go
```

2. Install dependencies:
```bash
go mod download
```

3. Set up the database:
- Create a PostgreSQL database named `taskdb`
- Update the database connection string in `config/config.go` if needed

## Running the Application

Start the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Tasks

- `POST /tasks` - Create a new task
- `GET /tasks` - Get all tasks
- `GET /tasks/{id}` - Get a task by ID
- `PATCH /tasks/{id}` - Update a task
- `DELETE /tasks/{id}` - Delete a task

### Example Request

Create a new task:
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Complete project",
    "status": "Pending"
  }'
```

## Testing

Run all tests:
```bash
go test ./...
```

Run specific test:
```bash
go test ./controller
go test ./service
go test ./repository
```

## Architecture

The project follows clean architecture principles:

- **Model Layer**: Defines the data structures
- **Repository Layer**: Handles data persistence
- **Service Layer**: Implements business logic
- **Controller Layer**: Manages HTTP requests and responses

## Dependencies

- [Gorilla Mux](https://github.com/gorilla/mux) - HTTP router
- [GORM](https://gorm.io/) - ORM library
- [TestContainers](https://golang.testcontainers.org/) - Testing with Docker
- [Testify](https://github.com/stretchr/testify) - Testing framework

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 