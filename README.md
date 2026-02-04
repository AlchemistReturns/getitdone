# GetItDone - Go API Boilerplate

This project is a clean, minimal boilerplate for creating a REST API using Go (Golang). It serves as an educational resource and a starting point for developers looking to understand the basics of building secure, scalable web services with Go, Gin, and GORM.

It demonstrates key concepts such as:
- **Project Structure**: A standard Go project layout (`cmd`, `internal`, etc.).
- **Authentication**: JWT-based authentication using HTTP-only cookies.
- **Database**: PostgreSQL integration using GORM (ORM), including migrations.
- **Routing**: structured routing with the Gin web framework.

## Project Structure

The project follows a standard Go project layout to ensure modularity and maintainability.

```
backend/
├── cmd/
│   └── api/
│       └── main.go           # Application entry point. Initializes the server and routes.
├── database/
│   ├── db.go                 # Database connection logic (GORM + PostgreSQL).
│   └── migrate.go            # Auto-migration logic to keep DB schema in sync with models.
├── internal/
│   ├── auth/                 # Authentication logic (Register, Login, Token generation).
│   ├── middlewares/          # HTTP middlewares (e.g., RequireAuth for JWT verification).
│   ├── models/               # Data structures (structs) mapping to database tables.
│   ├── protected/            # Handlers for protected routes (require authentication).
│   └── public/               # Handlers for public routes (accessible by everyone).
├── .env                      # Environment variables (DB credentials, JWT secret).
├── docker-compose.yaml       # Docker configuration for running the app and DB.
├── Dockerfile                # Instructions to build the Go application container.
├── go.mod                    # Go module definition and dependencies.
└── go.sum                    # Checksums for dependencies.
```

## API Endpoints

All API routes are prefixed with `/api/v1`.

### Public Routes

#### 1. Home Check
Returns a simple welcome message to verify the API is running.

- **URL**: `/api/v1/`
- **Method**: `GET`
- **Response**:
```json
{
  "message": "Home"
}
```

#### 2. Register
Create a new user account.

- **URL**: `/api/v1/register`
- **Method**: `POST`
- **Body**:
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "secretpassword"
}
```
- **Response** (200 OK):
```json
{
  "user": {
    "ID": 1,
    "CreatedAt": "2024-02-04T12:00:00Z",
    "UpdatedAt": "2024-02-04T12:00:00Z",
    "DeletedAt": null,
    "name": "John Doe",
    "email": "john@example.com",
    "password": "" // Password is created but not returned in response for security
  }
}
```

#### 3. Login
Authenticate a user and receive a secure HTTP-only cookie containing the JWT.

- **URL**: `/api/v1/login`
- **Method**: `POST`
- **Body**:
```json
{
  "email": "john@example.com",
  "password": "secretpassword"
}
```
- **Response** (200 OK):
```json
{}
```
*Note: A cookie named `Authorization` is set in the response headers.*

### Protected Routes
These routes require a valid `Authorization` cookie obtained from the Login endpoint.

#### 1. Get Profile
Retrieve the currently logged-in user's information. The user details are extracted directly from the JWT token.

- **URL**: `/api/v1/protected/profile`
- **Method**: `GET`
- **Headers**: Cookie: `Authorization=<jwt_token>`
- **Response**:
```json
{
  "email": "john@example.com",
  "name": "John Doe"
}
```

## Getting Started

### Prerequisites
- Docker & Docker Compose
- Go 1.25+ (if running locally without Docker)

### Running with Docker (Recommended)

1.  Clone the repository.
2.  Run the application and database:
    ```bash
    docker-compose up --build
    ```
3.  The API will be available at `http://localhost:8080`.

### Running Locally

1.  Ensure PostgreSQL is running.
2.  Create a `.env` file in the `backend` directory with your database credentials.
3.  Run the application:
    ```bash
    cd backend
    go run cmd/api/main.go
    ```
