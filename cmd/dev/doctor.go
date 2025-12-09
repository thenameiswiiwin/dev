package main

import (
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system health and dependencies",
	Long:  `Verify that all required tools and dependencies are properly installed`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Running system diagnostics...")
		log.Info("\n=== System Information ===")
		log.Info("OS: %s", osInfo.OS)
		log.Info("Package Manager: %s", osInfo.PackageManager)
		log.Info("")

		// Check core tools
		log.Info("=== Core Tools ===")
		checkTool("git")
		checkTool("curl")
		checkTool("zsh")
		checkTool("tmux")
		checkTool("nvim")
		checkTool("ripgrep")
		checkTool("fd")
		checkTool("fzf")

		log.Info("")
		log.Info("=== Package Managers ===")
		if osInfo.HasBrew {
			log.Success("Homebrew installed")
		} else {
			log.Warn("Homebrew not found")
		}

		if osInfo.HasPacman {
			log.Success("Pacman installed")
		}

		if osInfo.HasYay {
			log.Success("Yay installed")
		}

		if osInfo.HasApt {
			log.Success("APT installed")
		}
	},
}

func checkTool(name string) {
	if commandExists(name) {
		log.Success("%s installed", name)
	} else {
		log.Warn("%s not found", name)
	}
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
