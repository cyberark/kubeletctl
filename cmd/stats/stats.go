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
package stats

import (
	"fmt"
	cmd "kubeletctl/cmd"
	"kubeletctl/pkg/api"

	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Return statistical information for the resources in the node.",
	Long: `Description:
  Return statistical information for the resources in the node.
  
  HTTP requests: 
    GET /stats
    GET /stats/summary
    GET /stats/summary?only_cpu_and_memory=true
    GET /stats/container
    GET /stats/{namespace}/{podName}/{uid}/{containerName}
    GET /stats/{podName}/{containerName}
  
  Example for usage:
    kubeletctl stats 
  
    // Stats of the contianers
    kubeletctl stats container
    
    // Query only the cpu and memory fields
    kubeletctl stats summary --only-cpu-mem
  
  With curl:
    curl -k https://<node_ip>:10250/stats
    curl -k https://<node_ip>:10250/stats/summary
    curl -k https://<node_ip>:10250/stats/summary?only_cpu_and_memory=true
    curl -k https://<node_ip>:10250/stats/container`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("stats called")

		var apiPath string
		if cmd.NamespaceFlag != "" && cmd.PodFlag != "" && cmd.PodUidFlag != "" && cmd.ContainerFlag != "" {
			apiPath = fmt.Sprintf("%s%s/%s/%s/%s/%s", cmd.ServerFullAddressGlobal, api.STATS, cmd.NamespaceFlag, cmd.PodFlag, cmd.PodUidFlag, cmd.ContainerFlag)
		} else if cmd.PodUidFlag != "" && cmd.ContainerFlag != "" {
			apiPath = fmt.Sprintf("%s%s/%s/%s", cmd.ServerFullAddressGlobal, api.STATS, cmd.PodUidFlag, cmd.ContainerFlag)
		} else {
			apiPath = fmt.Sprintf("%s%s", cmd.ServerFullAddressGlobal, api.STATS)
		}

		resp, err := api.GetRequest(api.GlobalClient, apiPath)
		cmd.PrintPrettyHttpResponse(resp, err)
	},
}

func init() {
	cmd.RootCmd.AddCommand(statsCmd)
}
