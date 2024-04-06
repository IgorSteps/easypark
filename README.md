# EasyPark

[Trello board](https://trello.com/invite/b/lGdfavnm/ATTI15a8afbd1ced04b229e8f2380279ac156CE4A0AF/easypark)

Easpark is a REST API built with Clean Architecture. It provides a set of endpoints for drivers and admins to use and manage parking.

## API Specification

API Specification can be found [here](docs/API_SPEC.md).

## Running locally

### Prerequisites

#### Mac

- Linux environment(normal terminal on MacOS)
- VS Code with:
  - [GoLang extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- [Docker Desktop for Mac](https://docs.docker.com/desktop/install/mac-install/)
- [GoLang](https://go.dev/doc/install) (Follow instructions for Mac)

#### Windows

- VS Code with:
  - [GoLang extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
  - [Remote extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) (need to connect VS code to WSL2)
- Linux environment([WSL2](https://learn.microsoft.com/en-us/windows/wsl/install))
- [Docker Desktop for Windows](https://docs.docker.com/desktop/install/windows-install/) (Download for Windows and [enable integration with WSL2 in the settings](https://docs.docker.com/desktop/wsl/))
- [GoLang](https://go.dev/doc/install) (Follow instructions for Linux and install in your WSL2)

### Setting up environment

Git clone to your Linux environment using `git clone https://github.com/IgorSteps/easypark.git`.

Open the project in VS Code:

- On Mac: just open it like you would any project.
- On Windows, use your VS Code Remote Extension to connect to your WSL2 and locate your cloned project there.

From project root in your Linux Environment, run:

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
