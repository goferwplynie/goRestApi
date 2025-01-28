# GoRestAPI Project

## Overview

GoRestAPI is a simple RESTful API built using the Gin framework in Go. It provides basic CRUD functionality to manage user data. The project demonstrates how to set up endpoints, handle JSON data, and work with files for persistent storage.

## Features

- **GET /users**: Retrieve all users.
- **GET /users/:id**: Retrieve a specific user by their ID.
- **POST /users**: Add a new user.
- **PATCH /users/:id**: Update user by his id.
- **DELETE /users/:id**: Delete a user by their ID.

## Prerequisites

- Go (1.16 or higher)
- Git

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/goferpwlynie/goRestApi.git
   cd goRestApi
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. create .env file
4. add `DATABASE_URL` variable to .env
5. add `ADMIN_LOGIN` and `ADMIN_PASSWORD` variables to .env
6. add `JWT_SECRET` variable to .env

## Usage

1. Start the server:

   ```bash
   go run main.go
   ```

2. The server will start on `http://localhost:8080`.

3. Use a tool like [Postman](https://www.postman.com/) or `curl` to interact with the API endpoints.

## Endpoints

### GET /users

Retrieve a list of all users.

```bash
curl -X GET localhost:8080/users
```

### GET /users/:id

Retrieve a specific user by their ID.

```bash
curl -X GET localhost:8080/users/1
```

### POST /login

login as admin and return jwt token

```bash
curl -X POST localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"name": admin_login, "password": password}'
```

### POST /users (protected)

Add a new user by sending a JSON payload.

```bash
curl -X POST localhost:8080/users \
-H "Authorization: jwt_token" \
-H "Content-Type: application/json" \
-d '{"Id": 1, "Name": "John Doe", "birthYear": 1999}'
```

### DELETE /users/:id (protected)

Delete a user by their ID.

```bash
curl -H "Authorization: jwt_token" \
-X DELETE localhost:8080/users/1
```

### PATCH /users/:id (protected)

Update user by his id.

```bash
curl -X PATCH localhost:8080/users/1 \
-H "Authorization: jwt_token" \
-H "Content-Type: application/json"\
-d '{"birthYear":1971}' -v
```

## File Storage

project uses PostgreSQL database hosted on [neon](https://neon.tech)

## Authentication

this project uses jwt tokens to authorize admins to endpoints that can change some data.

## Dependencies

- [jwt](github.com/golang-jwt/jwt/v5): jwt library for go
- [Gin](https://github.com/gin-gonic/gin): HTTP web framework
- [pgx](https://github.com/jackc/pgx): PostgreSQL driver

## License

This project is licensed under the MIT License.
