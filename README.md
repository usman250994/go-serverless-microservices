![Work in Progress](https://img.shields.io/badge/Status-Work_in_Progress-orange?style=for-the-badge&logo=github)

# 🏗️ Go Serverless Microservices

[![Go Version](https://img.shields.io/badge/Go-1.22+-blue.svg)](https://go.dev)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](./LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)]()
[![Architecture](https://img.shields.io/badge/Clean%20Architecture-DDD-orange.svg)]()

> **Scalable, maintainable, and production-ready microservices in Go, powered by Clean Architecture & DDD.**

---

## ✨ Key Features

- 🧹 **Clean Architecture**: Handlers → Services → Models
- 🏛️ **Domain-Driven Design (DDD)**: Modular, business-focused code
- 📦 **Scalable Structure**: Add new domains/services with zero coupling
- 🔌 **Flexible Integration**: REST, gRPC, GraphQL ready
- 🧪 **Testable Services**: Clear boundaries for easy testing
- ⚡ **Minimal Boilerplate**: Production-ready from the start

---

## 📂 Project Structure

```text
cmd/
 └── users/
	└── main.go        # Entry point for Users microservice
internal/
 └── users/
	├── routes.go      # Routes & handler mapping
	├── handler.go     # HTTP handlers
	├── service.go     # Business logic
	└── model.go       # Domain models/entities
	└── repo.go       # crud operations for db
 	└── type.go       # request response types for validation
```

---

## 🚀 How It Works

- **main.go**: Starts the service & initializes dependencies
- **routes.go**: Registers endpoints & maps to handlers
- **handler.go**: Processes requests, calls services
- **type.go**:  Contains requests' DTOs, for validation
- **service.go**: Contains business logic & workflows
- **model.go**: Defines domain entities & data structures
- **repo.go**: Defines db crud operations using Dynamo

---

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **Cloud**: AWS (tested with Cognito, Api gateway and lambdas)
- **Architecture**: Clean Architecture + DDD
- **Routing**: chi
- **Persistence**: Extendable (DynamoDB)
- **Validation**:  go-playground/validator
- **Authentication**:  using jwt for extracting userId

---

## 🧭 Principles

- ✅ **Single Responsibility**: Each layer does one job
- ✅ **Dependency Inversion**: Inner layers never depend on outer
- ✅ **Explicit Boundaries**: Clear contracts between layers
- ✅ **Scalability**: Easily add new domains

---

## 🚀 Getting Started

```sh
# Clone repo
$ git clone https://github.com/usman250994/go-serverless-microservices.git

# Run users service
$ cd cmd/users
$ go run main.go
```

---

## 🚧 Upcoming Features

- 🗄️ **Database Abstraction**: Internal package for DynamoDB access—cleaner, decoupled repo logic.
- 📡 **Event Streaming**: Integrate AWS SNS, SQS, and DynamoDB Streams for real-time event sourcing.
- 🔍 **Query Wrapper**: Generic query builder to simplify and standardize DynamoDB queries.

---


## 🤝 Contributing

Contributions are welcome! Open issues or PRs to improve structure or add features.

---

## 📜 License

MIT License – use this boilerplate for your own projects.
