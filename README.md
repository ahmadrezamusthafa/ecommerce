# eCommerce API
- [Getting Started](#getting-started)
  - [Installation](#database-migration)
  - [Database Migration](#installation)
  - [Running the Service](#running-the-server)
- [API Endpoints](#api-endpoints)

## Getting Started
### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ahmadrezamusthafa/ecommerce.git
   ```
2. Navigate to the project directory:
   ```bash
   cd ecommerce
   ```
3. Install dependencies (if using Go modules):
   ```bash
   go mod download
   ```
4. Start infrastructure:
   ```bash
   make start-infra
   ```
5. Execute migration:
   ```bash
   make run-migration
   ```
6. Run the service:
   ```bash
   make run-service
   ```
   The service will start on `http://localhost:8005`.

## API Endpoints

Below is a list of available API endpoints. For detailed documentation, including request and response payloads, refer to the [API Documentation](https://documenter.getpostman.com/view/7913952/2sAYdcqXVz).

| **Endpoint**                     | **Method** | **Description**                       |
|----------------------------------|------------|---------------------------------------|
| `/api/v1/user/register`          | POST       | Register a new user                   |
| `/api/v1/user/login`             | POST       | Login a user                          |
| `/api/v1/user/update`            | PUT        | Update user profile                   |
| `/api/v1/products`               | GET        | Get all products                      |
| `/api/v1/products/{id}`          | GET        | Get a specific product by ID          |
| `/api/v1/products/search`        | GET        | Search products by query              |
| `/api/v1/cart`                   | GET        | Get the users cart                    |
| `/api/v1/cart/items`             | POST       | Add an item to the cart               |
| `/api/v1/cart/items/{id}`        | DELETE     | Remove an item from the cart          |
| `/api/v1/orders`                 | POST       | Submit an order                       |
| `/api/v1/orders/top-customers`   | GET        | Get top customers                     |
| `/api/v1/orders`                 | GET        | Get all orders                        |
| `/api/v1/account/deposit`        | POST       | Deposit funds into the users account  |
| `/api/v1/account/withdraw`       | POST       | Withdraw funds from the users account |
| `/api/v1/account`                | GET        | Get the users account balance         |
