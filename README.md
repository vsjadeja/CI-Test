# CI-Test

Test project for continuous integration for build process.

## Overview

This project demonstrates CI/CD practices with automated testing and code coverage validation through GitHub Actions.

## Code Coverage Requirements

The project enforces a **minimum code coverage threshold of 60%**. 

### GitHub Actions Behavior

- **Coverage > 60%**: The workflow continues to the next steps (deployment, publishing, etc.)
- **Coverage ≤ 60%**: The GitHub Action fails and stops the workflow

This ensures code quality standards are maintained before any changes are merged or deployed.

## Setup

### Prerequisites

- Go 1.16 or higher
- Git

### Running Tests

To run the test suite locally:

```bash
go test -v ./src
```

To check code coverage:

```bash
go test -coverprofile=coverage.out ./src
go tool cover -html=coverage.out
```

## Project Structure

```
.
├── src/
│   ├── helpers.go       # Main helper functions
│   ├── helpers_test.go  # Test suite for helpers
│   ├── main.go          # Main application
│   └── go.mod           # Go module definition
├── build/               # Build artifacts
└── README.md            # This file
```

## Helper Functions

- `add(a, b int) int`: Returns the sum of two integers
- `subtract(a, b int) int`: Returns the difference of two integers
- `multiply(a, b int) int`: Returns the product of two integers
- `divide(a, b int) (int, error)`: Returns the quotient of two integers; returns error if dividing by zero
- `modulus(a, b int) (int, error)`: Returns the remainder of division; returns error if dividing by zero

## CI/CD Pipeline

The GitHub Actions workflow automatically:
1. Runs all tests with verbose output
2. Calculates code coverage percentage
3. Validates coverage meets the 60% threshold
4. **Builds the binary** if coverage threshold is met
5. **Commits and pushes** the build artifact to the `build/` directory
6. Fails and notifies if coverage is insufficient

### Build Output

- **Binary location**: `build/ci-test`
- **Automatically committed**: Yes (when coverage > 60%)
- **Build trigger**: Every push to `main` or `develop` branches

### Running the Built Binary

After the workflow completes successfully, you can find the compiled binary in the `build/` directory:

```bash
./build/ci-test
```

## Workflow Status

The workflow file is located at `.github/workflows/ci.yml` and is triggered on:
- Push to `main` or `develop` branches
- Pull requests to `main` or `develop` branches
