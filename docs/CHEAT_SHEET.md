# Command Cheat Sheet

Quick reference for all commands in the Dev Environment Manager.

## 📦 Installation

```bash
# Clone repository
git clone https://github.com/thenameiswiiwin/dev.git ~/workspace/github.com/thenameiswiiwin/dev

# Go to directory
cd ~/workspace/github.com/thenameiswiiwin/dev

# Build CLI
make build

# Add to PATH (optional, makes commands shorter)
echo 'export PATH="$HOME/workspace/github.com/thenameiswiiwin/dev/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

## 🎮 Basic Commands

| Command | What It Does | When To Use |
|---------|-------------|-------------|
| `dev bootstrap` | Install essential tools | First time setup |
| `dev preset list` | Show available presets | See what you can install |
| `dev preset apply <preset>` | Install language tools | Setup for Python, Go, etc. |
| `dev dotfiles sync` | Sync your settings | Configure terminal/editor |
| `dev doctor` | Check installation | Troubleshooting |
| `dev build` | Build containers | Advanced: Docker setup |
| `dev k8s render <preset>` | Show K8s manifests | Advanced: Kubernetes |

## 🚩 Common Flags

| Flag | What It Does | Example |
|------|-------------|---------|
| `--dry-run` | Preview without changes | `dev preset apply python --dry-run` |
| `--verbose` or `-v` | Show detailed output | `dev bootstrap --verbose` |
| `--help` or `-h` | Show help | `dev preset --help` |

## 🐍 Python Setup

```bash
# Complete Python setup
dev bootstrap                    # Install base tools
dev preset apply python          # Install Python tools
dev dotfiles sync               # Sync settings
dev doctor                      # Verify

# Test Python
python3 --version
pip --version
```

## 🔷 Go Setup

```bash
# Complete Go setup
dev bootstrap                    # Install base tools
dev preset apply go              # Install Go tools
dev dotfiles sync               # Sync settings
dev doctor                      # Verify

# Test Go
go version
```

## 🦀 Rust Setup

```bash
# Complete Rust setup
dev bootstrap                    # Install base tools
dev preset apply rust            # Install Rust tools
dev dotfiles sync               # Sync settings
dev doctor                      # Verify

# Test Rust
rustc --version
cargo --version
```

## 🌐 Web Development Setup

```bash
# Complete Web setup
dev bootstrap                    # Install base tools
dev preset apply web             # Install Node.js tools
dev dotfiles sync               # Sync settings
dev doctor                      # Verify

# Test Node
node --version
npm --version
pnpm --version
```

## 🔧 Development Commands

```bash
# Build the CLI
make build

# Run tests
make test

# Lint code
make lint

# Format code
make fmt

# Clean build artifacts
make clean

# Build and install globally
make install
```

## 🐳 Container Commands

```bash
# Build all containers
dev build

# Build specific preset
dev build python

# Build with tests
dev build python --test

# Build with custom tag
dev build python --tag v1.0

# Run a container
docker run -it dev-env-python:latest
```

## ☸️ Kubernetes Commands

```bash
# Render manifests
dev k8s render python

# Apply to cluster (after rendering)
kubectl apply -k deploy/python/

# Check deployment
kubectl get pods
kubectl get services

# Port forward
kubectl port-forward svc/dev-env-python 8000:8000
```

## 📁 File Locations

| What | Where |
|------|-------|
| Built CLI | `bin/dev` |
| Source code | `cmd/dev/` |
| Manifests | `manifests/` |
| Dotfiles | `env/` |
| Containers | `containers/` |
| K8s configs | `deploy/` |
| Documentation | `docs/` |

## 🔍 Troubleshooting Commands

```bash
# Check system status
dev doctor

# Preview what would change
dev preset apply python --dry-run

# See detailed output
dev preset apply python --verbose

# Get help
dev --help
dev preset --help
dev preset apply --help

# Check versions
python3 --version
go version
node --version
rustc --version
```

## 💡 Pro Tips

### 1. Always Preview First
```bash
dev preset apply python --dry-run
```

### 2. Use Verbose for Debugging
```bash
dev bootstrap --verbose
```

### 3. Install Multiple Presets
```bash
dev preset apply python
dev preset apply go
dev preset apply web
```

### 4. Check Before Installing
```bash
dev preset list
dev doctor
```

### 5. Backup Before Syncing
```bash
# Dry-run shows what will be backed up
dev dotfiles sync --dry-run
```

## 🆘 Common Issues

### "Command not found: dev"
```bash
# Use full path
./bin/dev --help

# Or add to PATH
echo 'export PATH="$HOME/workspace/github.com/thenameiswiiwin/dev/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

### "Permission denied"
```bash
chmod +x ./bin/dev
```

### "Not in git repository"
```bash
cd ~/workspace/github.com/thenameiswiiwin/dev
```

### Installation failed
```bash
# Check what's missing
dev doctor

# Try verbose mode
dev preset apply python --verbose
```

## 📚 Quick Links

- [Beginner's Guide](BEGINNER_GUIDE.md) - Detailed tutorial
- [Quick Start](QUICK_START.md) - Fast setup
- [README](../README.md) - Main documentation

## 🎯 Common Workflows

### First Time Setup
```bash
1. git clone <repo> && cd dev
2. make build
3. dev bootstrap
4. dev preset apply python
5. dev dotfiles sync
6. dev doctor
```

### Add New Language
```bash
1. dev preset list
2. dev preset apply <preset> --dry-run
3. dev preset apply <preset>
4. dev doctor
```

### Daily Development
```bash
# Already set up, just use your tools!
python3 my_script.py
go run main.go
cargo run
npm run dev
```

### Update Everything
```bash
cd ~/workspace/github.com/thenameiswiiwin/dev
git pull
make build
dev doctor
```

---

**Print this page** and keep it handy while learning! 📄
