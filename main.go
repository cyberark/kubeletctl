package main

import (
	"kubeletctl/cmd"
	_ "kubeletctl/cmd/log"
	_ "kubeletctl/cmd/metrics"
	_ "kubeletctl/cmd/proxy"
	_ "kubeletctl/cmd/proxy/debug"
	_ "kubeletctl/cmd/proxy/healthz"
	_ "kubeletctl/cmd/scan"
	_ "kubeletctl/cmd/spec"
	_ "kubeletctl/cmd/stats"
)

// build for release go build -ldflags "-s -w" (no symbols and debug info)
// TODO: Add tests
// TODO: Use vendor folder
func main() {
	cmd.Execute()
}