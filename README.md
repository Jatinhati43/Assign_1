# Assign_1 – Second Highest Number Finder in Go

This repository contains a Go program that finds the **second highest number** in an array of integers. It follows the standard Go project layout as recommended by [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

## 🧠 Functionality

The application:
- Accepts a predefined array of integers.
- Identifies the second highest unique number.
- Prints the result to the terminal.

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or later

### Running the Application

```bash
go run ./cmd

Example Output

Array: [3, 10, 6, 1, 20]
Second highest number: 10

🧪 Running Tests

go test ./internal/...

✅ Running with Coverage

go test -cover ./internal/...
