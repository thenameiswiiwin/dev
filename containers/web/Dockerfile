# Web development container (Node.js/TypeScript/React)
FROM node:22-bookworm-slim

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive \
    DEV_PRESET=web \
    NODE_ENV=development

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

# Install pnpm and bun
RUN npm install -g pnpm && \
    curl -fsSL https://bun.sh/install | bash

# Install global npm packages
RUN npm install -g \
    typescript \
    tsx \
    eslint \
    prettier \
    vite

# Add bun to PATH
ENV PATH="/root/.bun/bin:${PATH}"

# Create workspace directory
WORKDIR /workspace

# Expose common ports (dev server + debug)
EXPOSE 3000 5173 9229

# Set default shell
SHELL ["/bin/zsh", "-c"]
CMD ["/bin/zsh"]
