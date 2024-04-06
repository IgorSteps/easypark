# API Specification

This document provides an overview of available API endpoints with example request and response bodies.

Endpoints are split into sections for each domain entity they operate on.

## User

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
      "ID":"3149cc94-3f9e-49f3-8dc8-516cdd03852d",
      "Username":"johndoe",
      "Email":"john.doe@example.com",
      "Password":"securepassword",
      "FirstName":"John",
      "LastName":"Doe",
      "Status":"active",
      "Role":"driver"
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

**Responses**:

- **200 OK***

  ```json
  {
    "message":"successfully updated user status"
  }
  ```

## Parking request

### 1. Create Parking Request API Endpoint

**Endpoint**: `POST /drivers/{id}/parking-requests`

**Description**: Creates a parking request for the driver with status 'pending'.

**Request Body**:

Ensure that the JSON sent in the REST request uses the RFC 3339 date/time format for StartTime and EndTime

```bash
curl -X POST http://localhost:8080/drivers/{id}/parking-requests \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <DRIVER_TOKEN>" \
-d '{
    "destination": "d56541a1-335b-49cb-84bb-672009ecf580",
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
    "meaningful error message"
  }
  ```

- **500 INTERNAL SERVER ERROR**

  ```json
  {
    "meaningful error message"
  }
  ```

### 2. Get All Parking Request API Endpoint

**Endpoint**: `GET /parking-requests`

**Description**: Gets all parking requests.

**Request Body**:

```bash
curl -H "Authorization: Bearer <ADMIN_TOKEN>"  http://localhost:8080/parking-requests
```

**Responses**:

- **200 OK**

```json
[
  {
    "ID":"8275277b-0ec2-4cbf-9129-79a76194fe2e",
    "UserID":"413662b8-0214-4935-a022-175438e6c4f1",
    "ParkingSpaceID":null,
    "DestinationParkingLotID":"714a2875-d358-423b-83b2-72a701a82492",
    "StartTime":"2024-04-06T16:59:09.441792+01:00",
    "EndTime":"2024-04-06T16:59:09.441792+01:00",
    "Status":"pending"
  },
  {
    "ID":"e071153e-9f5b-409e-8451-457e64dba8a2",
    "UserID":"413662b8-0214-4935-a022-175438e6c4f1",
    "ParkingSpaceID":null,
    "DestinationParkingLotID":"2fa0a013-9d29-451c-9b04-0e65c5c82990",
    "StartTime":"2024-04-06T16:59:09.443992",
    "EndTime":"2024-04-06T16:59:09.441792+01:00",
    "Status":"pending"
  },
  {"others requests..."}
]
```

- **500 INTERNAL SERVER ERROR**

```json
  {
   "meaningful error message"
  }
```

### 3. Get All Parking Requests for Driver API Endpoint

**Endpoint**: `GET /drivers/{id}/parking-requests`

**Description**: Gets all parking requests for a particular driver.

**Request Body**:

```bash
curl -H "Authorization: Bearer <DRIVER_TOKEN>"  http://localhost:8080/drivers/{id}/parking-requests
```

**Response**:

- **200 OK**

  ```json
    [
      {
        "ID":"8275277b-0ec2-4cbf-9129-79a76194fe2e",
        "UserID":"drivers id",
        "ParkingSpaceID":null,
        "DestinationParkingLotID":"714a2875-d358-423b-83b2-72a701a82492",
        "StartTime":"2024-04-06T16:59:09.441792+01:00",
        "EndTime":"2024-04-06T16:59:09.441792+01:00",
        "Status":"pending"
      },
      {
        "ID":"e071153e-9f5b-409e-8451-457e64dba8a2",
        "UserID":"drivers id",
        "ParkingSpaceID":null,
        "DestinationParkingLotID":"2fa0a013-9d29-451c-9b04-0e65c5c82990",
        "StartTime":"2024-04-06T16:59:09.443992",
        "EndTime":"2024-04-06T16:59:09.441792+01:00",
        "Status":"pending"
      },
      {"their others requests..."}
    ]
  ```

- **500 INTERNAL SERVER ERROR**

```json
  {
    "meaningful error message"
  }
```

### 4. Update Parking Request Status API Endpoint

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
   "meaningful error message"
  }
  ```

- **500 INTERNAL SERVER ERROR**

  ```json
  {
   "meaningful error message"
  }
  ```

### 5. Update Parking Request Space API Endpoint

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

**Responses**:

- **400 BAD REQUEST**

  ```json
  {
   "meaningful error message"
  }
  ```

- **500 INTERNAL SERVER**

  ```json
  {
   "meaningful error message"
  }
  ```

- **400 BAD REQUEST**

  ```json
  {
   "start time cannot be after the end time"
  }
  ```

## Parking Lot

### 1. Create Parking Lot API Endpoint

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
