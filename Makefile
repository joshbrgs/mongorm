# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=mongorm
BINARY_UNIX=$(BINARY_NAME)_unix

# Build the project
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

# Run tests
test: 
	$(GOTEST) -v ./...

# Clean the project
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Cross compilation for Linux
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

# Run the application
run: 
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)

# Format the code
fmt:
	$(GOCMD) fmt ./...

# Lint the code
lint:
	golangci-lint run ./...

# Format, lint, and test
validate-build: fmt lint test

.PHONY: build test clean run fmt lint validate-build
