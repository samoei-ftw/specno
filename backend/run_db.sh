#!/bin/bash
# Starts PostgreSQL in a local Docker container and runs migrations
# Author: Samoei Oloo
# Created: 2025-03-28
#
# This script sets up a PostgreSQL container and applies migrations
# using SQL files located in the database directory.

if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

docker network create specno-network || true

if [ "$(docker ps -aq -f name=$DB_CONTAINER_NAME)" ]; then
    docker stop $DB_CONTAINER_NAME && docker rm $DB_CONTAINER_NAME
fi

docker run -d \
    --name $DB_CONTAINER_NAME \
    --network specno-network \
    -e POSTGRES_USER=$DB_USER \
    -e POSTGRES_PASSWORD=$DB_PASSWORD \
    -e POSTGRES_DB=$DB_NAME \
    -p $DB_PORT:5432 \
    -v $DB_VOLUME:/var/lib/postgresql/data \
    postgres:latest

until docker exec -i $DB_CONTAINER_NAME pg_isready -U $DB_USER >/dev/null 2>&1; do
    sleep 2
done

echo "Performing migrations..."
for file in $DB_MIGRATIONS_DIR/*.sql; do
    echo "Applying migration: $file"
    docker exec -i $DB_CONTAINER_NAME psql -U $DB_USER -d $DB_NAME < "$file"
done

echo "All migrations applied successfully."

while true; do
  sleep 60
done