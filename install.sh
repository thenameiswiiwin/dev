#!/usr/bin/env bash
set -e

# Dev Environment Manager Installation Script
# Usage: curl -fsSL https://raw.githubusercontent.com/thenameiswiiwin/dev/main/install.sh | bash

REPO_URL="https://github.com/thenameiswiiwin/dev.git"
INSTALL_DIR="${HOME}/workspace/github.com/thenameiswiiwin/dev"

echo "================================"
echo "Dev Environment Manager Installer"
echo "================================"
echo ""

# Check if git is installed
if ! command -v git &>/dev/null; then
  echo "Error: git is not installed"
  echo "Please install git first:"
  echo "  macOS: brew install git"
  echo "  Debian/Ubuntu: sudo apt-get install git"
  echo "  Arch: sudo pacman -S git"
  exit 1
fi

# Check if Go is installed
if ! command -v go &>/dev/null; then
  echo "Error: Go is not installed"
  echo "Please install Go 1.24+ first:"
  echo "  macOS: brew install go"
  echo "  Debian/Ubuntu: sudo apt-get install golang-go"
  echo "  Arch: sudo pacman -S go"
  echo ""
  echo "Or visit: https://go.dev/dl/"
  exit 1
fi

# Create parent directory
PARENT_DIR=$(dirname "$INSTALL_DIR")
mkdir -p "$PARENT_DIR"

# Clone or update repository
if [ -d "$INSTALL_DIR" ]; then
  echo "Repository already exists at $INSTALL_DIR"
  echo "Updating..."
  cd "$INSTALL_DIR"
  git pull origin main
else
  echo "Cloning repository to $INSTALL_DIR..."
  git clone "$REPO_URL" "$INSTALL_DIR"
  cd "$INSTALL_DIR"
fi

echo ""
echo "Building CLI..."
make build

echo ""
echo "================================"
echo "Installation complete!"
echo "================================"
echo ""
echo "Next steps:"
echo "  1. Add to PATH (optional):"
echo "     echo 'export PATH=\"$INSTALL_DIR/bin:\$PATH\"' >> ~/.zshrc"
echo "     source ~/.zshrc"
echo ""
echo "  2. Bootstrap your system:"
echo "     $INSTALL_DIR/bin/dev bootstrap"
echo ""
echo "  3. Apply a preset:"
echo "     $INSTALL_DIR/bin/dev preset list"
echo "     $INSTALL_DIR/bin/dev preset apply <preset>"
echo ""
echo "  4. Sync dotfiles:"
echo "     $INSTALL_DIR/bin/dev dotfiles sync"
echo ""
echo "Documentation: $INSTALL_DIR/README.md"
