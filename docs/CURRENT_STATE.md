# Current State
- Repo purpose is to bootstrap a personal dev machine and dotfiles primarily through shell scripts. A tiny Go HTTP server (`main.go`) serves the embedded `resources/setup` installer script; Dockerfile builds that server.
- Bootstrap flow: `resources/setup` runs `apt update`, installs Homebrew, clones this repo into `~/workspace/github.com/thenameiswiiwin/dev`, then calls `./run`.
- `run` requires `DEV_ENV` pointing at the repo and executes every executable in `runs/` (optionally filtered by substring). Each script installs tools via Homebrew or `apt` (e.g., dev tools, Docker, language runtimes, UI apps) without sequencing or dependency checks.
- `dev-env` copies dotfiles from `env/` into `$HOME` and `$XDG_CONFIG_HOME`, wipes existing directories under `.config`/`.local`, and installs helper scripts (e.g., `tmux-sessionizer`).
- Vendored assets include full Lua 5.1.5 and LuaRocks 3.11.1 trees plus font binaries under `resources/fonts/`.

# Constraints & Problems
- Mixed package managers and OS assumptions: `resources/setup` calls both `apt` (Debian/Ubuntu) and Homebrew (macOS), and many `runs/` scripts use `brew` while others use `apt`, so the flow breaks on the wrong host.
- Non-idempotent, destructive home changes: `dev-env` removes `$XDG_CONFIG_HOME/*` subdirs and overwrites dotfiles without backups; `runs/` scripts clone into `~/workspace` and delete existing projects.
- Unreliable orchestration: `run` iterates `find ... -perm +111` (non-POSIX), runs scripts in filesystem order, and logs an undefined `$env`. There is no error handling or OS/feature detection before running package installs or UI commands (`open -a Docker`).
- Large vendored tarballs and fonts inflate the repo and may have licensing implications; they are copied wholesale instead of fetched on demand.
- Go module declares many unused dependencies and defaults `GO_VERSION=1` in the Dockerfile, which will not match the `go.mod` toolchain (1.24.x) and produces a non-reproducible build.
- Security/usability gaps: curl-piped installers without checksum, blanket `sudo` calls, no signature checking, and no validation/tests (no linting, shellcheck, or CI).

# Target Architecture (proposed)
- Replace ad-hoc scripts with a single Go or Bash CLI (`dev`) that offers `install`, `sync-dotfiles`, `doctor`, and `cleanup` commands with clear logging and exit codes.
- Define OS-specific package manifests (e.g., YAML/JSON) for macOS vs Debian/Ubuntu and drive installers from that data; gate commands on detected OS/package manager and make them idempotent.
- Manage dotfiles via symlinks (`stow`, `chezmoi`, or a small Go linker) instead of destructive copies; back up or skip existing files and allow profile selection.
- Move large artifacts (Lua/LuaRocks tarballs, fonts) to on-demand downloads or Git LFS/releases with checksum verification; keep only minimal bootstrap logic in the repo.
- Simplify bootstrap: a single curlable script that detects OS, installs prerequisites (Homebrew/apt packages), pulls the repo, and invokes the CLI; deprecate the HTTP server and Dockerfile unless a containerized bootstrap is truly needed.
- Add validation: shellcheck/shfmt for scripts, Go lint/tests for the CLI, and a small CI workflow to ensure manifests stay consistent; provide a dry-run mode to show planned changes before applying them.
