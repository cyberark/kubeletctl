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

// https://kubernetes.io/docs/tasks/access-application-cluster/port-forward-access-application-cluster/#forward-a-local-port-to-a-port-on-the-pod
// portForwardCmd represents the portForward command
var portForwardCmd = &cobra.Command{
	Use:   "portForward <command> -c <container> -p <pod> -n <namespace>",
	Short: "Attach to a container",
	Long: `Description:
  Attach to a container.
  
  HTTP requests:
    GET  /portForward/{podNamespace}/{podID}/{containerName}
    POST /portForward/{podNamespace}/{podID}/{containerName}
    GET  /portForward/{podNamespace}/{podID}/{uid}/{containerName}
    POST /portForward/{podNamespace}/{podID}/{uid}/{containerName}

  Example for usage:
    // The default namespace is "default"
    kubeletctl portForward -p <pod> -c <container> -n <namespace>
  
  
  With curl:
    curl -k -X POST https://<node_ip>:10250/portForward/default/mypod/nginx?input=1&output=1&tty=1`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("portForward called")
		// TODO: test it, where is the port number should be specified ?
		var apiPath string
		if cmd.PodUidFlag == "" {
			apiPath = fmt.Sprintf("%s%s/%s/%s/%s", cmd.ServerFullAddressGlobal, api.PORT_FORWARD, cmd.NamespaceFlag, cmd.PodFlag, cmd.ContainerFlag)
		} else {
			apiPath = fmt.Sprintf("%s%s/%s/%s/%s/%s", cmd.ServerFullAddressGlobal, api.PORT_FORWARD, cmd.NamespaceFlag, cmd.PodFlag, cmd.PodUidFlag, cmd.ContainerFlag)
		}

		apiPath = apiPath + "?input=1&output=1&tty=1"
		resp, err := api.GetRequest(api.GlobalClient, apiPath)
		cmd.PrintPrettyHttpResponse(resp, err)
	},
}

func init() {
	cmd.RootCmd.AddCommand(portForwardCmd)
}
