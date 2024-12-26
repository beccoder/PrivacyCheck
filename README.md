# Project Documentation

## Overview
This project is a service designed to manage user registration, authentication, and searching for user-related leak data. The application interacts with a PostgreSQL database and utilizes external APIs for enhanced functionality. Below is a detailed guide on how to set up and use this project.

---

## Requirements
Before running the project, ensure the following tools are installed:
- [Golang](https://golang.org/)
- [Docker](https://www.docker.com/)

---

## Setup and Run Instructions
1. Clone the repository.
2. Navigate to the project directory.
    - Create a `.env` file based on the `example.env` file.
3. Start the Docker environment:
   ```bash
   make docker-start
   ```
    - This command will:
        - Initialize a PostgreSQL database.
        - Apply all migration files to set up the database schema.
4. Start the application:
      ```bash
      go run cmd/main.go
      ```
5. Open the Swagger documentation in your browser:
    - [Swagger UI](http://localhost:8080/api/v1/swagger/index.html)

---

## APIs
The service provides three main APIs:

### 1. Register
- **Endpoint:** `POST /api/v1/auth/register`
- **Description:** Registers a new user into the system.
- **Request Body:**
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```
- **Response:**
  ```json
  {
    "message": "User registered successfully"
  }
  ```

### 2. Login
- **Endpoint:** `POST /api/v1/auth/login`
- **Description:** Authenticates an existing user and returns a token.
- **Request Body:**
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```
- **Response:**
  ```json
  {
    "token": "jwt-token"
  }
  ```

### 3. Search Leak Data
- **Endpoint:** `GET /api/v1/search-my-leak-data`
- **Description:** Searches for leak data associated with the authenticated user.
- **Headers:**
  ```text
  Authorization: Bearer <jwt-token>
  ```
- **Response:**
  ```json
  {
    "status": "FOUND",
    "data": [
      {
        "id": 1,
        "ip": "42.48.100.32",
        "age": 28
      }
    ]
  }
  ```

---

## Database
The system uses a PostgreSQL database with the following tables:

### 1. Users Table
- **Purpose:** Stores user authentication data.
- **Columns:**
    - `id`: Unique identifier.
    - `firstname` User's first name 
    - `lastname` User's last name
    - `email`: User's email address.
    - `password`: Hashed password.

### 2. Leak Data Table
- **Purpose:** Stores information about user leak data.
- **Columns:**
    - `id`: Unique identifier.
    - `user_id`: Reference to the user.
    - `status`: Status of the data (e.g., FOUND or NOT_FOUND).
    - `data`: JSONB column containing detailed leak data.

---

## Search API Workflow
1. The Search API first checks the local database for leak data associated with the user.
2. If no data is found locally, it queries the external DummyJson API using the user's email.
3. The response from DummyJson is:
    - Saved into the database with a status of `FOUND` or `NOT_FOUND`.
    - Returned to the user.
4. On subsequent searches, the system retrieves the saved data directly from the local database.

---

## Notes
- This implementation focuses on core functionality and working logic.
- Comprehensive error handling has not been fully implemented as the primary goal was to complete the technical task.

---