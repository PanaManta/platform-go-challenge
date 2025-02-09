GO = go
BINARY_NAME = app
SRC_DIR = ./cmd/api
BUILD_DIR = ./build
DOCS_DIR=docs
SWAGGER=github.com/swaggo/swag/cmd/swag

.PHONY: all
all: build

.PHONY: build
build:
	@echo "Building the Go application..."
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)

.PHONY: run
run: build
	@echo "Running the Go application..."
	$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: clean
clean:
	@echo "Cleaning up build files..."
	rm -rf $(BUILD_DIR)

.PHONY: lint
lint:
	@echo "Running Go lint..."
	golangci-lint run

.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test ./...

.PHONY: install
install:
	@echo "Installing dependencies..."
	$(GO) mod tidy

.PHONY: docs
docs:
	@echo "Generating docs..."
	go run $(SWAGGER) init -g cmd/api/main.go -o $(DOCS_DIR)