# easypark

[Trello board](https://trello.com/invite/b/lGdfavnm/ATTI15a8afbd1ced04b229e8f2380279ac156CE4A0AF/easypark)

## Running locally

### Prerequisites

- Docker
- LTS Go version

### Setting up environment

1. Run `docker-compose up -d` to create required PostgreSQL image.
2. Run `make setup-db` to truncate existing tables and start fresh.

### Starting the app

1. Build the app, run `make build`.
2. To run the app, run `make run`.

## Testing

To generate mocks, run `make mocks`.

To run unit tests, run `make unit`.

## App Spec

1. Create Driver:

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

## Useful things

### Connecting to PgAdmin

PgAdming provides a nice UI for db management and debugging.

1. Go to `http://localhost:5050` to access PgAdmin
2. Log in with the `PGADMIN_DEFAULT_EMAIL` and `PGADMIN_DEFAULT_PASSWORD` in the docker-compose.yml file
3. To connect to our PostgreSQL database from PgAdmin:
    - Right-click on "Servers" in the left panel and choose "Create > Server".
    - In the "Create Server" dialog, go to the "Connection" tab.
    - Set "Hostname/address" to `database`, which is the name of our PostgreSQL service defined in our docker-compose.yml.
    - Fill in the "Username" and "Password" fields with the POSTGRES_USER and POSTGRES_PASSWORD specified in docker-compose.yml.
    - Click "Save" to establish the connection.
