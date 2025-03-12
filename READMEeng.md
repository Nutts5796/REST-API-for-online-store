![Скриншот](images/Go.png)
# 🛒 Mini REST API for Online Store

This is a REST API for a simple online store developed with Go + Fiber, using PostgreSQL as the database. The main goal is to provide functionality for users and products, including registration, authentication, and retrieving product and user data.


# 📌 Tech Stack

Go + Fiber — fast and convenient web framework.
PostgreSQL — reliable relational database.
Docker + Compose — for easy deployment and service management.
JWT Auth — secure user authentication.
Makefile — simplifies commands for running and development.

# 🚀 Quick Start

✅ Prerequisites
Make sure you have installed:

Docker
Docker Compose
Go

# 🔐 Authentication and User Management

The application implements user authentication via JWT tokens. It includes basic user management features:

✅ User Registration /register

Expects a JSON payload:
    {
  "username": "user_name",
  "email": "email@domain.com",
  "password": "password"
}

Hashes the password.
Saves the user to the PostgreSQL users table.
Returns a 201 Created status and a success message.

✅ User Login /login

Expects a JSON payload:
    {
  "email": "email@domain.com",
  "password": "password"
}

Verifies if a user exists with the provided email.
Compares the hashed password with the given password.
On successful validation, generates a JWT token (currently returns a dummy token, but can easily be replaced with a real JWT).
Returns the token in JSON format:
    {
  "token": "jwt-token"
}

✅ Logout /logout and Refresh Token /refresh-token

Stub endpoints: return 200 OK and a simple message.
Can be extended to include revoked tokens or refresh token functionality.

# 📦 Product Management

Provides the ability to retrieve product listings from the PostgreSQL database.

✅ Get Product List /products

Returns an array of products in JSON format:
    [
  {
    "id": 1,
    "name": "Product Name",
    "description": "Product Description",
    "price": 99.99
  },
  ...
]

Data comes from the products table.
Open access (no authentication required).

# 👥 User Management

Allows retrieval of all registered users. This is a protected endpoint, available only to authenticated users.

✅ Get User List /users

Returns JSON with an array of users:
    {
  "users": [
    {
      "id": 1,
      "username": "user1",
      "email": "user1@example.com"
    },
    ...
  ]
}

Data comes from the users table.
In the future, roles can be added (e.g., user, admin).

# 🏗 Project Structure

The project has a clean and scalable code structure:

├── Dockerfile                # Build description for the Go application image
├── Makefile                  # Command automation for run, down, migrate, lint, test
├── READMEeng.md              # Documentation in English
├── READMErus.md              # Documentation in Russian
├── cmd
│   └── main.go               # Application entry point: server initialization and DB connection
├── docker-compose.yml        # Docker Compose configuration (Go API + PostgreSQL)
├── go.mod                    # Go module dependencies
├── go.sum                    # Checksum files for modules
├── internal
│   ├── handler               # HTTP request handlers package
│   │   ├── auth.go           # Registration, login, logout, token refresh handlers
│   │   ├── product.go        # Product list retrieval handler
│   │   └── routes.go         # Application route registration
│   └── middleware
│       └── jwt.go            # JWT token verification middleware (authentication)
├── migrations
│   └── init.sql              # Initial migration script for users and products tables
└── model
    └── user.go               # User model structure description (User)


# ⚡ How It Works Step by Step

When you run docker-compose up --build:

PostgreSQL database container starts.
The Go application starts and connects to the database.
The migration (init.sql) runs and creates tables if necessary.
A user can:
Register by sending a POST /register request.
Login and receive a JWT token via POST /login.
Make requests to public or protected endpoints.

# 🔧 Possible Improvements (Roadmap)

✅ Implement full JWT authentication with refresh tokens.
✅ Add user roles (admin, user).
✅ Add CRUD endpoints for managing products:

POST /products — add a product.
PUT /products/:id — update a product.
DELETE /products/:id — delete a product.
✅ Integrate with Swagger for automatic API documentation generation.
✅ Add integration tests using testcontainers-go.

# 🎯 Summary: Why This Application?

This is a minimalist yet thoughtfully designed backend for an online store, featuring a simple architecture and basic functionality:

A foundation for a real project with Go + Fiber.
A solid example of a REST API architecture with JWT and Docker.
A ready-to-expand base for more complex business logic.