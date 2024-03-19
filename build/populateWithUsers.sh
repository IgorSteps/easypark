#!/bin/sh

curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{
    "Username": "user1",
    "Email": "user1@example.com",
    "Password": "securepassword",
    "FirstName": "test",
    "LastName": "user1"
}'

curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{
    "Username": "user2",
    "Email": "user2@example.com",
    "Password": "securepassword",
    "FirstName": "test",
    "LastName": "user2"
}'

curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{
    "Username": "user3",
    "Email": "user3@example.com",
    "Password": "securepassword",
    "FirstName": "test",
    "LastName": "user3"
}'