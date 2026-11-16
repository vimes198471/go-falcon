# go-falcon

![CI](https://img.shields.io/github/actions/workflow/status/vimes198471/go-falcon/ci.yml?branch=main&label=CI) ![Go](https://img.shields.io/badge/go-1.21+-00ADD8) ![GoDoc](https://img.shields.io/badge/godoc-reference-00ADD8) ![License](https://img.shields.io/github/license/vimes198471/go-falcon?color=blue)

A distributed-ready Go service with built-in health checks, structured request logging, and panic recovery middleware. go-falcon uses gorilla/mux for routing, provides a thread-safe in-memory item store with RWMutex, and supports graceful shutdown via signal handling. Includes a multi-stage Dockerfile (Alpine-based final image) and a GitHub Actions CI workflow.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Running Tests](#running-tests)
- [Contributing](#contributing)
- [Changelog](#changelog)
- [License](#license)

## Features

- Single binary deployment -- no external runtime dependencies
- Request logging middleware with method, path, and duration
- HTTP server with gorilla/mux routing and standard net/http handlers
- Graceful shutdown on SIGINT/SIGTERM via signal.Notify
- /api/health endpoint returning JSON status and timestamp

## Prerequisites

- Go 1.21 or higher
- Docker (optional, for containerized deployment)

## Installation

```bash
go get https://github.com/vimes198471/go-falcon
# or build from source
git clone https://github.com/vimes198471/go-falcon.git
cd go-falcon && go build -o bin/server cmd/server/main.go
# or run with Docker
docker build -t go-falcon . && docker run -p 8080:8080 go-falcon
```

## Usage

```bash
# Run the server (defaults to :8080)
go run cmd/server/main.go
# => Server starting on :8080

# Build a production binary
go build -o bin/server cmd/server/main.go

# Run with a custom port
PORT=3000 ./bin/server

# Health check
curl http://localhost:8080/api/health
# => {"status":"ok","timestamp":"2024-01-15T10:30:00Z"}

# Run tests
go test ./... -v

# Build with Docker
docker build -t go-falcon .
docker run -p 8080:8080 go-falcon
```

## Project Structure

```
go-falcon/
├── cmd/server/main.go
├── internal/
│   ├── handler/handler.go
│   └── middleware/middleware.go
├── go.mod  go.sum  Dockerfile  Makefile  .env.example
├── README.md  LICENSE  CHANGELOG.md  CONTRIBUTING.md  CODE_OF_CONDUCT.md
├── .github/
│   ├── workflows/ci.yml
│   ├── ISSUE_TEMPLATE/
│   │   ├── bug_report.md
│   │   └── feature_request.md
│   └── PULL_REQUEST_TEMPLATE.md
└── .gitignore
```

## Running Tests

```bash
go test ./... -v
# with coverage
go test ./... -cover -coverprofile=coverage.out
```

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

Quick start:

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Make your changes and run tests (see [Running Tests](#running-tests) above)
4. Commit with a clear message: `git commit -am 'Add: my feature'`
5. Push to your fork: `git push origin feature/my-feature`
6. Open a Pull Request

Please ensure all tests pass before submitting a pull request.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for a list of changes and version history.

## License

This project is licensed under the Apache-2.0 -- see the [LICENSE](LICENSE) file for details.

Copyright (c) 2026 ivan smith.

## Acknowledgements

- Built with Go, gorilla/mux, and the standard library
- CI/CD powered by [GitHub Actions](https://github.com/features/actions)
- Badges by [shields.io](https://shields.io)
- Inspired by best practices from the open-source community
