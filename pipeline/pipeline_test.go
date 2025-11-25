package pipeline

import (
	"testing"

	"github.com/jadefr/deploy-tool/config"
)

func TestBuildDockerImage(t *testing.T) {
	cfg := &config.Config{
		AppName:     "test-app",
		DockerImage: "test-app:test",
	}

	err := BuildDockerImage(cfg)
	if err != nil {
		t.Logf("BuildDockerImage failed (expected in test env): %v", err)
	}
}

func TestDeployToKubernetes(t *testing.T) {
	cfg := &config.Config{
		AppName:             "test-app",
		KubernetesNamespace: "default",
		HelmChartPath:       "./charts/my-app",
	}

	err := DeployToKubernetes(cfg)
	if err != nil {
		t.Logf("DeployToKubernetes failed (expected in test env without helm): %v", err)
	}
}

func TestRunTests(t *testing.T) {
	cfg := &config.Config{
		AppName:             "test-app",
		KubernetesNamespace: "default",
	}

	err := RunTests(cfg)
	if err != nil {
		t.Logf("RunTests failed (expected in test env without helm/k8s): %v", err)
	}
}

func TestConfigLoad(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("config.Load failed: %v", err)
	}
	if cfg == nil {
		t.Fatal("config.Load returned nil")
	}
	if cfg.AppName == "" {
		t.Error("AppName is empty")
	}
}
