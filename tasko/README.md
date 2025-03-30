# Specno Technical Assessment Submission
This is the central repository for the Specno Technical Assessment

## Pre-requisites
Before running the project, ensure you have the following installed:
	•	Docker: For running the services in containers.
	•	DBeaver (or any other SQL client): To interact with the PostgreSQL database.
## Running the project locally
There are two shell scripts to help you set up the project locally.
	1.	Running the PostgreSQL Database:
First, run the run_db.sh script. This script will start a local PostgreSQL server inside a Docker container. To run it, execute the following command:
image.png
This will bring up the PostgreSQL container, ensuring that the backend service can connect to the database.

	2.	Running the Backend Service:
After the database is running, run the run.sh script in the root of the repository. This script will:
	•	Build the Docker image from the Dockerfile.
	•	Tag the image as specno-backend.
	•	Run the backend service in a container, accessible on port 8080.
To run the backend service, execute:
image.png
Once both scripts have run successfully, the backend service will be accessible at:
http://localhost:8080.
How it Works:
	•	run_db.sh: Starts a PostgreSQL container with environment variables such as POSTGRES_USER, POSTGRES_PASSWORD, and POSTGRES_DB defined in the .env file. It waits for the database to be ready and applies SQL migrations found in the migrations directory.
	•	run.sh: Builds the Go backend service Docker image from the Dockerfile, makes the Go binary executable, and starts the service in a Docker container.
You can now interact with the backend API through this endpoint.
