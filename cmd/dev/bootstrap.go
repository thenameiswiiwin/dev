package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap the development environment",
	Long: `Detects the operating system and installs core dependencies:
- Homebrew (on macOS or if not present on Linux)
- Core tools: git, curl, zsh, tmux
- Basic development tools`,
	Run: runBootstrap,
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)
}

func runBootstrap(cmd *cobra.Command, args []string) {
	log.Info("Starting bootstrap process...")
	log.Info("Detected OS: %s", osInfo.OS)
	log.Info("Package Manager: %s", osInfo.PackageManager)

	// Install Homebrew if needed
	if !osInfo.HasBrew && osInfo.OS == "macos" {
		installHomebrew()
	} else if osInfo.HasBrew {
		log.Success("Homebrew already installed")
	}

	// Install core dependencies
	installCoreDependencies()

	log.Success("Bootstrap completed successfully!")
	log.Info("\nNext steps:")
	log.Info("  1. Run 'dev preset list' to see available presets")
	log.Info("  2. Run 'dev preset apply <preset>' to install a preset")
	log.Info("  3. Run 'dev dotfiles sync' to sync dotfiles")
}

func installHomebrew() {
	log.Action("Installing Homebrew")

	if dryRun {
		return
	}

	// Use the official Homebrew install script
	cmd := exec.Command("bash", "-c",
		`/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		log.Error("Failed to install Homebrew: %v", err)
		os.Exit(1)
	}

	log.Success("Homebrew installed")
}

func installCoreDependencies() {
	coreTools := []string{"git", "curl", "zsh", "tmux", "ripgrep", "fd", "fzf"}

	log.Info("\nInstalling core dependencies...")

	for _, tool := range coreTools {
		if commandExists(tool) {
			log.Debug("%s already installed", tool)
			continue
		}

		log.Action("Installing %s", tool)

		if dryRun {
			continue
		}

		var cmd *exec.Cmd
		switch osInfo.PackageManager {
		case "brew":
			cmd = exec.Command("brew", "install", tool)
		case "apt":
			cmd = exec.Command("sudo", "apt-get", "install", "-y", tool)
		case "pacman":
			cmd = exec.Command("sudo", "pacman", "-S", "--noconfirm", tool)
		case "yay":
			cmd = exec.Command("yay", "-S", "--noconfirm", tool)
		default:
			log.Error("Unsupported package manager: %s", osInfo.PackageManager)
			continue
		}

		if verbose {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}

		if err := cmd.Run(); err != nil {
			log.Warn("Failed to install %s: %v", tool, err)
		} else {
			log.Success("Installed %s", tool)
		}
	}
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	if verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd.Run()
}

func runCommandOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// Helper to check if package is installed
func isPackageInstalled(pkg string) bool {
	switch osInfo.PackageManager {
	case "brew":
		output, err := runCommandOutput("brew", "list", pkg)
		return err == nil && output != ""
	case "apt":
		output, err := runCommandOutput("dpkg", "-s", pkg)
		return err == nil && output != ""
	case "pacman", "yay":
		output, err := runCommandOutput("pacman", "-Q", pkg)
		return err == nil && output != ""
	default:
		return false
	}
}

func installPackage(pkg string) error {
	log.Action("Installing package: %s", pkg)

	if dryRun {
		return nil
	}

	var cmd *exec.Cmd
	switch osInfo.PackageManager {
	case "brew":
		cmd = exec.Command("brew", "install", pkg)
	case "apt":
		cmd = exec.Command("sudo", "apt-get", "install", "-y", pkg)
	case "pacman":
		cmd = exec.Command("sudo", "pacman", "-S", "--noconfirm", pkg)
	case "yay":
		cmd = exec.Command("yay", "-S", "--noconfirm", pkg)
	default:
		return fmt.Errorf("unsupported package manager: %s", osInfo.PackageManager)
	}

	if verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}
