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
package spec

import (
	cmd "kubeletctl/cmd"
	"kubeletctl/pkg/api"

	"github.com/spf13/cobra"
)

/*
Spec command is calling getSpec(..) in server.go. Inside this function it calls GetCachedMachineInfo()
The function GetCachedMachineInfo() is in kubelet_getters.go and returns Kubelet's machineInfo which is cached MachineInfo returned by cadvisor

 */

// specCmd represents the spec command
var specCmd = &cobra.Command{
	Use:   "spec",
	Short: "Cached MachineInfo returned by cadvisor",
	Long: `Description:
  Cached MachineInfo returned by cadvisor.
  
  HTTP request: 
    GET /spec
  
  Example for usage:
    kubeletctl spec
  
  With curl:
  curl -k https://<node_ip>:10250/spec`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("spec called")
		apiPathUrl := cmd.ServerFullAddressGlobal + api.SPEC
		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
		cmd.PrintPrettyHttpResponse(resp, err)
		if resp != nil && resp.StatusCode == 404 {
		    println("[*] kubelet exposes the endpoint /spec only if cadvisor endpoints are enabled when the kubelet starts")
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(specCmd)
}
