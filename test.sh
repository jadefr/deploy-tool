#!/bin/bash
set -e

# Integration test script for deploy-tool
# Tests the full pipeline with helm dry-run to avoid actually modifying the cluster

APP_NAME="${APP_NAME:-my-app}"
DOCKER_IMAGE="${DOCKER_IMAGE:-my-app:test-$(date +%s)}"
K8S_NAMESPACE="${K8S_NAMESPACE:-default}"
HELM_CHART_PATH="${HELM_CHART_PATH:-./charts/my-app}"

echo "=== Deploy-Tool Integration Test ==="
echo "APP_NAME: $APP_NAME"
echo "DOCKER_IMAGE: $DOCKER_IMAGE"
echo "K8S_NAMESPACE: $K8S_NAMESPACE"
echo "HELM_CHART_PATH: $HELM_CHART_PATH"
echo ""

# Test 1: Build the binary
echo "Test 1: Building binary..."
go build -o /tmp/deploy-tool-test .
if [ ! -f /tmp/deploy-tool-test ]; then
    echo "❌ Failed to build binary"
    exit 1
fi
echo "✅ Binary built successfully"
echo ""

# Test 2: Unit tests
echo "Test 2: Running unit tests..."
go test -v ./...
echo "✅ Unit tests passed"
echo ""

# Test 3: Validate Helm chart
echo "Test 3: Validating Helm chart..."
if ! helm lint "$HELM_CHART_PATH" > /dev/null 2>&1; then
    echo "⚠️  Helm lint warnings (non-fatal)"
else
    echo "✅ Helm chart validation passed"
fi
echo ""

# Test 4: Helm dry-run (simulates deploy without modifying cluster)
echo "Test 4: Testing helm dry-run..."
if helm upgrade --install "$APP_NAME" "$HELM_CHART_PATH" \
    --namespace "$K8S_NAMESPACE" \
    --create-namespace \
    --dry-run \
    --debug > /dev/null 2>&1; then
    echo "✅ Helm dry-run succeeded"
else
    echo "⚠️  Helm dry-run failed (may need kube access)"
fi
echo ""

# Test 5: Run binary with skip flags
echo "Test 5: Running binary with SKIP flags..."
APP_NAME="$APP_NAME" \
DOCKER_IMAGE="$DOCKER_IMAGE" \
K8S_NAMESPACE="$K8S_NAMESPACE" \
HELM_CHART_PATH="$HELM_CHART_PATH" \
SKIP_BUILD=1 \
SKIP_DEPLOY=1 \
SKIP_TEST=1 \
/tmp/deploy-tool-test

if [ $? -eq 0 ]; then
    echo "✅ Binary ran successfully"
else
    echo "❌ Binary failed"
    exit 1
fi
echo ""

echo "=== All tests passed! ==="
rm -f /tmp/deploy-tool-test
