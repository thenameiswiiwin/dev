# Beginner's Guide to Dev Environment Manager

Welcome! This guide will help you understand and use this project, even if you're new to coding.

## What is This Project?

Think of this project as a **smart installer** for your computer that:
- Sets up programming tools automatically
- Keeps all your settings in one place
- Works the same way on different computers
- Lets you switch between different programming languages easily

**Real-world analogy:** It's like having a recipe book that not only tells you what ingredients you need, but also goes shopping for you and sets up your kitchen!

## What You'll Need

### 1. A Computer
- **macOS** (Mac computers)
- **Linux** (Ubuntu, Debian, or Arch Linux)

### 2. Basic Tools (we'll help you install these)
- **Terminal** - A text-based way to talk to your computer (already on your Mac/Linux)
- **Git** - A tool to download code from the internet
- **Go** - A programming language (needed to build our tool)

## Part 1: Understanding the Terminal

### What is the Terminal?

The **terminal** (also called "command line" or "shell") is a way to control your computer by typing commands instead of clicking buttons.

### How to Open the Terminal

**On macOS:**
1. Press `Command (⌘) + Space` to open Spotlight
2. Type "Terminal" and press Enter
3. A window with text will appear - this is your terminal!

**On Linux:**
1. Press `Ctrl + Alt + T` (works on most Linux systems)
2. Or search for "Terminal" in your applications

### Basic Terminal Concepts

When you see something like this in a guide:
```bash
ls
```

It means: "Type `ls` in the terminal and press Enter"

