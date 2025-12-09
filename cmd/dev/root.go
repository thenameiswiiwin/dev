package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thenameiswiiwin/dev/internal/logger"
	"github.com/thenameiswiiwin/dev/internal/osdetect"
)

var (
	dryRun  bool
	verbose bool
	log     *logger.Logger
	osInfo  *osdetect.Info
)

var rootCmd = &cobra.Command{
	Use:   "dev",
	Short: "Dev environment management tool",
	Long: `A CLI tool for managing development environments, dotfiles, and presets.
Supports host installation, containers, devcontainers, and Kubernetes deployments.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Initialize logger
		log = logger.New(dryRun, verbose)

		// Detect OS (skip for commands that don't need it)
		if cmd.Name() != "version" && cmd.Name() != "help" {
			var err error
			osInfo, err = osdetect.Detect()
			if err != nil {
				log.Error("Failed to detect OS: %v", err)
				os.Exit(1)
			}

			if verbose {
				log.Debug("OS: %s, Package Manager: %s", osInfo.OS, osInfo.PackageManager)
			}
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Show what would be done without making changes")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
