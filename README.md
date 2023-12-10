# Crud-live


This is a simple CRUD (Create, Read, Update, Delete) API implemented in Go using PostgreSQL as the database. The project provides RESTful endpoints for managing user data.

## Requirements

Make sure you have the following installed before running the project:

- Go programming language
- PostgreSQL
- Required Go packages: `github.com/gorilla/mux`, `github.com/lib/pq`

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/AliSinaDevelo/crud-live.git
   cd crud-live

2. Set up the PostgreSQL database and provide the connection URL via the DATABASE_URL environment variable.

3. Install dependancies:
   ```bash
   go get -u github.com/gorilla/mux
   go get -u github.com/lib/pq

4. Build and run the project:
   ```bash
    go build
    ./crud-live

## API Endpoints

GET /users: Retrieve a list of all users.

GET /users/{id}: Retrieve details of a specific user by ID.

POST /users: Create a new user. Send a JSON payload with the name and email fields.

PUT /users/{id}: Update an existing user by ID. Send a JSON payload with the updated name and email fields.

DELETE /users/{id}: Delete a user by ID.




