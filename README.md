
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

## API Endpoints

| Method | Endpoint               | Description                     |
|--------|------------------------|---------------------------------|
|POST    |`/signup `              |Create user                      |
|POST    |`/login/`               |Authenticate user                |
|POST    |`/pay`                  |Process a payment                |
|GET     |`/pay/message`          |Verify a payment                 |
|GET     |`/wallet`               |Get wallet details               |
|POST    |`/wallet/create`        |Create wallet                    |
|POST    |`/wallet/create/card`   |Add new card                     |
|DELETE  |`/wallet/delete`        |Delete wallet                    |
|DELETE  |`/wallet/delete/card`   |Delete card                      |
|POST    |`/verify`               |Verify user                      |


