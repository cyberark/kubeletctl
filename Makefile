.PHONY: all linux-arm64 clean docker docker-release

all: linux-arm64

linux-arm64:
	go mod vendor
	go fmt ./...
	mkdir -p build
	GOARCH=arm64 GOOS=linux go build -v -o build/kubeletctl_linux_arm64

docker:
	docker build . -t kubeletctl:latest

docker-release:
	docker build -t kubeletctl:release -f Dockerfile.latest

clean:
	rm -rf build
