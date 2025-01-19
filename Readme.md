# GoRestAPI Project

## Overview
GoRestAPI is a simple RESTful API built using the Gin framework in Go. It provides basic CRUD functionality to manage user data. The project demonstrates how to set up endpoints, handle JSON data, and work with files for persistent storage.

## Features
- **GET /users**: Retrieve all users.
- **GET /users/:id**: Retrieve a specific user by their ID.
- **POST /users**: Add a new user.
- **PUT /users/:id**: Replace a specific user by id
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
curl -X GET http://localhost:8080/users
```

### GET /users/:id
Retrieve a specific user by their ID.
```bash
curl -X GET http://localhost:8080/users/1
```

### POST /users
Add a new user by sending a JSON payload.
```bash
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{"Id": 1, "Name": "John Doe", "birthYear": 30}'
```

### DELETE /users/:id
Delete a user by their ID.
```bash
curl -X DELETE http://localhost:8080/users/1
```

### PUT /users/:id
Replace a specific user by id
```bash
curl -X PUT 127.0.0.1:8080/users/1 -H "Content-Type: application/json" -d '{"name":"John", "surname":"Doe", "birthYear":1979}' -v
```

## File Storage
The project uses a JSON file to persist user data. The `jsonTools` package handles reading from and writing to this file. Ensure the JSON file exists and is properly formatted before starting the application.

## Dependencies
- [Gin](https://github.com/gin-gonic/gin): HTTP web framework

## License
This project is licensed under the MIT License.


