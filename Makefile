BINARY_NAME=tailwindcss-palette
VERSION?=$(shell grep 'Version *=' ./internal/version/version.go | sed -E 's/.*"([^"]+)".*/\1/')
BUILD_DIR=build
MAIN_FILE=cmd/tailwindcss-palette/tailwindcss-palette.go
GIT_COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
LDFLAGS=-ldflags "-X github.com/claytonchew/tailwindcss-palette-go/internal/version.Version=$(VERSION) -X github.com/claytonchew/tailwindcss-palette-go/internal/version.CommitHash=$(GIT_COMMIT) -X github.com/claytonchew/tailwindcss-palette-go/internal/version.BuildDate=$(BUILD_DATE)"

PLATFORMS=linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64

all: clean test build

build:
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

deps:
	$(GOGET) -v -d ./...

# Build for all platforms
cross-platform:
	mkdir -p $(BUILD_DIR)
	$(eval TEMP_DIR := $(shell mktemp -d))

	@for platform in $(PLATFORMS); do \
		echo "Building for $$platform..."; \
		GOOS=$$(echo $$platform | cut -d/ -f1); \
		GOARCH=$$(echo $$platform | cut -d/ -f2); \
		OUTPUT_NAME=$(BINARY_NAME); \
		if [ $$GOOS = "windows" ]; then \
			OUTPUT_NAME+='.exe'; \
		fi; \
		GOOS=$$GOOS GOARCH=$$GOARCH $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_$${GOOS}_$${GOARCH}$${EXT} $(MAIN_FILE) || exit 1; \
	done; \
	rm -rf $(TEMP_DIR)

.PHONY: all build test clean run deps cross-platform
