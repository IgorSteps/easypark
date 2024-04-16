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
      "error": "Invalid credentials"
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
      "error": "User already exists"
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

**Description**: Updates driver status in the system(can only be `ban` status for now).

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

Ensure that the JSON sent in the REST request uses the RFC 3339 date/time format for StartTime and EndTime. And destination being a desired ParkingLot ID.

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
    "destination": "d56541a1-335b-49cb-84bb-672009ecf580",
    "startTime": "2024-05-01T09:00:00Z",
    "endTime": "2024-05-01T17:00:00Z",
    "status": "pending"
  }
  ```

- **400 BAD REQUEST**

  ```json
  {"error": "meaningful error message"}
  ```

- **500 INTERNAL SERVER ERROR**

  ```json
  {"error": "meaningful error message"}
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
  {"error": "meaningful error message"}
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
  {"error": "meaningful error message"}
  ```

### 4. Update Parking Request Status API Endpoint

**Endpoint**: `PATCH /parking-requests/{id}/status`

**Description**: Updated a parking request status to `rejected` or `pending`.

**Request Body**:

Ensure that the status is one of the following: `rejected` or `pending`.

```bash
curl -X PATCH http://localhost:8080/parking-requests/{id}/status \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <ADMIN_TOKEN>" \
-d '{
    "status": "rejected"
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
  {"error": "meaningful error message"}
  ```

- **500 INTERNAL SERVER ERROR**

  ```json
  {"error": "meaningful error message"}
  ```

### 5. Assign Parking Request a Space API Endpoint

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
  {"error": "meaningful error message"}
  ```

- **500 INTERNAL SERVER**

  ```json
  {"error": "meaningful error message"}
  ```

## Parking Lot

### 1. Create Parking Lot API Endpoint

**Endpoint**: `POST /parking-lots`

**Description**: Create a parking lot, and its parking spaces.

**Request Body**:

```bash
curl -X POST http://localhost:8080/parking-lots \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <ADMIN_TOKEN>" \
-d '{
    "name": "cmp",
    "capacity": 10
}'
```

**Responses**:

- **201 CREATED**

```json
  {
    "id":"83601bb8-9ad1-45a8-a3f4-21dd219fd054",
    "name":"cmp",
    "capacity":10,
    "parkingSpaces":[
        {
          "ID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b",
          "ParkingLotID":"83601bb8-9ad1-45a8-a3f4-21dd219fd054",
          "Name":"cmp-1",
          "Status":"blocked",
          "ParkingRequests":[{"approved parking requests assigned to this parking space"}]
        },
        {"other spaces..."}
      ]
  }
```

- **400 BAD REQUEST**

  ```json
  {"error": "meaningful error message"}
  ```

- **500 INTERNAL SERVER ERROR**

  ```json
  {"error": "meaningful error message"}
  ```

### 2. Get All Parking Lots API Endpoint

**Endpoint**: `GET /parking-lots`

**Description**: Gets all parking lots and their statistics(total, occupied, available, blocked and reserved spaces).

**Request Body**:

```bash
curl -H "Authorization: Bearer <ADMIN_TOKEN>" http://localhost:8080/parking-lots
```

**Response**:

- **200 OK**

  ```json
    [
      {
        "ID":"bb8625ea-8c80-484c-8a75-3386649eef25",
        "Name":"cmp",
        "Capacity":10,
        "ParkingSpaces":[
          {
            "ID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b",
            "ParkingLotID":"bb8625ea-8c80-484c-8a75-3386649eef25",
            "Name":"cmp-1",
            "Status":"blocked",
            "ParkingRequests":[{"approved parking requests assigned to this parking space"}]
          },
          {"other parking spaces..."}
        ],
        "Available":0,
        "Occupied":0,
        "Reserved":0,
        "Blocked":0
      },
      {"other parking lots..."}
    ]
  ```

- **500 INTERNAL SERVER ERROR**
  
  ```json
  {"error": "meaningful error message"}
  ```

### 3. Delete a Parking Lot API Endpoint

**Endpoint**: `DELETE /parking-lots/{id}`

**Description**: Deletes a parking lots with given ID, as well as its parking spaces and parking requests referencing these spaces.

**Request Body**:

```bash
curl -X DELETE http://localhost:8080/parking-lots/{id} \
-H "Authorization: Bearer <ADMIN_TOKEN>"
```

**Response**:

- **200 OK**

  ```json
  {"successfully deleted parking lot"}
  ```

- **400 BAD REQUEST**

  ```json
  {"error": "meaningful error message"}
  ```

- **500 INTERNAL SERVER ERROR**

  ```json
  {"error": "meaningful error message"}
  ```

## Parking Space

### 1. Get Single Parking Space API Endpoint

**Endpoint**: `GET /parking-spaces/{id}`

**Description**: Gets a parking space status with the given ID.

**Request Body**:

```bash
curl -H "Authorization: Bearer <USER_TOKEN>"  http://localhost:8080/parking-spaces/{id}
```

**Response**:

