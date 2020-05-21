.PHONY: all

all:
	go mod vendor;
	go fmt ./...;
	mkdir -p build;
	GOFLAGS=-mod=vendor gox -ldflags "-s -w" --osarch="linux/amd64" --osarch="windows/amd64" --osarch="darwin/amd64" -output "build/kubeletctl_{{.OS}}_{{.Arch}}"
