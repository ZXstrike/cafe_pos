# Microservices Cafe POS Backend Using Golang

## Project Structure

cafe-pos/
├── api-gateway/                  # API Gateway to route requests to services
│   ├── main.go                   # Entry point for the API Gateway
│   ├── routes/                   # Routing configuration for services
│   ├── middlewares/              # Common middlewares (e.g., auth, rate limit)
│   └── config/                   # Configurations (e.g., ports, environment)
├── services/
│   ├── user-service/             # Handles user management (staff login, roles)
│   │   ├── main.go               # Entry point for user service
│   │   ├── db/                   
│   │   ├── handler/              # Business logic for user operations
│   │   ├── models/               # User model and database schema
│   │   ├── routes/               # Route definitions
│   │   ├── utils/                # 
│   │   └── config/               # Configuration (e.g., DB connection, JWT secret)
│   ├── menu-service/             # Manages menu items (view, add, update items)
│   │   ├── main.go               # Entry point for menu service
│   │   ├── db/                   
│   │   ├── handler/              # Business logic for menu operations
│   │   ├── models/               # Menu item schema
│   │   ├── routes/               # Route definitions
│   │   ├── utils/                # 
│   │   └── config/               # Configuration (e.g., DB connection)
│   ├── order-service/            # Handles orders and order items
│   │   ├── main.go               # Entry point for order service
│   │   ├── db/                   
│   │   ├── handler/              # Business logic for order operations
│   │   ├── models/               # Order schema
│   │   ├── routes/               # Route definitions
│   │   ├── utils/                # 
│   │   └── config/               # Configuration (e.g., DB connection)
│   ├── payment-service/          # Manages payment processing and statuses
│   │   ├── main.go               # Entry point for payment service
│   │   ├── db/                   
│   │   ├── handler/              # Business logic for payment operations
│   │   ├── models/               # Payment schema
│   │   ├── routes/               # Route definitions
│   │   ├── utils/                # 
│   │   └── config/               # Configuration (e.g., payment gateway credentials)
│   └── guest-service/            # Manages guest sessions and orders from guests
│       ├── main.go               # Entry point for guest service
│       ├── db/                   
│       ├── handler/              # Guest-specific ordering logic
│       ├── models/               # Guest session and order schema
│       ├── routes/               # Route definitions
│       ├── utils/                # 
│       └── config/               # Configuration (e.g., Redis connection)
├── libs/ ```