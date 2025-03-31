#!/bin/sh
# Starts PostgreSQL in a local Docker container and runs migrations
# Author: Samoei Oloo
# Created: 2025-03-28
#
# This script sets up a docker container to run the backend project locally

# Create a custom network if it doesn't exist
docker network create specno-network || true

# Start the PostgreSQL container (assuming it's already set up in run_db.sh)
# You can either source the `run_db.sh` script here or just call it after ensuring it's been set up
./run_db.sh

# Build the user service image
echo "Building user service Docker image..."
docker build -t specno-user-service -f Dockerfile .

# Run the user service container on the custom network
echo "Starting user service container..."
docker run -d \
  -p 8080:8080 \
  --name specno-user-service \
  --network specno-network \
  specno-user-service

echo "User service is running on port 8080."