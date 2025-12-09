package installer

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/thenameiswiiwin/dev/internal/config"
	"github.com/thenameiswiiwin/dev/internal/logger"
	"github.com/thenameiswiiwin/dev/internal/osdetect"
)

// Installer handles package installation
type Installer struct {
	OSInfo  *osdetect.Info
	DryRun  bool
	Verbose bool
	Log     *logger.Logger
}

// New creates a new installer
func New(osInfo *osdetect.Info, dryRun, verbose bool, log *logger.Logger) *Installer {
	return &Installer{
		OSInfo:  osInfo,
		DryRun:  dryRun,
		Verbose: verbose,
		Log:     log,
	}
}

// InstallPreset installs all packages from a preset manifest
func (i *Installer) InstallPreset(manifest *config.Manifest) error {
	i.Log.Info("Installing preset: %s", manifest.Name)
	i.Log.Info("Description: %s", manifest.Description)
	i.Log.Info("")

	totalSteps := len(manifest.Packages)
	if len(manifest.PostInstall) > 0 {
		totalSteps++
	}
	if len(manifest.NPMGlobal) > 0 {
		totalSteps++
	}

	currentStep := 0

	// Install packages
	for _, pkg := range manifest.Packages {
		currentStep++
		i.Log.Step(currentStep, totalSteps, "Installing %s", pkg.Name)

		if err := i.InstallPackage(&pkg); err != nil {
			i.Log.Error("Failed to install %s: %v", pkg.Name, err)
			// Continue with other packages
		}
	}

	// Install npm global packages
	if len(manifest.NPMGlobal) > 0 {
		currentStep++
		i.Log.Step(currentStep, totalSteps, "Installing npm global packages")
		for _, pkg := range manifest.NPMGlobal {
			if err := i.installNPMGlobal(pkg); err != nil {
				i.Log.Error("Failed to install npm package %s: %v", pkg, err)
			}
		}
	}

	// Run post-install commands
	if len(manifest.PostInstall) > 0 {
		currentStep++
		i.Log.Step(currentStep, totalSteps, "Running post-install commands")
		for _, cmd := range manifest.PostInstall {
			if err := i.runCommand(cmd); err != nil {
				i.Log.Error("Failed to run post-install command: %v", err)
			}
		}
	}

	i.Log.Success("Preset %s installed successfully!", manifest.Name)
	return nil
}

// InstallPackage installs a single package
func (i *Installer) InstallPackage(pkg *config.Package) error {
	// Check if already installed
	if i.isInstalled(pkg) {
		i.Log.Debug("%s already installed", pkg.Name)
		return nil
	}

	i.Log.Action("Installing %s", pkg.Name)

	if i.DryRun {
		return nil
	}

	// Handle different install methods
	switch pkg.InstallMethod {
	case "go_install":
		return i.installGoPackage(pkg.GoPackage)
	case "pipx":
		return i.installPipx(pkg.Name)
	case "pip":
		return i.installPip(pkg.Name)
	case "npm":
		return i.installNPM(pkg.NPMPackage)
	case "script":
		return i.installScript(pkg.ScriptURL)
	default:
		return i.installSystemPackage(pkg)
	}
}

// installSystemPackage installs a package using the system package manager
func (i *Installer) installSystemPackage(pkg *config.Package) error {
	pkgName := pkg.GetPackageName(i.OSInfo.PackageManager)
	if pkgName == "" {
		return fmt.Errorf("package %s not available for %s", pkg.Name, i.OSInfo.PackageManager)
	}

	var cmd *exec.Cmd
	switch i.OSInfo.PackageManager {
	case "brew":
		cmd = exec.Command("brew", "install", pkgName)
	case "apt":
		cmd = exec.Command("sudo", "apt-get", "install", "-y", pkgName)
	case "pacman":
		cmd = exec.Command("sudo", "pacman", "-S", "--noconfirm", pkgName)
	case "yay":
		cmd = exec.Command("yay", "-S", "--noconfirm", pkgName)
	default:
		return fmt.Errorf("unsupported package manager: %s", i.OSInfo.PackageManager)
	}

	if i.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}

// installGoPackage installs a Go package using go install
func (i *Installer) installGoPackage(pkgPath string) error {
	cmd := exec.Command("go", "install", pkgPath)
	if i.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd.Run()
}

// installPipx installs a package using pipx
func (i *Installer) installPipx(pkg string) error {
	cmd := exec.Command("pipx", "install", pkg)
	if i.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd.Run()
}

// installPip installs a package using pip
func (i *Installer) installPip(pkg string) error {
	cmd := exec.Command("pip3", "install", pkg)
	if i.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd.Run()
}

// installNPM installs a package using npm
func (i *Installer) installNPM(pkg string) error {
	cmd := exec.Command("npm", "install", "-g", pkg)
	if i.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd.Run()
}

// installNPMGlobal installs a global npm package
func (i *Installer) installNPMGlobal(pkg string) error {
	i.Log.Action("Installing npm package: %s", pkg)
	if i.DryRun {
		return nil
	}
	return i.installNPM(pkg)
}

// installScript installs by running a script from a URL
func (i *Installer) installScript(url string) error {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("curl -fsSL %s | bash", url))
	cmd.Stdin = os.Stdin
	if i.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd.Run()
}

// runCommand runs a shell command
func (i *Installer) runCommand(cmdStr string) error {
	i.Log.Action("Running: %s", cmdStr)
	if i.DryRun {
		return nil
	}

	cmd := exec.Command("bash", "-c", cmdStr)
	if i.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd.Run()
}

// isInstalled checks if a package is already installed
func (i *Installer) isInstalled(pkg *config.Package) bool {
	// Try to find the command in PATH
	_, err := exec.LookPath(pkg.Name)
	return err == nil
}
