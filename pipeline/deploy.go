package pipeline

import (
	"fmt"
	"os/exec"

	"github.com/jadefr/deploy-tool/config"
)

func DeployToKubernetes(cfg *config.Config) error {
	helmCmd := exec.Command("helm", "upgrade", "--install", cfg.AppName, cfg.HelmChartPath, "--namespace", cfg.KubernetesNamespace)
	output, err := helmCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("helm deploy failed: %v, output: %s", err, string(output))
	}
	fmt.Printf("🚢 Deployed to Kubernetes: %s\n", string(output))
	return nil
}