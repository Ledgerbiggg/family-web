APP := chat-bot

ifeq ($(OS), Windows_NT)
    RM = del /F
    GOOS = windows
	APP_NAME = $(APP).exe
else
    RM = rm -f
	GOOS = linux
	APP_NAME = $(APP)
endif

all: build

.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	go env -w GOOS=$(GOOS)
	go build -o $(APP_NAME) .

.PHONY: test
test:
	@echo "Running tests with coverage..."
	go test -cover ./test/... --coverprofile=coverage
	go tool cover -func=coverage -o coverage

.PHONY: test-services
test-services:
	@echo "Running tests for services..."
	go test ./test -run 'TestService'

.PHONY: clean
clean:
	@echo "Cleaning up..."
	$(RM) -f $(APP_NAME)

help:
	@echo "Available targets:"
	@echo "  make build              Build the project"
	@echo "  make test               Run all tests"
	@echo "  make test-with-coverage Run tests with coverage report"
	@echo "  make test-services      Run tests for a specific service (use SERVICE=<ServiceName>)"
	@echo "  make clean              Clean the build artifacts"
