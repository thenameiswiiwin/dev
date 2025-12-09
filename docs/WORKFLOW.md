# Visual Workflow Guide

This guide shows how to use the Dev Environment Manager with visual diagrams.

## Complete Setup Flow

```
┌─────────────────────────────────────────────────────────────┐
│                    FRESH COMPUTER                           │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  Step 1: Install Prerequisites                              │
│  • Git                                                       │
│  • Go 1.24+                                                  │
│  • Package manager (brew/apt/pacman)                         │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  Step 2: Clone & Build                                       │
│  $ git clone <repo>                                          │
│  $ cd dev                                                    │
│  $ make build                                                │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  Step 3: Bootstrap System                                    │
│  $ dev bootstrap                                             │
│  ✓ Installs: git, zsh, tmux, neovim, ripgrep, etc.         │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  Step 4: Choose Your Path                                    │
│                                                              │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐   │
│  │  Python  │  │    Go    │  │   Rust   │  │   Web    │   │
│  └──────────┘  └──────────┘  └──────────┘  └──────────┘   │
└─────────────────────────────────────────────────────────────┘
           │            │            │            │
           ▼            ▼            ▼            ▼
┌─────────────────────────────────────────────────────────────┐
│  Step 5: Apply Preset                                        │
│  $ dev preset apply <python|go|rust|web>                    │
│  ✓ Installs language-specific tools                         │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  Step 6: Sync Dotfiles (Optional)                           │
│  $ dev dotfiles sync                                         │
│  ✓ Configures terminal and editor                           │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│  Step 7: Verify Installation                                 │
│  $ dev doctor                                                │
│  ✓ Checks everything is working                             │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                   READY TO CODE! 🚀                          │
└─────────────────────────────────────────────────────────────┘
```

## What Each Preset Installs

### Python Preset

```
dev preset apply python
         │
         ├──> Python 3.12 (Language)
         ├──> uv (Fast package installer)
         ├──> pipx (Tool installer)
         ├──> pyright (LSP for code completion)
         ├──> ruff (Linter & formatter)
         ├──> pytest (Testing framework)
         └──> debugpy (Debugger)
```

### Go Preset

```
dev preset apply go
         │
         ├──> Go 1.24 (Language)
         ├──> gopls (LSP for code completion)
         ├──> gofumpt (Code formatter)
         ├──> golangci-lint (Linter)
         └──> delve (Debugger)
```

### Rust Preset

```
dev preset apply rust
         │
         ├──> rustup (Toolchain manager)
         ├──> Rust stable (Language)
         ├──> rust-analyzer (LSP)
         ├──> rustfmt (Code formatter)
         ├──> clippy (Linter)
         └──> lldb (Debugger)
```

### Web Preset

```
dev preset apply web
         │
         ├──> Node.js 22 (Runtime)
         ├──> npm, pnpm, bun (Package managers)
         ├──> TypeScript (Language)
         ├──> ESLint (Linter)
         ├──> Prettier (Formatter)
         ├──> typescript-language-server (LSP)
         ├──> tailwindcss-language-server (LSP)
         └──> volar (Vue LSP)
```

## Command Flow

### Bootstrap Flow

```
$ dev bootstrap
      │
      ├─> Detect OS (macOS/Linux)
      │
      ├─> Check for package manager
      │   ├─> brew (macOS)
      │   ├─> apt (Debian/Ubuntu)
      │   └─> pacman/yay (Arch)
      │
      ├─> Install core tools
      │   ├─> git
      │   ├─> curl
      │   ├─> zsh
      │   ├─> tmux
      │   ├─> neovim
      │   ├─> ripgrep
      │   ├─> fd
      │   └─> fzf
      │
      └─> ✓ Bootstrap complete!
```

### Preset Apply Flow

```
$ dev preset apply python
      │
      ├─> Load manifest (manifests/presets/python.yaml)
      │
      ├─> Check what's already installed
      │
      ├─> Install missing packages
      │   ├─> python3 (via brew/apt/pacman)
      │   ├─> uv (via brew/apt/pacman)
      │   ├─> pipx (via brew/apt/pacman)
      │   ├─> pyright (via pipx)
      │   ├─> ruff (via pipx)
      │   └─> pytest (via pip)
      │
      ├─> Run post-install commands (if any)
      │
      └─> ✓ Preset installed!
           │
           └─> Next steps shown to user
```

### Dotfiles Sync Flow

```
$ dev dotfiles sync
      │
      ├─> Find dotfiles (env/ directory)
      │
      ├─> For each file/directory:
      │   │
      │   ├─> Check if destination exists
      │   │   │
      │   │   ├─> Yes: Create backup
      │   │   │   (file.backup-20231209-120000)
      │   │   │
      │   │   └─> No: Continue
      │   │
      │   └─> Create symlink
      │       (home → repo/env/file)
      │
      └─> ✓ Dotfiles synced!
```

## Project Structure Flow

