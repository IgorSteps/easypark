# EasyPark

[Task tracking](https://trello.com/invite/b/lGdfavnm/ATTI15a8afbd1ced04b229e8f2380279ac156CE4A0AF/easypark)

EasyPark is a backend built with [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) principles. It provides REST API endpoints for drivers and admins to be able to use and manage parking through our [UI](https://github.com/IgorSteps/easypark-ui).

## API Specification and Design

Before reading the API Specification, read our [design docs](./docs/DESIGN.MD) to understand the states, rules and assumptions regarding entities within the EasyPark system.

API Specification can be found [here](docs/API_SPEC.md).

## Running locally

### Prerequisites

#### Mac

- Linux environment(normal terminal on MacOS)
- [Docker Desktop for Mac](https://docs.docker.com/desktop/install/mac-install/)
- [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers).

#### Windows

- Linux environment([WSL2](https://learn.microsoft.com/en-us/windows/wsl/install))
- [Docker Desktop for Windows](https://docs.docker.com/desktop/install/windows-install/) (Download for Windows and [enable integration with WSL2 in the settings](https://docs.docker.com/desktop/wsl/))
- [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers).

### Setting up dev environment

Git clone to your Linux environment using `git clone https://github.com/IgorSteps/easypark.git`.

Run `make dev-env` from project root to bring up the `igorsteps/go-dev:latest` development container with all required tooling. Subsequently you can run `make dev-env-down` to bring it down.

Attach to it using VS Code's [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers).

### Starting the app

From project root, run:

1. Build the app, run `make build`.
2. To run the app, run `make run`.

If changes to dependency graph have been made, you must edit `wire.go` file and run `make wire` to regenerate dependency injection code(`wire_gen.go` file).

### Troubleshooting

TODO

## Testing

- To regenerate mocks for unit tests, run `make mocks`.

- To run unit tests, run `make unit`.
  - To see unit test coverage, run `make coverage-report`. This will output a `unit-test-coverage.out` report that can be viewed in the browser using `make coverage-interactive`.

- To run functional tests, run `make functional`.

## Useful things

### CI pipeline (GitHub Actions)

For every commit you make to your PRs, a GitHub [actions workflow](.github/workflows/go.yml) will get triggered that automatically runs tasks such as:

- Building and running the app.
- Unit and functionally testing the app.
- Testing unit test coverage is above 70%.

If any of the steps fail, the checks will fail and you will not be able to merge your PR unti you fix the issue.

### Config Management

Config values are specified in [here](./config.yaml). If edit it to include new key-value pairs, you must mirror that as respective struct fields in [here](/internal/drivers/config/config.go). Note that the keys in the yaml must match struct field names.

### Cleaning database tables

Run `make clean-db` to truncate existing tables. Note, you need to add new table names to [script](./build/cleandb.sh).

### Creating admin user

Run the `make create-admin` runs this [script](./build/createadmin.sh) to create an admin directly in the database and give you a JWT back.

This creates an admin with the following details:

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

### Connecting to PgAdmin

PgAdming provides a nice UI for DB management and debugging.

1. Go to `http://localhost:5050` to access PgAdmin
2. Log in with the `PGADMIN_DEFAULT_EMAIL` and `PGADMIN_DEFAULT_PASSWORD` in the [docker-compose.yml](./docker-compose.yml) file
3. To connect to our PostgreSQL database from PgAdmin:
    - Right-click on "Servers" in the left panel and choose "Create > Server".
    - In the "Create Server" dialog, go to the "Connection" tab.
    - Set "Hostname/address" to `database`, which is the name of our PostgreSQL service defined in our [docker-compose.yml](./docker-compose.yml).
    - Fill in the "Username" and "Password" fields with the POSTGRES_USER and POSTGRES_PASSWORD specified in [docker-compose.yml](./docker-compose.yml).
    - Click "Save" to establish the connection.
