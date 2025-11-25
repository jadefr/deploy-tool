# Makefile for building and running deploy-tool locally or in the debug container

APP_NAME := deploy-tool
DEBUG_IMAGE := deploy-tool-debug

.PHONY: all build run run-host build-debug run-debug run-debug-shell clean

all: build

# Build binary for host
build:
	go build -o $(APP_NAME) .

# Run the binary on the host (uses host docker/helm)
run-host: build
	./$(APP_NAME)

# Build the debug image (includes docker CLI, helm, kubectl)
build-debug:
	docker build -f Dockerfile.debug -t $(DEBUG_IMAGE) .

# Run the debug container, mount docker socket and repo; skip deploy/tests by default
# Example: make run-debug SKIP_DEPLOY=0 SKIP_TEST=0
run-debug: build-debug
	@echo "Running $(DEBUG_IMAGE) with mounted repo and docker socket"
	docker run --rm \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v "$(shell pwd)":/app \
		-e APP_NAME=my-app \
		-e DOCKER_IMAGE=my-app:local \
		$(if $(SKIP_BUILD),-e SKIP_BUILD=$(SKIP_BUILD),) \
		$(if $(SKIP_DEPLOY),-e SKIP_DEPLOY=$(SKIP_DEPLOY),-e SKIP_DEPLOY=1) \
		$(if $(SKIP_TEST),-e SKIP_TEST=$(SKIP_TEST),-e SKIP_TEST=1) \
		$(DEBUG_IMAGE)

# Run a shell in the debug container (for interactive debugging)
run-debug-shell: build-debug
	docker run --rm -it \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v "$(shell pwd)":/app \
		-e APP_NAME=my-app \
		-e DOCKER_IMAGE=my-app:local \
		$(DEBUG_IMAGE) /bin/bash

test:
	bash test.sh

test-helm-dryrun: build
	SKIP_BUILD=1 SKIP_TEST=1 \
	HELM_CHART_PATH=./charts/my-app \
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)
