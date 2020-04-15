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
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"

	"github.com/spf13/cobra"
)

// runningpodsCmd represents the runningpods command
var runningpodsCmd = &cobra.Command{
	Use:   "runningpods",
	Short: "Returns all pods running on kubelet from looking at the container runtime cache.",
	Long: `Description:
  
  Returns all pods running on kubelet from looking at the container runtime cache.
  HTTP request: 
    GET /runningpods
  
  Example for usage:
    kubeletctl.exe runningpods
  
  With curl:
    curl -k https://<node_ip>:10250/runningpods`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("runningpods called")
		apiPathUrl := cmd.ServerFullAddressGlobal + api.RUNNING_PODS
		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
		cmd.PrintPrettyHttpResponse(resp, err)
	},
}

func init() {
	cmd.RootCmd.AddCommand(runningpodsCmd)
}
