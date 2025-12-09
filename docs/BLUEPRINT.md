# Section 1: Architecture Overview
- Layers
  - Host bootstrap: detect OS, install Homebrew (if missing), ensure core tools (git, curl, zsh, tmux, minimal Go/Bash CLI).
  - Preset installers: data-driven manifests per preset (python, go, rust, web) that install language toolchains + LSP/formatter/debugger + editor integration; reusable across host and containers.
  - Containers: one Dockerfile per preset; build from minimal base, install manifest-defined tools, expose common ports (app + debug).
  - Devcontainers: devcontainer.json templates referencing preset images, volume mounts, postCreate hooks to sync dotfiles and install optional extras.
  - DevPod: consumes devcontainer.json directly; presets become workspaces with matching tools.
  - Kubernetes (optional): thin deploy/<preset>/ overlays for running preset images; health checks and env (PORT, DATABASE_URL).
- Flow ASCII
```
git clone dev-env
      |
      v
  ./dev bootstrap (OS detect -> install brew/pacman+yay -> core deps)
      |
      v
  ./dev preset list/select (python|go|rust|web)
      |
      v
  Choice A: host install   Choice B: container/devcontainer
        |                          |
 ./dev preset apply            devcontainer.json -> DevPod/VS Code
        |                          |
        v                          v
   zsh/tmux/LazyVim ready   Isolated preset container with tooling
```

# Section 2: Preset Definitions
- Common to all
  - Shell/editor: zsh, tmux + tmux-sessionizer, Neovim with LazyVim base; per-preset plugins/LSPs auto-enabled; ripgrep, fd, fzf, git.
  - DB: psql client.
  - Debug tools: delve (Go), lldb/gdb (others), netcat/htop/btop minimal set.
- Python
  - Toolchain: uv or pyenv + python3.x, pipx.
  - LSP/format/lint: pyright, ruff (fmt+lint), black (optional), debugpy.
  - Test: pytest.
  - Nvim: ensure mason installs pyright/ruff/debugpy; preset snippet loads python keymaps/debug profile.
- Go
  - Toolchain: Go toolchain (version pinned per manifest).
  - LSP/format/lint: gopls, gofumpt, golangci-lint; dlv for debug.
  - Test: go test.
  - Nvim: mason ensures gopls/dlv; gofmt/gofumpt on save.
- Rust
  - Toolchain: rustup, minimal profile.
  - LSP/format/lint: rust-analyzer, rustfmt, clippy; lldb/gdb optional.
  - Test: cargo test.
  - Nvim: mason ensures rust-analyzer; formatter hook uses rustfmt.
- Web (React/Next.js/TS/Tailwind)
  - Toolchain: node via volta or fnm; npm/yarn/pnpm; bun optional for scripts; turbo optional.
  - LSP/format/lint: typescript-language-server, eslint, prettier, stylelint (optional), vscode-js-debug; tailwindcss-language-server.
  - Test: vitest/jest; playwright optional toggle.
  - Nvim: mason ensures tsserver/volar/tailwindcss/eslint_d/prettierd; format on save with prettierd.
- Zsh/tmux deltas
  - Minimal preset-specific aliases (e.g., `ptest` -> pytest -q, `gtest` -> go test ./...).
  - tmux-sessionizer config aware of preset workspace roots; optional session commands per preset.

# Section 3: Cross-Platform Install Rules
- Order of preference: Homebrew (macOS/Linux) -> pacman + yay (Arch) if brew absent.
- OS detection lives in a single bootstrap script/CLI command (`./dev bootstrap`) that exports normalized facts (OS=macOS|arch, HAS_BREW=true/false).
- Install logic: use manifest keyed by package manager; each step is idempotent (check installed version before install; no blind rm -rf). Support `--dry-run`.
- Never mix package managers; per-runner selects exactly one backend. For Arch with brew installed, use brew; otherwise pacman+yay.
- Dotfiles: use symlink-based sync (or stow-like helper) with backup of existing files; no wholesale deletion of config dirs.

# Section 4: Containers / Devcontainers / DevPod
- Dockerfiles: `containers/<preset>/Dockerfile` built from slim base (debian-slim or archlinux for parity), install base shell/tmux/neovim + preset manifest tools; optional `ARG PRESET_VERSION`.
- Devcontainer: root `devcontainer.json` with `"name": "dev-env"`, `"features": {}` and `"customizations"`; `"variant"` or `"image"` points to preset image; workspace mount to `/workspace`. Provide separate `devcontainer.<preset>.json` or a single parametrized template using `dockerComposeFile`/`runArgs` for port sets.
- Post-create: hook to run `./dev dotfiles sync --preset <preset>` and install Neovim plugins headlessly.
- DevPod: reuse devcontainer.json; DevPod templates point to the same images and mounts. Provide `devpod-presets/<preset>.yaml` referencing the devcontainer config and image tag.

# Section 5: Kubernetes Scaffolding
- Layout: `deploy/<preset>/kustomization.yaml`, `deployment.yaml`, `service.yaml`.
- Deployment: single container using preset image; env: `PORT` (default 8080), `DATABASE_URL` optional; probes: HTTP GET `/healthz` on PORT; resources small defaults.
- Service: ClusterIP exposing PORT. Add `configmap.yaml` for minimal env defaults if needed.
- Kustomize base overlays allow namespace/replica tweaks without Helm.

# Section 6: Draft AGENTS.md Content
- See AGENTS.md in repo; key points: keep scripts minimal/dry-run-friendly, prefer brew, avoid curl|sudo without checks, containers first, run lint/tests and minimal image build for touched presets, keep images lean.

# Section 7: Milestones + Taskboard
- Milestone 1: Foundation (do first)
  - Create `./dev` CLI scaffold with OS detection, dry-run, logging.
  - Add manifests for package mappings (brew/pacman+yay) and base tools; add dotfile symlink helper.
  - Add shellcheck/shfmt + CI for scripts; gofmt/go test for CLI.
- Milestone 2: Presets (host)
  - Define preset manifests (python/go/rust/web) with toolchains and LSP/format/debug.
  - Implement `./dev preset apply --preset <name> [--host]` using manifests.
  - Wire Neovim LazyVim layering to auto-install per-preset tools; minimal tmux/zsh aliases.
- Milestone 3: Containers
  - Add `containers/<preset>/Dockerfile` built from manifests; add `./dev build --preset <name>`.
  - Add smoke test script to run container, check tools versions, and exit cleanly.
- Milestone 4: Devcontainers & DevPod
  - Add `devcontainer.<preset>.json` pointing to preset images; postCreate to sync dotfiles.
  - Add DevPod templates referencing devcontainer configs; document workflow.
- Milestone 5: Kubernetes
  - Add `deploy/<preset>/` scaffolds with kustomization, deployment, service; include health checks/env defaults.
  - Add minimal `./dev k8s render` to validate kustomize build.
- Milestone 6: Hardening
  - Add version pinning/checksums for downloads; tighten sudo usage.
  - Add lint/test/image build CI matrix across presets; add docs for troubleshooting.
