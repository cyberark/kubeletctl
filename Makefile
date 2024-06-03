.PHONY: all linux-arm64 clean docker docker-release

all: linux-arm64

linux-arm64:
	@echo "Running go mod vendor"
	go mod vendor
	@echo "Running go fmt"
	go fmt ./...
	@echo "Creating build directory"
	mkdir -p build
	@echo "Running gox for linux/arm64"
	gox -verbose -osarch="linux/arm64" -output="build/kubeletctl_{{.OS}}_{{.Arch}}"

docker:
	docker build . -t kubeletctl:latest

docker-release:
	docker build -t kubeletctl:release -f Dockerfile.latest

clean:
	rm -rf build
