package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"
	"log"
	"net"
	"net/http"
	"time"

	v1 "k8s.io/api/core/v1"
)

func FindContainersWithRCE(nodeIP string) Node {
	/*hosts, Err := getHosts(cidr)
	if Err != nil {
		log.Fatal(Err)
	}*/

	node := getNodeWithPodsRCECheck(nodeIP)
	return node
}

// TODO: improve to multi-threaded
func checkPodsForRCE(nodeIP string, pods v1.PodList) []Pod {
	command := "?cmd=ls"
	var nodePods []Pod

	for _, pod := range pods.Items {
		var podContainers []Container
		for _, container := range pod.Spec.Containers {
			containerRCERun := false
			apiPathUrl := fmt.Sprintf("%s://%s:%s%s/%s/%s/%s%s", cmd.ProtocolScheme, nodeIP, cmd.PortFlag, api.RUN, pod.Namespace, pod.Name, container.Name, command)
			resp, err := api.PostRequest(api.GlobalClient, apiPathUrl, []byte{})

			// TODO: check if this check is enough
			if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
				containerRCERun = true
			}
			podContainers = append(podContainers, Container{
				Name:    container.Name,
				RCEExec: false,
				RCERun:  containerRCERun,
			})
		}

		nodePods = append(nodePods, Pod{
			Name:       pod.Name,
			Namespace:  pod.Namespace,
			Containers: podContainers,
		})
	}

	return nodePods
}

// TODO: use this function inside getNodeWithPodsRCECheck() and move checkPodsForRCE outside
func GetPodListFromNodeIP(nodeIP string) (v1.PodList, error) {
	var pods v1.PodList

	kubeletPodsUrl := fmt.Sprintf("%s://%s:%s%s", cmd.ProtocolScheme, nodeIP, cmd.PortFlag, api.PODS)
	resp, err := api.GetRequest(api.GlobalClient, kubeletPodsUrl)

	// TODO: maybe send the error from this function
	if err != nil {
		fmt.Printf("[*] Failed to get pods from: %s\n", nodeIP)
	} else {

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("[*] Failed to read pods JSON from: %s\n", nodeIP)
		} else {
			err = json.Unmarshal(bodyBytes, &pods)
			if err != nil {
				fmt.Printf("[*] Failed to parse pods JSON from: %s\n", nodeIP)
			}
		}
	}

	return pods, err
}

func getNodeWithPodsRCECheck(nodeIP string) Node {
	node := Node{
		IPAddress: nodeIP,
		Pods:      nil,
	}
	var pods v1.PodList
	kubeletPodsUrl := fmt.Sprintf("%s://%s:%s%s", cmd.ProtocolScheme, nodeIP, cmd.PortFlag, api.PODS)
	resp, err := api.GetRequest(api.GlobalClient, kubeletPodsUrl)

	// TODO: maybe send the error from this function
	if err != nil {
		fmt.Printf("[*] Failed to get pods from: %s\n", nodeIP)
	} else {

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("[*] Failed to read pods JSON from: %s\n", nodeIP)
		} else {
			err = json.Unmarshal(bodyBytes, &pods)
			if err != nil {
				fmt.Printf("[*] Failed to parse pods JSON from: %s\n", nodeIP)
			} else {
				nodePods := checkPodsForRCE(nodeIP, pods)
				node.Pods = nodePods
			}
		}
	}

	return node
}

// TODO: add an option for the user to change it
const CONCURRENCY_DEFAULT_LIMIT = 50

func FindOpenedKubeletOnNodes(cidr string) []string {
	hosts, err := getHosts(cidr)
	if err != nil {
		log.Fatal(err)
	}

	vulnerableAddresses := boundedParallelGet(hosts, CONCURRENCY_DEFAULT_LIMIT, SCAN_MODE_HEALTH)

	return vulnerableAddresses
}

