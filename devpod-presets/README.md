# DevPod Presets

This directory contains DevPod workspace templates for each development preset.

## Usage

### Using DevPod CLI

```bash
# Start a Python development workspace
devpod up . --devcontainer-path .devcontainer/devcontainer.python.json

# Start a Go development workspace
devpod up . --devcontainer-path .devcontainer/devcontainer.go.json

# Start a Rust development workspace
devpod up . --devcontainer-path .devcontainer/devcontainer.rust.json

# Start a Web development workspace
devpod up . --devcontainer-path .devcontainer/devcontainer.web.json
```

### Using DevPod Desktop

1. Open DevPod Desktop
2. Click "New Workspace"
3. Select "From Git Repository"
4. Enter the repository URL
5. Choose the appropriate devcontainer configuration:
   - `.devcontainer/devcontainer.python.json` for Python
   - `.devcontainer/devcontainer.go.json` for Go
   - `.devcontainer/devcontainer.rust.json` for Rust
   - `.devcontainer/devcontainer.web.json` for Web

## Available Presets

### Python
- Python 3.12
- uv, pipx, pip
- pyright, ruff, black
- pytest, debugpy
- Exposed ports: 8000 (app), 5678 (debug)

### Go
- Go 1.24
- gopls, gofumpt, golangci-lint
- delve debugger
- Exposed ports: 8080 (app), 2345 (debug)

### Rust
- Rust (latest stable)
- rust-analyzer, clippy, rustfmt
- lldb debugger
- Exposed ports: 8080 (app), 5000 (debug)

### Web
- Node.js 22
- TypeScript, ESLint, Prettier
- Tailwind CSS, Vite
- pnpm, bun
- Exposed ports: 3000 (Next.js/React), 5173 (Vite), 9229 (debug)

## Features

All presets include:
- zsh shell with tmux
- Neovim with LazyVim
- Modern CLI tools (ripgrep, fd, fzf, bat)
- Git integration
- PostgreSQL client

## Building Images Locally

To build the images locally before using them:

```bash
# Build all presets
dev build

# Build specific preset
dev build python

# Build and test
dev build --test
```

## Customization

To customize a preset:

1. Copy the desired devcontainer config
2. Modify the settings as needed
3. Point DevPod to your custom config

Example:
```bash
cp .devcontainer/devcontainer.python.json .devcontainer/devcontainer.custom.json
# Edit .devcontainer/devcontainer.custom.json
devpod up . --devcontainer-path .devcontainer/devcontainer.custom.json
```
