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
	"github.com/spf13/cobra"
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"
)

// criCmd represents the cri command
var criCmd = &cobra.Command{
	Use:   "cri",
	Short: "Run commands inside a container through the Container Runtime Interface (CRI)",
	Long: `Description:
  Run commands inside a container through the Container Runtime Interface (CRI).
  This is already implemented as part of the "/exec". 
  This might be *deprecated*, but in the past after executing the "/exec" command you received a 302 redirection id,
  you then use this id to together with the command you want. 
  
  In the past you would needed to run something like that to receive 302 redirection value:
    curl -k -v -H "X-Stream-Protocol-Version: v2.channel.k8s.io" \
    		   -H "X-Stream-Protocol-Version: channel.k8s.io" \
    		   -X POST "https://<node_ip>:10250/exec/<namespace>/<pod>/<container>?command=env&input=1&output=1&tty=1"
  
  It opened a stream which you can access using wscat:
    wscat -c "https://<node_ip>:10250/cri/exec/<valueFrom302>" --no-check
  
  You can also use it with new versions of cURL:
    curl -k --include \
    	 --no-buffer \
    	 --header "Connection: Upgrade" \
    	 --header "Upgrade: websocket" \
    	 --header "Sec-WebSocket-Key: <base64_key>" \
    	 --header "Sec-WebSocket-Version: 13" \
    	 https://<node_ip>:10250/cri/exec/<valueFrom302>
  
  HTTP request: 
    GET /cri/exec/{valueFrom302}?cmd={command}
  
  Example for usage:
    kubeletctl cri <command> -value302 <id>
    kubeletctl cri "env" -value302 <id>
  
  With curl:
    curl -k https://<node_ip>:10250/cri/exec/<valuefrom302>?cmd=<command>
    curl -k https://<node_ip>:10250/cri/exec/123456avcdef?cmd=echo+foo`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("cri called")

		apiPathUrl := cmd.ServerFullAddressGlobal + api.CRI + "/" + value302 + "?" + commandMaker("cmd", args)
		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
		cmd.PrintPrettyHttpResponse(resp, err)
	},
}

var value302 string

func init() {
	cmd.RootCmd.AddCommand(criCmd)
	pf := criCmd.PersistentFlags()
	pf.StringVarP(&value302, "value302", "", "", "302 redirection ID from using /exec")
	cobra.MarkFlagRequired(pf, "value302")
}
