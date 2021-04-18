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
package healthz

import (
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"
	"kubeletctl/pkg/utils"

	"github.com/spf13/cobra"
)

// healthzCmd represents the healthz command
var healthzCmd = &cobra.Command{
	Use:   "healthz <command>",
	Short: "Check the state of the node",
	Long: `Description:
  Check the state of the node. If everything ok it should return "ok". 
  HTTP request: 
    GET /healthz
    GET /healthz/log
    GET /healthz/ping
    GET /healthz/syncloop

  Example for usage:
    kubeletctl healthz
    kubeletctl healthz ping

  With curl:
    curl -k https://<node_ip>:10250/healthz
    curl -k https://<node_ip>:10250/log
    curl -k https://<node_ip>:10250/ping
    curl -k https://<node_ip>:10250/syncloop`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("healthz called")

		inputArgs := ""
		if utils.IsNotArgsEmpty(args) {
			inputArgs = "/" + args[0]
		}

		// We didn't define the commands as cobra's commands, the user can type any commands for scalability.
		apiPathUrl := cmd.ServerFullAddressGlobal + api.HEALTHZ + inputArgs
		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
		cmd.PrintPrettyHttpResponse(resp, err)
	},
}

func init() {
	cmd.RootCmd.AddCommand(healthzCmd)
}
