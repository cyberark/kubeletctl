package api

// To use the go-client imports, you need to install it based on k8s
// https://github.com/kubernetes/client-go/blob/master/INSTALL.md#enabling-go-modules
// set GO111MODULE=on
// go mod init
// go get k8s.io/client-go@master
// go build

import (
	"crypto/tls"
	"fmt"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"net/http"
	"net/url"
	"os"
	"time"
)

// To investiage and compare against kubectl, use the --v=9 in kubectl:
//https://stackoverflow.com/a/51859036

// https://stackoverflow.com/questions/43314689/example-of-exec-in-k8ss-pod-by-using-go-client/54317689
func Exec(serverIp string, serverPort string, serverFullAddress string, apiPath string, queryCommands string, method string) {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}

	config := &restclient.Config{
		//Host:                "https://<ip>:10250",
		Host: serverFullAddress,
		//APIPath:             "/exec/default/<pod_name>/<contianer_name>",
		APIPath: apiPath,
		TLSClientConfig: restclient.TLSClientConfig{
			Insecure: true,
		},
		Transport: tr,
	}

	urlObject := &url.URL{
		Scheme: "https",
		Opaque: "",
		User:   nil,
		//Host:       "<ip>>:10250",
		Host: fmt.Sprintf("%s:%s", serverIp, serverPort),
		//Path:       "/exec/default/<pod_name>/<contianer_name>",
		Path:    apiPath,
		RawPath: "",
		//RawQuery:   "command=ls&command=/&input=1&output=1&tty=1",
		RawQuery: fmt.Sprintf("%s&input=1&output=1&tty=1", queryCommands),
	}

	exec, err := remotecommand.NewSPDYExecutor(config, method, urlObject)
	if err != nil {
		fmt.Println(err)
	}

	// Credit to this blog post https://www.henryxieblogs.com/2019/05/
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Tty:    true,
	})
}
