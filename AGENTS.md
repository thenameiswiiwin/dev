# Agents Guide

- Scope: keep this repo lean; avoid installing anything not needed for the requested preset or task.
- Modes: default to dry-run where supported; when writing installers, add `--dry-run` or `DRY_RUN=1`.
- Package managers: prefer Homebrew (macOS and Linux). On Arch, fall back to pacman + yay; never mix apt with Arch.
- Idempotency: scripts must be re-runnable without breaking existing setups; guard destructive steps, prompt or back up before overwriting.
- Containers first: prefer adding or updating preset Dockerfiles and devcontainer configs over host-global installs.
- Dependencies: avoid global language installs unless required for bootstrap; prefer toolchain managers (uv/pyenv, go toolchain, rustup, nvm/volta/bun).
- Security: no curl | sudo without checksum/signature; pin versions where practical; keep images minimal.
- Editors/shell: zsh + tmux + LazyVim are first-class; keep tmux-sessionizer available and configured per preset.
- Checks: run lint/tests plus a minimal image build for touched presets; shellcheck/shfmt for scripts, gofmt/go test for Go CLI, docker build for relevant Dockerfiles.
- Logging: keep CLI/script output concise; include clear success/failure codes; avoid noisy progress bars unless needed.
