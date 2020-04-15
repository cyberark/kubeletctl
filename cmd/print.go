package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/tidwall/pretty"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	"log"
	"net/http"
	"os"
)

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

func isJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil
}

func PrintPods(podList v1.PodList) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Pod", "Namespace", "Containers"})

	for _, pod := range podList.Items {
		var containersString string
		for _, container := range pod.Spec.Containers {
			containersString += container.Name + "\n"
		}

		tw.AppendRow([]interface{}{pod.Name, pod.Namespace, containersString})
	}

	tw.SetTitle("Pods from Kubelet")
	tw.SetStyle(table.StyleLight)
	tw.Style().Title.Align = text.AlignCenter
	tw.SetAutoIndex(true)
	tw.Style().Options.SeparateRows = true
	fmt.Println(tw.Render())
}

func PrintPrettyHttpResponse(resp *http.Response, err error) {
	if err != nil {
		fmt.Printf("[*] Failed to run HTTP request with error: %s\n", err)
		os.Exit(1)
	}
	
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)
	
	if resp.StatusCode == http.StatusOK {

		// TODO: consider changing it by checking the first and the last byte "{..}".
		// Notice that there is line feed (byte 10 in decimal) that need to consider
		// if isJSON(jsonStringData){
		if  isJSON(bodyString) {
			jsonByteData := pretty.Pretty(bodyBytes)
			fmt.Println(string(jsonByteData))
		} else {
			// failed to parse JSON
			fmt.Println(string(bodyBytes))
		}
	} else {
		fmt.Printf("[*] The reponse failed with status: %d\n", resp.StatusCode)
		fmt.Printf("[*] Message: %s\n", bodyString)
	}
}