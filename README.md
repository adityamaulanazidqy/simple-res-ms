# Go Simple Microservice - Restaurant Management System

A practical example of microservice architecture implementation using Go language. This project demonstrates core microservice concepts including service decomposition, containerization, and independent deployment.

## Project Overview

This repository contains a complete restaurant management system built using microservice architecture patterns. Each service is independently deployable and follows single responsibility principles.

### Services Included

1. **Authentication Service** (Port 8081)
   - User registration and login functionality
   - JWT-based authentication (planned)
   - User management endpoints

2. **Order Service** (Port 8082)  
   - Order creation and management
   - User-order relationship handling
   - Order status tracking

3. **Product Service** (Port 8083)
   - Full CRUD operations for products
   - Product catalog management
   - Price and inventory tracking

## Learning Objectives

This project serves as an educational resource for understanding:

- **Microservice Architecture**: Service decomposition and independent deployment
- **Go HTTP Server**: Building RESTful APIs with Go's standard library
- **Docker Containerization**: Packaging and deploying Go applications
- **Service Communication**: Inter-service communication patterns
- **Testing Strategies**: Integration testing for REST services
- **Project Structure**: Organizing multi-service Go projects

## Technology Stack

- **Language**: Go 1.25
- **Containerization**: Docker & Docker Compose
- **HTTP Framework**: Standard `net/http` package
- **Data Storage**: In-memory storage (easily replaceable with databases)
- **Testing**: Go testing package with HTTP client testing

## Quick Start

### Prerequisites
- Go 1.25 installed
- Docker and Docker Compose installed

### Running with Docker (Recommended)

```bash
# Clone the repository
git clone <repository-url>
cd restaurant

# Start all services
docker-compose up --build

# Or for development with live reloading
docker-compose -f docker-compose.dev.yml up
```

### Manual Setup

```bash
# Start each service individually
cd services/authentication-service && go run main.go
cd services/order-service && go run main.go  
cd services/product-service && go run main.go
```

## Service Endpoints

### Authentication Service
- `POST /register` - Register new user
- `POST /login` - User authentication

### Order Service
- `POST /order` - Create new order
- `GET /order` - Retrieve all orders

### Product Service
- `POST /product` - Create new product
- `GET /product` - Get all products
- `GET /product/{id}` - Get product by ID
- `PUT /product/{id}` - Update product
- `DELETE /product/{id}` - Delete product

## Project Structure

```
restaurant/
├── model/                    # Shared data models
│   ├── user.go
│   ├── order.go
│   └── product.go
├── services/                 # Individual microservices
│   ├── authentication-service/
│   │   ├── main.go
│   │   ├── authentication_test.go
│   │   └── Dockerfile
│   ├── order-service/
│   │   ├── main.go
│   │   ├── order_test.go
│   │   └── Dockerfile
│   └── product-service/
│       ├── main.go
│       ├── product_test.go
│       └── Dockerfile
├── storage/                  # Shared storage layer
│   └── restaurant_db.go
├── docker-compose.yml        # Production deployment
├── docker-compose.dev.yml    # Development setup
└── go.mod                   # Go module configuration
```

## Educational Features

### Microservice Concepts Demonstrated

1. **Service Independence**: Each service can be developed, tested, and deployed separately
2. **Container Isolation**: Services run in isolated Docker containers
3. **Network Communication**: Services communicate through Docker networks
4. **Separation of Concerns**: Each service handles specific business domain
5. **Scalability**: Services can be scaled independently based on demand

### Go-Specific Learning Points

1. **HTTP Server Implementation**: Using Go's standard `net/http` package
2. **JSON Handling**: Request/response marshaling and unmarshaling
3. **Error Handling**: Proper HTTP status codes and error responses
4. **Testing**: Integration testing with HTTP clients
5. **Project Organization**: Multi-module Go project structure

## Testing

Each service includes comprehensive integration tests:

```bash
# Run tests for a specific service
cd services/authentication-service && go test -v
cd services/order-service && go test -v
cd services/product-service && go test -v
```

**Note**: Services must be running for integration tests to pass.

## Development Workflow

1. **Code Changes**: Modify service code in respective directories
2. **Testing**: Run integration tests to verify functionality
3. **Containerization**: Use Docker for consistent deployment
4. **Iteration**: Services can be updated independently

## Extending the Project

This project provides a solid foundation that can be extended with:

- Database integration (PostgreSQL, MongoDB)
- Message queues for service communication
- API Gateway implementation
- Service discovery mechanisms
- Monitoring and logging
- Authentication/Authorization enhancements
- CI/CD pipeline integration

## Contributing

This project is designed for learning purposes. Feel free to:
- Fork and modify for your learning
- Add new services or features
- Improve existing implementations
- Share your extensions with the community

## License

This project is available for educational use and modification.