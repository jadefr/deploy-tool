package pipeline

import (
	"fmt"
	"os/exec"

	"github.com/jadefr/deploy-tool/config"
)

func BuildDockerImage(cfg *config.Config) error {
	dockerCmd := exec.Command("docker", "build", "-t", cfg.DockerImage, ".")
	output, err := dockerCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("docker build failed: %v, output: %s", err, string(output))
	}
	fmt.Printf("Docker image built: %s\n", cfg.DockerImage)
	return nil
}