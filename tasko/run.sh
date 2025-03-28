#!/bin/sh

cd cmd/user-service
go run main.go &
echo "User service started"

cd ../project-service
go run main.go &
echo "Project service started"

cd ../task-service
go run main.go &
echo "Task service started"

wait