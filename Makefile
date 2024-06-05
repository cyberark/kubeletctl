.PHONY: all clean linux windows darwin docker docker-release

BUILD_DIR=build

all:
	go mod vendor;
	go fmt ./...;
	mkdir -p $(BUILD_DIR);
	GOARCH=386 GOOS=linux go build -v -o $(BUILD_DIR)/kubeletctl_linux_386;
	GOARCH=amd64 GOOS=linux go build -v -o $(BUILD_DIR)/kubeletctl_linux_amd64;
	GOARCH=arm64 GOOS=linux go build -v -o $(BUILD_DIR)/kubeletctl_linux_arm64;
	GOARCH=386 GOOS=windows go build -v -o $(BUILD_DIR)/kubeletctl_windows_386.exe;
	GOARCH=amd64 GOOS=windows go build -v -o $(BUILD_DIR)/kubeletctl_windows_amd64.exe;
	# GOARCH=386 GOOS=darwin go build -v -o $(BUILD_DIR)/kubeletctl_darwin_386; # Commented as Darwin 386 builds may not be supported
	GOARCH=amd64 GOOS=darwin go build -v -o $(BUILD_DIR)/kubeletctl_darwin_amd64;
	GOARCH=arm64 GOOS=darwin go build -v -o $(BUILD_DIR)/kubeletctl_darwin_arm64;
	
linux: linux_386 linux_amd64 linux_arm64

linux_386:
	GOARCH=386 GOOS=linux go build -v -o $(BUILD_DIR)/kubeletctl_linux_386;

linux_amd64:
	GOARCH=amd64 GOOS=linux go build -v -o $(BUILD_DIR)/kubeletctl_linux_amd64;

linux_arm64:
	GOARCH=arm64 GOOS=linux go build -v -o $(BUILD_DIR)/kubeletctl_linux_arm64;

windows: windows_386 windows_amd64

windows_386:
	GOARCH=386 GOOS=windows go build -v -o $(BUILD_DIR)/kubeletctl_windows_386.exe;

windows_amd64:
	GOARCH=amd64 GOOS=windows go build -v -o $(BUILD_DIR)/kubeletctl_windows_amd64.exe;

darwin: darwin_amd64 darwin_arm64

# darwin_386: # Commented out as macOS 386 builds may not be applicable
# 	GOARCH=386 GOOS=darwin go build -v -o $(BUILD_DIR)/kubeletctl_darwin_386;

darwin_amd64:
	GOARCH=amd64 GOOS=darwin go build -v -o $(BUILD_DIR)/kubeletctl_darwin_amd64;

darwin_arm64:
	GOARCH=arm64 GOOS=darwin go build -v -o $(BUILD_DIR)/kubeletctl_darwin_arm64;

docker:
	docker build . -t kubeletctl:latest

docker-release:
	docker build -t kubeletctl:release -f Dockerfile.latest .

clean:
	rm -rf $(BUILD_DIR)/
