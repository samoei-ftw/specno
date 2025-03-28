#!/bin/sh
# Starts PostgreSQL in a local Docker container and runs migrations
# Author: Samoei Oloo
# Created: 2025-03-28
#
# This script sets up a docker container to run the backend project locally
docker build -t specno-backend .

docker run -d -p 8080:8080 --name specno-be specno-backend