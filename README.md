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

# ⚡ How It Works Step by Step

When you run docker-compose up --build:

PostgreSQL database container starts.
The Go application starts and connects to the database.
The migration (init.sql) runs and creates tables if necessary.
A user can:
Register by sending a POST /register request.
Login and receive a JWT token via POST /login.
Make requests to public or protected endpoints.

# 🎯 Summary: Why This Application?

This is a minimalist yet thoughtfully designed backend for an online store, featuring a simple architecture and basic functionality:

A foundation for a real project with Go + Fiber.
A solid example of a REST API architecture with JWT and Docker.
A ready-to-expand base for more complex business logic.
