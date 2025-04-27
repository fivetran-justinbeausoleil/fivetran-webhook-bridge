# Project settings
APP_NAME = webhook-server

# Default target
.PHONY: all
all: build

# Build binary
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME) main.go

# Run locally (for dev/testing)
.PHONY: run
run:
	@echo "Running $(APP_NAME)..."
	go run main.go

# Cross compile for Mac (Intel)
.PHONY: build-mac
build-mac:
	@echo "Building $(APP_NAME) for Mac (amd64)..."
	GOOS=darwin GOARCH=amd64 go build -o $(APP_NAME)-mac main.go

# Cross compile for Mac (Apple Silicon M1/M2)
.PHONY: build-mac-arm
build-mac-arm:
	@echo "Building $(APP_NAME) for Mac (arm64)..."
	GOOS=darwin GOARCH=arm64 go build -o $(APP_NAME)-mac-arm main.go

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME) $(APP_NAME)-mac $(APP_NAME)-mac-arm