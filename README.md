
# Payment System

## Overview
This payment system is built using Go and designed to manage user wallets, payment processing, invoicing, and other related operations. The system allows users to manage their cards and transactions efficiently, ensuring secure and fast payment workflows.

## Features
- **User Authentication**: Secure login and token-based authentication.
- **Wallet Management**: Multiple cards per wallet, transaction history.
- **Invoice Generation**: Track payments and generate invoices.
- **Notifications**: Real-time updates for payment statuses.
- **Reports**: Generate user payment and wallet reports.

## Technologies
- **Backend**: Golang (Gin Framework)
- **Database**: SQLite
- **Authentication**: JWT-based token system
- **Containerization**: Docker

## Project Structure
```plaintext
├── cmd/server         # Application entry point
├── internal           # Business logic and services
│   ├── auth           # Authentication services
│   ├── payment        # Payment processing services
│   ├── wallet         # Wallet management services
│   └── invoice        # Invoice services
├── pkg                # Helper packages
│   ├── config         # Configuration management
│   ├── database       # Database connection
│   └── middleware     # Custom middleware
├── api/v1             # API routes and handlers
├── test               # Unit tests
├── go.mod             # Module dependencies
└── Dockerfile         # Docker container setup
```

## Setup Instructions

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Abdulaxad7/payment-system.git
   cd payment-system
   ```

2. **Install dependencies**:
   Ensure Go and Docker are installed. Then, run:
   ```bash
   go mod download
   ```

3. **Run the application**:
   You can run the project locally or via Docker:
   ```bash
   go run cmd/server/main.go
   ```
   Or build and run it using Docker:
   ```bash
   docker build -t payment-system .
   docker run -p 8080:8080 payment-system
   ```

## API Endpoints

| Method | Endpoint               | Description                     |
|--------|------------------------|---------------------------------|
| POST   | `/login`                | Authenticate user               |
| GET    | `/wallets/{id}`         | Get wallet details              |
| POST   | `/payments`             | Process a payment               |
| POST   | `/invoices`             | Generate an invoice             |
| GET    | `/reports/{user_id}`    | Generate user report            |

## Future Improvements
- Add support for external payment gateways.
- Implement multi-currency support.
- Add frontend integration.

## License
This project is licensed under the MIT License.
