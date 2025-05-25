# m8 – Go Database Migration Tracker

A simple CLI tool to create and apply database schema migrations, built in Go using clean architecture principles.

---

## Features (MVP)

- `make` — Create new up/down migration SQL files
- `up` — Apply all pending migrations
- `down` — Revert the last applied migration
- `status` — Show which migrations are applied and which are pending

---

## Getting Started

### Prerequisites

- Go 1.24 or later

### Running

1. Clone the repository:

```bash
git clone https://github.com/ShivangSrivastava/m8.git
cd m8
````

2. Build or run the CLI commands:

```bash
go run main.go make add_users_table
go run main.go up
go run main.go down
go run main.go status
```

---

## Configuration

Currently, database connection details and migration settings are hardcoded in the source code.

> **TODO:** Add support for external configuration files (YAML/ENV) for database credentials, migration directory, and other settings.

---

## Project Structure

* `cmd/` — CLI commands entry points (`make`, `up`, `down`, `status`)
* `internal/app/` — Core migration logic
* `internal/infra/` — Database and filesystem implementations
* `internal/core/` — Core domain entities and types
* `migrations/` — Auto-generated SQL migration files

---

## Next Steps (TODO)

* Improve code quality with consistent naming conventions
* Remove code duplication and improve maintainability
* Add more comprehensive tests with mocking for DB interactions
* Replace hardcoded configuration with external config files
* Write detailed documentation and better comments manually