```
dev/
├─ cmd/dev/              ──→  CLI commands
│  ├─ main.go            ──→  Entry point
│  ├─ bootstrap.go       ──→  Bootstrap command
│  ├─ preset.go          ──→  Preset commands
│  └─ ...
│
├─ internal/             ──→  Internal packages
│  ├─ osdetect/          ──→  Detect OS type
│  ├─ logger/            ──→  Logging utilities
│  ├─ config/            ──→  Load manifests
│  ├─ installer/         ──→  Install packages
│  └─ dotfiles/          ──→  Sync dotfiles
│
├─ manifests/            ──→  Package definitions
│  ├─ base.yaml          ──→  Core tools
│  └─ presets/           ──→  Language-specific
│     ├─ python.yaml
│     ├─ go.yaml
│     ├─ rust.yaml
│     └─ web.yaml
│
├─ containers/           ──→  Docker images
│  ├─ python/Dockerfile
│  ├─ go/Dockerfile
│  ├─ rust/Dockerfile
│  └─ web/Dockerfile
│
├─ env/                  ──→  Your dotfiles
│  ├─ .config/           ──→  App configs
│  ├─ .local/            ──→  Scripts
│  ├─ .zshrc             ──→  Shell config
│  └─ ...
│
└─ bin/                  ──→  Built binary
   └─ dev                ──→  The CLI tool
```

## How Manifests Work

```
manifests/presets/python.yaml
         │
         │  Define what to install
         ▼
┌─────────────────────────────────────┐
│ name: python                        │
│ packages:                           │
│   - name: python3                   │
│     brew: python@3.12               │
│     apt: python3.12                 │
│     pacman: python                  │
│   - name: uv                        │
│     brew: uv                        │
│     ...                             │
└─────────────────────────────────────┘
         │
         │  Loaded by CLI
         ▼
internal/config/manifest.go
         │
         │  Parsed into data structure
         ▼
internal/installer/installer.go
         │
         │  Checks OS type
         ▼
┌─────────────────────────────────────┐
│ If macOS → Use brew package name    │
│ If Ubuntu → Use apt package name    │
│ If Arch → Use pacman package name   │
└─────────────────────────────────────┘
         │
         │  Install command
         ▼
$ brew install python@3.12
$ apt-get install python3.12
$ pacman -S python
```

## Container Workflow (Advanced)

```
$ dev build python
      │
      ├─> Read Dockerfile (containers/python/Dockerfile)
      │
      ├─> Build image with Docker
      │   (Installs Python + tools inside container)
      │
      ├─> Tag as dev-env-python:latest
      │
      └─> Optionally run smoke test
          (Check if Python works in container)

$ docker run -it dev-env-python:latest
      │
      └─> Opens shell in container with all tools ready
```

## Kubernetes Workflow (Advanced)

```
$ dev k8s render python
      │
      ├─> Read manifests (deploy/python/)
      │   ├─ kustomization.yaml
      │   ├─ deployment.yaml
      │   ├─ service.yaml
      │   └─ configmap.yaml
      │
      ├─> Use kustomize to combine them
      │
      └─> Output complete K8s manifest

$ kubectl apply -k deploy/python/
      │
      ├─> Deploys to Kubernetes cluster
      │
      └─> Creates:
          ├─> Pods (running containers)
          ├─> Service (networking)
          └─> ConfigMap (configuration)
```

## Decision Tree: Which Command To Use?

```
                     START
                       │
                       ▼
              First time setup?
                    /  \
                  Yes   No
                  │      │
                  ▼      ▼
           dev bootstrap  Already have tools?
                          /  \
                        Yes   No
                        │     │
                        ▼     ▼
                  dev doctor  Need a language?
                              /  \
                            Yes   No
                            │     │
                            ▼     ▼
                  dev preset apply  Configure editor?
                                    /  \
                                  Yes   No
                                  │     │
                                  ▼     ▼
                          dev dotfiles  Check status?
                          sync          │
                                       ▼
                                  dev doctor
```

## Typical User Journey

### Beginner Developer

```
Day 1:  Install prerequisites (Git, Go)
        Clone repository
        Run: make build
        Run: dev bootstrap

Day 2:  Choose Python preset
        Run: dev preset apply python
        Write first "Hello World" program

Week 1: Learn Python basics
        Use installed tools (python3, pytest)

Week 2: Customize editor
        Run: dev dotfiles sync
        Modify configs in env/

Month 1: Try another language
         Run: dev preset apply web
         Build first website
```

### Experienced Developer

```
Day 1:  git clone && make build
        dev bootstrap
        dev preset apply go
        dev preset apply python
        dev preset apply web
        dev dotfiles sync

Day 2:  Start coding immediately
        All tools ready to go!

Later:  dev build (create containers)
        dev k8s render (deploy to K8s)
```

## Summary Diagram

```
┌──────────────────────────────────────────────────────────────┐
│                    Dev Environment Manager                    │
│                                                               │
│  Commands           What They Do                              │
│  ─────────          ────────────                              │
│  bootstrap     →    Install base tools                        │
│  preset        →    Install language tools                    │
│  dotfiles      →    Sync configurations                       │
│  build         →    Create containers                         │
│  k8s           →    Deploy to Kubernetes                      │
│  doctor        →    Check health                              │
│                                                               │
│  Presets: Python | Go | Rust | Web                           │
│  Platforms: macOS | Linux (Ubuntu, Debian, Arch)             │
│                                                               │
└──────────────────────────────────────────────────────────────┘
```

---

**See Also:**

- [Beginner's Guide](BEGINNER_GUIDE.md) - Detailed explanations
- [Quick Start](QUICK_START.md) - Fast setup
- [Cheat Sheet](CHEAT_SHEET.md) - Command reference
