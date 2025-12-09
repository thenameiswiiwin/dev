package osdetect

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// OS represents the operating system type
type OS string

const (
	MacOS  OS = "macos"
	Arch   OS = "arch"
	Debian OS = "debian"
	Ubuntu OS = "ubuntu"
	Linux  OS = "linux"
)

// Info contains detected OS information
type Info struct {
	OS             OS
	HasBrew        bool
	HasPacman      bool
	HasYay         bool
	HasApt         bool
	PackageManager string
}

// Detect performs OS detection and returns system information
func Detect() (*Info, error) {
	info := &Info{}

	// Detect OS
	goos := runtime.GOOS
	switch goos {
	case "darwin":
		info.OS = MacOS
	case "linux":
		info.OS = detectLinuxDistro()
	default:
		return nil, fmt.Errorf("unsupported OS: %s", goos)
	}

	// Check for package managers
	info.HasBrew = commandExists("brew")
	info.HasPacman = commandExists("pacman")
	info.HasYay = commandExists("yay")
	info.HasApt = commandExists("apt-get")

	// Determine primary package manager
	info.PackageManager = determinePackageManager(info)

	return info, nil
}

// detectLinuxDistro detects the Linux distribution
func detectLinuxDistro() OS {
	if fileExists("/etc/arch-release") {
		return Arch
	}

	if fileExists("/etc/os-release") {
		content := readFile("/etc/os-release")
		if strings.Contains(content, "ID=ubuntu") {
			return Ubuntu
		}
		if strings.Contains(content, "ID=debian") {
			return Debian
		}
	}

	return Linux
}

// determinePackageManager selects the primary package manager based on OS and availability
func determinePackageManager(info *Info) string {
	// Prefer Homebrew if available
	if info.HasBrew {
		return "brew"
	}

	// Fall back to OS-specific managers
	switch info.OS {
	case MacOS:
		return "brew"
	case Arch:
		if info.HasYay {
			return "yay"
		}
		if info.HasPacman {
			return "pacman"
		}
	case Debian, Ubuntu:
		if info.HasApt {
			return "apt"
		}
	}

	return "unknown"
}

// commandExists checks if a command is available in PATH
func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// fileExists checks if a file exists
func fileExists(path string) bool {
	cmd := exec.Command("test", "-f", path)
	return cmd.Run() == nil
}

// readFile reads a file and returns its content
func readFile(path string) string {
	cmd := exec.Command("cat", path)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(output)
}

// String returns a string representation of the OS
func (o OS) String() string {
	return string(o)
}
