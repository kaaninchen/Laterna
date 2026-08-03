MAIN_PKG :=./cmd/server
BIN_NAME ?= laterna
BIN_DIR := bin
CONFIG_FILE := config.json

.DEFAULT_GOAL := build

build:
	@mkdir -p $(BIN_DIR)
	@echo "🔧 Building binary..."
	@cp $(CONFIG_FILE) $(BIN_DIR)/
	@go build -o $(BIN_DIR)/$(BIN_NAME) $(MAIN_PKG)
	@echo "Build complete: $(BIN_DIR)/$(BIN_NAME)"
	
run:
	@echo "Running server..."
	@go run $(MAIN_PKG)
clean:
	@echo "🧹 Cleaning up..."
	@rm -rf $(BIN_DIR) logs/
	@go mod tidy	
help:
	@echo "Makefile targets:"
	@echo "  build     – Build the binary"
	@echo "  run       – Run the development server"
	@echo "  clean     – Remove build artifacts and tidy modules"
	@echo "  help      – Show this help message"