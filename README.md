# Persistent Go REST API Server

A lightweight, production-grade CRUD REST API built using **Go (Golang)** and **SQLite3**. This project features a completely decoupled modular architecture, custom logging middleware, strict input data validation, and persistent local database storage.

---

## Key Features

* **Modular Architecture:** Clean structural separation of concerns across `/config`, `/handlers`, and `/middleware` packages.
* **Database Persistence:** Integrated with SQLite3 via the standard `database/sql` pool to store records across runtime sessions.
* **Custom Middleware Logger:** Tracks incoming HTTP request methods, server execution timestamps, and exact routing paths natively.
* **Input Validation Gatekeeper:** Strict data handling that prevents malformed body data, negative integers, or missing parameters before database insertion.

---

## Project Architecture Layout

```text
My Server/
├── go.mod                 # Go module definition
├── go.sum                 # Dependency checksum hashes
├── main.go                # Central routing & initialization entry point
├── storage.db             # Auto-generated SQLite database file
├── config/
│   └── db.go              # Database connection initialization & table migration
├── middleware/
│   └── logger.go          # Custom HTTP request tracking pipeline
└── handlers/
    └── users.go           # Core HTTP endpoint handler functions (CRUD)
