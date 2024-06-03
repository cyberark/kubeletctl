.PHONY: all linux-arm64 clean docker docker-release

all: linux-arm64

linux-arm64:
	go mod vendor
	go fmt ./...
	mkdir -p build
	gox -verbose -osarch="linux/arm64" -output="build/kubeletctl_{{.OS}}_{{.Arch}}"

docker:
	docker build . -t kubeletctl:latest

docker-release:
	docker build -t kubeletctl:release -f Dockerfile.latest

clean:
	rm -rf build
