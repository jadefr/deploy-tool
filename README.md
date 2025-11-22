# Deploy Tool

A lightweight continuous deployment utility written in Go that automates testing, Docker image builds, image publishing, and Helm-based deployments to Kubernetes clusters. It’s opinionated, easy to extend, and designed to be integrated into local workflows or CI pipelines.

## Features
- Detect and run pipeline stages: test → build → push → deploy
- Build Docker images and optionally push to a registry
- Deploy using Helm (install / upgrade) into a Kubernetes namespace
- Configuration via environment variables for simple integration
- Extensible pipeline: add scanning, notifications, rollbacks, multi-cluster steps

## Quickstart
### Prerequisites
- Go 1.18+
- Docker (for build/push)
- Helm (for deploy)
- kubectl configured for the target cluster (if deploying to Kubernetes)

### Clone and prepare
```
git clone https://jadefr/deploy-tool.git
cd deploy-tool
go mod tidy
```

### Set environment variables (eg)
```
export APP_NAME=my-app
export DOCKER_IMAGE=my-registry/my-app:latest
export K8S_NAMESPACE=default
export HELM_CHART_PATH=./charts/my-app
```

### Run locally
```
go run main.go
```

### Build a binary
```
go build -o deploy-tool .
./deploy-tool
```

### Run pipeline steps manually (debug)
```
# tests
go test ./...

# build
docker build -t $DOCKER_IMAGE .

# push
docker push $DOCKER_IMAGE

# helm deploy
helm upgrade --install $APP_NAME $HELM_CHART_PATH --namespace $K8S_NAMESPACE
```


## Configuration
The tool uses environment variables (defaults shown):
- APP_NAME (default: meu-app) — release name for Helm
- DOCKER_IMAGE (default: meu-app:latest) — image tag to build/push
- KUBE_NAMESPACE (default: default) — target Kubernetes namespace
- HELM_CHART_PATH (default: ./charts/meu-app) — path to Helm chart

You can change Load() in config/config.go to support .env files or CLI flags.

## Extending the tool (ideas)
- Stream command stdout/stderr and add structured logging
- Add registry authentication (Docker Hub, ECR, GCR)
- Add Helm value injection and templating per environment
- Implement rollback and release history management using Helm APIs
- Create a small web UI or gRPC API to trigger and monitor runs
- Add a GitHub Actions / GitLab CI example that triggers the tool in CI

## Contributing
- Fork the repo, create a branch, implement features or fixes, open a PR.
- Keep go.mod tidy and include tests for new behavior.
- Don’t commit secrets or kubeconfigs; add them to .gitignore
