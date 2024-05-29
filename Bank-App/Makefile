# Makefile for Go Bank-App

# Variables
BINARY_NAME=Bank-App
BUILD_DIR=bin

# Targets
build:
	@echo "Building the application..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME)

run: build
	@echo "Running the application..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

test:
	@echo "Running tests..."
	@go test -v ./...

clean:
	@echo "Cleaning up..."
	@rm -f $(BUILD_DIR)/$(BINARY_NAME)

