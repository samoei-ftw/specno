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