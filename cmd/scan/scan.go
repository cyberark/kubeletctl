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
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"
	"kubeletctl/pkg/utils"
	"net/http"
)

func printVulnerableNodes(nodesUrls []string) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Node IP"})

	for _, node := range nodesUrls {
		tw.AppendRow([]interface{}{node})
	}

	tw.SetTitle("Nodes with opened Kubelet API")
	tw.SetStyle(table.StyleLight)
	tw.Style().Title.Align = text.AlignCenter
	tw.Style().Options.SeparateRows = true
	tw.SetAutoIndex(true)
	fmt.Println(tw.Render())
}

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scans for nodes with opened kubelet API",
	Long: `Description:
  Scans for nodes with opened kubelet's API.

  Examples:
    // It will find all nodes that have opened kubelet API
    kubeletctl scan --cidr "123.123.123.123/24"`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("scan called")
		//api.GlobalClient.Timeout = 3 * time.Second

		if cidrFlag != "" {
			nodesIPs := utils.FindOpenedKubeletOnNodes(cidrFlag)
			if len(nodesIPs) > 0 {
				printVulnerableNodes(nodesIPs)
			}
		} else if cmd.ServerIpAddressFlag != "" {
			apiPathUrl := cmd.ServerFullAddressGlobal + api.HEALTHZ
			resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
			if err == nil && resp.StatusCode == http.StatusOK {
				printVulnerableNodes([]string{cmd.ServerIpAddressFlag})
			}
		} else {
			fmt.Println("[*] No flag was chosen")
		}
	},
}

var cidrFlag string

func init() {
	cmd.RootCmd.AddCommand(scanCmd)
	cmd.RootCmd.PersistentFlags().StringVarP(&cidrFlag, "cidr", "", "", "A network of IP addresses (Example: x.x.x.x/24)")
}
