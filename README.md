# Zota API Integration

This project integrates with the Zota API for handling payment deposits and status checks. The implementation follows Domain-Driven Design (DDD) principles and uses Go.

## Project Structure

```plaintext
project-root/
├── cmd/
│   └── main.go
├── domain/
│   ├── model/
│   │   └── payment.go
│   ├── repository/
│   │   └── payment_repository.go
│   └── service/
│       ├── payment_service.go
│       └── payment_service_test.go
├── infrastructure/
│   └── zota/
│       └── zota_adapter.go
├── interfaces/
│   └── http/
│       ├── handler/
│       │   ├── deposit_handler.go
│       │   ├── status_handler.go
│       │   ├── deposit_handler_test.go
│       │   ├── status_handler_test.go
│       │   └── mock_payment_service.go
│       └── router.go
├── .env
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
└── docker-compose.yml
```

## Requirements

- Go 1.22
- Docker
- Docker Compose

## Installation

1. Clone the repository:
    ```sh
    git clone <repository-url>
    cd project-root
    ```

2. Copy `.env.example` to `.env` and fill in your Zota API credentials.

3. Build and run the Docker container:
    ```sh
    docker-compose up --build
    ```

## Usage

### Endpoints

- **POST /deposit**
    - **Description**: Initiates a deposit request.
    - **Query Parameters**:
        - `amount` (string): The amount for the deposit.
        - `currency` (string): The currency for the deposit.

- **GET /status**
    - **Description**: Checks the status of a deposit.
    - **Query Parameters**:
        - `merchantOrderID` (string): The merchant order ID to check the status for.

### Environment Variables

- `BASE_URL`: Base URL for the Zota API.
- `ENDPOINT_ID`: Endpoint ID for the Zota API.
- `API_SECRET_KEY`: API secret key for the Zota API.
- `MERCHANT_ID`: Merchant ID for the Zota API.

### Example `.env` File

```env
BASE_URL=https://api.zota.com
ENDPOINT_ID=your-endpoint-id
API_SECRET_KEY=your-api-secret-key
MERCHANT_ID=your-merchant-id
```

## Testing

To run tests:

```sh
go test ./...
```

## Project Structure Details

- `cmd/main.go`: The entry point of the application.
- `domain/model/`: Contains the data models for the project.
- `domain/repository/`: Defines the repository interfaces.
- `domain/service/`: Implements the business logic.
- `infrastructure/zota/`: Contains the Zota API adapter.
- `interfaces/http/handler/`: HTTP handlers for deposit and status endpoints.
- `interfaces/http/router.go`: Configures the HTTP router.

## License

This project is licensed under the MIT License.

