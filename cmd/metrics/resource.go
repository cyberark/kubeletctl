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
	"kubeletctl/pkg/utils"

	"github.com/spf13/cobra"
)

// resourceCmd represents the resource command
var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Return information about node resources",
	Long: `Description:
  Return information about node resources.
  
  HTTP requests: 
    GET /metrics/resource
    GET /metrics/resource/v1alpha1

  Example for usage:
    kubeletctl metrics resource
    kubeletctl metrics resource v1alpha1
  
  With curl:
    curl -k https://<node_ip>:10250/metrics/resource
    curl -k https://<node_ip>:10250/metrics/resource/v1alpha1`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("resource called")
		inputArgs := ""
		if utils.IsNotArgsEmpty(args) {
			inputArgs = args[0]
		}

		apiPathUrl := cmd.ServerFullAddressGlobal + api.METRICS_RESOURCE + inputArgs
		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
		cmd.PrintPrettyHttpResponse(resp, err)
	},
}

func init() {
	metricsCmd.AddCommand(resourceCmd)
}
