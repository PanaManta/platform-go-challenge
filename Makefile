GO = go
BINARY_NAME = app
SRC_DIR = ./cmd/api
BUILD_DIR = ./build
DOCS_DIR=docs
SWAGGER=github.com/swaggo/swag/cmd/swag

.PHONY: all
all: build

.PHONY: install
install:
	@echo "Installing dependencies..."
	$(GO) mod tidy

.PHONY: build
build: install
	@echo "Building the Go application..."
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)

.PHONY: run
run: build
	@echo "Running the Go application..."
	$(BUILD_DIR)/$(BINARY_NAME)

# Directly start the go application
.PHONY: start
start:
	@echo "Starting the Go application..."
	$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: clean
clean:
	@echo "Cleaning up build files..."
	rm -rf $(BUILD_DIR)

.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test ./test/...

.PHONY: docs
docs:
	@echo "Generating docs..."
	go run $(SWAGGER) init -g cmd/api/main.go -o $(DOCS_DIR)
