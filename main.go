package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jadefr/deploy-tool/config"
	"github.com/jadefr/deploy-tool/pipeline"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("🚀 Initializing the deploy pipeline: %+v\n", cfg)

	if os.Getenv("SKIP_BUILD") == "" {
		if err := pipeline.BuildDockerImage(cfg); err != nil {
			log.Fatalf("❌ build step failed: %v", err)
		}
	} else {
		fmt.Println("⏭️  SKIP_BUILD set, skipping build step")
	}

	if os.Getenv("SKIP_DEPLOY") == "" {
		if err := pipeline.DeployToKubernetes(cfg); err != nil {
			log.Fatalf("❌ deploy step failed: %v", err)
		}
	} else {
		fmt.Println("⏭️  SKIP_DEPLOY set, skipping deploy step")
	}

	if os.Getenv("SKIP_TEST") == "" {
		if err := pipeline.RunTests(cfg); err != nil {
			log.Fatalf("❌ test step failed: %v", err)
		}
	} else {
		fmt.Println("⏭️  SKIP_TEST set, skipping test step")
	}

	fmt.Printf("🎉 Deployment pipeline completed successfully!")
}
