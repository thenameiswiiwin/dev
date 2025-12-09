# Dev Environment Manager

A comprehensive development environment management system with support for multiple language presets, containers, devcontainers, and Kubernetes deployments.

## 👋 New to Coding?

**Start here:** [Beginner's Guide](docs/BEGINNER_GUIDE.md) - Detailed step-by-step instructions with explanations

**Just want to get started?** [Quick Start Guide](docs/QUICK_START.md) - 5-minute setup

## Features

- 🚀 **Preset-based system** - Python, Go, Rust, Web (React/Next.js/TypeScript)
- 🔧 **CLI tool** - Manage everything from the command line
- 🐳 **Container support** - Docker images for each preset
- 📦 **Devcontainer integration** - VS Code and DevPod compatible
- ☸️ **Kubernetes ready** - Deploy presets to K8s clusters
- 🔄 **Cross-platform** - macOS, Arch Linux, Debian, Ubuntu
- 🛡️ **Idempotent operations** - Safe to run multiple times
- 🧪 **Dry-run mode** - Test before applying changes

## Quick Start

### Prerequisites

- Git
- One of: Homebrew (macOS/Linux), APT (Debian/Ubuntu), or Pacman+Yay (Arch)
- Go 1.24+ (for building the CLI)

### Installation

```bash
# Clone the repository
git clone https://github.com/thenameiswiiwin/dev.git
cd dev

# Build the CLI
make build

# Bootstrap your system
./bin/dev bootstrap

# Apply a preset
./bin/dev preset apply python

# Sync dotfiles
./bin/dev dotfiles sync
```

## Available Commands

```bash
dev bootstrap        # Install core dependencies
dev preset list      # Show available presets
dev preset apply     # Install a language preset
dev dotfiles sync    # Sync dotfiles with symlinks
dev build           # Build container images
dev k8s render      # Render Kubernetes manifests
dev doctor          # Check system health
```

## Presets

### Python

- Python 3.12, uv, pipx
- LSP: pyright, ruff
- Testing: pytest
- Debugging: debugpy

### Go

- Go 1.24
- LSP: gopls
- Formatting: gofumpt
- Linting: golangci-lint
- Debugging: delve

### Rust

- Rust (latest stable via rustup)
- LSP: rust-analyzer
- Formatting: rustfmt
- Linting: clippy

### Web

- Node.js 22, pnpm, bun
- TypeScript, ESLint, Prettier
- Tailwind CSS support
- LSP: typescript-language-server, volar

## Project Structure

```
dev/
├── cmd/dev/              # CLI source code
├── internal/             # Internal packages
│   ├── osdetect/        # OS detection
│   ├── logger/          # Logging utilities
│   ├── config/          # Manifest loading
│   ├── installer/       # Package installation
│   └── dotfiles/        # Dotfile syncing
├── manifests/            # Package manifests
│   ├── base.yaml
│   └── presets/
├── containers/           # Docker images
├── .devcontainer/       # Devcontainer configs
├── deploy/              # Kubernetes manifests
├── env/                 # Dotfiles
├── docs/                # Documentation
└── tmux-sessionizer/    # Tmux session manager
```

## Usage Examples

### Host Installation

```bash
# Install Python development environment
dev preset apply python

# Install Go development environment
dev preset apply go --verbose

# Test what would be installed (dry-run)
dev preset apply rust --dry-run
```

### Container Usage

```bash
# Build all container images
dev build

# Build specific preset with smoke tests
dev build python --test

# Run a container
docker run -it dev-env-python:latest
```

### Devcontainer Usage

```bash
# With DevPod CLI
devpod up . --devcontainer-path .devcontainer/devcontainer.python.json

# With VS Code
# Open folder → Reopen in Container → Select devcontainer config
```

### Kubernetes Deployment

```bash
# Render manifests
dev k8s render python

# Apply to cluster
kubectl apply -k deploy/python/

# Access the service
kubectl port-forward svc/dev-env-python 8000:8000
```

## Development

### Building

```bash
make build      # Build the CLI
make test       # Run tests
make lint       # Run linters
make fmt        # Format code
```

### CI/CD

GitHub Actions workflows automatically:

- Lint Go code and shell scripts
- Run tests
- Build and test all container images
- Validate Kubernetes manifests

## Documentation

### Getting Started

- 🎓 [Beginner's Guide](docs/BEGINNER_GUIDE.md) - Complete tutorial for newcomers
- ⚡ [Quick Start Guide](docs/QUICK_START.md) - Fast 5-minute setup

### Technical Documentation

- [Implementation Summary](docs/IMPLEMENTATION_SUMMARY.md) - Complete overview
- [Blueprint](docs/BLUEPRINT.md) - Architecture and design
- [Version Pinning](docs/VERSION_PINNING.md) - Security and versions
- [Reorganization](docs/REORGANIZATION.md) - Cleanup and restructuring
- [Agents Guide](AGENTS.md) - Development guidelines

## Philosophy

This project follows these principles:

- **Container-first**: Same images for development and deployment
- **Data-driven**: Manifests define what gets installed
- **Idempotent**: Safe to run commands multiple times
- **Cross-platform**: Works on multiple operating systems
- **Well-tested**: Automated testing and validation
- **Documented**: Clear guides for every feature

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run `make lint` and `make test`
5. Submit a pull request

## License

See LICENSE file for details.

## Support

- Documentation: [docs/](docs/)
- Issues: https://github.com/thenameiswiiwin/dev/issues
