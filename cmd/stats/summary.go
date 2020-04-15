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

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("summary called")
		apiPathUrl := cmd.ServerFullAddressGlobal + api.STATS_SUMMARY
		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
		cmd.PrintPrettyHttpResponse(resp, err)
	},
}

var onlyCpuAndMemoryFlag bool
func init() {
	statsCmd.AddCommand(summaryCmd)
	summaryCmd.PersistentFlags().BoolVarP(&onlyCpuAndMemoryFlag, "only-cpu-mem", "", false, "Added the query \"?only_cpu_and_memory=true\" to the request")
}
