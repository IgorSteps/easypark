#!/bin/bash

docker exec -it database sh -c "psql -U devUser -d easypark -c \"INSERT INTO users (id, username, email, password, first_name, last_name, role) VALUES ('a131a9a0-8d09-4166-b6fc-f8a08ba549e9', 'adminUsername', 'admin@example.com', 'securePassword', 'Admin', 'User', 'admin');\""