# Specno Technical Assessment Submission

This repository contains the full solution for the Specno Technical Assessment.

---

## 📦 Prerequisites

Ensure the following tools are installed on your system before running the project:

- **Node.js**
- **npm** (or an alternative like `yarn`)
- **Docker** – for running services in containers.
- **DBeaver** (or any SQL client) – to interact with the PostgreSQL database.

### Configuration

- A `.env` file is required with all necessary environment variables.
- **TODO:** Add any additional configuration details here.

---

## 🚀 Running the Project Locally

### 🛠️ Backend and API Gateway

There are two shell scripts to streamline the local setup.

#### 1. Start the PostgreSQL Database

Use the `db.sh` script to spin up a PostgreSQL container:

```bash
./db.sh
```

This script starts the PostgreSQL service in Docker, making the database accessible to the backend.

#### 2. Run the Backend Microservices

Use the run.sh script from the root of the repository to stop existing containers, rebuild, and launch all microservices:
This launches the full stack, including:
	•	user_service
	•	task_service
	•	project_service
	•	NGINX (API gateway)

#### 🔍 What Each Script Does
	•	db.sh:
	•	Spins up a PostgreSQL container.
	•	Loads environment variables like POSTGRES_USER, POSTGRES_PASSWORD, and POSTGRES_DB from the .env file.
	•	Waits for the database to initialize.
	•	Applies migrations from the migrations directory.
	•	run.sh:
	•	Stops and removes all existing Docker containers and volumes.
	•	Rebuilds Docker images for the backend services.
	•	Starts the full Dockerized microservice stack using docker-compose.

Once everything is running, the backend API should be accessible through the API Gateway (NGINX).

⸻

### 🎨 Frontend

To get the frontend running locally:
	1.	Navigate to the frontend directory:
  ```bash
  cd frontend
  ```
  2.	Install dependencies:
  ```bash
  npm install
  ```
  3.	Start the development server:
  ```bash
  npm run dev
  ```
  4. Open your browser and navigate to:
  [http://localhost:5173](http://localhost:5173)

  ---

## 🧪 Testing

This project includes a Postman workspace with pre-configured requests for testing all endpoints.

### Access the Workspace

You can access the Postman collection here:  
[🔗 Specno Assessment – Postman Workspace](https://app.postman.com/workspaces/0efc949a-70c2-4083-9a79-bd7bf77c1907)

### Running Tests

1. Ensure the backend services and database are running (`./db.sh` and `./run.sh`).
2. Open Postman and import the collection or access it through the shared link.
3. Use the collection to:
   - TODO- add list of tests here

### Notes

- Ensure your `.env` values match the API URLs and ports used in Postman.
- If you're using JWTs, tokens will need to be copied into the `Authorization` headers after login.