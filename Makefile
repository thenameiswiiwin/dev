.PHONY: help build test lint fmt clean install

# Default target
help:
	@echo "Available targets:"
	@echo "  build     - Build the dev CLI"
	@echo "  test      - Run Go tests"
	@echo "  lint      - Run all linters (Go, shell scripts, manifests)"
	@echo "  fmt       - Format code (gofmt, shfmt)"
	@echo "  clean     - Clean build artifacts"
	@echo "  install   - Install the dev CLI to /usr/local/bin"

# Build the CLI
build:
	@echo "Building dev CLI..."
	go build -o bin/dev ./cmd/dev
	@echo "✓ Built bin/dev"

# Run tests
test:
	@echo "Running Go tests..."
	go test -v ./...

# Run all linters
lint: lint-go lint-shell lint-manifests

# Lint Go code
lint-go:
	@echo "Running Go linters..."
	@echo "Checking gofmt..."
	@if [ -n "$$(gofmt -l .)" ]; then \
		echo "Go code is not formatted. Run 'make fmt' to fix:"; \
		gofmt -l .; \
		exit 1; \
	fi
	@echo "Running go vet..."
	go vet ./...
	@echo "✓ Go code looks good"

# Lint shell scripts
lint-shell:
	@echo "Running shell script linters..."
	@command -v shellcheck >/dev/null 2>&1 || { echo "shellcheck not installed. Install with: brew install shellcheck"; exit 1; }
	@command -v shfmt >/dev/null 2>&1 || { echo "shfmt not installed. Install with: brew install shfmt"; exit 1; }
	@echo "Running shellcheck..."
	@find . -name "*.sh" -o -path "./runs/*" -type f -executable | while read -r file; do \
		if file "$$file" | grep -q "shell script"; then \
			echo "  Checking $$file"; \
			shellcheck "$$file" || exit 1; \
		fi; \
	done
	@echo "✓ Shell scripts look good"

# Lint manifests
lint-manifests:
	@echo "Validating manifest YAML files..."
	@if command -v yamllint >/dev/null 2>&1; then \
		yamllint manifests/; \
	elif python3 -c "import yaml" 2>/dev/null; then \
		for file in manifests/*.yaml manifests/presets/*.yaml; do \
			if [ -f "$$file" ]; then \
				echo "  Validating $$file"; \
				python3 -c "import yaml; yaml.safe_load(open('$$file'))" || exit 1; \
			fi; \
		done; \
	else \
		echo "⚠ Neither yamllint nor PyYAML found, skipping YAML validation"; \
		echo "  Install with: pip install pyyaml or brew install yamllint"; \
	fi
	@echo "✓ Manifests check complete"

# Format code
fmt:
	@echo "Formatting Go code..."
	gofmt -w .
	@echo "Formatting shell scripts..."
	@command -v shfmt >/dev/null 2>&1 && find . -name "*.sh" -type f -exec shfmt -w {} \; || echo "shfmt not installed, skipping"
	@echo "✓ Code formatted"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	@echo "✓ Clean complete"

# Install CLI
install: build
	@echo "Installing dev CLI to /usr/local/bin..."
	sudo cp bin/dev /usr/local/bin/dev
	@echo "✓ Installed /usr/local/bin/dev"
