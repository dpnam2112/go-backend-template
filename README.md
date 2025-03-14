# Go API template

A simple template for API development in Go.

## Features

- RESTful API with `Gin`
- PostgreSQL integration with `sqlc`
- Database migrations using `golang-migrate`
- Structured project layout
- Dockerized deployment
- Environment variable management using `viper`
- Logging with `logrus`

## Project Structure

```
biglead-chatbots/
│── cmd/                  # Main application entry points
│   ├── server/           # Main service entrypoint
│── internal/             # Internal packages (business logic)
│   ├── config/           # Configuration management
│   ├── database/         # Database connection and queries
│   ├── handlers/         # HTTP handlers/controllers
│   ├── services/         # Business logic layer
│   ├── repositories/     # Data access layer
│   ├── providers/        # Dependency injection
│   └── middleware/       # Middleware (logging, auth, etc.)
│── api/                  # OpenAPI specifications
│── migrations/           # Database migrations
│── test/                 # Integration tests
│── .env                  # Environment variables
│── Dockerfile            # Containerization
│── Makefile              # Automation scripts
│── go.mod                # Go module dependencies
│── go.sum                # Dependency checksums
│── main.go               # Entry point
```

## Prerequisites

- Go 1.21+
- PostgreSQL 13+
- `sqlc` installed (`go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`)
- `golang-migrate` installed (`go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`)
- Docker (optional for containerized deployment)

## Setup & Installation

### 1. Clone the Repository
```sh
git clone https://github.com/your-org/aiverse-project.git
cd github.com/dpnam2112/go-backend-template
```

### 2. Initialize Dependencies
```sh
go mod tidy
```

### 3. Configure the Environment
Create a `.env` file:
```sh
PORT=8080
DATABASE_URL=postgres://user:password@localhost:5432/mydb?sslmode=disable
LOG_LEVEL=debug
```

### 4. Run Migrations
Ensure your database is running and apply migrations:
```sh
migrate -path migrations -database "$DATABASE_URL" up
```

To rollback the last migration:
```sh
migrate -path migrations -database "$DATABASE_URL" down 1
```

### 5. Generate SQL Code
Run `sqlc` to generate database access code:
```sh
sqlc generate
```

### 6. Run the Server

Create a config file `config.env` to contain environment variable (You can change the file's name in the package `config`:
```env
# config.env
PORT=8080
POSTGRES_URI=postgres://nam:123@localhost:5432/dbname
LOG_LEVEL=debug
```

Then run the server:
```sh
go run main.go
```

### 7. Test API Endpoints
#### Create a User
```sh
curl -X POST "http://localhost:8080/users" -H "Content-Type: application/json" -d '{
  "name": "Alice",
}'
```

#### Get User by ID
```sh
curl -X GET "http://localhost:8080/users/:id"
```

## Docker Deployment

### 1. Build the Docker Image
```sh
docker build -t go-backend-template -f docker/Dockerfile.local .
```

### 2. Run the Container
```sh
docker run -p 8080:8080 --env-file config.env go-backend-template
```

## Generate Swagger documentation
Run the following command:

```bash
swag init
```

## License
This project is licensed under the MIT License.

## Contributors
- [Nam](https://github.com/dpnam2112)

---
For issues and contributions, open a pull request or submit an issue on GitHub.