func isPortOpen(host string, ports []string) bool {
	isOpen := false
	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			//fmt.Println("Connecting error:", Err)
		}
		if conn != nil {
			defer conn.Close()
			isOpen = true
			//fmt.Println("Opened", net.JoinHostPort(host, port))
		}
	}

	return isOpen
}

/*
func isHostReachable(Url string) bool {
	timeout := time.Duration(2 * time.Second)
	isReach := false
	_, Err := net.DialTimeout("tcp", Url, timeout)
	if Err == nil {
		isReach = true
	}

	return isReach
}*/

// a struct to hold the Result from each request including an index
// which will be used for sorting the results after they come in
type Result struct {
	StatusCode int
	Url        string
	Err        error
	Node       Node
}

type Node struct {
	IPAddress string
	Pods      []Pod
}

type Pod struct {
	Name       string
	Namespace  string
	Containers []Container
}

type Container struct {
	Name    string
	RCEExec bool
	RCERun  bool
}

type ScanMode string

const (
	SCAN_MODE_RCE    = "RCE"
	SCAN_MODE_HEALTH = "HEALTH"
)

// https://gist.github.com/thedevsaddam/69eacea490a1e9b432f9e1ada1d79106
// boundedParallelGet sends requests in parallel but only up to a certain
// limit, and furthermore it's only parallel up to the amount of CPUs but
// is always concurrent up to the concurrency limit
func boundedParallelGet(urls []string, concurrencyLimit int, scanMode ScanMode) []string {
	// make a slice to hold the results we're expecting
	var vulnerableNodes []string

	// this buffered channel will block at the concurrency limit
	semaphoreChan := make(chan struct{}, concurrencyLimit)

	// this channel will not block and collect the http request results
	resultsChan := make(chan *Result)

	// make sure we close these channels when we're done with them
	defer func() {
		close(semaphoreChan)
		close(resultsChan)
	}()

	// keen an index and loop through every Url we will send a request to
	for i, url := range urls {

		// start a go routine with the index and Url in a closure
		go func(i int, ipAddress string) {

			// this sends an empty struct into the semaphoreChan which
			// is basically saying add one to the limit, but when the
			// limit has been reached block until there is room
			semaphoreChan <- struct{}{}

			var err error
			var resp *http.Response
			var kubeletUrlAddress string

			statusCode := 0
			kubeletUrlAddress = ipAddress
			node := Node{}

			// TODO: should we change 10250 to variable in case the kubelet will be in different port ?
			// If not we should consider a constant
			if isPortOpen(ipAddress, []string{cmd.PortFlag}) {

				if scanMode == SCAN_MODE_RCE {
					node = FindContainersWithRCE(ipAddress)
				} else {
					kubeletUrlAddress = fmt.Sprintf("%s://%s:%s%s", cmd.ProtocolScheme, ipAddress, cmd.PortFlag, api.HEALTHZ)
					// send the request and put the response in a Result struct
					// along with the index so we can sort them later along with
					// any error that might have occoured
					resp, err = api.GetRequest(api.GlobalClient, kubeletUrlAddress)
					if resp != nil {
						statusCode = resp.StatusCode
						kubeletUrlAddress = fmt.Sprintf("%s://%s:%s", cmd.ProtocolScheme, ipAddress, cmd.PortFlag)
					}
				}
			}

			result := &Result{statusCode, kubeletUrlAddress, err, node}

			// now we can send the Result struct through the resultsChan
			resultsChan <- result
			// once we're done it's we read from the semaphoreChan which
			// has the effect of removing one from the limit and allowing
			// another goroutine to start
			<-semaphoreChan

		}(i, url)
	}

	// start listening for any results over the resultsChan
	// once we get a Result append it to the Result slice
	var count int
	for {
		result := <-resultsChan
		count += 1

		if result.StatusCode == http.StatusOK {
			vulnerableNodes = append(vulnerableNodes, result.Url)
		}

		// if we've reached the expected amount of urls then stop
		if count == len(urls) {
			break
		}
	}

	// now we're done we return the results
	return vulnerableNodes
}