**Common symbols you'll see:**
- `$` or `>` at the start of a line = This is just showing it's a command (don't type this part)
- `#` = A comment explaining something (don't type this part)

**Example:**
```bash
# This lists files in the current folder
ls

# This shows where you are on your computer
pwd
```

## Part 2: Installing the Project

### Step 1: Check if Git is Installed

In your terminal, type:
```bash
git --version
```

**What you should see:**
```
git version 2.x.x
```

**If you see an error:**
- **macOS:** Type `xcode-select --install` and follow prompts
- **Ubuntu/Debian:** Type `sudo apt-get install git`
- **Arch Linux:** Type `sudo pacman -S git`

### Step 2: Check if Go is Installed

In your terminal, type:
```bash
go version
```

**What you should see:**
```
go version go1.24.x
```

**If you see an error, install Go:**
- **macOS:** `brew install go` (if you have Homebrew)
- **Ubuntu/Debian:** `sudo apt-get install golang-go`
- **Arch Linux:** `sudo pacman -S go`
- **Or visit:** https://go.dev/dl/ and download the installer

### Step 3: Download the Project

Copy and paste this entire command into your terminal:

```bash
git clone https://github.com/thenameiswiiwin/dev.git ~/workspace/github.com/thenameiswiiwin/dev
```

**What this does:**
- Downloads the project from the internet
- Puts it in a folder at `~/workspace/github.com/thenameiswiiwin/dev`
- The `~` symbol means "your home folder"

**What you should see:**
```
Cloning into '~/workspace/github.com/thenameiswiiwin/dev'...
remote: Enumerating objects: ...
remote: Counting objects: ...
Receiving objects: 100% ...
```

### Step 4: Go to the Project Folder

```bash
cd ~/workspace/github.com/thenameiswiiwin/dev
```

**What `cd` means:** "Change Directory" - it's like clicking on a folder to open it

**To confirm you're in the right place, type:**
```bash
pwd
```

**You should see:**
```
/Users/YourName/workspace/github.com/thenameiswiiwin/dev
```

### Step 5: Build the Tool

```bash
make build
```

**What this does:**
- Compiles the Go code into a program you can run
- Creates a program called `dev` in the `bin` folder

**What you should see:**
```
Building dev CLI...
go build -o bin/dev ./cmd/dev
✓ Built bin/dev
```

**If you see errors:**
- Make sure you're in the project folder (Step 4)
- Make sure Go is installed correctly (Step 2)

### Step 6: Test the Tool

```bash
./bin/dev --help
```

**What you should see:**
```
A CLI tool for managing development environments...

Usage:
  dev [command]

Available Commands:
  bootstrap   Bootstrap the development environment
  preset      Manage development presets
  ...
```

**Success!** You've installed the tool! 🎉

## Part 3: Understanding the Commands

Our tool has several commands. Think of each command as a button that does something specific:

### `dev bootstrap`
**What it does:** Installs essential tools that most programmers need
**When to use it:** First time setting up a new computer
**Tools it installs:**
- Git (for downloading code)
- zsh (a better terminal)
- tmux (run multiple programs in one terminal window)
- neovim (a powerful text editor)
- Various helpful utilities

### `dev preset list`
**What it does:** Shows you what programming languages you can set up
**Output example:**
```
Available presets:
  • python  - Python development
  • go      - Go development
  • rust    - Rust development
  • web     - Web development (JavaScript/TypeScript)
```

### `dev preset apply <preset>`
**What it does:** Installs everything you need for a specific programming language
**Example:** `dev preset apply python`
**What it installs for Python:**
- Python programming language
- Tools to help you write Python code
- Debugging tools
- Testing tools

### `dev dotfiles sync`
**What it does:** Copies your settings and configurations
**Example:** Your editor preferences, terminal colors, shortcuts
**Important:** This backs up your existing settings before changing anything

### `dev doctor`
**What it does:** Checks if everything is installed correctly
**Use this to:** Troubleshoot problems

### `dev build`
**What it does:** Creates Docker containers (like virtual computers)
**Advanced feature:** You probably won't need this when starting out

## Part 4: Your First Setup (Step by Step)

Let's set up a Python development environment as an example:

### Step 1: Bootstrap Your System

```bash
./bin/dev bootstrap
```

**What happens:**
1. Checks what operating system you're using
2. Installs Homebrew (on Mac) or updates your package manager
3. Installs essential tools
4. Shows progress as it works

**This might take:** 5-15 minutes

**What you'll see:**
```
Starting bootstrap process...
Detected OS: macos
Package Manager: brew
✓ Homebrew already installed
→ Installing git
✓ Installed git
→ Installing zsh
✓ Installed zsh
...
✓ Bootstrap completed successfully!
```

### Step 2: See What You Can Install

```bash
./bin/dev preset list
```

**Output:**
```
Available presets:
  • python  - Python development (uv/pyenv, pyright, ruff, pytest)
  • go      - Go development (gopls, gofumpt, golangci-lint, delve)
  • rust    - Rust development (rustup, rust-analyzer, clippy)
  • web     - Web development (Node.js, TypeScript, React, Tailwind)
```

### Step 3: Install Python Tools

```bash
./bin/dev preset apply python
```

**What happens:**
1. Installs Python 3.12 (the programming language)
2. Installs `uv` (fast package installer for Python)
3. Installs `pyright` (helps you write better Python code)
4. Installs `ruff` (checks your code for mistakes)
5. Installs `pytest` (for testing your code)

**This might take:** 5-10 minutes

**What you'll see:**
```
Installing preset: python
Description: Python development with uv/pyenv, pyright, ruff, and pytest

[1/8] Installing python3
→ Installing python3
✓ Installed python3
[2/8] Installing uv
→ Installing uv
✓ Installed uv
...
✓ Preset python installed successfully!

Next steps:
  1. Open Neovim and run :MasonInstallAll to install LSPs
  2. Restart your shell to load new aliases
  3. Run 'dev doctor' to verify installation
```

### Step 4: Verify Everything Works

```bash
./bin/dev doctor
```

**What you'll see:**
```
Running system diagnostics...

=== System Information ===
OS: macos
Package Manager: brew

=== Core Tools ===
✓ git installed
✓ curl installed
✓ zsh installed
✓ tmux installed
✓ nvim installed
✓ python3 installed
...
```

### Step 5: Sync Your Dotfiles (Optional)

```bash
./bin/dev dotfiles sync --dry-run
```

**The `--dry-run` flag:** Shows what would happen WITHOUT actually doing it (like a preview)

**If everything looks good, run it for real:**
```bash
./bin/dev dotfiles sync
```

## Part 5: Understanding Presets

### What is a Preset?

A **preset** is a collection of tools for a specific programming language. Think of it like a "starter pack" for different types of coding.

### Available Presets:

#### 🐍 Python Preset
**Best for:**
- Learning to code
- Data science and AI
- Web applications
- Automation scripts

**What you get:**
- Python 3.12
- Code editor support
- Testing tools
- Debugging tools

**Try it:** `dev preset apply python`

#### 🔷 Go Preset
**Best for:**
- Building web servers
- Command-line tools
- Cloud applications

**What you get:**
- Go programming language
- Fast code formatting
- Powerful debugging

**Try it:** `dev preset apply go`

#### 🦀 Rust Preset
**Best for:**
- System programming
- High-performance applications
- Safe, fast code

**What you get:**
- Rust compiler
- Code analysis tools
- Clippy (helpful suggestions)

**Try it:** `dev preset apply rust`

#### 🌐 Web Preset
**Best for:**
- Websites
- Web applications
- Mobile apps with React Native

**What you get:**
- Node.js (JavaScript runtime)
- TypeScript support
- React framework
- Tailwind CSS
- Multiple package managers (npm, pnpm, bun)

**Try it:** `dev preset apply web`

## Part 6: Common Questions

### Q: Can I install multiple presets?

**A:** Yes! You can install as many as you want:
```bash
dev preset apply python
dev preset apply web
dev preset apply go
```

### Q: What if something goes wrong?

**A:** Use the `--dry-run` flag to test first:
```bash
dev preset apply python --dry-run
```

This shows what WOULD happen without actually doing it.

### Q: How do I uninstall?

**A:** Currently, you'd need to remove packages manually. We're working on a cleanup command!

### Q: Can I see more details while it's working?

**A:** Yes! Use the `--verbose` flag:
```bash
dev preset apply python --verbose
```

### Q: Where are my dotfiles stored?

**A:** In the `env/` folder of this project:
- `env/.config/` - Application settings
- `env/.local/` - Scripts and local files
- `env/.zshrc` - Terminal settings

### Q: Is this safe?

**A:** Yes!
- The `--dry-run` flag lets you preview changes
- Dotfiles are symlinked (not copied), so you can easily undo
- Your original files are backed up with timestamps

## Part 7: What to Do After Setup

### 1. Restart Your Terminal

Some changes need a fresh terminal to take effect:

**Close and reopen your terminal**, or type:
```bash
source ~/.zshrc
```

### 2. Test Your New Tools

**For Python:**
```bash
python3 --version
pip --version
```

**For Go:**
```bash
go version
```

**For Node/Web:**
```bash
node --version
npm --version
```

### 3. Write Your First Program

**Python example:**
```bash
# Create a file
echo 'print("Hello, World!")' > hello.py

# Run it
python3 hello.py
```

**Expected output:**
```
Hello, World!
```

### 4. Open Your Code Editor

**If you have VS Code:**
```bash
code .
```

**If you want to use Neovim (terminal editor):**
```bash
nvim hello.py
```

## Part 8: Helpful Tips

### Tip 1: Use Tab Completion

While typing a command, press `Tab` to auto-complete:
```bash
dev pre[TAB]  → dev preset
```

### Tip 2: Check Command History

Press the `↑` (up arrow) key to see previous commands

### Tip 3: Get Help for Any Command

Add `--help` to any command:
```bash
dev --help
dev preset --help
dev bootstrap --help
```

### Tip 4: Make Commands Shorter

Add this to your shell config to type `dev` instead of `./bin/dev`:

```bash
echo 'export PATH="$HOME/workspace/github.com/thenameiswiiwin/dev/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

Now you can just type:
```bash
dev doctor
dev preset list
```

## Part 9: Troubleshooting

### Problem: "Command not found: dev"

**Solution:** Make sure you're running `./bin/dev` (with the `./bin/` part)

Or add it to your PATH (see Tip 4 above)

### Problem: "Permission denied"

**Solution:** The file isn't executable. Fix it:
```bash
chmod +x ./bin/dev
```

### Problem: "No such file or directory"

**Solution:** Make sure you're in the project folder:
```bash
cd ~/workspace/github.com/thenameiswiiwin/dev
pwd  # Should show the project path
```

### Problem: Installation is slow

**Solution:** This is normal! Installing development tools takes time. Get a coffee ☕

### Problem: I want to start over

**Solution:** Delete the project and reinstall:
```bash
cd ~
rm -rf ~/workspace/github.com/thenameiswiiwin/dev
# Then follow installation steps again
```

## Part 10: Learning More

### Recommended Learning Path

1. **Week 1:** Install and explore
   - Run `dev bootstrap`
   - Try `dev doctor`
   - Look at the files in `env/`

2. **Week 2:** Install your first preset
   - Choose Python or Web (easiest for beginners)
   - Write a "Hello World" program
   - Explore the installed tools

3. **Week 3:** Customize
   - Edit files in `env/.config/`
   - Try different terminal colors
   - Create your own aliases

4. **Month 2+:** Advanced features
   - Try Docker containers
   - Learn about Kubernetes
   - Contribute improvements

### External Resources

**Learn Terminal Basics:**
- https://www.learnenough.com/command-line-tutorial
- https://overthewire.org/wargames/bandit/

**Learn Git:**
- https://learngitbranching.js.org/
- https://try.github.io/

**Learn Python:**
- https://www.python.org/about/gettingstarted/
- https://realpython.com/

**Learn Web Development:**
- https://developer.mozilla.org/en-US/docs/Learn
- https://www.freecodecamp.org/

## Glossary

**Terms you'll see often:**

- **CLI (Command Line Interface):** Using text commands instead of clicking buttons
- **Repository (Repo):** A folder containing code and files
- **Git:** A tool to track changes in your code
- **Package Manager:** A tool that installs other tools (like an app store for developers)
- **Homebrew:** The most popular package manager for macOS
- **Terminal:** A text-based way to control your computer
- **Shell:** The program that runs in your terminal (like zsh or bash)
- **Dotfiles:** Configuration files that start with a dot (like `.zshrc`)
- **PATH:** A list of folders where your computer looks for programs
- **Preset:** A collection of tools for a specific programming language
- **LSP (Language Server Protocol):** Helps your editor understand code better
- **Dry-run:** Preview mode that shows what would happen without doing it
- **Bootstrap:** Set up the basic tools you need to get started
- **Container:** A lightweight virtual computer (Docker)
- **Manifest:** A file that lists what should be installed

## Getting Help

**Need more help?**

1. **Read the error message carefully** - It often tells you what's wrong
2. **Use `--help` flag** - Every command has a help page
3. **Check the docs** - Look in the `docs/` folder
4. **Run `dev doctor`** - Diagnose installation issues
5. **Use `--dry-run`** - Preview before running
6. **Ask for help** - File an issue on GitHub

## Summary

You've learned how to:
- ✅ Open and use the terminal
- ✅ Install the dev environment manager
- ✅ Run basic commands
- ✅ Install programming language tools
- ✅ Troubleshoot common problems
- ✅ Understand key concepts

**Remember:**
- Start with `dev bootstrap`
- Use `--dry-run` to preview
- Use `--help` when stuck
- Take it one step at a time

Happy coding! 🚀

---

**Next Steps:**
1. Complete the setup (Part 4)
2. Write your first program (Part 7)
3. Customize your environment (env/ folder)
4. Learn your chosen programming language

**Questions?** Check the main [README.md](../README.md) or [IMPLEMENTATION_SUMMARY.md](IMPLEMENTATION_SUMMARY.md)
