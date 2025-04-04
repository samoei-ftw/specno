#!/bin/sh
# Starts PostgreSQL in a local Docker container and runs migrations
# Author: Samoei Oloo
# Created: 2025-03-28

echo "Stopping all containers and removing volumes..."
docker-compose down -v

echo "Rebuilding and starting containers with docker-compose..."
docker-compose up --build -d

echo "Waiting for PostgreSQL to be ready..."
until docker-compose exec -T tasko-database pg_isready -U postgres > /dev/null 2>&1; do
  sleep 2
done