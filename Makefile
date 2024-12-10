APP := family-web

ifeq ($(OS), Windows_NT)
    RM = del /F
    GOOS = windows
	APP_NAME = $(APP).exe
else
    RM = rm -f
	GOOS = linux
	APP_NAME = $(APP)
endif

all: re-build

.PHONY: re-build
re-build:
	@echo "Cleaning up..."
	@if [ -f "$(APP_NAME)" ]; then \
		echo "Removing $(APP_NAME)..."; \
		$(RM) -f $(APP_NAME); \
	else \
		echo "$(APP_NAME) does not exist, skipping removal."; \
	fi
	@echo "Building $(APP_NAME)..."
	go env -w GOOS=$(GOOS)
	go build -o $(APP_NAME) .


.PHONY: test
test:
	@echo "Running tests with coverage..."
	go test -cover ./test/... --coverprofile=.coverage
	go tool cover -func=coverage -o .coverage

swag:
	@echo "Generating swagger docs..."
	swag init

help:
	@echo "Available targets:"

