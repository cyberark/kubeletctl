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
	"github.com/spf13/cobra"
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"
)

// attachCmd represents the attach command
var attachCmd = &cobra.Command{
	Use:   "attach <command> -c <container> -p <pod> -n <namespace>",
	Short: "Attach to a container",
	Long: `Description:
  Attach to a container.
  
  HTTP requests:
    GET  /attach/{podNamespace}/{podID}/{containerName}
    POST /attach/{podNamespace}/{podID}/{containerName}
    GET  /attach/{podNamespace}/{podID}/{uid}/{containerName}
    POST /attach/{podNamespace}/{podID}/{uid}/{containerName}

  Example for usage:
    // The default namespace is "default"
    kubeletctl attach -p <pod> -c <container> -n <namespace>
  
  
  With curl:
    curl -k https://<node_ip>:10250/attach/default/mypod/nginx?input=1&output=1&tty=1`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("attach called")
		// TODO: should we use the Exec api for this too?
		var apiPath string
		if cmd.PodUidFlag == "" {
			apiPath = fmt.Sprintf("%s%s/%s/%s/%s", cmd.ServerFullAddressGlobal, api.ATTACH, cmd.NamespaceFlag, cmd.PodFlag, cmd.ContainerFlag)
		} else {
			apiPath = fmt.Sprintf("%s%s/%s/%s/%s/%s", cmd.ServerFullAddressGlobal, api.ATTACH, cmd.NamespaceFlag, cmd.PodFlag, cmd.PodUidFlag, cmd.ContainerFlag)
		}

		apiPath = apiPath + "?input=1&output=1&tty=1"
		resp, err := api.GetRequest(api.GlobalClient, apiPath)
		cmd.PrintPrettyHttpResponse(resp, err)

	},
}

func init() {
	cmd.RootCmd.AddCommand(attachCmd)
}
