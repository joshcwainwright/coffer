image := "coffer"
registry := "ghcr.io/joshcwainwright"
env_file := "$HOME/.config/coffer/.env"

# List available recipes
default:
    @just --list

# Run natively against the production env file
run: (_run env_file)

# Run natively against the sandbox env file
run-sandbox: (_run (env_file + ".sandbox"))

# Shared runner: source an env file, then run from the repo root so ./data/coffer.db is stable
_run file:
    #!/usr/bin/env bash
    set -euo pipefail
    if [ ! -f "{{file}}" ]; then
        echo "missing env file: {{file}}" >&2
        echo "copy .env.example to it, then chmod 600" >&2
        exit 1
    fi
    set -a && . "{{file}}" && set +a
    exec go run ./cmd/coffer

# Run the Go tests
test:
    go test ./...

# Vet and compile-check without producing a binary
check:
    go vet ./... && go build ./...

# Tidy go.mod / go.sum
tidy:
    go mod tidy

# Build the Docker image for this machine's architecture
build:
    docker build -t {{image}}:dev .

# Run the built image against the local named volume
run-image: build
    docker run --rm -v coffer-data:/data {{image}}:dev

# Run the tests inside the build stage (Linux parity check)
test-docker:
    docker build --target build -t {{image}}:test .
    docker run --rm -v coffer-gomod:/go/pkg/mod {{image}}:test go test ./...

# Build and push a multi-arch release image, e.g. `just release v1`
release version:
    docker buildx build --platform linux/amd64,linux/arm64 \
      -t {{registry}}/{{image}}:{{version}} --push .
