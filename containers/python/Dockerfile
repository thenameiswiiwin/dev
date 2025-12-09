# Python development container
FROM python:3.12-slim

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive \
    PYTHONUNBUFFERED=1 \
    PIP_NO_CACHE_DIR=1 \
    DEV_PRESET=python

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
    postgresql-client \
    && rm -rf /var/lib/apt/lists/*

# Install Neovim from releases
RUN curl -LO https://github.com/neovim/neovim/releases/latest/download/nvim-linux64.tar.gz \
    && tar -xzf nvim-linux64.tar.gz \
    && mv nvim-linux64/bin/nvim /usr/local/bin/ \
    && mv nvim-linux64/share/nvim /usr/local/share/ \
    && mv nvim-linux64/lib/nvim /usr/local/lib/ \
    && rm -rf nvim-linux64 nvim-linux64.tar.gz

# Install Python tools
RUN pip install --upgrade pip && \
    pip install uv pipx

# Install Python LSP and formatters via pipx
RUN pipx install pyright && \
    pipx install ruff && \
    pipx install black && \
    pipx install debugpy && \
    pipx install pytest

# Add pipx binaries to PATH
ENV PATH="/root/.local/bin:${PATH}"

# Create workspace directory
WORKDIR /workspace

# Expose common ports (app + debug)
EXPOSE 8000 5678

# Set default shell
SHELL ["/bin/zsh", "-c"]
CMD ["/bin/zsh"]
