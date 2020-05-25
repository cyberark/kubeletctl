package api

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"net/http"
	"os"
	"time"
	"fmt"
)

var GlobalClient *http.Client

// a struct to hold the result from each request including an index
// which will be used for sorting the results after they come in
type Result struct {
	Url      string
	res      http.Response
	HttpVerb HTTPVerb
	err      error
}

type HTTPVerb string

const (
	GET    HTTPVerb = "GET"
	POST   HTTPVerb = "POST"
	DELETE HTTPVerb = "DELETE"
)

func InitHttpClient() {

	insecure := true
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	config, err := kubeConfig.ClientConfig()
	if err != nil && len(os.Getenv(clientcmd.RecommendedConfigPathEnvVar)) > 0 {
		fmt.Print("[*] There is a problem with the file in KUBECONFIG environment variable")
		panic(err.Error())
	}

	var tr *http.Transport

	if config != nil {
		fmt.Print("[*] Using KUBECONFIG environment variable")
		tr = getHttpTransportWithCertificates(config, insecure)
	} else {
		tr = &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
		}
	}

	GlobalClient = &http.Client{
		Transport: tr,
		Timeout:   time.Second * 20,
	}
}

func getHttpTransportWithCertificates(config *restclient.Config, insecure bool) *http.Transport {
	var cert tls.Certificate
	var err error
	// TODO: need to handle a rare case where cert is file and key is data and the opposite.
	// Load client cert
	if config.TLSClientConfig.CertFile != "" {
		cert, err = tls.LoadX509KeyPair(config.TLSClientConfig.CertFile, config.TLSClientConfig.KeyFile)
	} else {
		cert, err = tls.X509KeyPair(config.TLSClientConfig.CertData, config.TLSClientConfig.KeyData)
	}

	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	var caCert []byte
	if config.TLSClientConfig.CAFile != "" {
		caCert, err = ioutil.ReadFile(config.TLSClientConfig.CAFile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		caCert = config.TLSClientConfig.CAData
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			RootCAs:            caCertPool,
			InsecureSkipVerify: insecure},
	}

	return tr
}

func InitGlobalClientFromFile(kubeconfig string) {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		panic(err.Error())
	}

	insecure := true
	tr := getHttpTransportWithCertificates(config, insecure)
	GlobalClient = &http.Client{
		Transport: tr,
		Timeout:   time.Second * 20,
	}

}

func InitGlobalClientFromCertificatesFiles(serverAddress string, caFile string, certFile string, keyFile string) {

	config := restclient.Config{
		Host: serverAddress,

		TLSClientConfig: restclient.TLSClientConfig{
			Insecure: false,
			CertFile: certFile,
			KeyFile:  keyFile,
			CAFile:   caFile,
		},
	}

	insecure := true
	tr := getHttpTransportWithCertificates(&config, insecure)
	GlobalClient = &http.Client{
		Transport: tr,
		Timeout:   time.Second * 20,
	}
}

func GetRequest(client *http.Client, url string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)

	//req.Header.Set("Authorization", "Bearer " + BEARER_TOKEN)
	resp, err := (*client).Do(req)

	return resp, err
}

func PutRequest(client *http.Client, url string, bodyData []byte) (*http.Response, error) {
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(bodyData))

	req.Header.Set("Content-Type", "text/plain")
	resp, err := (*client).Do(req)

	return resp, err
}

func PostRequest(client *http.Client, url string, bodyData []byte) (*http.Response, error) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bodyData))
	//req.Header.Set("Authorization", "Bearer " + BEARER_TOKEN)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := (*client).Do(req)

	return resp, err
}

/*
func PostRequest2(client *http.Client, url string, bodyData []byte) (*http.Response, error){
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bodyData))
	req.Header.Add("X-Stream-Protocol-Version", "v2.channel.k8s.io")
	req.Header.Add("X-Stream-Protocol-Version", "channel.k8s.io")
	req.Header.Add("Upgrade", "SPDY/3.1")
	req.Header.Add("Connection","upgrade")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := (*client).Do(req)

	return resp, err
}*/
