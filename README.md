# easypark

[Trello board](https://trello.com/invite/b/lGdfavnm/ATTI15a8afbd1ced04b229e8f2380279ac156CE4A0AF/easypark)

Easpark is a REST API built with clean architecture principles. This backend ensures scalability, maintainability, and flexibility, allowing for integration with our frontend and other services services.

## API Spec

### 1. User Login API Endpoint

**Endpoint**: `POST /login`

**Description**: Authenticates a user and returns a JWT for authorised access.

**Request**:

```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
    "Username": "johndoe",
    "Password": "securepassword"
}'
```

**Responses**:

- **200 OK**

    ```json
    {
      "message": "User logged in successfully",
      "token": "jwt token here redacted"
    }
    ```

- **400 Bad Request**

    ```json
    {
      "error": "invalid request body"
    }
    ```

- **401 Unauthorized**
  
    ```json
    {
      "error": "Invalid credentials" // or "User not found"
    }
    ```

- **500 Internal Server Error**

    ```json
    {
      "error": "An unexpected error occurred"
    }
    ```

### 2. User Creation API Endpoint

**Endpoint**: `POST /register`

**Description**: Registers a new user in the system.

**Request Body**:

```bash
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{
    "Username": "johndoe",
    "Email": "john.doe@example.com",
    "Password": "securepassword",
    "FirstName": "John",
    "LastName": "Doe"
}'
```

**Responses**:

- **201 Created**

    ```json
    {
      "message": "user created successfully"
    }
    ```

- **400 Bad Request**

    ```json
    {
      "error": "User already exists" // or "invalid request body"
    }
    ```

- **500 Internal Server Error**

    ```json
    {
      "error": "An unexpected error occurred"
    }
    ```

## Running locally

### Prerequisites

- Linux environment
- VS Code
- Docker
- Golang (LTS version)

### Setting up environment

From project root, run:

1. Run `docker-compose up -d` to create required PostgreSQL image and optional PgAdmin image for DB user interface.

### Starting the app

From project root, run:

1. Build the app, run `make build`.
2. To run the app, run `make run`.

If changes to dependecy graph have been made, you must edit `wire.go` file and run `make wire` to regenerate dependecy injection code(`wire_gen.go` file).

### Troubleshooting

Will be edited once problems appear.

## Testing

To regenerate mocks for unit tests, run `make mocks`.

To run unit tests, run `make unit`.

To run functional tests, run `make functional`.

## Useful things

### Cleaning database tables

Run `make clean-db` to truncate existing tables. Note, you need to add new table names to `./build/clean-bd.sh` script.

### Creating admin user

Run the follwoing psql command in the database container:

```bash
psql -U devUser -d easypark -c "INSERT INTO users (id, username, email, password, first_name, last_name, role) VALUES ('a131a9a0-8d09-4166-b6fc-f8a08ba549e9', 'adminUsername', 'admin@example.com', 'securePassword', 'Admin', 'User', 'admin');"
```

To get JWT for this admin, run:

```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
    "Username": "adminUsername",
    "Password": "securePassword"
}'
```

To access admin routes:

```bash
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOiJhMTMxYTlhMC04ZDA5LTQxNjYtYjZmYy1mOGEwOGJhNTQ5ZTkiLCJ1c2VybmFtZSI6ImFkbWluVXNlcm5hbWUiLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE3MTA1Mzk4NjR9.7qjL67k50MUtqEzZ2IHm5-TdRvcWaQRiILUDwVbv7XI" http://localhost:8080/admin
```

### Connecting to PgAdmin

PgAdming provides a nice UI for DB management and debugging.

1. Go to `http://localhost:5050` to access PgAdmin
2. Log in with the `PGADMIN_DEFAULT_EMAIL` and `PGADMIN_DEFAULT_PASSWORD` in the docker-compose.yml file
3. To connect to our PostgreSQL database from PgAdmin:
    - Right-click on "Servers" in the left panel and choose "Create > Server".
    - In the "Create Server" dialog, go to the "Connection" tab.
    - Set "Hostname/address" to `database`, which is the name of our PostgreSQL service defined in our docker-compose.yml.
    - Fill in the "Username" and "Password" fields with the POSTGRES_USER and POSTGRES_PASSWORD specified in docker-compose.yml.
    - Click "Save" to establish the connection.
