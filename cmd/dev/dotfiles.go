package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/thenameiswiiwin/dev/internal/config"
	"github.com/thenameiswiiwin/dev/internal/dotfiles"
)

var dotfilesCmd = &cobra.Command{
	Use:   "dotfiles",
	Short: "Manage dotfiles",
	Long:  `Sync dotfiles using symlinks (non-destructive, with backups)`,
}

var dotfilesSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync dotfiles to home directory",
	Long: `Create symlinks from env/ to $HOME and $XDG_CONFIG_HOME.
Backs up existing files instead of overwriting them.`,
	Run: func(cmd *cobra.Command, args []string) {
		repoRoot, err := config.FindRepoRoot()
		if err != nil {
			log.Error("Failed to find repository root: %v", err)
			os.Exit(1)
		}

		syncer := dotfiles.New(repoRoot, dryRun, log)
		if err := syncer.Sync(); err != nil {
			log.Error("Failed to sync dotfiles: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(dotfilesCmd)
	dotfilesCmd.AddCommand(dotfilesSyncCmd)
}
