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
package metrics

import (
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"

	"github.com/spf13/cobra"
)

// metricsCmd represents the metrics command
var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Return information about node CPU and memory usage",
	Long: `Description:
  Return information about node CPU and memory usage.
  
  HTTP requests: 
    GET /metrics
    GET /metrics/cadvisor
    GET /metrics/probes
  
  Example for usage:
    kubeletctl metrics
    kubeletctl metrics cadvisor
    kubeletctl metrics probes
    kubeletctl metrics resource v1alpha1
  
  With curl:
    curl -k https://<node_ip>:10250/metrics
    curl -k https://<node_ip>:10250/metrics/cadvisor
    curl -k https://<node_ip>:10250/probes
    curl -k https://<node_ip>:10250/metrics/resource/v1alpha1`,

	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("metrics called")

		apiPathUrl := cmd.ServerFullAddressGlobal + api.METRICS
		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
		cmd.PrintPrettyHttpResponse(resp, err)
	},
}

func init() {
	cmd.RootCmd.AddCommand(metricsCmd)
}
