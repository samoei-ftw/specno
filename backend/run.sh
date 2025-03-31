#!/bin/sh
# Starts PostgreSQL in a local Docker container and runs migrations
# Author: Samoei Oloo
# Created: 2025-03-28
#
# This script sets up a docker container to run the backend project locally

docker network create specno-network || true

echo "Removing existing containers..."
docker rm -f specno-user-service || true

docker rm -f specno-postgres || true
docker volume rm specno-postgres-data || true

echo "Removing existing Docker image(s)..."
docker rmi -f specno-user-service || true

docker exec -it specno-db psql -U postgres -d postgres -c "
DO \$\$ DECLARE
    r RECORD;
BEGIN
    FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
        EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.tablename) || ' CASCADE';
    END LOOP;
END \$\$;"

./run_db.sh

echo "Building Docker image..."
docker build -t specno-user-service -f Dockerfile .

echo "Starting container..."
docker run -d \
  -p 8080:8080 \
  --name specno-user-service \
  --network specno-network \
  specno-user-service

echo "Backend server is running on port 8080."