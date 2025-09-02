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
```

---

## 🚀 How It Works

- **main.go**: Starts the service & initializes dependencies
- **routes.go**: Registers endpoints & maps to handlers
- **handler.go**: Processes requests, calls services
- **service.go**: Contains business logic & workflows
- **model.go**: Defines domain entities & data structures

---

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **Architecture**: Clean Architecture + DDD
- **Routing**: gorilla-mux
- **Persistence**: Extendable (DynamoDB)

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
$ git clone https://github.com/yourusername/go-microservice-starter.git

# Run users service
$ cd cmd/users
$ go run main.go
```

---

## 🤝 Contributing

Contributions are welcome! Open issues or PRs to improve structure or add features.

---

## 📜 License

MIT License – use this boilerplate for your own projects.
