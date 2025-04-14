    #!/bin/sh
    # Rebuilds and starts the full microservices stack (API gateway + services) using Docker Compose
    # Author: Samoei Oloo
    # Updated: 2025-04-14

    echo "Stopping all containers and removing volumes..."
    docker-compose down -v

    echo "Rebuilding and starting containers with docker-compose..."
    docker-compose up --build -d
