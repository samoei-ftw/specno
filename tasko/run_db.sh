#!/bin/bash
# Starts PostgreSQL in a local Docker container and runs migrations
# Author: Samoei Oloo
# Created: 2025-03-28
#
# This script sets up a PostgreSQL container and applies migrations
# using SQL files located in the database directory.

# Env vars
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

# Create a custom network if it doesn't exist
docker network create specno-network || true

# Stop and remove any existing PostgreSQL container
if [ "$(docker ps -aq -f name=$DB_CONTAINER_NAME)" ]; then
    echo "Stopping and removing existing PostgreSQL container..."
    docker stop $DB_CONTAINER_NAME && docker rm $DB_CONTAINER_NAME
fi

# Start the PostgreSQL container on the custom network
echo "Starting PostgreSQL container..."
docker run -d \
    --name $DB_CONTAINER_NAME \
    --network specno-network \
    -e POSTGRES_USER=$DB_USER \
    -e POSTGRES_PASSWORD=$DB_PASSWORD \
    -e POSTGRES_DB=$DB_NAME \
    -p $DB_PORT:5432 \
    -v $DB_VOLUME:/var/lib/postgresql/data \
    postgres:latest

# Wait for PostgreSQL to be ready
echo "Waiting for PostgreSQL to be ready..."
until docker exec -i $DB_CONTAINER_NAME pg_isready -U $DB_USER >/dev/null 2>&1; do
    sleep 2
done

# Run migrations on PostgreSQL
echo "PostgreSQL is ready. Running migrations..."
for file in $DB_MIGRATIONS_DIR/*.sql; do
    echo "Applying migration: $file"
    docker exec -i $DB_CONTAINER_NAME psql -U $DB_USER -d $DB_NAME < "$file"
done

echo "All migrations applied successfully."