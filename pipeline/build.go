package pipeline

import (
	"fmt"
	"os/exec"
)

func BuildDockerImage(cfg *Config) error {
	dockerCmd := exec.Command("docker", "build", "-t", cfg.DockerImage, ".")
	output, err := dockerCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("docker build failed: %v, output: %s", err, string(output))
	}
	fmt.Printf("Docker image built: %s\n", cfg.DockerImage)
	return nil
}