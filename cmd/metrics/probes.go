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
	"github.com/spf13/cobra"
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"
)

// probesCmd represents the probes command
var probesCmd = &cobra.Command{
	Use:   "probes",
	Short: "Return information about node probes",
	Long: `Description:
  Return information about node probes.
  
  HTTP requests: 
    GET /metrics/resource/probes

  Example for usage:
    kubeletctl metrics resource probes

  With curl:
    curl -k https://<node_ip>:10250/metrics/resource/probes`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("probes called")

		apiPathUrl := cmd.ServerFullAddressGlobal + api.METRICS_PROBES
		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
		cmd.PrintPrettyHttpResponse(resp, err)
	},
}

func init() {
	metricsCmd.AddCommand(probesCmd)
}
