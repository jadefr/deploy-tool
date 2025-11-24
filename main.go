package main

import (
	"fmt"
	"log"

	"github.com/jadefr/deploy-tool/config"
	"github.com/jadefr/deploy-tool/pipeline"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("🚀 Initializing the deploy pipeline: %+v\n", cfg)

	if err := pipeline.BuildDockerImage(cfg); err != nil {
		log.Fatalf("❌ build step failed: %v", err)
	}

	if err := pipeline.DeployToKubernetes(cfg); err != nil {
		log.Fatalf("❌ deploy step failed: %v", err)
	}

	if err := pipeline.RunTests(cfg); err != nil {
		log.Fatalf("❌ test step failed: %v", err)
	}

	fmt.Printf("🎉 Deployment pipeline completed successfully!")
}