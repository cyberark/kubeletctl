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
package log

import (
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"
	"kubeletctl/pkg/utils"

	"github.com/spf13/cobra"
)

// logsCmd represents the log command
var logsCmd = &cobra.Command{
	Use:   "log",
	Short: "Return the log from the node.",
	Long: `Description:
  Returns the log from the node. 
  HTTP request: 
    GET /log
    GET /log/{subpath}
  Example for usage:
    kubeletctl pods

  With curl:
    curl -k https://<node_ip>:10250/log`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("log called")

		inputArgs := ""
		if utils.IsNotArgsEmpty(args) {
			inputArgs = args[0]
		}

		// TODO: add support to /{logpath:*}
		apiPathUrl := cmd.ServerFullAddressGlobal + api.LOGS + "/" + inputArgs
		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
		cmd.PrintPrettyHttpResponse(resp, err)

	},
}

func init() {
	cmd.RootCmd.AddCommand(logsCmd)
}
