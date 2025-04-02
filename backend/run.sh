#!/bin/sh
# Starts PostgreSQL in a local Docker container and runs migrations
# Author: Samoei Oloo
# Created: 2025-03-28
#
# This script sets up a docker container to run the backend project locally

CONTAINER_ONE="user-api"
CONTAINER_TWO="specno-db"
cleanup() {
  if [ "$(docker ps -q -f name=$CONTAINER_ONE)" ]; then
    docker stop $CONTAINER_ONE
  fi
  if [ "$(docker ps -q -f name=$CONTAINER_TWO)" ]; then
    docker stop $CONTAINER_TWO
  fi
  docker rm $CONTAINER_ONE
  docker rm $CONTAINER_TWO
  docker volume rm specno-db || true
  docker network rm specno-network || true
}

trap cleanup INT

docker network inspect specno-network > /dev/null 2>&1 || docker network create specno-network

echo "Removing existing containers..."
docker rm -f user-api || true
docker rm -f specno-postgres || true
docker rmi -f user-api || true

docker exec -it specno-db psql -U postgres -d postgres -c "
DO \$\$ DECLARE
    r RECORD;
BEGIN
    FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
        EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.tablename) || ' CASCADE';
    END LOOP;
END \$\$;"

echo "Building Docker image for user-api..."
docker build -t user-api -f ./user-service/Dockerfile .

echo "Starting user container..."
docker run -d \
  -p 8080:8080 \
  --name specno-user-service \
  --network specno-network \
  -v $(pwd)/.env:/app/.env \
  user-api

timeout=1
start_time=$(date +%s)

while true; do
  current_time=$(date +%s)
  elapsed_time=$((current_time - start_time))
  
  if [ "$elapsed_time" -ge "$timeout" ]; then
    echo "Exiting..."
    break
  fi
  # Check if the user-api service is up and running
  if docker logs $CONTAINER_ONE 2>&1 | grep -q "Server started"; then
    echo "User API service is up and running."
    break
  fi

  sleep 1
done