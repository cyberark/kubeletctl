.PHONY: all linux-arm64 docker docker-release

all:
	go mod vendor
	go fmt ./...
	mkdir -p build
	GOFLAGS=-mod=vendor gox -ldflags "-s -w" --osarch="linux/386" --osarch="linux/amd64" --osarch="linux/arm64" --osarch="windows/386" --osarch="windows/amd64" --osarch="darwin/386" --osarch="darwin/amd64" --osarch="darwin/arm64" -output "build/kubeletctl_{{.OS}}_{{.Arch}}"

linux-arm64:
	go mod vendor
	go fmt ./...
	mkdir -p build
	GOFLAGS=-mod=vendor gox -ldflags "-s -w" --osarch="linux/arm64" -output "build/kubeletctl_linux_arm64"

docker:
	docker build . -t kubeletctl:latest

docker-release:
	docker build -t kubeletctl:release -f Dockerfile.latest .
