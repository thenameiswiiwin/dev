package dotfiles

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/thenameiswiiwin/dev/internal/logger"
)

// Syncer handles dotfile synchronization
type Syncer struct {
	RepoRoot string
	DryRun   bool
	Log      *logger.Logger
}

// New creates a new dotfile syncer
func New(repoRoot string, dryRun bool, log *logger.Logger) *Syncer {
	return &Syncer{
		RepoRoot: repoRoot,
		DryRun:   dryRun,
		Log:      log,
	}
}

// Sync syncs dotfiles from env/ directory to home directory
func (s *Syncer) Sync() error {
	s.Log.Info("Syncing dotfiles...")

	envDir := filepath.Join(s.RepoRoot, "env")
	if _, err := os.Stat(envDir); os.IsNotExist(err) {
		return fmt.Errorf("env directory not found: %s", envDir)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		configHome = filepath.Join(homeDir, ".config")
	}

	// Sync .config directory
	configSrc := filepath.Join(envDir, ".config")
	if _, err := os.Stat(configSrc); err == nil {
		if err := s.syncDirectory(configSrc, configHome); err != nil {
			return fmt.Errorf("failed to sync .config: %w", err)
		}
	}

	// Sync .local directory
	localSrc := filepath.Join(envDir, ".local")
	localDest := filepath.Join(homeDir, ".local")
	if _, err := os.Stat(localSrc); err == nil {
		if err := s.syncDirectory(localSrc, localDest); err != nil {
			return fmt.Errorf("failed to sync .local: %w", err)
		}
	}

	// Sync dotfiles in env/ root (like .zshrc, .zsh_profile)
	if err := s.syncDotfiles(envDir, homeDir); err != nil {
		return fmt.Errorf("failed to sync dotfiles: %w", err)
	}

	s.Log.Success("Dotfiles synced successfully")
	return nil
}

// syncDirectory syncs a directory by creating symlinks
func (s *Syncer) syncDirectory(src, dest string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the root directory itself
		if path == src {
			return nil
		}

		// Calculate relative path
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(dest, relPath)

		if info.IsDir() {
			return s.createDirectory(destPath)
		}

		return s.createSymlink(path, destPath)
	})
}

// syncDotfiles syncs dotfiles (files starting with .) in the root directory
func (s *Syncer) syncDotfiles(src, dest string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		name := entry.Name()

		// Skip directories and non-dotfiles
		if entry.IsDir() || !strings.HasPrefix(name, ".") {
			continue
		}

		// Skip .config and .local (handled separately)
		if name == ".config" || name == ".local" {
			continue
		}

		srcPath := filepath.Join(src, name)
		destPath := filepath.Join(dest, name)

		if err := s.createSymlink(srcPath, destPath); err != nil {
			s.Log.Error("Failed to sync %s: %v", name, err)
		}
	}

	return nil
}

// createDirectory creates a directory if it doesn't exist
func (s *Syncer) createDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		s.Log.Action("Creating directory: %s", path)

		if !s.DryRun {
			if err := os.MkdirAll(path, 0755); err != nil {
				return err
			}
		}
	}
	return nil
}

// createSymlink creates a symlink, backing up existing files
func (s *Syncer) createSymlink(src, dest string) error {
	// Check if destination already exists
	destInfo, err := os.Lstat(dest)
	if err == nil {
		// Check if it's already a symlink pointing to the right place
		if destInfo.Mode()&os.ModeSymlink != 0 {
			target, err := os.Readlink(dest)
			if err == nil && target == src {
				s.Log.Debug("Symlink already exists: %s -> %s", dest, src)
				return nil
			}
		}

		// Back up existing file/symlink
		if err := s.backupFile(dest); err != nil {
			return fmt.Errorf("failed to backup %s: %w", dest, err)
		}
	}

	s.Log.Action("Creating symlink: %s -> %s", dest, src)

	if !s.DryRun {
		// Ensure parent directory exists
		parentDir := filepath.Dir(dest)
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			return err
		}

		// Create symlink
		if err := os.Symlink(src, dest); err != nil {
			return fmt.Errorf("failed to create symlink: %w", err)
		}
	}

	return nil
}

// backupFile backs up an existing file with a timestamp
func (s *Syncer) backupFile(path string) error {
	timestamp := time.Now().Format("20060102-150405")
	backupPath := fmt.Sprintf("%s.backup-%s", path, timestamp)

	s.Log.Warn("Backing up existing file: %s -> %s", path, backupPath)

	if !s.DryRun {
		if err := os.Rename(path, backupPath); err != nil {
			return err
		}
	}

	return nil
}
