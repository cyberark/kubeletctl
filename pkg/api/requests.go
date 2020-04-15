package api

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"time"
)

var GlobalClient *http.Client

// a struct to hold the result from each request including an index
// which will be used for sorting the results after they come in
type Result struct {
	Url string
	res   http.Response
	HttpVerb HTTPVerb
	err   error
}

type HTTPVerb string
const (
	GET HTTPVerb = "GET"
	POST HTTPVerb = "POST"
	DELETE HTTPVerb = "DELETE"
)



func InitHttpClient() {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	GlobalClient = &http.Client{
		Transport: tr,
		Timeout: time.Second * 20,
	}
}

func GetRequest(client *http.Client, url string) (*http.Response, error){
	req, _ := http.NewRequest("GET", url, nil)

	//req.Header.Set("Authorization", "Bearer " + BEARER_TOKEN)
	resp, err := (*client).Do(req)

	return resp, err
}

func PutRequest(client *http.Client, url string, bodyData []byte) (*http.Response, error){
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(bodyData))

	req.Header.Set("Content-Type", "text/plain")
	resp, err := (*client).Do(req)

	return resp, err
}

func PostRequest(client *http.Client, url string, bodyData []byte) (*http.Response, error){
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