# RFZ Developer CLI - Makefile
# Build, test, and development commands

.PHONY: build run test lint clean dev golden-update

# Build the CLI binary
build:
	go build -o bin/rfz-cli ./cmd/rfz-cli

# Run the CLI in development mode
run:
	go run ./cmd/rfz-cli

# Run all tests
test:
	go test -v -race ./...

# Run visual regression tests with golden file comparison
test-golden:
	go test -v ./internal/ui/... -run Golden

# Update golden files (use after intentional UI changes)
golden-update:
	go test ./internal/ui/... -run Golden -update

# Run linter
lint:
	golangci-lint run ./...

# Format code
fmt:
	gofmt -w .
	goimports -w .

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Install development dependencies
dev-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest

# Build for all platforms
build-all:
	GOOS=darwin GOARCH=amd64 go build -o bin/rfz-darwin-amd64 ./cmd/rfz-cli
	GOOS=darwin GOARCH=arm64 go build -o bin/rfz-darwin-arm64 ./cmd/rfz-cli
	GOOS=linux GOARCH=amd64 go build -o bin/rfz-linux-amd64 ./cmd/rfz-cli
	GOOS=windows GOARCH=amd64 go build -o bin/rfz-windows-amd64.exe ./cmd/rfz-cli

# Run with mock data (for UI development)
run-mock:
	RFZ_MOCK=true go run ./cmd/rfz-cli

# Help
help:
	@echo "RFZ Developer CLI - Available commands:"
	@echo "  make build        - Build the CLI binary"
	@echo "  make run          - Run the CLI"
	@echo "  make test         - Run all tests"
	@echo "  make test-golden  - Run visual regression tests"
	@echo "  make golden-update- Update golden files"
	@echo "  make lint         - Run linter"
	@echo "  make fmt          - Format code"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make dev-deps     - Install development dependencies"
	@echo "  make build-all    - Build for all platforms"
	@echo "  make run-mock     - Run with mock data"
