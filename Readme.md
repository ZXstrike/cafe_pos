# Microservices Cafe POS Backend Using Golang

## Overview
This project is a backend system for a cafe point-of-sale (POS) application built using Golang. The architecture follows a microservices approach to ensure modularity, scalability, and ease of maintenance.

## Project Structure
```plaintext
cafe-pos/
├── api-gateway/                  # API Gateway to route requests to services
│   ├── main.go                   # Entry point for the API Gateway
│   ├── routes/                   # Routing configuration for services
│   ├── middlewares/              # Common middlewares (e.g., authentication, rate limiting)
│   └── config/                   # Configuration files (e.g., ports, environment variables)
├── services/                     # Microservices directory
│   ├── user-service/             # Handles user management (e.g., staff login, roles)
│   │   ├── main.go               # Entry point for the user service
│   │   ├── db/                   # Database initialization and interactions
│   │   ├── handler/              # Business logic for user operations
│   │   ├── models/               # User models and database schemas
│   │   ├── routes/               # API route definitions
│   │   ├── utils/                # Utility functions for the user service
│   │   └── config/               # Configuration (e.g., DB connection, JWT secret)
│   ├── menu-service/             # Manages menu items (e.g., view, add, update)
│   │   ├── main.go               # Entry point for the menu service
│   │   ├── db/                   # Database interactions
│   │   ├── handler/              # Business logic for menu operations
│   │   ├── models/               # Menu item models and schemas
│   │   ├── routes/               # API route definitions
│   │   ├── utils/                # Utility functions for the menu service
│   │   └── config/               # Configuration (e.g., DB connection)
│   ├── order-service/            # Handles orders and order items
│   │   ├── main.go               # Entry point for the order service
│   │   ├── db/                   # Database interactions
│   │   ├── handler/              # Business logic for order operations
│   │   ├── models/               # Order models and schemas
│   │   ├── routes/               # API route definitions
│   │   ├── utils/                # Utility functions for the order service
│   │   └── config/               # Configuration (e.g., DB connection)
│   ├── payment-service/          # Manages payment processing and statuses
│   │   ├── main.go               # Entry point for the payment service
│   │   ├── db/                   # Database interactions
│   │   ├── handler/              # Business logic for payment operations
│   │   ├── models/               # Payment models and schemas
│   │   ├── routes/               # API route definitions
│   │   ├── utils/                # Utility functions for the payment service
│   │   └── config/               # Configuration (e.g., payment gateway credentials)
│   └── guest-service/            # Manages guest sessions and orders from guests
│       ├── main.go               # Entry point for the guest service
│       ├── db/                   # Database interactions
│       ├── handler/              # Guest-specific ordering logic
│       ├── models/               # Guest session and order models
│       ├── routes/               # API route definitions
│       ├── utils/                # Utility functions for the guest service
│       └── config/               # Configuration (e.g., Redis connection)
├── libs/                         # Shared libraries or utilities across services
```

## Key Features
- **API Gateway:** Centralized entry point for routing requests to the respective microservices.
- **User Management:** Staff login and role management.
- **Menu Management:** Add, view, and update menu items.
- **Order Management:** Process orders and manage order items.
- **Payment Processing:** Handle payment transactions and track statuses.
- **Guest Ordering:** Enable guests to place orders through QR code-based sessions.

## Technologies Used
- **Programming Language:** Golang
- **Database:** PostgreSQL (or other relational databases as configured)
- **Message Broker:** RabbitMQ (optional, for inter-service communication)
- **API Gateway:** Custom implementation using Golang
- **Authentication:** JWT-based authentication
- **Caching:** Redis (for guest session and other transient data)

## Getting Started
### Prerequisites
- Go 1.19 or later
- Docker and Docker Compose
- PostgreSQL (or preferred database)
- Redis (optional, for caching)

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/cafe-pos.git
   cd cafe-pos
   ```
2. Build the project:
   ```bash
   go build ./...
   ```
3. Configure the environment variables for each service (e.g., database connection strings, JWT secrets).

### Running with Docker Compose
1. Make sure Docker and Docker Compose are installed on your system.
2. Navigate to the project directory containing the `docker-compose.yml` file.
3. Start all services using:
   ```bash
   docker-compose up --build
   ```
4. To stop the services, use:
   ```bash
   docker-compose down
   ```

### Running the Services Individually
- Start the API Gateway:
  ```bash
  cd api-gateway
  go run main.go
  ```
- Start individual services (e.g., user service):
  ```bash
  cd services/user-service
  go run main.go
  ```

### Testing
Use tools like Postman or cURL to test API endpoints.

## Future Improvements
- Implement gRPC for inter-service communication.
- Add observability with tools like Prometheus and Grafana.
- Enhance CI/CD pipelines for automated testing and deployment.

## Contributing
Feel free to submit issues, fork the repository, and send pull requests. For major changes, please open a discussion first.

## License
This project is licensed under the [MIT License](LICENSE).

