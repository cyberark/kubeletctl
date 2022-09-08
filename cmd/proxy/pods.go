/*
Copyright (c) 2020 CyberArk Software Ltd. All rights reserved

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"
	"log"
	"os"
)

// podsCmd represents the pods command
var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Get list of pods on the node",
	Long: `Description:
  Get list of pods on the node. 
  HTTP request: GET /pods
  Example for usage:
  kubeletctl.exe pods
  
  With curl:
  curl -k https://<node_ip>:10250/pods
`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("pods called")

		apiPathUrl := cmd.ServerFullAddressGlobal + api.PODS
		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)

		if cmd.RawFlag {
			cmd.PrintPrettyHttpResponse(resp, err)
		} else {
			if err != nil {
				fmt.Printf("[*] Failed to run HTTP request with error: %s\n", err)
				os.Exit(1)
			}

			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			var pods v1.PodList
			err = json.Unmarshal(bodyBytes, &pods)

			if err != nil {
				// TODO: this function does some of the previous checks, consider modification
				cmd.PrintPrettyHttpResponse(resp, err)
			} else {
				// TODO: consider using the "--namespace" flag to filter the result and print only from specific namespace
				cmd.PrintPods(pods)
			}
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(podsCmd)
}
