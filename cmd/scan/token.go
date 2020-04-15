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
package scan

import (
	"fmt"
	"github.com/spf13/cobra"
	"kubeletctl/cmd"
	"kubeletctl/pkg/utils"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token --server <node_ip>",
	Short: "Scans for for all the tokens in a given Node.",
	Long: `Description:
  Searching for all the pods and containers in a given node. 
  For each of the containers it will search for the token and decoded it.

  Examples:
    kubeletctl scan token --server 123.123.123.123"`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("token called")

		// TODO: in some container the path might start with /run/secrets.. need to add such case
		if cmd.ServerIpAddressFlag != "" {
			utils.GetTokensFromAllPods(cmd.ServerIpAddressFlag)
		} else {
			fmt.Println("[*] Server IP address was not specified")
		}
	},
}

func init() {
	scanCmd.AddCommand(tokenCmd)
}
