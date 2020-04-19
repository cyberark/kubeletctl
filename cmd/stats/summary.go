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
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"

	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Return performance stats summary of node, pods and containers.",
	Long: `Description:
  Return performance stats summary of node, pods and containers.
  
  HTTP requests:
    GET /stats/summary
    GET /stats/summary?only_cpu_and_memory=true
  
  Example for usage:
    // Query only the cpu and memory fields
    kubeletctl stats summary --only-cpu-mem
  
  With curl:
    curl -k https://<node_ip>:10250/stats/summary
    curl -k https://<node_ip>:10250/stats/summary?only_cpu_and_memory=true`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("summary called")
		apiPathUrl := cmd.ServerFullAddressGlobal + api.STATS_SUMMARY
		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
		cmd.PrintPrettyHttpResponse(resp, err)
	},
}

var onlyCpuAndMemoryFlag bool
func init() {
	statsCmd.AddCommand(summaryCmd)
	// TODO: maybe instead of using the flag, let the use of arguments and it will be scalability
	// Is this is the only flag or there are more ?
	// https://github.com/kubernetes/kubernetes/blob/4fda1207e347af92e649b59d60d48c7021ba0c54/pkg/kubelet/metrics/metrics.go#L34-L87
	summaryCmd.PersistentFlags().BoolVarP(&onlyCpuAndMemoryFlag, "only-cpu-mem", "", false, "Added the query \"?only_cpu_and_memory=true\" to the request")
}
