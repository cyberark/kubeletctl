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
	"kubeletctl/pkg/utils"
)

func printNodesWithRCEContainers(nodes []utils.Node) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Node IP", "Pods", "Namespace", "Containers", "RCE"})
	tw.AppendHeader(table.Row{"", "", "", "", "RUN"})
	for _, node := range nodes {
		printNodeOnce := true
		nodeIpTemp := ""
		if printNodeOnce {
			nodeIpTemp = node.IPAddress
			printNodeOnce = false
		}

		for _, pod := range node.Pods {
			printOnce := true
			for _, container := range pod.Containers {

				podTemp := ""
				namespaceTemp := ""
				if printOnce {
					podTemp = pod.Name
					namespaceTemp = pod.Namespace
					printOnce = false
				}

				// TODO: support exec check also
				//rceExec := "-"
				rceRun := "-"

				//if container.RCEExec {
				//	rceExec = "+"
				//}

				if container.RCERun {
					rceRun = "+"
				}
				tw.AppendRow(table.Row{nodeIpTemp, podTemp, namespaceTemp, container.Name, rceRun})
				nodeIpTemp = ""
			}
		}
	}

	tw.SetTitle("Node with pods vulnerable to RCE")
	tw.SetStyle(table.StyleLight)
	tw.Style().Title.Align = text.AlignCenter
	tw.Style().Options.SeparateRows = true
	tw.SetAutoIndex(true)
	fmt.Println(tw.Render())
}

// rceCmd represents the rce command
var rceCmd = &cobra.Command{
	Use:   "rce",
	Short: "Scans for nodes with opened kubelet API",
	Long: `Description:
  Scans for nodes with opened kubelet's API and remote code execution on their containers.

  Examples:
	// Searching for containers with RCE over one node
    kubeletctl scan rce --server 123.123.123.123

	// Searching for containers with RCE over multiple nodes
    kubeletctl scan rce --cidr "123.123.123.123/24"`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("rce called")

		if cmd.ServerIpAddressFlag != "" {
			node := utils.FindContainersWithRCE(cmd.ServerIpAddressFlag)
			nodes := []utils.Node{node}
			printNodesWithRCEContainers(nodes)
		} else {
			fmt.Println("[*] No flag was chosen")
		}
	},
}

func init() {
	scanCmd.AddCommand(rceCmd)
}
