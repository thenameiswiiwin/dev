package main

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/thenameiswiiwin/dev/internal/config"
)

var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Kubernetes deployment operations",
	Long:  `Manage Kubernetes deployments for presets`,
}

var k8sRenderCmd = &cobra.Command{
	Use:   "render [preset]",
	Short: "Render Kubernetes manifests",
	Long:  `Use kustomize to render Kubernetes manifests for a preset`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		preset := args[0]

		repoRoot, err := config.FindRepoRoot()
		if err != nil {
			log.Error("Failed to find repository root: %v", err)
			os.Exit(1)
		}

		deployPath := filepath.Join(repoRoot, "deploy", preset)
		if _, err := os.Stat(deployPath); os.IsNotExist(err) {
			log.Error("No Kubernetes manifests found for preset: %s", preset)
			log.Info("Available presets: python, go, rust, web")
			os.Exit(1)
		}

		log.Info("Rendering Kubernetes manifests for preset: %s", preset)

		if dryRun {
			log.DryRunMsg("Would render manifests from: %s", deployPath)
			return
		}

		// Check if kustomize is available
		if !commandExists("kustomize") && !commandExists("kubectl") {
			log.Error("Neither kustomize nor kubectl found")
			log.Info("Install kustomize: https://kubectl.docs.kubernetes.io/installation/kustomize/")
			os.Exit(1)
		}

		// Try kustomize first, fall back to kubectl
		var cmdExec *exec.Cmd
		if commandExists("kustomize") {
			cmdExec = exec.Command("kustomize", "build", deployPath)
		} else {
			cmdExec = exec.Command("kubectl", "kustomize", deployPath)
		}

		cmdExec.Stdout = os.Stdout
		cmdExec.Stderr = os.Stderr

		if err := cmdExec.Run(); err != nil {
			log.Error("Failed to render manifests: %v", err)
			os.Exit(1)
		}

		log.Success("Manifests rendered successfully")
	},
}

func init() {
	rootCmd.AddCommand(k8sCmd)
	k8sCmd.AddCommand(k8sRenderCmd)
}
