# easypark

[Trello board](https://trello.com/invite/b/lGdfavnm/ATTI15a8afbd1ced04b229e8f2380279ac156CE4A0AF/easypark)

Easpark is a REST API built with Clean Architecture. It provides a set of endpoints for drivers and admins to use and manage parking.

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

### 3. Get All Drivers API Endpoint

**Endpoint**: `GET /drivers`

**Description**: Gets all driver users in the system.

**Request Body**:

```bash
curl -H "Authorization: Bearer <ADMIN_TOKEN> http://localhost:8080/drivers
```

**Responses**:

- **200 OK**

    ```json
    [
      {
        "ID":"910e78c8-d2eb-41e8-aec6-c70a33b692df",
        "Username":"user1",
        "Email":"user1@example.com",
        "Password":"securepassword",
        "FirstName":"test",
        "LastName":"user1",
        "Role":"driver",
        "Status":"active",
      },
      {"other users..."}
    ]
    ```

- **500 Internal Server Error**

    ```json
    {
      "error": "An unexpected error occurred"
    }
    ```

### 4. Update Driver Status API Endpoint

**Endpoint**: `PATCH /drivers/{id}/status`

**Description**: Updates driver status in the system.

**Request Body**:

- Banning:
  
  ```bash
  curl -X PATCH http://localhost:8080/drivers/{id}/status \
      -H "Content-Type: application/json" \
      -H "Authorization: Bearer <ADMIN_TOKEN>" \
      -d '{"status":"ban"}'
  ```

- Unbanning: TODO

**Responses**:

- **200 OK***

  ```json
  {
    "message":"successfully updated user status"
  }
  ```

### 5. Create Parking Request API Endpoint

**Endpoint**: `POST /drivers/{id}/parking-requests`

**Description**: Creates a parkign request for the driver.

**Request Body**:

Ensure that the JSON sent in the REST request uses the RFC 3339 date/time format for StartTime and EndTime

```bash
curl -X POST http://localhost:8080/drivers/{id}/parking-requests \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <DRIVER_TOKEN>" \
-d '{
    "destination": "science",
    "startTime": "2024-05-01T09:00:00Z",
    "endTime": "2024-05-01T17:00:00Z"
}'
```

**Responses**:

- **201 CREATED**

  ```json
  {
    "destination": "science",
    "startTime": "2024-05-01T09:00:00Z",
    "endTime": "2024-05-01T17:00:00Z",
    "status": "pending"
  }
  ```

- **400 BAD REQUEST**

  ```json
  {
   "start time cannot be after the end time"
  }
  ```

- **500 INTERNAL SERVER ERROR**

  ```json
  {
    "Intenal error: some error message"
  }
  ```

### 6. Update Parking Request Status API Endpoint

**Endpoint**: `PATCH /parking-requests/{id}/status`

**Description**: Updated a parking request status.

**Request Body**:

Ensure that the status is one of the following: `approved, rejected, pending`.

```bash
curl -X PATCH http://localhost:8080/parking-requests/{id}/status \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <ADMIN_TOKEN>" \
-d '{
    "status": "approved"
}'
```

**Responses**:

- **200 OK**

  ```json
  {
    "message": "successfully updated the parking request status"
  }
  ```

- **400 BAD REQUEST**

  ```json
  {
   "meaningful message"
  }
  ```

- **500 INTERNAL SERVER ERROR**

  ```json
  {
    "Intenal error: some error message"
  }
  ```

### 7. Update Parking Request Space API Endpoint

**Endpoint**: `PATCH /parking-requests/{id}/space`

**Description**: Update a parking request with a parking space.

**Request Body**:

```bash
curl -X PATCH http://localhost:8080/parking-requests/{id}/space \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <ADMIN_TOKEN>" \
-d '{
    "parkingSpaceID": "desired parking space id"
}'
```

### 8. Create Parking Lot API Endpoint

**Endpoint**: `POST /parking-lots`

**Description**: Create a parking lot.

**Request Body**:

```bash
curl -X POST http://localhost:8080/parking-lots \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <ADMIN_TOKEN>" \
-d '{
    "name": "boom",
    "capacity": 10
}'
```

**Responses**:

- **201 CREATED**

```json
  {
    "id":"83601bb8-9ad1-45a8-a3f4-21dd219fd054",
    "name":"science",
    "capacity":10,
    "parkingSpaces":[
        {
          "ID":"a2856678-3bfb-4ce2-a041-b5d0048d8993","ParkingLotID":"83601bb8-9ad1-45a8-a3f4-21dd219fd054","Name":"science-1",
          "Status":"available",
          "FreeAt":"0001-01-01T00:00:00Z","OccupiedAt":"0001-01-01T00:00:00Z",
          "UserID":null,
          "ParkingRequests":null
        },
        {"other spaces..."}
      ]
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

Run `make clean-db` to truncate existing tables. Note, you need to add new table names to `./build/clean-db.sh` script.

### Creating admin user

Run the `./build/createadmin.sh` script to create a user with admin role in the database. This creates an admin with the following details:

```json
{
  "id":"a131a9a0-8d09-4166-b6fc-f8a08ba549e9",
  "username":"adminUsername",
  "email":"admin@example.com", 
  "password":"securePassword",
  "firstname":"Admin",
  "lastname":"User",
  "role":"admin"
}
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

For example, as an admin, you can curl `drivers` endpoint to get all drivers:

```bash
curl -H "Authorization: Bearer <ADMIN_TOKEN>" http://localhost:8080/drivers
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
