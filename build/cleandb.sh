#!/bin/bash

# DB credentials, same as in docker-compose.
DB_USER="devUser"
DB_PASSWORD="devPassword"
DB_NAME="easypark"
DB_HOST="localhost"
DB_PORT="5432"

TABLES="users"

for TABLE in $TABLES; do
    PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -c "TRUNCATE TABLE $TABLE;"
done

echo "Tables truncated successfully."