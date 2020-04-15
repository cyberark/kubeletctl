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
	"strings"
)


func commandMaker(command string, args []string) string{
	//commandString := "command="
	commandString := command + "="
	var fullCommand string

	commands := strings.Split(args[0], " ")
	for index, command := range commands {
		fullCommand += commandString + command

		if (len(commands)-1) != index{
			fullCommand += "&"
		}
	}

	return fullCommand
}



// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec <command> -c <container> -p <pod> -n <namespace>",
	Short: "Run commands inside a container",
	Long: `Description:
  Run commands inside a container.
  
  HTTP requests:
    GET  /exec/{podNamespace}/{podID}/{containerName}?command={command}/&input=1&output=1&tty=1"
    POST /exec/{podNamespace}//{containerName}?command={command}/&input=1&output=1&tty=1"
    GET  /exec/{podNamespace}/{podID}/{uid}/{containerName}?command={command}/&input=1&output=1&tty=1"
    POST /exec/{podNamespace}/{podID}/{uid}/{containerName}?command={command}/&input=1&output=1&tty=1"
  
  
  Example for usage:
    // The default namespace is "default"
    kubeletctl exec "ls /" -p <pod> -c <container> -n <namespace>
  
  
  With curl:
    curl -k -H "Connection: Upgrade" 
            -H "Upgrade: SPDY/3.1" 
            -H "X-Stream-Protocol-Version: v2.channel.k8s.io" 
            -H "X-Stream-Protocol-Version: channel.k8s.io" 
            -X POST "https://<node_ip>:10250/exec/<podNamespace>/<podID>/<containerName>?command=ls&command=/&input=1&output=1&tty=1"
  
  This request will return status 101 Switching Protocols, therefore you will need to use client that can handle SPDY client.`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("exec called")
		// https://bugzilla.redhat.com/show_bug.cgi?id=1509228

		var apiPath string
		if cmd.PodUidFlag == "" {
			apiPath = fmt.Sprintf("%s/%s/%s/%s", api.EXEC, cmd.NamespaceFlag, cmd.PodFlag, cmd.ContainerFlag)
		} else {
			apiPath = fmt.Sprintf("%s/%s/%s/%s/%s", api.EXEC, cmd.NamespaceFlag, cmd.PodFlag, cmd.PodUidFlag, cmd.ContainerFlag)
		}

		api.Exec(cmd.ServerIpAddressFlag, cmd.PortFlag, cmd.ServerFullAddressGlobal, apiPath,  commandMaker("command", args),"POST")
	},
}


var execCommandLineFlag string

func init() {
	cmd.RootCmd.AddCommand(execCmd)
}