- **200 OK**

  ```json
  {
    "ID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b",
    "ParkingLotID":"6404407e-6729-4a7b-9f5e-22059233a030",
    "Name":"cmp-1",
    "Status":"blocked",
    "ParkingRequests":[{"approved parking requests assigned to this parking space"}]
  }
  ```

  - **400 BAD REQUEST**

  ```json
  {"error": "meaningful error message"}
  ```

- **500 INTERNAL SERVER ERROR**

  ```json
  {"error": "meaningful error message"}
  ```

### 2. Update Parking Space Status API Endpoint

**Endpoint**: `PATCH /parking-spaces/{id}/status`

**Description**: Updates a parking space status with the given ID.

**Request Body**:

Ensure that the status is one of the following: `available, occupied, blocked, reserved`.

```bash
curl -X PATCH http://localhost:8080/parking-spaces/{id}/status \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <ADMIN_TOKEN>" \
-d '{
    "status": "blocked"
}'
```

**Response**:

- **200 OK**

  ```json
  {
    "ID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b",
    "ParkingLotID":"6404407e-6729-4a7b-9f5e-22059233a030",
    "Name":"cmp-1",
    "Status":"blocked",
    "ParkingRequests":[{"approved parking requests assigned to this parking space"}]
  }
  ```

- **400 BAD REQUEST**

  ```json
  {"error": "meaningful error message"}
  ```

- **500 INTERNAL SERVER ERROR**

  ```json
  {"error": "meaningful error message"}
  ```

## Notification

### 1. Create Notification API Endpoint

**Endpoint**: `POST /drivers/{id}/notification`

**Description**: Creates a notification for a driver with the given id. Notification type is `Arrival = 0` and `Departure = 1`.

**Request Body**:

```bash
curl -X POST http://localhost:8080/drivers/{id}/notifications \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <DRIVER_TOKEN>" \
-d '{
    "requestID": "parking request ID to which the parking space is allocated",
    "parkingSpaceID": "allocated parking space id",
    "location": "cmp-1",
    "notificationType": 0
}'
```

**Responses**:

- **200 OK**

  ```json
  {
    "ID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b",
    "Type": 0,
    "DriverID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b",
    "ParkingSpaceID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b",
    "Location":"cmp-1",
    "Timestamp":"0000-12-31T23:58:45-00:01",
  }
  ```

- **400 BAD REQUEST**

    ```json
  {"error": "meaningful error message"}
  ```

- **500 INTERNAL SERVER ERROR**

    ```json
  {"error": "meaningful error message"}
  ```

### 2. Get All Notifications API Endpoint

**Endpoint**: `GET /notifications`

**Description**: Gets all notifications.

**Request Body**:

```bash
curl -H "Authorization: Bearer <ADMIN_TOKEN>"  http://localhost:8080/notifications
```

**RESPONSES**:

- **200 OK**

 ```json
  [ 
    {
      "ID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b",
      "Type": 0,
      "DriverID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b",
      "ParkingSpaceID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b",
      "Location":"cmp-1",
      "Timestamp":"0000-12-31T23:58:45-00:01",
    },
    {"other notifications"}
  ]
 ```

- **500 INTERNAL SERVER ERROR**

    ```json
  {"error": "meaningful error message"}
  ```
  
## Alert

### 1. Get Single Alert API Endpoint

**Endpoint**: `GET /alerts/{id}`

**Description**: Gets single alert using its id.

**Request Body**:

```bash
curl -H "Authorization: Bearer <ADMIN_TOKEN>"  http://localhost:8080/alerts/{id}
```

**RESPONSES**:

Returns an error of one of these types: `0 - Location mismatch alert`.

- **200 OK**

  ```json
  {
      "ID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b",
      "Type": 0,
      "Message": "some message",
      "UserID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b",
      "ParkingSpaceID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b"
  },
  ```

- **400 BAD REQUEST**

  ```json
  {"error": "meaningful error message"}
  ```

- **500 INTERNAL SERVER**

  ```json
  {"error": "meaningful error message"}
  ```

### 2. Check Late Arrivals API Endpoint

**Endpoint**: `POST /alerts/late-arrivals`

**Description**: Runs a workflow to check if arrival notifications have not been received withing given threshold. Returns created alerts as result of the check. Note, that this check is performed automatically by the [Scheduler](./DESIGN.MD#alerts) at constant intervals. This endpoint is if the admin wants to do it manually.

**Request Body**:

```bash
curl -X POST http://localhost:8080/alerts/late-arrivals \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <ADMIN_TOKEN>" \
-d '{
    "threshold": "1h"
}'
```

**Responses**:

- **200 OK**

  ```json
    [
      {
      "ID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b",
      "Type": 0,
      "Message": "some message",
      "UserID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b",
      "ParkingSpaceID":"a678f5a6-9731-4741-ad0b-de5efbbffc9b"
      },
      {"other alerts..."}
    ]
  ```
- **400 BAD REQUEST**

  ```json
  {"error": "meaningful error message"}
  ```

- **500 INTERNAL SERVER**

  ```json
  {"error": "meaningful error message"}
  ```

