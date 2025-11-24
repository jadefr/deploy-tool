package pipeline

import (
    "fmt"
    "os/exec"

    "github.com/jadefr/deploy-tool/config"
)

func RunTests(cfg *config.Config) error {
    testCmd := exec.Command("helm", "test", cfg.AppName, "--namespace", cfg.KubernetesNamespace)
    output, err := testCmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("tests failed: %v, output: %s", err, string(output))
    }
    fmt.Printf("Tests passed: %s\n", string(output))
    return nil
}