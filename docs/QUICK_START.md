# Quick Start Guide

**⏱️ 5 Minutes to Get Started**

## 🎯 What This Does

Sets up your computer with programming tools automatically.

## 📋 Prerequisites

- A Mac or Linux computer
- Internet connection
- 15 minutes of time

## 🚀 Installation (Copy & Paste)

### Step 1: Install Go

**macOS:**
```bash
brew install go
```

**Ubuntu/Debian:**
```bash
sudo apt-get update && sudo apt-get install golang-go
```

**Arch Linux:**
```bash
sudo pacman -S go
```

### Step 2: Download and Build

```bash
# Download the project
git clone https://github.com/thenameiswiiwin/dev.git ~/workspace/github.com/thenameiswiiwin/dev

# Go to the folder
cd ~/workspace/github.com/thenameiswiiwin/dev

# Build it
make build

# Add to PATH (makes commands shorter)
echo 'export PATH="$HOME/workspace/github.com/thenameiswiiwin/dev/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

## 🎮 Basic Commands

```bash
# Install essential tools
dev bootstrap

# See what you can install
dev preset list

# Install Python tools
dev preset apply python

# Check everything works
dev doctor

# Sync your settings
dev dotfiles sync
```

## 🐍 Example: Setup for Python

```bash
# Step 1: Install basics
dev bootstrap

# Step 2: Install Python tools
dev preset apply python

# Step 3: Test it
python3 --version

# Step 4: Write code!
echo 'print("Hello World")' > test.py
python3 test.py
```

## 🌐 Example: Setup for Web Development

```bash
# Step 1: Install basics
dev bootstrap

# Step 2: Install web tools
dev preset apply web

# Step 3: Test it
node --version
npm --version

# Step 4: Create a project!
npx create-next-app my-app
```

## 🆘 Help Commands

```bash
dev --help              # Show all commands
dev preset --help       # Help with presets
dev doctor              # Check what's installed
dev preset list         # Show available presets
```

## 🔧 Useful Flags

```bash
--dry-run    # Preview without making changes
--verbose    # Show detailed output
--help       # Get help for any command
```

## 📚 Learn More

- **New to coding?** Read [BEGINNER_GUIDE.md](BEGINNER_GUIDE.md)
- **Want details?** Read [IMPLEMENTATION_SUMMARY.md](IMPLEMENTATION_SUMMARY.md)
- **Main docs:** [README.md](../README.md)

## ⚡ Common Tasks

### Install Multiple Languages

```bash
dev preset apply python
dev preset apply go
dev preset apply web
```

### Preview Before Installing

```bash
dev preset apply python --dry-run
```

### See Detailed Output

```bash
dev preset apply python --verbose
```

### Fix Installation Issues

```bash
dev doctor
```

## 🎯 Next Steps

1. ✅ Complete installation
2. ✅ Run `dev bootstrap`
3. ✅ Pick a preset and install it
4. 📖 Read [BEGINNER_GUIDE.md](BEGINNER_GUIDE.md) for detailed explanation
5. 💻 Start coding!

---

**Total time:** ~10-15 minutes
**Questions?** See [BEGINNER_GUIDE.md](BEGINNER_GUIDE.md)
