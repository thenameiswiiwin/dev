package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/thenameiswiiwin/dev/internal/config"
)

var (
	buildTest bool
	buildTag  string
)

var buildCmd = &cobra.Command{
	Use:   "build [preset]",
	Short: "Build container image for a preset",
	Long:  `Build Docker/OCI container images for presets using containers/<preset>/Dockerfile`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoRoot, err := config.FindRepoRoot()
		if err != nil {
			log.Error("Failed to find repository root: %v", err)
			os.Exit(1)
		}

		if len(args) == 0 {
			// Build all presets
			presets := []string{"python", "go", "rust", "web"}
			for i, preset := range presets {
				log.Step(i+1, len(presets), "Building %s container", preset)
				if err := buildPreset(repoRoot, preset); err != nil {
					log.Error("Failed to build %s: %v", preset, err)
				}
			}
		} else {
			preset := args[0]
			log.Info("Building container for preset: %s", preset)
			if err := buildPreset(repoRoot, preset); err != nil {
				log.Error("Failed to build: %v", err)
				os.Exit(1)
			}
		}

		if buildTest {
			log.Info("\nRunning smoke tests...")
			if len(args) == 0 {
				for _, preset := range []string{"python", "go", "rust", "web"} {
					smokeTest(preset)
				}
			} else {
				smokeTest(args[0])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().BoolVar(&buildTest, "test", false, "Run smoke tests after building")
	buildCmd.Flags().StringVar(&buildTag, "tag", "latest", "Image tag")
}

func buildPreset(repoRoot, preset string) error {
	dockerfilePath := filepath.Join(repoRoot, "containers", preset, "Dockerfile")
	if _, err := os.Stat(dockerfilePath); os.IsNotExist(err) {
		return fmt.Errorf("Dockerfile not found for preset: %s", preset)
	}

	imageName := fmt.Sprintf("dev-env-%s:%s", preset, buildTag)
	log.Action("Building image: %s", imageName)

	if dryRun {
		return nil
	}

	cmd := exec.Command("docker", "build",
		"-t", imageName,
		"-f", dockerfilePath,
		filepath.Join(repoRoot, "containers", preset))

	if verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker build failed: %w", err)
	}

	log.Success("Built %s", imageName)
	return nil
}

func smokeTest(preset string) {
	imageName := fmt.Sprintf("dev-env-%s:%s", preset, buildTag)
	log.Info("Smoke testing %s...", imageName)

	if dryRun {
		log.DryRunMsg("Would run smoke test for %s", preset)
		return
	}

	// Check version commands based on preset
	var checkCmd string
	switch preset {
	case "python":
		checkCmd = "python3 --version && pip --version"
	case "go":
		checkCmd = "go version"
	case "rust":
		checkCmd = "rustc --version && cargo --version"
	case "web":
		checkCmd = "node --version && npm --version"
	default:
		log.Warn("Unknown preset: %s", preset)
		return
	}

	// Run smoke test
	cmd := exec.Command("docker", "run", "--rm", imageName, "/bin/bash", "-c", checkCmd)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Error("Smoke test failed for %s: %v", preset, err)
		if verbose {
			log.Info("Output: %s", string(output))
		}
	} else {
		log.Success("Smoke test passed for %s", preset)
		if verbose {
			log.Info("Output: %s", string(output))
		}
	}
}
