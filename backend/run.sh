#!/bin/sh
# Starts PostgreSQL in a local Docker container and runs migrations
# Author: Samoei Oloo
# Created: 2025-03-28
#
# This script sets up a docker container to run the backend project locally

CONTAINER_ONE="specno-user-service"
CONTAINER_TWO="specno-db"
cleanup() {
  if [ "$(docker ps -q -f name=$CONTAINER_ONE)" ]; then
    echo "Stopping container: $CONTAINER_ONE"
    docker stop $CONTAINER_ONE
  fi
  if [ "$(docker ps -q -f name=$CONTAINER_TWO)" ]; then
    echo "Stopping container: $CONTAINER_TWO"
    docker stop $CONTAINER_TWO
  fi
  docker rm $CONTAINER_ONE
  docker rm $CONTAINER_TWO
  #docker kill specno-user-service specno-db || true
  docker volume rm specno-db || true
  docker network rm specno-network || true
}

trap cleanup INT

docker network create specno-network || true

echo "Removing existing containers..."
docker rm -f specno-user-service || true
docker rm -f specno-postgres || true
docker volume rm specno-db || true=
docker rmi -f specno-user-service || true

docker exec -it specno-db psql -U postgres -d postgres -c "
DO \$\$ DECLARE
    r RECORD;
BEGIN
    FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
        EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.tablename) || ' CASCADE';
    END LOOP;
END \$\$;"

docker build -t specno-user-service -f Dockerfile .

echo "Starting container..."
docker run -d \
  -p 8080:8080 \
  --name specno-user-service \
  --network specno-network \
  specno-user-service

timeout=1
start_time=$(date +%s)

while true; do
  current_time=$(date +%s)
  elapsed_time=$((current_time - start_time))
  
  if [ "$elapsed_time" -ge "$timeout" ]; then
    echo "Exiting..."
    break
  fi

  sleep 1
done