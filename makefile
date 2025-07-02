include .env

# Run the client
run: ## Run the command line executable
	@echo "\033[2m→ Running the command line executable...\033[0m"
	@go run cmd/skene/main.go

# Lint
lint:
	@echo "\033[2m→ Running linter...\033[0m"
	@golangci-lint run --config .golangci.yaml

# Tests
test:
	@echo "\033[2m→ Running tests...\033[0m"
	@go test ./...

# Builds
build: ## Build para o sistema atual
	@echo "\033[2m→ Detectando sistema...\033[0m"
	@OS=$$(go env GOOS); \
	ARCH=$$(go env GOARCH); \
	OUT="bin/skene-$${OS}-$${ARCH}"; \
	EXT=""; \
	if [ "$${OS}" = "windows" ]; then EXT=".exe"; fi; \
	echo "\033[2m→ Building for $${OS} $${ARCH}...\033[0m"; \
	GOOS=$${OS} GOARCH=$${ARCH} go build -o "$${OUT}$${EXT}" cmd/skene/main.go

build-darwin-amd64: ## Build for macOS Intel
	@echo "\033[2m→ Building for macOS Intel...\033[0m"
	@GOOS=darwin GOARCH=amd64 go build -o bin/skene-darwin-amd64 cmd/skene/main.go

build-darwin-arm64: ## Build for macOS Apple Silicon
	@echo "\033[2m→ Building for macOS Apple Silicon...\033[0m"
	@GOOS=darwin GOARCH=arm64 go build -o bin/skene-darwin-arm64 cmd/skene/main.go

build-linux-amd64: ## Build for Linux amd64
	@echo "\033[2m→ Building for Linux amd64...\033[0m"
	@GOOS=linux GOARCH=amd64 go build -o bin/skene-linux-amd64 cmd/skene/main.go

build-linux-arm64: ## Build for Linux ARM64
	@echo "\033[2m→ Building for Linux ARM64...\033[0m"
	@GOOS=linux GOARCH=arm64 go build -o bin/skene-linux-arm64 cmd/skene/main.go

build-linux-amd64-static: ## Build static Linux amd64
	@echo "\033[2m→ Building static Linux amd64...\033[0m"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/skene-linux-amd64-static cmd/skene/main.go

build-windows-amd64: ## Build for Windows amd64
	@echo "\033[2m→ Building for Windows amd64...\033[0m"
	@GOOS=windows GOARCH=amd64 go build -o bin/skene-windows-amd64.exe cmd/skene/main.go

build-windows-arm64: ## Build for Windows ARM64
	@echo "\033[2m→ Building for Windows ARM64...\033[0m"
	@GOOS=windows GOARCH=arm64 go build -o bin/skene-windows-arm64.exe cmd/skene/main.go

build-windows-386: ## Build for Windows 386 (32-bit)
	@echo "\033[2m→ Building for Windows 386...\033[0m"
	@GOOS=windows GOARCH=386 go build -o bin/skene-windows-386.exe cmd/skene/main.go

build-linux-386: ## Build for Linux 386 (32-bit)
	@echo "\033[2m→ Building for Linux 386...\033[0m"
	@GOOS=linux GOARCH=386 go build -o bin/skene-linux-386 cmd/skene/main.go

# Build all targets
build-all: build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-linux-arm64 build-linux-amd64-static build-windows-amd64 build-windows-arm64 build-windows-386 build-linux-386
	@echo "\033[2m→ Built for all platforms\033[0m"

# Clean binaries
clean:
	@echo "\033[2m→ Cleaning binaries...\033[0m"
	@rm -rf bin/*

.PHONY: run lint test clean build-all build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-linux-arm64 build-linux-amd64-static build-windows-amd64 build-windows-arm64 build-windows-386 build-linux-386
