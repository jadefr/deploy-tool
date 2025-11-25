# Deploy Tool

A lightweight continuous deployment utility written in Go that automates the full pipeline: build Docker images, and deploy to Kubernetes using Helm. Run it locally, in a container, or from CI/CD. Designed to be simple, extensible, and production-ready.

## Features

- **Three-stage pipeline**: build Docker image → deploy to Kubernetes → run tests
- **Selective execution**: skip any stage via env vars (`SKIP_BUILD`, `SKIP_DEPLOY`, `SKIP_TEST`)
- **Helm-native**: uses Helm for declarative, idempotent deployments
- **Configuration via env**: simple integration into any automation tool
- **Containerized**: includes `Dockerfile` for production and `Dockerfile.debug` for local testing
- **Testable**: unit tests + integration test script with helm dry-run support
- **No CI/CD required**: runs standalone; add to Jenkins, GitHub Actions, GitLab CI, or cron jobs

## Prerequisites

- **Go 1.25.4+** (for building)
- **Docker** (for building images, and to run the debug container)
- **Helm 3.x** (for deployments)
- **kubectl** (to interact with clusters)
- **kind** or **minikube** (optional, for local testing)

## Quick Start

### 1. Clone and build

```bash
git clone https://github.com/jadefr/deploy-tool.git
cd deploy-tool
go build -o deploy-tool .
```

### 2. (Optional) Create a local kind cluster

```bash
kind create cluster
kubectl cluster-info
```

### 3. Build and load a test image into kind

```bash
docker build -t my-app:local .
kind load docker-image my-app:local
```

### 4. Deploy using deploy-tool

```bash
export APP_NAME=my-app
export DOCKER_IMAGE=my-app:local
export K8S_NAMESPACE=default
export HELM_CHART_PATH=./charts/my-app

# Skip build (already done) and test (not set up); run deploy only
SKIP_BUILD=1 SKIP_TEST=1 ./deploy-tool
```

### 5. Verify deployment

```bash
kubectl get deployments -n default
kubectl get pods -n default
```

## Running Tests

### Unit tests and integration tests

```bash
make test
```

This runs:
- Binary build
- Unit tests (4 test functions)
- Helm chart validation
- Helm dry-run (simulates deploy without modifying cluster)
- Full pipeline with all stages skipped (config loading only)

### Run Helm deploy (dry-run or real)

```bash
# Dry-run only (no cluster changes)
make test-helm-dryrun

# Real deploy to kind cluster (if configured)
SKIP_BUILD=1 SKIP_TEST=1 ./deploy-tool
```

## Using Make targets

```bash
make build              # Build binary
make run-host           # Build and run binary on host (uses host docker/helm)
make build-debug        # Build Dockerfile.debug image with Docker CLI, Helm, kubectl
make run-debug          # Run debug container with repo and docker socket mounted
make run-debug-shell    # Interactive shell in debug container
make test               # Run full integration tests
make test-helm-dryrun   # Test helm dry-run
make clean              # Remove binary
```

### Example: Deploy in container

```bash
# Build debug image
make build-debug

# Run deploy in container (mounts docker socket to use host Docker)
make run-debug SKIP_BUILD=1 SKIP_TEST=1
```

## Configuration

Environment variables (defaults shown in `config/config.go`):

| Variable | Default | Purpose |
|----------|---------|---------|
| `APP_NAME` | `my-app` | Release name for Helm |
| `DOCKER_IMAGE` | `my-app:latest` | Docker image tag to build/push |
| `K8S_NAMESPACE` | `default` | Target Kubernetes namespace |
| `HELM_CHART_PATH` | `./charts/my-app` | Path to Helm chart directory |
| `SKIP_BUILD` | (not set) | Set to any value to skip build step |
| `SKIP_DEPLOY` | (not set) | Set to any value to skip deploy step |
| `SKIP_TEST` | (not set) | Set to any value to skip test step |

### Customizing config

Edit `config/config.go` to:
- Load from `.env` files
- Parse CLI flags
- Read from vaults or config servers

## Project Structure

```
deploy-tool/
├── main.go                 # Entry point
├── go.mod                  # Go module definition
├── Dockerfile              # Production image (distroless)
├── Dockerfile.debug        # Debug image (debian, includes docker/helm/kubectl)
├── Makefile                # Build and run targets
├── test.sh                 # Integration test script
├── config/
│   └── config.go           # Configuration loading
├── pipeline/
│   ├── build.go            # Docker build step
│   ├── deploy.go           # Helm deploy step
│   ├── test.go             # Helm test step
│   └── pipeline_test.go    # Unit tests
└── charts/
    └── my-app/             # Example Helm chart for testing
        ├── Chart.yaml
        ├── values.yaml
        └── templates/
            └── deployment.yaml
```

## Examples

### Local deployment to kind

```bash
# Create kind cluster
kind create cluster

# Build and load image
docker build -t my-app:local .
kind load docker-image my-app:local

# Deploy
export DOCKER_IMAGE=my-app:local
SKIP_BUILD=1 SKIP_TEST=1 ./deploy-tool

# Verify
kubectl get pods
kubectl get deployments my-app
```

### Deploy in Docker (containerized)

```bash
make build-debug
make run-debug SKIP_BUILD=1 SKIP_TEST=1
```

### Use in CI/CD (GitHub Actions example)

```yaml
- name: Deploy with deploy-tool
  env:
    APP_NAME: my-app
    DOCKER_IMAGE: my-registry/my-app:${{ github.sha }}
    K8S_NAMESPACE: production
    HELM_CHART_PATH: ./charts/my-app
  run: |
    go build -o deploy-tool .
    SKIP_TEST=1 ./deploy-tool
```

### Use in cron job

```bash
0 2 * * * /home/user/deploy-tool/deploy-tool >> /var/log/deploy-tool.log 2>&1
```

## Troubleshooting

### "could not import deploy-tool/config"

Ensure `go.mod` uses the correct module path: `module github.com/jadefr/deploy-tool`

### Docker build fails in Dockerfile

Ensure `Dockerfile` builder uses `golang:1.25.4` (matches `go.mod`)

### Helm deploy fails with "kubernetes cluster unreachable"

Configure `kubectl`: `kubectl config current-context` and `kubectl cluster-info`

### Container deploy fails to find docker socket

Mount the socket: `-v /var/run/docker.sock:/var/run/docker.sock`

## Extending the tool

- **Add pre-deploy checks**: smoke tests, image scanning
- **Add post-deploy validation**: health checks, canary rollouts
- **Add registry auth**: support for Docker Hub, ECR, GCR
- **Add templating**: inject env-specific Helm values
- **Add webhooks**: notify Slack/PagerDuty on deploy
- **Add UI/API**: simple dashboard or gRPC trigger

See pipeline functions in `pipeline/*.go` for extension points.

## Contributing

- Fork the repo, create a branch, implement your feature or fix
- Run `make test` to validate
- Keep `go.mod` tidy: `go mod tidy`
- Add tests for new behavior
- Submit a PR

## License

MIT
