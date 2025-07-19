
          
# Backend Absensi - Golang Fiber API

A RESTful API for an attendance management system built with Golang, Fiber framework, and MySQL database.

## Features

- Employee management
- Department management
- Attendance tracking (clock-in and clock-out)
- JWT Authentication
- Swagger API documentation

## Tech Stack

- [Go](https://golang.org/) - Programming language
- [Fiber](https://gofiber.io/) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [MySQL](https://www.mysql.com/) - Database
- [JWT](https://github.com/golang-jwt/jwt) - Authentication
- [Swagger](https://github.com/gofiber/swagger) - API documentation

## Prerequisites

- Go 1.16 or higher
- MySQL 5.7 or higher
- Git

## Project Structure

```
├── config/             # Database configuration
├── controller/         # HTTP request handlers
├── docs/               # Swagger documentation
├── dto/                # Data Transfer Objects
├── middleware/         # Custom middleware
├── model/              # Database models
├── repository/         # Database operations
├── routes/             # API routes
├── service/            # Business logic
├── utils/              # Utility functions
├── .env                # Environment variables
├── .gitignore          # Git ignore file
├── main.go             # Application entry point
└── README.md           # Project documentation
```

## Environment Setup

1. Clone the repository

```bash
git clone <repository-url>
cd backend-absensi
```

2. Create a `.env` file in the root directory with the following variables:

```
# Application
APP_PORT=3030

# Database
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=your_password
DB_NAME=absensi_db

# JWT
JWT_SECRET=your_jwt_secret_key
```

3. Install dependencies

```bash
go mod tidy
```

## Database Setup

1. Create a MySQL database

```bash
mysql -u root -p
```

```sql
CREATE DATABASE absensi_db;
EXIT;
```

2. The application will automatically migrate the database schema when started

## Running the Application

1. Start the server

```bash
go run main.go
```

2. The server will start at `http://localhost:3030`
3. Swagger documentation will be available at `http://localhost:3030/swagger/`

## API Endpoints

### Authentication
- `POST /api/login/:employeeId` - Login with employee ID

### Employees (Protected Routes)
- `GET /api/employee` - Get all employees
- `GET /api/employee/:id` - Get employee by ID
- `POST /api/employee` - Create a new employee
- `PUT /api/employee/:id` - Update an employee
- `DELETE /api/employee/:id` - Delete an employee

### Departments (Protected Routes)
- `GET /api/departement` - Get all departments
- `GET /api/departement/:id` - Get department by ID
- `POST /api/departement` - Create a new department
- `PUT /api/departement/:id` - Update a department
- `DELETE /api/departement/:id` - Delete a department

### Attendance (Protected Routes)
- `POST /api/attendance/clock-in` - Clock in
- `PUT /api/attendance/clock-out/:id` - Clock out
- `GET /api/attendance/logs` - Get attendance logs

## Authentication

The API uses JWT for authentication. To access protected routes, include the JWT token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```

## Development

### Generating Swagger Documentation

This project uses Swagger for API documentation. To update the documentation:

1. Install swag

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

2. Generate documentation

```bash
swag init
```

## License

This project is licensed under the MIT License.

        