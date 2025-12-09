package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/thenameiswiiwin/dev/internal/config"
	"github.com/thenameiswiiwin/dev/internal/installer"
)

var presetCmd = &cobra.Command{
	Use:   "preset",
	Short: "Manage development presets",
	Long:  `Install and manage language-specific development presets (python, go, rust, web)`,
}

var presetListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available presets",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Available presets:")
		log.Info("  • python  - Python development (uv/pyenv, pyright, ruff, pytest)")
		log.Info("  • go      - Go development (gopls, gofumpt, golangci-lint, delve)")
		log.Info("  • rust    - Rust development (rustup, rust-analyzer, clippy)")
		log.Info("  • web     - Web development (Node.js, TypeScript, React, Tailwind)")
		log.Info("\nUsage: dev preset apply <preset>")
	},
}

var presetApplyCmd = &cobra.Command{
	Use:   "apply [preset]",
	Short: "Apply a preset to the host system",
	Long:  `Install toolchains, LSPs, formatters, and editor configuration for a preset`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		preset := args[0]

		// Find repo root
		repoRoot, err := config.FindRepoRoot()
		if err != nil {
			log.Error("Failed to find repository root: %v", err)
			os.Exit(1)
		}

		// Load preset manifest
		manifest, err := config.LoadPresetManifest(repoRoot, preset)
		if err != nil {
			log.Error("Failed to load preset manifest: %v", err)
			log.Info("\nAvailable presets: python, go, rust, web")
			os.Exit(1)
		}

		// Create installer
		inst := installer.New(osInfo, dryRun, verbose, log)

		// Install preset
		if err := inst.InstallPreset(manifest); err != nil {
			log.Error("Failed to install preset: %v", err)
			os.Exit(1)
		}

		// Print next steps
		log.Info("\nNext steps:")
		if manifest.Nvim != nil {
			log.Info("  1. Open Neovim and run :MasonInstallAll to install LSPs")
		}
		if len(manifest.Aliases) > 0 {
			log.Info("  2. Restart your shell to load new aliases")
		}
		log.Info("  3. Run 'dev doctor' to verify installation")
	},
}

func init() {
	rootCmd.AddCommand(presetCmd)
	presetCmd.AddCommand(presetListCmd)
	presetCmd.AddCommand(presetApplyCmd)
}
