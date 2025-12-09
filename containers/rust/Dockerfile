# Rust development container
FROM rust:1.82-slim-bookworm

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive \
    DEV_PRESET=rust \
    CARGO_HOME=/usr/local/cargo \
    RUSTUP_HOME=/usr/local/rustup

# Install base system packages
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    curl \
    git \
    ca-certificates \
    zsh \
    tmux \
    ripgrep \
    fd-find \
    fzf \
    bat \
    lldb \
    postgresql-client \
    pkg-config \
    libssl-dev \
    && rm -rf /var/lib/apt/lists/*

# Install Neovim from releases
RUN curl -LO https://github.com/neovim/neovim/releases/latest/download/nvim-linux64.tar.gz \
    && tar -xzf nvim-linux64.tar.gz \
    && mv nvim-linux64/bin/nvim /usr/local/bin/ \
    && mv nvim-linux64/share/nvim /usr/local/share/ \
    && mv nvim-linux64/lib/nvim /usr/local/lib/ \
    && rm -rf nvim-linux64 nvim-linux64.tar.gz

# Install Rust components and tools
RUN rustup component add rust-analyzer && \
    rustup component add rustfmt && \
    rustup component add clippy

# Create workspace directory
WORKDIR /workspace

# Expose common ports (app + debug)
EXPOSE 8080 5000

# Set default shell
SHELL ["/bin/zsh", "-c"]
CMD ["/bin/zsh"]
