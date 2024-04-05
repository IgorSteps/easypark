#!/bin/sh

DB_USER="devUser"
DB_PASSWORD="devPassword"
DB_NAME="easypark"
DB_HOST="localhost"
DB_PORT="5432"

# Insert a Parking Lot
PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -c "INSERT INTO parking_lots (id, name, location, capacity) VALUES ('06876ca4-69c2-46e8-b387-43b84c013a98', 'Lot A', '123 Main St', 100)"

# Insert a Parking Space
PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -c  "INSERT INTO parking_spaces (id, parking_lot_id, name, status, reserved_for, occupied_at, user_id) VALUES ('a4e23008-6ce6-44b1-9d05-de26d1fb67be', '06876ca4-69c2-46e8-b387-43b84c013a98', 'Space 1', 'available', NULL, NULL, NULL);"
