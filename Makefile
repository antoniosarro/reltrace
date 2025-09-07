.PHONY: build dev-deps

# Build variables
BINARY_NAME=reltrace
BUILD_DIR=bin
MAIN_PATH=cmd/reltrace
VERSION?=dev
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

# Build flags
LDFLAGS=-X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.BuildTime=$(BUILD_TIME)

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=1 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME) ./$(MAIN_PATH)

# Build for multiple platforms
build-all: build-linux build-darwin build-windows

build-linux:
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)-linux ./$(MAIN_PATH)

build-darwin:
	@echo "Building for macOS..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin ./$(MAIN_PATH)

build-windows:
	@echo "Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)-windows.exe ./$(MAIN_PATH)

# Install dependencies
dev-deps:
	@echo "Installing development dependencies..."
	go mod download
	go mod tidy

# Format the code
format:
	@echo "Formatting code..."
	go fmt ./...

# Run the application in development mode
dev: build
	./$(BUILD_DIR)/$(BINARY_NAME)