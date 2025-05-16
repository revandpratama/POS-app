# POS Application

A Point of Sale (POS) system with RESTful API endpoints for user management, authentication, product management, and cashier operations.

## Features

- **User Management**: Create, read, update and delete users
- **Authentication**: Secure login and registration
- **Product Management**: Inventory handling with CRUD operations
- **Cashier Operations**: Process transactions and manage sales

## API Documentation

### Base URL

All API requests should be directed to:
```
http://localhost:8080/api
```

### Authentication

The API uses JWT Bearer token authentication for protected endpoints.

#### Login
```
POST /login
```

Request body:
```json
{
    "email": "user@example.com",
    "password": "password"
}
```

#### Register
```
POST /register
```

Request body:
```json
{
    "name": "user 1"
    "email": "user@example.com",
    "password": "password"
}
```

### User Endpoints

| Method | Endpoint    | Description      | Authentication |
|--------|-------------|------------------|---------------|
| GET    | /users      | Get all users    | Bearer token           |
| GET    | /users/:id  | Get a single user| Bearer token           |
| PUT    | /users/:id  | Update a user    | Bearer token            |
| DELETE | /users/:id  | Delete a user    | Bearer token            |

### Product Endpoints

| Method | Endpoint       | Description        | Authentication |
|--------|----------------|--------------------|---------------|
| GET    | /products      | Get all products   | Bearer token  |
| GET    | /products/:id  | Get single product | Bearer token  |
| POST   | /products      | Create product     | Bearer token  |
| PUT    | /products/:id  | Update product     | Bearer token  |
| DELETE | /products/:id  | Delete product     | Bearer token  |

Example product creation:
```json
{
    "name": "Product Name",
    "price": 2000,
    "quantity": 200
}
```

### Cashier Endpoints

| Method | Endpoint                | Description           | Authentication |
|--------|-------------------------|-----------------------|---------------|
| GET    | /cashier/transactions   | Get transaction list  | No           |
| POST   | /cashier/transactions   | Create transaction    | No           |
| PUT   | /cashier/transactions   | Update stock          | No           |

## Installation

1. Clone the repository
2. Install dependencies
3. Configure environment variables
4. Start the server

```bash
# Example commands
git clone https://github.com/yourusername/pos-app.git
cd pos-app
go mod install
go run main.go
```

## Usage

1. Register a user account
2. Log in to obtain your authentication token
3. Use the token in the Authorization header for authenticated requests
4. Manage products and process transactions

## Development

The API is built using:
- Golang (Echo)
- JWT for authentication
- RESTful API design principles
