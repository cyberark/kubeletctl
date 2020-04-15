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

// containerCmd represents the container command
var containerCmd = &cobra.Command{
	Use:   "container",
	Short: "Return performance stats of containers.",
	Long: `Description:
  Return performance stats of containers.
  
  HTTP requests:
    GET /stats/container
  
  Example for usage:
    // Stats of the contianers
    kubeletctl stats container

  With curl:
    curl -k https://<node_ip>:10250/stats/container`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("container called")

		apiPathUrl := cmd.ServerFullAddressGlobal + api.STATS_CONTAINER
		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
		cmd.PrintPrettyHttpResponse(resp, err)
	},
}

func init() {
	statsCmd.AddCommand(containerCmd)
}
