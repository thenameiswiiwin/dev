package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Package represents a single package with OS-specific names
type Package struct {
	Name          string `yaml:"name"`
	Brew          string `yaml:"brew,omitempty"`
	Apt           string `yaml:"apt,omitempty"`
	Pacman        string `yaml:"pacman,omitempty"`
	Description   string `yaml:"description,omitempty"`
	InstallMethod string `yaml:"install_method,omitempty"`
	GoPackage     string `yaml:"go_package,omitempty"`
	NPMPackage    string `yaml:"npm_package,omitempty"`
	ScriptURL     string `yaml:"script_url,omitempty"`
}

// NvimConfig represents Neovim-specific configuration
type NvimConfig struct {
	MasonEnsureInstalled []string `yaml:"mason_ensure_installed"`
	FormatOnSave         bool     `yaml:"format_on_save"`
	Formatter            string   `yaml:"formatter"`
}

// Manifest represents a complete manifest file (base or preset)
type Manifest struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Packages    []Package         `yaml:"packages"`
	PostInstall []string          `yaml:"post_install,omitempty"`
	NPMGlobal   []string          `yaml:"npm_global,omitempty"`
	Nvim        *NvimConfig       `yaml:"nvim,omitempty"`
	Aliases     map[string]string `yaml:"aliases,omitempty"`
}

// LoadManifest loads a manifest from a YAML file
func LoadManifest(path string) (*Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest: %w", err)
	}

	var manifest Manifest
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}

	return &manifest, nil
}

// LoadBaseManifest loads the base manifest
func LoadBaseManifest(repoRoot string) (*Manifest, error) {
	path := filepath.Join(repoRoot, "manifests", "base.yaml")
	return LoadManifest(path)
}

// LoadPresetManifest loads a preset manifest by name
func LoadPresetManifest(repoRoot, preset string) (*Manifest, error) {
	path := filepath.Join(repoRoot, "manifests", "presets", preset+".yaml")
	return LoadManifest(path)
}

// GetPackageName returns the appropriate package name for the given package manager
func (p *Package) GetPackageName(packageManager string) string {
	switch packageManager {
	case "brew":
		if p.Brew != "" {
			return p.Brew
		}
	case "apt":
		if p.Apt != "" {
			return p.Apt
		}
	case "pacman", "yay":
		if p.Pacman != "" {
			return p.Pacman
		}
	}
	return p.Name
}

// FindRepoRoot finds the repository root by looking for go.mod or .git
func FindRepoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		// Check for go.mod or .git
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find repository root")
		}
		dir = parent
	}
}
