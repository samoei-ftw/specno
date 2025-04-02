#!/bin/sh
# Starts PostgreSQL in a local Docker container and runs migrations
# Author: Samoei Oloo
# Created: 2025-03-28
#
# This script sets up a docker container to run the backend project locally

CONTAINER_USER="specno-user-service"
CONTAINER_PROJECT="specno-project-service"
CONTAINER_DB="specno-db"

cleanup() {
  #echo "Stopping and removing existing containers..."
  if [ "$(docker ps -q -f name=$CONTAINER_USER)" ]; then
    docker stop $CONTAINER_USER
  fi
  if [ "$(docker ps -q -f name=$CONTAINER_PROJECT)" ]; then
    docker stop $CONTAINER_PROJECT
  fi
  if [ "$(docker ps -q -f name=$CONTAINER_DB)" ]; then
    docker stop $CONTAINER_DB
  fi

  #docker rm -f $CONTAINER_USER $CONTAINER_PROJECT $CONTAINER_DB || true
  #docker volume rm specno-db || true
  #docker network rm specno-network || true
}

trap cleanup INT

docker network inspect specno-network > /dev/null 2>&1 || docker network create specno-network

echo "Removing existing containers..."
docker rm -f $CONTAINER_DB $CONTAINER_PROJECT || true
docker rmi -f user-api project-api || true

echo "Building Docker images..."
docker build -t user-api -f ./Dockerfile .
docker build -t project-api -f ./Dockerfile .

echo "Starting containers..."
docker run -d \
  -p 8080:8080 \
  --name $CONTAINER_USER \
  --network specno-network \
  --env-file .env \
  user-api

docker run -d \
  -p 8083:8083 \
  --name $CONTAINER_PROJECT \
  --network specno-network \
  --env-file .env \
  project-api

timeout=10
start_time=$(date +%s)

check_service() {
  local container=$1
  local service_name=$2

  while true; do
    current_time=$(date +%s)
    elapsed_time=$((current_time - start_time))

    if [ "$elapsed_time" -ge "$timeout" ]; then
      echo "Timeout reached for $service_name. Exiting..."
      break
    fi

    if docker logs $container 2>&1 | grep -q "Server started"; then
      echo "$service_name is up and running."
      break
    fi

    sleep 1
  done
}

check_service $CONTAINER_USER "User API"
check_service $CONTAINER_PROJECT "Project API"