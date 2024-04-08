#!/bin/sh

DB_USER="devUser"
DB_PASSWORD="devPassword"
DB_NAME="easypark"
DB_HOST="localhost"
DB_PORT="5432"


PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -c "INSERT INTO users (id, username, email, password, first_name, last_name, role) VALUES ('a131a9a0-8d09-4166-b6fc-f8a08ba549e9', 'adminUsername', 'admin@example.com', 'securePassword', 'Admin', 'User', 'admin');"

curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
    "Username": "adminUsername",
    "Password": "securePassword"
}'
