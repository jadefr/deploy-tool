package config

import (
	"os"
)

type Config struct {
	AppName             string
	DockerImage         string
	KubernetesNamespace string
	HelmChartPath       string
}

func Load() (*Config, error) {
	return &Config{
		AppName:             getEnv("APP_NAME", "my-app"),
		DockerImage:         getEnv("DOCKER_IMAGE", "my-app:latest"),
		KubernetesNamespace: getEnv("K8S_NAMESPACE", "default"),
		HelmChartPath:       getEnv("HELM_CHART_PATH", "./charts/my-app"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}