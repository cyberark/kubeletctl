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
	"kubeletctl/pkg/utils"
	"os"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run <command> -c <container> -p <pod> -n <namespace>",
	Short: "Run commands inside a container",
	Long: `Description:
  Run commands inside a container.
  
  HTTP requests:
    POST /run/{podNamespace}/{podID}/{containerName}
    POST /run/{podNamespace}/{podID}/{uid}/{containerName}
  
  The body of the HTTP request:
    "cmd={command}"
  
  Example for usage:
    // The default namespace is "default"
    kubeletctl run "ls /" -p <pod> -c <container> -n <namespace>
  
  
  With curl:
    curl -k -XPOST https://<node_ip>:10250/run/default/mypod/nginx -d "cmd=ls /"`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("run called")
		// run "ls /" -c nginx -p my-nginx-6b474476c4-rfgn2 -n default

		inputArgs := ""
		if utils.IsNotArgsEmpty(args) {
			inputArgs = args[0]
		}

		// TODO: check if it can handle multiple commands like: "bin/bash -c “/bin/bash”
		// https://www.openshift.com/blog/executing-commands-in-pods-using-k8s-api
		var apiPathUrl string
		if utils.AreNamespacePodAndContainerFlagsSet(cmd.NamespaceFlag, cmd.PodFlag, cmd.ContainerFlag) {
			apiPathUrl = fmt.Sprintf("%s%s/%s/%s/%s", cmd.ServerFullAddressGlobal, api.RUN, cmd.NamespaceFlag, cmd.PodFlag, cmd.ContainerFlag)
		} else if cmd.PodUidFlag != "" {
			apiPathUrl = fmt.Sprintf("%s%s/%s/%s/%s/%s", cmd.ServerFullAddressGlobal, api.RUN, cmd.NamespaceFlag, cmd.PodFlag, cmd.PodUidFlag, cmd.ContainerFlag)
		} else if allPodsFlag {
			fmt.Println("[*] Running command on all pods")
		} else {
			fmt.Println("[*] Missing some flags, exiting")
			os.Exit(1)
		}

		var command string
		if inputArgs != "" {
			command = "cmd=" + inputArgs
		} else {
			fmt.Println("[*] No command was set, setting default command 'ls /'")
			command = "cmd=ls /"
		}

		if allPodsFlag {
			utils.RunCommandOnAllPodsInANode(cmd.ServerIpAddressFlag, command)
		} else {
			resp, err := api.PostRequest(api.GlobalClient, apiPathUrl, []byte(command))
			cmd.PrintPrettyHttpResponse(resp, err)
		}
	},
}

var allPodsFlag bool
func init() {
	cmd.RootCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().BoolVarP(&allPodsFlag, "all-pods", "", false, "It will search for all the pods in the node and run the command on everyone")
}
