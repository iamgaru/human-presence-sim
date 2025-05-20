
BINARY_NAME=simulator
BUILD_DIR=bin
INSTALL_PATH=/usr/local/bin

.PHONY: all build install test race coverage clean help

all: build

build:
	@echo "🔨 Building $(BINARY_NAME)..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/simulator

install: build
	@echo "📦 Installing to $(INSTALL_PATH)..."
	sudo cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_PATH)/$(BINARY_NAME)

test:
	@echo "🧪 Running unit tests..."
	go test ./...

race:
	@echo "🏁 Running tests with race detection..."
	go test -race ./...

coverage:
	@echo "📈 Generating test coverage report..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clean:
	@echo "🧹 Cleaning up..."
	rm -rf $(BUILD_DIR) coverage.out

help:
	@echo ""
	@echo "Available targets:"
	@echo "  build     - Compile the CLI binary"
	@echo "  install   - Copy binary to $(INSTALL_PATH)"
	@echo "  test      - Run unit tests"
	@echo "  race      - Run tests with race detection"
	@echo "  coverage  - Run tests and output HTML coverage report"
	@echo "  clean     - Remove build and coverage artifacts"
	@echo "  help      - Show this message"
	@echo ""
