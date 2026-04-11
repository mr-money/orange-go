# CLAUDE.md

This file provides guidance for Claude Code when working with the orange-go project.

## Overview

**orange-go** is a microservice web framework built with Gin, following Domain-Driven Design (DDD) architecture. It supports multiple microservices, MySQL, MongoDB, Redis, and RabbitMQ.

- **Framework**: Gin web framework
- **Architecture**: DDD (Domain-Driven Design)
- **Layering**: Model → Repository → Service → App
- **Go Version**: ≥ 1.21

## Quick Navigation

- [Project Structure](#project-structure) - Directory layout and key files
- [Architecture](#architecture) - DDD layers and patterns
- [Setup & Installation](#setup--installation) - How to get started
- [Development Guidelines](#development-guidelines) - Coding conventions
- [Testing](#testing) - Test patterns and practices

## Project Structure

```
orange-go/
├── AbTest/           # A/B testing modules
├── App/              # Application layer (HTTP handlers)
│   ├── Api/          # API application handlers
│   └── Index/        # Index handlers
├── Config/           # Configuration (TOML format)
├── Container/        # Service containers/entry points
│   └── Api/          # API service entry (main.go)
├── Database/         # Database migrations and setup
├── Library/          # Reusable libraries
│   ├── Cache/        # Redis cache
│   ├── Gorm/         # GORM (MySQL)
│   ├── Handler/      # Utility handlers (JWT, hash, etc.)
│   ├── Logger/       # Zap logging
│   ├── MongoDB/      # MongoDB client
│   ├── Payment/      # Payment integration (WeChat)
│   └── Wechat/       # WeChat integration
├── MiddleWare/       # Gin middleware (CORS, CSRF, JWT auth)
├── Model/            # Domain models/entities
├── Queue/            # Queue system (Redis/RabbitMQ via Machinery)
│   └── Worker/       # Queue workers
├── Repository/       # Repository layer (data access)
├── Routes/           # Route definitions
├── Service/          # Service layer (business logic)
└── Test/             # Tests
    ├── Api/          # API tests
    ├── Log/          # Logging tests
    └── Queue/        # Queue tests
```

## Architecture

### Layer Flow

```
HTTP Request
    ↓
Routes (routes/api.go, routes/web.go)
    ↓
App Layer (App/Api/) - HTTP handlers, request/response
    ↓
Service Layer (Service/) - Business logic
    ↓
Repository Layer (Repository/) - Data access abstraction
    ↓
Model Layer (Model/) - Data models & ORM
    ↓
Database (MySQL/MongoDB/Redis)
```

### Key Architectural Files

| File                    | Purpose                 |
| ----------------------- | ----------------------- |
| `Container/Api/main.go` | API service entry point |
| `Routes/api.go`         | API route definitions   |
| `Routes/web.go`         | Web route definitions   |
| `Config/config.go`      | Configuration loader    |
| `Database/migration.go` | Database migrations     |
| `Queue/config.go`       | Queue configuration     |

## Setup & Installation

### Local Development

```bash
# Enable Go modules
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

# Install dependencies
go mod tidy

# Copy and configure
cp Config/toml/web.toml.default Config/toml/web.toml
# Edit Config/toml/web.toml with your settings

# Build and run
cd Container/Api
go build -o app
./app
```

### Docker

```bash
# Build (--build-arg image={microservice name})
docker build -t orange-go/api --build-arg image=Api .

# Run
docker run -dp 8080:8080 --name orange_go orange-go/api
```

## Development Guidelines

### Adding New Endpoints

1. **Model** - Define data model in `Model/`
2. **Repository** - Create data access in `Repository/{Entity}/`
3. **Service** - Add business logic in `Service/{Entity}/`
4. **App** - Create HTTP handler in `App/Api/{Entity}/`
5. **Routes** - Add route in `Routes/api.go`

### Configuration

Configuration uses TOML format. To add new config:

1. Copy `Config/toml/web.toml.default` → `Config/toml/web.toml`
2. Define struct in `Config/web.go`
3. Initialize in `Config/config.go` init()

### Database Models

Define models in `Model/` with GORM tags. Example:

```go
// Model/user.go
type User struct {
    gorm.Model
    Name string `gorm:"column:name"`
}

func (User) TableName() string {
    return "user"
}

func (User) GetOption(key string) interface{} {
    // return engine, comment, charset
}
```

Add migration in `Database/mysql.go`:

```go
func getMysqlMigrations() []map[string]interface{} {
    return append(mysqlMigrations,
        map[string]interface{}{
            "engine":  Model.User{}.GetOption("engine"),
            "comment": Model.User{}.GetOption("comment"),
            "charset": Model.User{}.GetOption("charset"),
            "model":   Model.User{},
        },
    )
}
```

### Queue Tasks

1. Define task function in `Queue/Worker/Api/{App}/`
2. Register in `Queue/tasks.go`
3. Add task with `Queue.AddTask(Queue.ExampleTaskFunc, params)`

### Middleware

Available middleware in `MiddleWare/`:

- `CORS()` - CORS handling
- `CSRF()` - CSRF validation
- `CSRFToken()` - CSRF token generation
- `Auth()` - JWT authentication

## Tech Stack

| Component     | Library                          |
| ------------- | -------------------------------- |
| Web Framework | github.com/gin-gonic/gin         |
| ORM           | gorm.io/gorm + MySQL driver      |
| Caching       | github.com/go-redis/redis/v8     |
| Queue         | github.com/RichardKnop/machinery |
| MongoDB       | go.mongodb.org/mongo-driver      |
| Logging       | go.uber.org/zap                  |
| Config        | github.com/BurntSushi/toml       |
| JWT           | github.com/dgrijalva/jwt-go      |
| CSRF          | github.com/gorilla/csrf          |

## Testing

Tests are in the `Test/` directory. Run tests with:

```bash
go test ./Test/...
```

### Test Patterns

- API tests in `Test/Api/`
- Logging tests in `Test/Log/`
- Queue tests in `Test/Queue/`
- Follow table-driven testing patterns
- Use `github.com/go-playground/assert/v2` for assertions
- Use `Test/test.go` for test utilities (GET/POST helpers)

## Common Patterns

### Repository Pattern

```go
// Repository/User/user.go
func GetUserById(id uint64) *Model.User {
    var user Model.User
    Model.UserModel().Take(&user, id)
    return &user
}
```

### Service Layer

```go
// Service/User/user.go
func GetUserInfo(id uint64) *Model.User {
    return Repository.User.GetUserById(id)
}
```

### App/Handler Layer

```go
// App/Api/User/user.go
func GetUserInfo(c *gin.Context) {
    // get params, call service, respond
}
```

### Redis Cache

```go
// Set cache
Cache.Redis.Set(Cache.Cxt, key, data, expiration)

// Get cache
result, err := Cache.Redis.Get(Cache.Cxt, key).Result()

// Generate key
key := Cache.SetKey("user", "profile", userId)
```

## Key Conventions

- **Language**: Go 1.21+
- **Configuration**: TOML files in `Config/toml/`
- **Error Handling**: Use `github.com/pkg/errors`
- **Logging**: Use `Library/Logger` (zap)
- **Code Style**: Follow standard Go conventions
- **File Naming**: LowerCamelCase for files, PascalCase for types
- **Test Files**: Use `_test.go` suffix, placed in `Test/` directory

## Entry Points

| Service | Entry File              |
| ------- | ----------------------- |
| API     | `Container/Api/main.go` |

## Remember

1. Always follow the DDD layers: Model → Repository → Service → App
2. Add database migrations when creating new models
3. Register queue tasks in `Queue/tasks.go`
4. Configure TOML files before running locally
5. Keep test code in `Test/` directory with `_test.go` suffix
6. Do NOT add test/demo endpoints to production routes
