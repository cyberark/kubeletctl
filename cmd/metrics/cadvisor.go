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

// cadvisorCmd represents the cadvisor command
var cadvisorCmd = &cobra.Command{
	Use:   "cadvisor",
	Short: "Return information about resource usage and performance characteristics of running containers",
	Long: `Description:
  Return information about resource usage and performance characteristics of running containers.
  
  HTTP requests:
    GET /metrics/cadvisor
  
  Example for usage:
    kubeletctl metrics cadvisor
  
  With curl:
    curl -k https://<node_ip>:10250/metrics/cadvisor`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("cadvisor called")

		apiPathUrl := cmd.ServerFullAddressGlobal + api.METRICS_CADVISOR
		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
		cmd.PrintPrettyHttpResponse(resp, err)
	},
}

func init() {
	metricsCmd.AddCommand(cadvisorCmd)
}
