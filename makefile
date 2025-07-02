include .env

# Constants
NAME := Skene
APPID := com.example.skene
ICON := assets/amphi.png

# Directories
BIN_DIR := bin
GUI_DIR := dist  # Changed from bin to dist for GUI apps

# Colors
NO_COLOR := \033[0m
OK_COLOR := \033[32;01m
INFO_COLOR := \033[36;01m
WARN_COLOR := \033[33;01m
ERROR_COLOR := \033[31;01m

# Help
.PHONY: help
help: ## Show this help
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  $(INFO_COLOR)%-20s$(NO_COLOR) %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development
.PHONY: run
run: ## Run the command line executable
	@printf "$(INFO_COLOR)→ Running the command line executable...$(NO_COLOR)\n"
	@go run cmd/skene/main.go

.PHONY: lint
lint: ## Run linter
	@printf "$(INFO_COLOR)→ Running linter...$(NO_COLOR)\n"
	@golangci-lint run --config .golangci.yaml

.PHONY: test
test: ## Run tests
	@printf "$(INFO_COLOR)→ Running tests...$(NO_COLOR)\n"
	@go test -race -cover ./...

# Packaging
.PHONY: package
package: ## Package the GUI application
	@mkdir -p $(GUI_DIR)
	@printf "$(INFO_COLOR)→ Packaging GUI application...$(NO_COLOR)\n"
	@cd $(GUI_DIR) && fyne package \
		-name "$(NAME)" \
		--app-id "$(APPID)" \
		-icon "$(abspath $(ICON))" \
		-source-dir "$(abspath cmd/skene)" \
		-release

# Package template
define package_target
.PHONY: package-$(1)-$(2)
package-$(1)-$(2): ## Package GUI application for $(1) $(2)
	@mkdir -p $(GUI_DIR)
	@printf "$(INFO_COLOR)→ Packaging for $(1) $(2)...$(NO_COLOR)\n"
	@cd $(GUI_DIR) && fyne package \
		-name "$(NAME)-$(1)-$(2)" \
		--app-id "$(APPID)" \
		-icon "$(abspath $(ICON))" \
		-source-dir "$(abspath cmd/skene)" \
		--target $(1)$(if $(filter-out amd64,$(2))/$(2),) \
		-release
endef

# Define all package targets
$(eval $(call package_target,darwin,amd64))
$(eval $(call package_target,darwin,arm64))
$(eval $(call package_target,linux,amd64))
$(eval $(call package_target,linux,arm64))
$(eval $(call package_target,windows,amd64))
$(eval $(call package_target,windows,arm64))

# Package for all platforms
.PHONY: package-all
package-all: ## Package the GUI application for all platforms
package-all: package-darwin-amd64 package-darwin-arm64 package-linux-amd64 \
             package-linux-arm64 package-windows-amd64 package-windows-arm64
	@printf "$(OK_COLOR)→ Packaged GUI application for all platforms$(NO_COLOR)\n"

# Build targets
.PHONY: build
build: ## Build for current system
	@$(eval OS := $(shell go env GOOS))
	@$(eval ARCH := $(shell go env GOARCH))
	@$(eval EXT := $(if $(filter windows,$(OS)),.exe,))
	@printf "$(INFO_COLOR)→ Building for $(OS) $(ARCH)...$(NO_COLOR)\n"
	@GOOS=$(OS) GOARCH=$(ARCH) go build -o $(BIN_DIR)/skene-$(OS)-$(ARCH)$(EXT) cmd/skene/main.go

# Build templates
define build_target
.PHONY: build-$(1)-$(2)
build-$(1)-$(2): ## Build for $(1) $(2)
	@printf "$(INFO_COLOR)→ Building for $(1) $(2)...$(NO_COLOR)\n"
	@GOOS=$(1) GOARCH=$(2) go build -o $(BIN_DIR)/skene-$(1)-$(2)$(if $(filter windows,$(1)),.exe,) cmd/skene/main.go
endef

# Static build template
define build_static_target
.PHONY: build-$(1)-$(2)-static
build-$(1)-$(2)-static: ## Build static binary for $(1) $(2)
	@printf "$(INFO_COLOR)→ Building static binary for $(1) $(2)...$(NO_COLOR)\n"
	@CGO_ENABLED=0 GOOS=$(1) GOARCH=$(2) go build -ldflags "-s -w" -o $(BIN_DIR)/skene-$(1)-$(2)-static$(if $(filter windows,$(1)),.exe,) cmd/skene/main.go
endef

# Define all build targets
$(eval $(call build_target,darwin,amd64))
$(eval $(call build_target,darwin,arm64))
$(eval $(call build_target,linux,amd64))
$(eval $(call build_target,linux,arm64))
$(eval $(call build_static_target,linux,amd64))
$(eval $(call build_target,windows,amd64))
$(eval $(call build_target,windows,arm64))
$(eval $(call build_target,windows,386))
$(eval $(call build_target,linux,386))

# Meta targets
.PHONY: build-all
build-all: ## Build for all platforms
build-all: build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-linux-arm64 \
           build-linux-amd64-static build-windows-amd64 build-windows-arm64 \
           build-windows-386 build-linux-386
	@printf "$(OK_COLOR)→ Built for all platforms$(NO_COLOR)\n"

.PHONY: clean
clean: ## Clean build artifacts
	@printf "$(INFO_COLOR)→ Cleaning binaries...$(NO_COLOR)\n"
	@rm -rf $(BIN_DIR)/*
	@printf "$(INFO_COLOR)→ Cleaning packages...$(NO_COLOR)\n"
	@rm -rf $(GUI_DIR)/*