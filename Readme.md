# Microservices Cafe POS Backend Using Golang




## Project Structure

```
cafe-pos/
├── api-gateway/                  # API Gateway to route requests to services
│   ├── main.go                   # Entry point for the API Gateway
│   ├── routes/                   # Routing configuration for services
│   ├── middlewares/              # Common middlewares (e.g., auth, rate limit)
│   └── config/                   # Configurations (e.g., ports, environment)
├── services/
│   ├── user-service/             # Handles user management (staff login, roles)
│   │   ├── main.go               # Entry point for user service
│   │   ├── controllers/          # Business logic for user operations
│   │   ├── models/               # User model and database schema
│   │   ├── routes/               # Route definitions
│   │   └── config/               # Configuration (e.g., DB connection, JWT secret)
│   ├── menu-service/             # Manages menu items (view, add, update items)
│   │   ├── main.go               # Entry point for menu service
│   │   ├── controllers/          # Business logic for menu operations
│   │   ├── models/               # Menu item schema
│   │   ├── routes/               # Route definitions
│   │   └── config/               # Configuration (e.g., DB connection)
│   ├── order-service/            # Handles orders and order items
│   │   ├── main.go               # Entry point for order service
│   │   ├── controllers/          # Business logic for order operations
│   │   ├── models/               # Order schema
│   │   ├── routes/               # Route definitions
│   │   └── config/               # Configuration (e.g., DB connection)
│   ├── payment-service/          # Manages payment processing and statuses
│   │   ├── main.go               # Entry point for payment service
│   │   ├── controllers/          # Business logic for payment operations
│   │   ├── models/               # Payment schema
│   │   ├── routes/               # Route definitions
│   │   └── config/               # Configuration (e.g., payment gateway credentials)
│   └── guest-service/            # Manages guest sessions and orders from guests
│       ├── main.go               # Entry point for guest service
│       ├── controllers/          # Guest-specific ordering logic
│       ├── models/               # Guest session and order schema
│       ├── routes/               # Route definitions
│       └── config/               # Configuration (e.g., Redis connection)
├── libs/                         # Shared libraries across services (e.g., JWT handling)
│   ├── jwt/                      # JWT generation and verification logic
│   ├── redis/                    # Redis connection handling
│   └── utils/                    # Common utilities and helpers
├── docker-compose.yml            # Docker Compose file to manage services
├── .env                          # Environment variables (e.g., DB_URL, Redis URL)
└── README.md                     # Project documentation
```