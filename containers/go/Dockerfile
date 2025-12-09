# Go development container
FROM golang:1.24-bookworm

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive \
    DEV_PRESET=go \
    CGO_ENABLED=1 \
    GOPROXY=https://proxy.golang.org,direct

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

# Install Go tools
RUN go install mvdan.cc/gofumpt@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install golang.org/x/tools/gopls@latest

# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin

# Create workspace directory
WORKDIR /workspace

# Expose common ports (app + debug)
EXPOSE 8080 2345

# Set default shell
SHELL ["/bin/zsh", "-c"]
CMD ["/bin/zsh"]
