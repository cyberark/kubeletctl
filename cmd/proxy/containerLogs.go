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
	"fmt"
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"

	"github.com/spf13/cobra"
)

// containerLogsCmd represents the containerLogs command
var containerLogsCmd = &cobra.Command{
	Use:   "containerLogs -p <pod> -c <container> -n <namespace>",
	Short: "Return container log",
	Long: `Description:
  Return container log. 

  HTTP request: 
    GET /containerLogs/{podNamespace}/{podID}/{containerName}
  Example for usage:
    kubeletctl containerLogs -p <pod> -c <container> -n <namespace>
  
  With curl:
    curl -k https://<node_ip>:10250/containerLogs/<podNamespace>/<podID>/<containerName>`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("containerLogs called")

		apiPathUrl := cmd.ServerFullAddressGlobal + api.CONTAINER_LOGS
		apiPathUrl = fmt.Sprintf("%s%s/%s/%s/%s", cmd.ServerFullAddressGlobal, api.CONTAINER_LOGS, cmd.NamespaceFlag, cmd.PodFlag, cmd.ContainerFlag)

		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
		cmd.PrintPrettyHttpResponse(resp, err)

	},
}

func init() {
	cmd.RootCmd.AddCommand(containerLogsCmd)
}
